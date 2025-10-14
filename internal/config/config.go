package config

import (
	"context"
	"net/http"
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

// Config holds all application configuration
type Config struct {
	HTTPAddr     string             `json:"httpAddr"`
	DatabaseURL  string             `json:"databaseUrl"`
	Server       ServerConfig       `json:"server"`
	JWT          JWTConfig          `json:"jwt"`
	CORS         CORSConfig         `json:"cors"`
	RateLimit    RateLimitConfig    `json:"rateLimit"`
	Interceptors InterceptorsConfig `json:"interceptors"`
	BizCtx       BizCtxConfig       `json:"bizCtx"`
	Environment  string             `json:"environment"`
}

// ServerConfig holds HTTP server configuration
type ServerConfig struct {
	ReadTimeout  time.Duration `json:"readTimeout"`
	WriteTimeout time.Duration `json:"writeTimeout"`
	IdleTimeout  time.Duration `json:"idleTimeout"`
}

// JWTConfig holds JWT configuration
type JWTConfig struct {
	SecretKey       string   `json:"secretKey"`
	PublicEndpoints []string `json:"publicEndpoints"`
	RequiredRole    string   `json:"requiredRole"`
}

// CORSConfig holds CORS configuration
type CORSConfig struct {
	AllowedOrigins   []string `json:"allowedOrigins"`
	AllowedMethods   []string `json:"allowedMethods"`
	AllowedHeaders   []string `json:"allowedHeaders"`
	AllowCredentials bool     `json:"allowCredentials"`
}

// RateLimitConfig holds rate limiting configuration
type RateLimitConfig struct {
	Enabled           bool `json:"enabled"`
	RequestsPerMinute int  `json:"requestsPerMinute"`
	BurstSize         int  `json:"burstSize"`
}

// InterceptorsConfig holds interceptor configuration
type InterceptorsConfig struct {
	RecoveryEnabled   bool `json:"recoveryEnabled"`
	LoggingEnabled    bool `json:"loggingEnabled"`
	MetricsEnabled    bool `json:"metricsEnabled"`
	RateLimitEnabled  bool `json:"rateLimitEnabled"`
	ValidationEnabled bool `json:"validationEnabled"`
	AuthEnabled       bool `json:"authEnabled"`
	CORSEnabled       bool `json:"corsEnabled"`
	BizCtxEnabled     bool `json:"bizCtxEnabled"`
}

// BizCtxConfig holds business context cookie configuration
type BizCtxConfig struct {
	SessionCookieName      string `json:"sessionCookieName"`
	OrganizationCookieName string `json:"organizationCookieName"`
	CookieDomain           string `json:"cookieDomain"`
	CookiePath             string `json:"cookiePath"`
	CookieMaxAgeHours      int    `json:"cookieMaxAgeHours"`
	CookieSecure           bool   `json:"cookieSecure"`
	CookieHttpOnly         bool   `json:"cookieHttpOnly"`
	CookieSameSite         string `json:"cookieSameSite"` // "strict", "lax", "none"
}

// Load loads configuration from various sources
func Load() *Config {
	ctx := context.Background()

	// Default configuration
	cfg := &Config{
		HTTPAddr:    getConfigString(ctx, "server.address", ":8000"),
		DatabaseURL: getConfigString(ctx, "database.url", "postgres://postgres:postgres@127.0.0.1:54322/postgres"),
		Environment: getConfigString(ctx, "environment", "development"),

		Server: ServerConfig{
			ReadTimeout:  getConfigDuration(ctx, "server.readTimeout", 15*time.Second),
			WriteTimeout: getConfigDuration(ctx, "server.writeTimeout", 15*time.Second),
			IdleTimeout:  getConfigDuration(ctx, "server.idleTimeout", 60*time.Second),
		},

		JWT: JWTConfig{
			SecretKey: getConfigString(ctx, "jwt.secretKey", "your-super-secret-key-change-in-production"),
			PublicEndpoints: getConfigStringSlice(ctx, "jwt.publicEndpoints", []string{
				"/v1consortium.auth.AuthService/Login",
				"/v1consortium.auth.AuthService/RegisterUser",
				"/health",
				"/swagger.json",
				"/enhanced-swagger.json",
				"/docs",
				"/docs/",
			}),
			RequiredRole: getConfigString(ctx, "jwt.requiredRole", ""),
		},

		CORS: CORSConfig{
			AllowedOrigins: getConfigStringSlice(ctx, "cors.allowedOrigins", []string{"*"}),
			AllowedMethods: getConfigStringSlice(ctx, "cors.allowedMethods", []string{
				"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD", "PATCH",
			}),
			AllowedHeaders: getConfigStringSlice(ctx, "cors.allowedHeaders", []string{
				"Accept", "Accept-Encoding", "Accept-Post", "Content-Encoding",
				"Content-Length", "Content-Type", "Connect-Protocol-Version",
				"Connect-Timeout-Ms", "Grpc-Timeout", "Grpc-Encoding",
				"Grpc-Accept-Encoding", "User-Agent", "X-User-Agent",
				"X-Grpc-Web", "Authorization",
			}),
			AllowCredentials: getConfigBool(ctx, "cors.allowCredentials", false),
		},

		RateLimit: RateLimitConfig{
			Enabled:           getConfigBool(ctx, "rateLimit.enabled", true),
			RequestsPerMinute: getConfigInt(ctx, "rateLimit.requestsPerMinute", 100),
			BurstSize:         getConfigInt(ctx, "rateLimit.burstSize", 10),
		},

		Interceptors: InterceptorsConfig{
			RecoveryEnabled:   getConfigBool(ctx, "interceptors.recoveryEnabled", true),
			LoggingEnabled:    getConfigBool(ctx, "interceptors.loggingEnabled", true),
			MetricsEnabled:    getConfigBool(ctx, "interceptors.metricsEnabled", true),
			RateLimitEnabled:  getConfigBool(ctx, "interceptors.rateLimitEnabled", true),
			ValidationEnabled: getConfigBool(ctx, "interceptors.validationEnabled", true),
			AuthEnabled:       getConfigBool(ctx, "interceptors.authEnabled", false), // Disabled by default for development
			CORSEnabled:       getConfigBool(ctx, "interceptors.corsEnabled", true),
			BizCtxEnabled:     getConfigBool(ctx, "interceptors.bizCtxEnabled", true),
		},

		BizCtx: BizCtxConfig{
			SessionCookieName:      getConfigString(ctx, "bizCtx.sessionCookieName", "v1_session_id"),
			OrganizationCookieName: getConfigString(ctx, "bizCtx.organizationCookieName", "v1_org_id"),
			CookieDomain:           getConfigString(ctx, "bizCtx.cookieDomain", ""),
			CookiePath:             getConfigString(ctx, "bizCtx.cookiePath", "/"),
			CookieMaxAgeHours:      getConfigInt(ctx, "bizCtx.cookieMaxAgeHours", 24),
			CookieSecure:           getConfigBool(ctx, "bizCtx.cookieSecure", true),
			CookieHttpOnly:         getConfigBool(ctx, "bizCtx.cookieHttpOnly", true),
			CookieSameSite:         getConfigString(ctx, "bizCtx.cookieSameSite", "strict"),
		},
	}

	return cfg
}

// GetJWTSecretBytes returns JWT secret as byte slice
func (c *Config) GetJWTSecretBytes() []byte {
	return []byte(c.JWT.SecretKey)
}

// IsDevelopment returns true if running in development
func (c *Config) IsDevelopment() bool {
	return c.Environment == "development" || c.Environment == "dev"
}

// IsProduction returns true if running in production
func (c *Config) IsProduction() bool {
	return c.Environment == "production" || c.Environment == "prod"
}

// Helper functions to get configuration values with defaults
func getConfigString(ctx context.Context, key, defaultValue string) string {
	if value := g.Cfg().MustGet(ctx, key); value.String() != "" {
		return value.String()
	}
	return defaultValue
}

func getConfigInt(ctx context.Context, key string, defaultValue int) int {
	if value := g.Cfg().MustGet(ctx, key); !value.IsEmpty() {
		return value.Int()
	}
	return defaultValue
}

func getConfigBool(ctx context.Context, key string, defaultValue bool) bool {
	if value := g.Cfg().MustGet(ctx, key); !value.IsEmpty() {
		return value.Bool()
	}
	return defaultValue
}

func getConfigDuration(ctx context.Context, key string, defaultValue time.Duration) time.Duration {
	if value := g.Cfg().MustGet(ctx, key); !value.IsEmpty() {
		if duration := value.Duration(); duration > 0 {
			return duration
		}
	}
	return defaultValue
}

func getConfigStringSlice(ctx context.Context, key string, defaultValue []string) []string {
	if value := g.Cfg().MustGet(ctx, key); !value.IsEmpty() {
		return value.Strings()
	}
	return defaultValue
}

// GetSameSiteMode converts string SameSite value to http.SameSite
func (c *BizCtxConfig) GetSameSiteMode() http.SameSite {
	switch c.CookieSameSite {
	case "strict":
		return http.SameSiteStrictMode
	case "lax":
		return http.SameSiteLaxMode
	case "none":
		return http.SameSiteNoneMode
	default:
		return http.SameSiteStrictMode
	}
}

// GetCookieMaxAge returns cookie max age as time.Duration
func (c *BizCtxConfig) GetCookieMaxAge() time.Duration {
	return time.Duration(c.CookieMaxAgeHours) * time.Hour
}
