package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"connectrpc.com/connect"
	"connectrpc.com/vanguard"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	authv1connect "v1consortium/api/auth/v1/v1connect"
	gatewayv1connect "v1consortium/api/gateway/v1/v1connect"
	servicesv1connect "v1consortium/api/services/v1/v1connect"
	"v1consortium/internal/config"
	authconnect "v1consortium/internal/controllerconnect/auth"
	gatewayconnect "v1consortium/internal/controllerconnect/gateway"
	servicesconnect "v1consortium/internal/controllerconnect/services"
	"v1consortium/internal/gateway"
	"v1consortium/internal/pkg/interceptors"
)

var (
	Combined = gcmd.Command{
		Name:  "combined",
		Usage: "combined",
		Brief: "start combined server with vanguard and connect",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			SetupInternalStartupData(ctx)
			return runCombinedServer(ctx)
		},
	}
)

func runCombinedServer(ctx context.Context) error {
	cfg := config.Load()

	log.Printf("🚀 Starting V1 Consortium server on %s", cfg.HTTPAddr)
	log.Printf("📊 Environment: %s", cfg.Environment)

	err := setupRiverDependentServices(ctx)
	if err != nil {
		panic(fmt.Errorf("failed to setup River services: %w", err))
	}

	// Initialize GoFrame server
	s := g.Server()
	s.SetAddr(cfg.HTTPAddr)
	s.SetServerAgent("V1-Consortium/1.0")

	// Setup Connect services and transcoder
	transcoder, err := setupConnectServices(ctx, cfg)
	if err != nil {
		return fmt.Errorf("failed to setup Connect services: %w", err)
	}

	// Setup routes
	setupRoutes(s, cfg, transcoder)

	log.Printf("✅ Server configured successfully")
	log.Println("📡 Protocols: HTTP/1.1, HTTP/2, gRPC, gRPC-Web, Connect")
	log.Printf("📖 Documentation: http://localhost%s/docs/", cfg.HTTPAddr)
	log.Printf("🔍 Health: http://localhost%s/health", cfg.HTTPAddr)

	// Setup graceful shutdown for River resources
	defer func() {
		log.Println("🧹 Cleaning up River dependencies...")
		CleanupRiverDependencies()
	}()

	// Start server (GoFrame handles graceful shutdown automatically)
	s.Run()
	return nil
}

// setupRoutes configures all HTTP routes using GoFrame server
func setupRoutes(s *ghttp.Server, cfg *config.Config, transcoder *vanguard.Transcoder) {
	// Add CORS middleware using GoFrame's native CORS handling
	s.Use(corsMiddleware(cfg))

	// Health check endpoint
	s.BindHandler("/health", func(r *ghttp.Request) {
		r.Response.Header().Set("Content-Type", "application/json")
		r.Response.WriteStatus(http.StatusOK)
		response := fmt.Sprintf(
			`{"status":"healthy","service":"v1-consortium-server","timestamp":"%s","protocols":["http","grpc","grpc-web","connect"],"version":"1.0.0"}`,
			time.Now().UTC().Format(time.RFC3339),
		)
		r.Response.Write(response)
	})

	// API routes - All Connect/gRPC traffic goes through transcoder
	s.BindHandler("/*", func(r *ghttp.Request) {
		transcoder.ServeHTTP(r.Response.ResponseWriter, r.Request)
	})

	// Setup OpenAPI/Swagger documentation
	setupSwaggerRoutes(s, cfg)

	log.Println("✅ Routes configured successfully")
}

// corsMiddleware uses GoFrame's native CORS handling
func corsMiddleware(cfg *config.Config) ghttp.HandlerFunc {
	return func(r *ghttp.Request) {
		// Convert config to GoFrame CORSOptions
		corsOptions := r.Response.DefaultCORSOptions()

		// Handle credentials and origin logic
		if cfg.CORS.AllowCredentials {
			// When credentials are allowed, we cannot use wildcard "*"
			// We must use specific domains
			if len(cfg.CORS.AllowDomain) > 0 {
				corsOptions.AllowDomain = cfg.CORS.AllowDomain
			} else {
				// Use the specific origin from config
				corsOptions.AllowOrigin = cfg.CORS.AllowOrigin
			}
			corsOptions.AllowCredentials = "true"
		} else {
			// When credentials are not allowed, we can use wildcard or specific domains
			if len(cfg.CORS.AllowDomain) > 0 {
				corsOptions.AllowDomain = cfg.CORS.AllowDomain
			} else {
				corsOptions.AllowOrigin = cfg.CORS.AllowOrigin
			}
			corsOptions.AllowCredentials = "false"
		}

		corsOptions.ExposeHeaders = cfg.CORS.ExposeHeaders
		corsOptions.MaxAge = cfg.CORS.MaxAge
		corsOptions.AllowMethods = cfg.CORS.AllowMethods
		corsOptions.AllowHeaders = cfg.CORS.AllowHeaders

		// Apply CORS
		r.Response.CORS(corsOptions)

		// Continue to next handler
		r.Middleware.Next()
	}
} // setupSwaggerRoutes configures OpenAPI/Swagger documentation routes
func setupSwaggerRoutes(s *ghttp.Server, cfg *config.Config) {
	// Find swagger file
	swaggerPath := findSwaggerFile()

	if swaggerPath == "" {
		log.Println("⚠️  Swagger file not found, documentation endpoints disabled")
		return
	}

	// Create OpenAPI enhancer
	openAPIEnhancer := gateway.NewOpenAPIEnhancer(
		swaggerPath,
		cfg.HTTPAddr,
		"1.0.0",
	)

	// Swagger JSON endpoints
	s.BindHandler("/swagger.json", func(r *ghttp.Request) {
		r.Response.Header().Set("Content-Type", "application/json")

		data, err := os.ReadFile(swaggerPath)
		if err != nil {
			log.Printf("Error reading swagger file: %v", err)
			r.Response.WriteStatus(http.StatusNotFound)
			r.Response.Write("Swagger file not found")
			return
		}
		r.Response.Write(data)
	})

	s.BindHandler("/enhanced-swagger.json", func(r *ghttp.Request) {
		openAPIEnhancer.ServeEnhancedSwaggerJSON().ServeHTTP(r.Response.ResponseWriter, r.Request)
	})

	// Swagger UI
	s.BindHandler("/docs/*", func(r *ghttp.Request) {
		openAPIEnhancer.ServeEnhancedSwaggerUI().ServeHTTP(r.Response.ResponseWriter, r.Request)
	})

	s.BindHandler("/docs", func(r *ghttp.Request) {
		r.Response.RedirectTo("/docs/", http.StatusMovedPermanently)
	})

	log.Println("✅ Swagger documentation endpoints configured")
}

// findSwaggerFile locates the swagger JSON file
func findSwaggerFile() string {
	candidates := []string{
		filepath.Join("api", "gen", "openapiv2", "api.swagger.json"),
		"gen/openapiv2/api.swagger.json",
		"manifest/openapi/api.swagger.json",
	}

	for _, path := range candidates {
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}
	return ""
}

// setupConnectServices creates Connect services and Vanguard transcoder
func setupConnectServices(ctx context.Context, cfg *config.Config) (*vanguard.Transcoder, error) {
	log.Println("🔧 Initializing Connect RPC services...")

	// Create service implementations
	authService := authconnect.NewAuthConnectService(ctx)
	gatewayService := gatewayconnect.NewGatewayConnectService(ctx)
	servicesService := servicesconnect.NewServicesConnectService(ctx)

	// Build interceptor chain
	interceptorChain := buildInterceptorChain(cfg)

	// Create Connect handlers
	services := []*vanguard.Service{}

	// Auth service
	authPath, authHandler := authv1connect.NewAuthServiceHandler(
		authService,

		connect.WithInterceptors(interceptorChain...),
	)
	services = append(services, vanguard.NewService(authPath, authHandler))

	// Gateway services
	configPath, configHandler := gatewayv1connect.NewConfigurationServiceHandler(
		gatewayService,
		connect.WithInterceptors(interceptorChain...),
	)
	services = append(services, vanguard.NewService(configPath, configHandler))

	monitoringPath, monitoringHandler := gatewayv1connect.NewMonitoringServiceHandler(
		gatewayService,
		connect.WithInterceptors(interceptorChain...),
	)
	services = append(services, vanguard.NewService(monitoringPath, monitoringHandler))

	// Business services
	businessHandlers := []struct {
		path    string
		handler http.Handler
	}{}

	// Organization Service
	orgPath, orgHandler := servicesv1connect.NewOrganizationServiceHandler(servicesService, connect.WithInterceptors(interceptorChain...))
	businessHandlers = append(businessHandlers, struct {
		path    string
		handler http.Handler
	}{orgPath, orgHandler})

	// Compliance Service
	compliancePath, complianceHandler := servicesv1connect.NewComplianceServiceHandler(servicesService, connect.WithInterceptors(interceptorChain...))
	businessHandlers = append(businessHandlers, struct {
		path    string
		handler http.Handler
	}{compliancePath, complianceHandler})

	// Drug Testing Service
	drugTestPath, drugTestHandler := servicesv1connect.NewDrugTestingServiceHandler(servicesService, connect.WithInterceptors(interceptorChain...))
	businessHandlers = append(businessHandlers, struct {
		path    string
		handler http.Handler
	}{drugTestPath, drugTestHandler})

	// Background Check Service
	bgCheckPath, bgCheckHandler := servicesv1connect.NewBackgroundCheckServiceHandler(servicesService, connect.WithInterceptors(interceptorChain...))
	businessHandlers = append(businessHandlers, struct {
		path    string
		handler http.Handler
	}{bgCheckPath, bgCheckHandler})

	// MVR Service
	mvrPath, mvrHandler := servicesv1connect.NewMVRServiceHandler(servicesService, connect.WithInterceptors(interceptorChain...))
	businessHandlers = append(businessHandlers, struct {
		path    string
		handler http.Handler
	}{mvrPath, mvrHandler})

	// DOT Physical Service
	dotPhysicalPath, dotPhysicalHandler := servicesv1connect.NewDOTPhysicalServiceHandler(servicesService, connect.WithInterceptors(interceptorChain...))
	businessHandlers = append(businessHandlers, struct {
		path    string
		handler http.Handler
	}{dotPhysicalPath, dotPhysicalHandler})

	// Document Service
	documentPath, documentHandler := servicesv1connect.NewDocumentServiceHandler(servicesService, connect.WithInterceptors(interceptorChain...))
	businessHandlers = append(businessHandlers, struct {
		path    string
		handler http.Handler
	}{documentPath, documentHandler})

	// Notification Service
	notificationPath, notificationHandler := servicesv1connect.NewNotificationServiceHandler(servicesService, connect.WithInterceptors(interceptorChain...))
	businessHandlers = append(businessHandlers, struct {
		path    string
		handler http.Handler
	}{notificationPath, notificationHandler})

	// Workflow Service
	workflowPath, workflowHandler := servicesv1connect.NewWorkflowServiceHandler(servicesService, connect.WithInterceptors(interceptorChain...))
	businessHandlers = append(businessHandlers, struct {
		path    string
		handler http.Handler
	}{workflowPath, workflowHandler})

	// Add business services to the list
	for _, svc := range businessHandlers {
		services = append(services, vanguard.NewService(svc.path, svc.handler))
	}

	// Create Vanguard transcoder
	transcoder, err := vanguard.NewTranscoder(services)
	if err != nil {
		return nil, fmt.Errorf("create transcoder: %w", err)
	}

	log.Printf("✅ Initialized %d Connect services", len(services))
	return transcoder, nil
}

// buildInterceptorChain creates the interceptor chain based on configuration
func buildInterceptorChain(cfg *config.Config) []connect.Interceptor {
	var interceptorChain []connect.Interceptor

	if cfg.Interceptors.RecoveryEnabled {
		interceptorChain = append(interceptorChain, interceptors.RecoveryInterceptor())
		log.Println("✅ Recovery interceptor enabled")
	}

	if cfg.Interceptors.LoggingEnabled {
		interceptorChain = append(interceptorChain, interceptors.LoggingInterceptor())
		log.Println("✅ Logging interceptor enabled")
	}

	if cfg.Interceptors.MetricsEnabled {
		interceptorChain = append(interceptorChain, interceptors.MetricsInterceptor())
		log.Println("✅ Metrics interceptor enabled")
	}

	if cfg.Interceptors.RateLimitEnabled && cfg.RateLimit.Enabled {
		interceptorChain = append(interceptorChain, interceptors.RateLimitInterceptor(interceptors.RateLimitConfig{
			RequestsPerMinute: cfg.RateLimit.RequestsPerMinute,
			BurstSize:         cfg.RateLimit.BurstSize,
		}))
		log.Printf("✅ Rate limiting enabled (%d req/min)", cfg.RateLimit.RequestsPerMinute)
	}

	if cfg.Interceptors.BizCtxEnabled {
		cookieConfig := interceptors.BizCtxCookieConfig{
			SessionCookieName:      cfg.BizCtx.SessionCookieName,
			OrganizationCookieName: cfg.BizCtx.OrganizationCookieName,
			CookieDomain:           cfg.BizCtx.CookieDomain,
			CookiePath:             cfg.BizCtx.CookiePath,
			CookieMaxAge:           cfg.BizCtx.GetCookieMaxAge(),
			CookieSecure:           cfg.BizCtx.CookieSecure && cfg.IsProduction(),
			CookieHttpOnly:         cfg.BizCtx.CookieHttpOnly,
			CookieSameSite:         cfg.BizCtx.GetSameSiteMode(),
		}
		interceptorChain = append(interceptorChain, interceptors.BizCtxCookieInterceptor(cookieConfig))
		log.Println("✅ BizCtx cookie interceptor enabled")
	}

	if cfg.Interceptors.ValidationEnabled {
		interceptorChain = append(interceptorChain, interceptors.ValidationInterceptor())
		log.Println("✅ Validation interceptor enabled")
	}

	if cfg.Interceptors.AuthEnabled {
		interceptorChain = append(interceptorChain, interceptors.AuthInterceptor(interceptors.JWTConfig{
			SecretKey:       cfg.GetJWTSecretBytes(),
			PublicEndpoints: cfg.JWT.PublicEndpoints,
			RequiredRole:    cfg.JWT.RequiredRole,
		}))
		log.Println("✅ JWT authentication interceptor enabled")
	} else {
		log.Println("⚠️  Authentication disabled (development mode)")
	}

	return interceptorChain
}
