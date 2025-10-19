package interceptors

import (
	"net/http"
	"time"

	"connectrpc.com/connect"
)

// ExampleInterceptorChain shows how to use the bizctx cookie interceptors
func ExampleInterceptorChain() []connect.Interceptor {
	// Configure the cookie settings
	cookieConfig := BizCtxCookieConfig{
		SessionCookieName:      "v1_session_id",
		OrganizationCookieName: "v1_org_id",
		CookieDomain:           ".yourdomain.com", // Set your domain
		CookiePath:             "/",
		CookieMaxAge:           24 * time.Hour,
		CookieSecure:           true, // Set to false for development
		CookieHttpOnly:         true,
		CookieSameSite:         http.SameSiteStrictMode,
	}

	// Create the interceptor chain
	return []connect.Interceptor{
		// Recovery should be first to catch any panics
		RecoveryInterceptor(),

		// Logging for debugging
		LoggingInterceptor(),

		// Read cookies and set bizctx values (early in the chain)
		BizCtxCookieReaderInterceptor(cookieConfig),

		// Your auth interceptor (can now use bizctx values)
		// AuthInterceptor(authConfig),

		// Rate limiting
		// RateLimitInterceptor(rateLimitConfig),

		// Metrics collection
		MetricsInterceptor(),

		// Write bizctx values to cookies (late in the chain, after business logic)
		BizCtxCookieWriterInterceptor(cookieConfig),

		// Validation
		ValidationInterceptor(),
	}
}

// ExampleCombinedInterceptor shows how to use the combined cookie interceptor
func ExampleCombinedInterceptor() []connect.Interceptor {
	cookieConfig := DefaultBizCtxCookieConfig()

	// Customize for your environment
	cookieConfig.SessionCookieName = "v1_session"
	cookieConfig.OrganizationCookieName = "v1_org"
	cookieConfig.CookieSecure = false // For development

	return []connect.Interceptor{
		RecoveryInterceptor(),
		LoggingInterceptor(),

		// This interceptor handles both reading and writing cookies
		BizCtxCookieInterceptor(cookieConfig),

		MetricsInterceptor(),
	}
}
