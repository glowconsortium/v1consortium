package interceptors

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
	"v1consortium/internal/service"

	"connectrpc.com/connect"
	"github.com/golang-jwt/jwt/v5"
)

// LoggingInterceptor logs Connect RPC calls with detailed information
func LoggingInterceptor() connect.UnaryInterceptorFunc {
	return connect.UnaryInterceptorFunc(func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			start := time.Now()

			// Extract client info from headers
			clientIP := ""
			userAgent := ""
			if headers := req.Header(); headers != nil {
				if xff := headers.Get("X-Forwarded-For"); xff != "" {
					clientIP = strings.Split(xff, ",")[0]
				} else if xri := headers.Get("X-Real-IP"); xri != "" {
					clientIP = xri
				}
				userAgent = headers.Get("User-Agent")
			}

			// Call the next handler
			resp, err := next(ctx, req)

			// Log the request with detailed info
			duration := time.Since(start)
			status := "OK"
			errorMsg := ""
			if err != nil {
				status = "ERROR"
				var connectErr *connect.Error
				if errors.As(err, &connectErr) {
					status = fmt.Sprintf("ERROR_%s", connectErr.Code().String())
					errorMsg = connectErr.Message()
				} else {
					errorMsg = err.Error()
				}
			}

			logLine := fmt.Sprintf("Connect RPC: %s %s %s %s",
				req.Spec().Procedure,
				status,
				duration,
				clientIP,
			)

			if errorMsg != "" {
				logLine += fmt.Sprintf(" error=%s", errorMsg)
			}

			if userAgent != "" {
				logLine += fmt.Sprintf(" user_agent=%q", userAgent)
			}

			log.Println(logLine)

			return resp, err
		})
	})
}

// JWTConfig holds JWT configuration for interceptors
type JWTConfig struct {
	SecretKey       []byte
	PublicEndpoints []string // Endpoints that don't require authentication
	RequiredRole    string   // Optional: required role for access
}

// AuthInterceptor validates JWT tokens for Connect RPC calls
func AuthInterceptor(config JWTConfig) connect.UnaryInterceptorFunc {
	return connect.UnaryInterceptorFunc(func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			procedure := req.Spec().Procedure

			// Check if endpoint is public
			for _, publicEndpoint := range config.PublicEndpoints {
				if procedure == publicEndpoint {
					log.Printf("Auth interceptor: allowing public endpoint %s", procedure)
					return next(ctx, req)
				}
			}

			log.Printf("Auth interceptor: checking auth for procedure %s", procedure)

			// Extract token from Authorization header
			authHeader := req.Header().Get("Authorization")
			if authHeader == "" {
				return nil, connect.NewError(
					connect.CodeUnauthenticated,
					fmt.Errorf("authorization header required"),
				)
			}

			// Check Bearer prefix
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				return nil, connect.NewError(
					connect.CodeUnauthenticated,
					fmt.Errorf("invalid authorization header format"),
				)
			}

			tokenString := parts[1]

			// Parse and validate token
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				// Validate signing method
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, connect.NewError(
						connect.CodeUnauthenticated,
						fmt.Errorf("invalid token signing method"),
					)
				}
				return config.SecretKey, nil
			})

			if err != nil || !token.Valid {
				return nil, connect.NewError(
					connect.CodeUnauthenticated,
					fmt.Errorf("invalid token: %w", err),
				)
			}

			// Extract claims
			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				return nil, connect.NewError(
					connect.CodeUnauthenticated,
					fmt.Errorf("invalid token claims"),
				)
			}

			// Check required role if specified
			if config.RequiredRole != "" {
				role, exists := claims["role"]
				if !exists || role != config.RequiredRole {
					return nil, connect.NewError(
						connect.CodePermissionDenied,
						fmt.Errorf("insufficient permissions"),
					)
				}
			}

			// Add user info to context for downstream handlers
			ctx = context.WithValue(ctx, "user_id", claims["sub"])
			ctx = context.WithValue(ctx, "user_email", claims["email"])
			if role, exists := claims["role"]; exists {
				ctx = context.WithValue(ctx, "user_role", role)
			}

			// Continue with authenticated context
			return next(ctx, req)
		})
	})
}

// RateLimitConfig holds rate limiting configuration
type RateLimitConfig struct {
	RequestsPerMinute int
	BurstSize         int
}

// Simple in-memory rate limiter (for production, use Redis or similar)
type rateLimiter struct {
	requests    map[string][]time.Time
	maxRequests int
	window      time.Duration
}

func newRateLimiter(requestsPerMinute int) *rateLimiter {
	return &rateLimiter{
		requests:    make(map[string][]time.Time),
		maxRequests: requestsPerMinute,
		window:      time.Minute,
	}
}

func (rl *rateLimiter) isAllowed(clientID string) bool {
	now := time.Now()

	// Clean old requests
	if requests, exists := rl.requests[clientID]; exists {
		var validRequests []time.Time
		cutoff := now.Add(-rl.window)

		for _, reqTime := range requests {
			if reqTime.After(cutoff) {
				validRequests = append(validRequests, reqTime)
			}
		}
		rl.requests[clientID] = validRequests
	}

	// Check if under limit
	if len(rl.requests[clientID]) >= rl.maxRequests {
		return false
	}

	// Add current request
	rl.requests[clientID] = append(rl.requests[clientID], now)
	return true
}

// RateLimitInterceptor implements rate limiting for Connect RPC calls
func RateLimitInterceptor(config RateLimitConfig) connect.UnaryInterceptorFunc {
	limiter := newRateLimiter(config.RequestsPerMinute)

	return connect.UnaryInterceptorFunc(func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			// Extract client identifier (IP or user ID)
			clientID := "unknown"

			// Try to get user ID from context first
			if userID := ctx.Value("user_id"); userID != nil {
				clientID = fmt.Sprintf("user:%s", userID)
			} else {
				// Fall back to IP address
				if xff := req.Header().Get("X-Forwarded-For"); xff != "" {
					clientID = fmt.Sprintf("ip:%s", strings.Split(xff, ",")[0])
				} else if xri := req.Header().Get("X-Real-IP"); xri != "" {
					clientID = fmt.Sprintf("ip:%s", xri)
				}
			}

			// Check rate limit
			if !limiter.isAllowed(clientID) {
				return nil, connect.NewError(
					connect.CodeResourceExhausted,
					fmt.Errorf("rate limit exceeded for client %s", clientID),
				)
			}

			return next(ctx, req)
		})
	})
}

// MetricsInterceptor collects basic metrics for Connect RPC calls
func MetricsInterceptor() connect.UnaryInterceptorFunc {
	return connect.UnaryInterceptorFunc(func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			start := time.Now()

			// Call the next handler
			resp, err := next(ctx, req)

			// Record metrics
			duration := time.Since(start)
			procedure := req.Spec().Procedure

			// Log metrics (in production, you'd send to a metrics service)
			status := "success"
			if err != nil {
				status = "error"
			}

			log.Printf("METRICS: procedure=%s status=%s duration=%s",
				procedure, status, duration)

			return resp, err
		})
	})
}

// ValidationInterceptor performs request validation
func ValidationInterceptor() connect.UnaryInterceptorFunc {
	return connect.UnaryInterceptorFunc(func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			// Basic request size validation
			const maxRequestSize = 1024 * 1024 // 1MB
			if req.Header().Get("Content-Length") != "" {
				// You could parse Content-Length here and validate
			}

			// Validate request message (if it implements a Validate method)
			if validator, ok := req.Any().(interface{ Validate() error }); ok {
				if err := validator.Validate(); err != nil {
					return nil, connect.NewError(
						connect.CodeInvalidArgument,
						fmt.Errorf("request validation failed: %w", err),
					)
				}
			}

			return next(ctx, req)
		})
	})
}

// RecoveryInterceptor handles panics and converts them to Connect errors
func RecoveryInterceptor() connect.UnaryInterceptorFunc {
	return connect.UnaryInterceptorFunc(func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (resp connect.AnyResponse, err error) {
			defer func() {
				if r := recover(); r != nil {
					log.Printf("PANIC in Connect RPC %s: %v", req.Spec().Procedure, r)
					err = connect.NewError(
						connect.CodeInternal,
						fmt.Errorf("internal server error"),
					)
				}
			}()

			return next(ctx, req)
		})
	})
}

// CORSConfig holds CORS configuration for Connect services
type CORSConfig struct {
	AllowedOrigins   []string
	AllowedHeaders   []string
	AllowCredentials bool
}

// DefaultCORSConfig returns sensible defaults for Connect RPC
func DefaultCORSConfig() CORSConfig {
	return CORSConfig{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{
			"Content-Type",
			"Authorization",
			"Connect-Protocol-Version",
			"Connect-Timeout-Ms",
			"Grpc-Timeout",
			"X-Grpc-Web",
			"X-User-Agent",
		},
		AllowCredentials: false,
	}
}

// CORSInterceptor handles CORS for Connect RPC (though this is usually handled at HTTP level)
func CORSInterceptor(config CORSConfig) connect.UnaryInterceptorFunc {
	return connect.UnaryInterceptorFunc(func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			// CORS is typically handled at the HTTP middleware level
			// This interceptor could add CORS-related context or validation

			origin := req.Header().Get("Origin")
			if origin != "" {
				// Validate origin against allowed origins
				allowed := false
				for _, allowedOrigin := range config.AllowedOrigins {
					if allowedOrigin == "*" || allowedOrigin == origin {
						allowed = true
						break
					}
				}

				if !allowed {
					return nil, connect.NewError(
						connect.CodePermissionDenied,
						fmt.Errorf("origin not allowed: %s", origin),
					)
				}
			}

			return next(ctx, req)
		})
	})
}

// BizCtxCookieConfig holds configuration for business context cookie management
type BizCtxCookieConfig struct {
	SessionCookieName      string        // Name of the session ID cookie
	OrganizationCookieName string        // Name of the organization ID cookie
	CookieDomain           string        // Domain for cookies
	CookiePath             string        // Path for cookies
	CookieMaxAge           time.Duration // Max age for cookies
	CookieSecure           bool          // Secure flag for cookies
	CookieHttpOnly         bool          // HttpOnly flag for cookies
	CookieSameSite         http.SameSite // SameSite attribute for cookies
}

// DefaultBizCtxCookieConfig returns sensible defaults for business context cookies
func DefaultBizCtxCookieConfig() BizCtxCookieConfig {
	return BizCtxCookieConfig{
		SessionCookieName:      "session_id",
		OrganizationCookieName: "organization_id",
		CookiePath:             "/",
		CookieMaxAge:           24 * time.Hour, // 24 hours
		CookieSecure:           true,
		CookieHttpOnly:         true,
		CookieSameSite:         http.SameSiteStrictMode,
	}
}

// BizCtxCookieReaderInterceptor reads values from cookies and sets them in bizctx
func BizCtxCookieReaderInterceptor(config BizCtxCookieConfig) connect.UnaryInterceptorFunc {
	return connect.UnaryInterceptorFunc(func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			// Parse cookies from the Cookie header
			cookieHeader := req.Header().Get("Cookie")
			if cookieHeader == "" {
				// No cookies present, continue without setting context
				return next(ctx, req)
			}

			cookies := parseCookies(cookieHeader)

			// Extract session ID from cookies and set in context
			if sessionID, exists := cookies[config.SessionCookieName]; exists && sessionID != "" {
				ctx = service.BizCtx().SetCurrentSessionID(ctx, sessionID)
				log.Printf("BizCtx Cookie Reader: Set session ID from cookie: %s", sessionID)
			}

			// Extract organization ID from cookies and set in context
			if orgID, exists := cookies[config.OrganizationCookieName]; exists && orgID != "" {
				ctx = service.BizCtx().SetCurrentOrganizationID(ctx, orgID)
				log.Printf("BizCtx Cookie Reader: Set organization ID from cookie: %s", orgID)
			}

			return next(ctx, req)
		})
	})
}

// BizCtxCookieWriterInterceptor reads values from bizctx and sets them as cookies in response
func BizCtxCookieWriterInterceptor(config BizCtxCookieConfig) connect.UnaryInterceptorFunc {
	return connect.UnaryInterceptorFunc(func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			// Call the next handler first
			resp, err := next(ctx, req)
			if err != nil {
				return resp, err
			}

			// Extract values from bizctx and set cookies
			var cookies []string

			// First try to get values from response headers (set by handlers)
			sessionID := resp.Header().Get("X-Session-ID")
			orgID := resp.Header().Get("X-Organization-ID")

			// If not in headers, try to get from context (fallback)
			if sessionID == "" {
				if sid, err := service.BizCtx().GetCurrentSessionID(ctx); err == nil && sid != "" {
					sessionID = sid
				}
			}

			if orgID == "" {
				if oid, err := service.BizCtx().GetCurrentOrganizationID(ctx); err == nil && oid != "" {
					orgID = oid
				}
			}

			// Set session ID cookie if available
			if sessionID != "" {
				cookie := buildCookie(config.SessionCookieName, sessionID, config)
				cookies = append(cookies, cookie)
				log.Printf("BizCtx Cookie Writer: Setting session ID cookie: %s", sessionID)
				// Remove the temporary header
				resp.Header().Del("X-Session-ID")
			}

			// Set organization ID cookie if available
			if orgID != "" {
				cookie := buildCookie(config.OrganizationCookieName, orgID, config)
				cookies = append(cookies, cookie)
				log.Printf("BizCtx Cookie Writer: Setting organization ID cookie: %s", orgID)
				// Remove the temporary header
				resp.Header().Del("X-Organization-ID")
			}

			// Set cookies in response headers
			if len(cookies) > 0 {
				// For Connect RPC, we need to set multiple Set-Cookie headers
				// Each cookie should be a separate Set-Cookie header
				for _, cookie := range cookies {
					resp.Header().Add("Set-Cookie", cookie)
				}
				log.Printf("BizCtx Cookie Writer: Set %d cookies in response", len(cookies))
			}

			return resp, err
		})
	})
}

// Helper function to parse cookies from Cookie header
func parseCookies(cookieHeader string) map[string]string {
	cookies := make(map[string]string)

	// Split by semicolon and parse each cookie
	parts := strings.Split(cookieHeader, ";")
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		// Split by equals sign
		keyValue := strings.SplitN(part, "=", 2)
		if len(keyValue) == 2 {
			key := strings.TrimSpace(keyValue[0])
			value := strings.TrimSpace(keyValue[1])
			cookies[key] = value
		}
	}

	return cookies
}

// Helper function to build cookie string
func buildCookie(name, value string, config BizCtxCookieConfig) string {
	cookie := fmt.Sprintf("%s=%s", name, value)

	if config.CookiePath != "" {
		cookie += fmt.Sprintf("; Path=%s", config.CookiePath)
	}

	if config.CookieDomain != "" {
		cookie += fmt.Sprintf("; Domain=%s", config.CookieDomain)
	}

	if config.CookieMaxAge > 0 {
		maxAge := int(config.CookieMaxAge.Seconds())
		cookie += fmt.Sprintf("; Max-Age=%d", maxAge)
	}

	if config.CookieSecure {
		cookie += "; Secure"
	}

	if config.CookieHttpOnly {
		cookie += "; HttpOnly"
	}

	switch config.CookieSameSite {
	case http.SameSiteStrictMode:
		cookie += "; SameSite=Strict"
	case http.SameSiteLaxMode:
		cookie += "; SameSite=Lax"
	case http.SameSiteNoneMode:
		cookie += "; SameSite=None"
	}

	return cookie
}

// BizCtxCookieInterceptor combines both reader and writer functionality
func BizCtxCookieInterceptor(config BizCtxCookieConfig) connect.UnaryInterceptorFunc {
	reader := BizCtxCookieReaderInterceptor(config)
	writer := BizCtxCookieWriterInterceptor(config)

	return connect.UnaryInterceptorFunc(func(next connect.UnaryFunc) connect.UnaryFunc {
		// Chain the reader and writer interceptors
		return reader(writer(next))
	})
}
