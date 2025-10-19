package gateway

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

// OpenAPIEnhancer enhances OpenAPI specifications and provides Swagger UI
type OpenAPIEnhancer struct {
	swaggerPath string
	serverAddr  string
	version     string
}

// NewOpenAPIEnhancer creates a new OpenAPI enhancer
func NewOpenAPIEnhancer(swaggerPath, serverAddr, version string) *OpenAPIEnhancer {
	return &OpenAPIEnhancer{
		swaggerPath: swaggerPath,
		serverAddr:  serverAddr,
		version:     version,
	}
}

// ServeEnhancedSwaggerJSON returns enhanced Swagger JSON with server info
func (e *OpenAPIEnhancer) ServeEnhancedSwaggerJSON() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// Read original swagger file
		data, err := os.ReadFile(e.swaggerPath)
		if err != nil {
			log.Printf("Error reading swagger file: %v", err)
			http.Error(w, "Swagger file not found", http.StatusNotFound)
			return
		}

		// Parse JSON
		var swagger map[string]interface{}
		if err := json.Unmarshal(data, &swagger); err != nil {
			log.Printf("Error parsing swagger JSON: %v", err)
			http.Error(w, "Invalid swagger JSON", http.StatusInternalServerError)
			return
		}

		// Enhance the swagger spec
		e.enhanceSwaggerSpec(swagger, r)

		// Return enhanced JSON
		enhanced, err := json.MarshalIndent(swagger, "", "  ")
		if err != nil {
			log.Printf("Error marshaling enhanced swagger: %v", err)
			http.Error(w, "Error generating enhanced swagger", http.StatusInternalServerError)
			return
		}

		w.Write(enhanced)
	}
}

// ServeEnhancedSwaggerUI serves an enhanced Swagger UI
func (e *OpenAPIEnhancer) ServeEnhancedSwaggerUI() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		// Get base URL for the API
		scheme := "http"
		if r.TLS != nil {
			scheme = "https"
		}
		baseURL := fmt.Sprintf("%s://%s", scheme, r.Host)

		html := fmt.Sprintf(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>V1 Consortium API Documentation</title>
    <link rel="stylesheet" type="text/css" href="https://unpkg.com/swagger-ui-dist@5.17.14/swagger-ui.css" />
    <link rel="icon" type="image/png" href="https://unpkg.com/swagger-ui-dist@5.17.14/favicon-32x32.png" sizes="32x32" />
    <style>
        html {
            box-sizing: border-box;
            overflow: -moz-scrollbars-vertical;
            overflow-y: scroll;
        }
        *, *:before, *:after {
            box-sizing: inherit;
        }
        body {
            margin:0;
            background: #fafafa;
            font-family: "Helvetica Neue", Helvetica, Arial, sans-serif;
        }
        .swagger-ui .topbar {
            background-color: #1f2937;
        }
        .swagger-ui .topbar .download-url-wrapper {
            display: none;
        }
        .custom-header {
            background: linear-gradient(135deg, #667eea 0%%, #764ba2 100%%);
            color: white;
            padding: 20px;
            text-align: center;
        }
        .custom-header h1 {
            margin: 0 0 10px 0;
            font-size: 2em;
        }
        .custom-header p {
            margin: 0;
            opacity: 0.9;
        }
        .protocol-badges {
            margin: 15px 0;
        }
        .protocol-badge {
            display: inline-block;
            background: rgba(255, 255, 255, 0.2);
            padding: 4px 12px;
            margin: 2px;
            border-radius: 16px;
            font-size: 0.8em;
            font-weight: bold;
        }
    </style>
</head>
<body>
    <div class="custom-header">
        <h1>V1 Consortium API</h1>
        <p>Compliance & Screening Services Platform</p>
        <div class="protocol-badges">
            <span class="protocol-badge">Connect RPC</span>
            <span class="protocol-badge">gRPC</span>
            <span class="protocol-badge">gRPC-Web</span>
            <span class="protocol-badge">HTTP/JSON</span>
        </div>
        <p>Version %s | Environment: Development</p>
    </div>
    <div id="swagger-ui"></div>
    
    <script src="https://unpkg.com/swagger-ui-dist@5.17.14/swagger-ui-bundle.js"></script>
    <script src="https://unpkg.com/swagger-ui-dist@5.17.14/swagger-ui-standalone-preset.js"></script>
    <script>
        window.onload = function() {
            const ui = SwaggerUIBundle({
                url: '%s/enhanced-swagger.json',
                dom_id: '#swagger-ui',
                deepLinking: true,
                presets: [
                    SwaggerUIBundle.presets.apis,
                    SwaggerUIStandalonePreset
                ],
                plugins: [
                    SwaggerUIBundle.plugins.DownloadUrl
                ],
                layout: "StandaloneLayout",
                tryItOutEnabled: true,
                requestInterceptor: function(req) {
                    // Add protocol version headers for Connect RPC
                    if (!req.headers) req.headers = {};
                    req.headers['Connect-Protocol-Version'] = '1';
                    req.headers['Content-Type'] = 'application/json';
                    return req;
                },
                onComplete: function() {
                    console.log('V1 Consortium API Documentation loaded');
                },
                validatorUrl: null,
                docExpansion: 'list',
                apisSorter: 'alpha',
                operationsSorter: 'alpha'
            });
            
            // Custom styling for better UX
            setTimeout(function() {
                const style = document.createElement('style');
                style.innerHTML = '%%
                    .swagger-ui .scheme-container { display: none; }
                    .swagger-ui .info { margin: 20px 0; }
                    .swagger-ui .info .title { font-size: 2em; color: #1f2937; }
                %%';
                document.head.appendChild(style);
            }, 1000);
        };
    </script>
</body>
</html>`, e.version, baseURL)

		w.Write([]byte(html))
	}
}

// enhanceSwaggerSpec enhances the swagger specification with runtime information
func (e *OpenAPIEnhancer) enhanceSwaggerSpec(swagger map[string]interface{}, r *http.Request) {
	// Set dynamic server information
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}

	host := r.Host
	if host == "" {
		host = strings.TrimPrefix(e.serverAddr, ":")
		if host == "" {
			host = "localhost:8000"
		} else {
			host = "localhost" + host
		}
	}

	// Update servers
	servers := []map[string]interface{}{
		{
			"url":         fmt.Sprintf("%s://%s", scheme, host),
			"description": "Current server",
		},
		{
			"url":         "http://localhost:8000",
			"description": "Local development server",
		},
	}
	swagger["servers"] = servers

	// Enhance info section
	if info, ok := swagger["info"].(map[string]interface{}); ok {
		info["title"] = "V1 Consortium API"
		info["description"] = `
## Compliance & Screening Services Platform

This API provides comprehensive compliance and screening services for transportation companies.

### Supported Protocols

- **Connect RPC**: Modern, type-safe RPC with HTTP/2
- **gRPC**: High-performance RPC for service-to-service communication  
- **gRPC-Web**: Browser-compatible gRPC over HTTP/1.1 and HTTP/2
- **HTTP/JSON**: RESTful JSON API for web applications

### Authentication

Most endpoints require JWT authentication via the Authorization header:
` + "```" + `
Authorization: Bearer <your-jwt-token>
` + "```" + `

### Rate Limiting

API requests are rate-limited to prevent abuse. Default limits:
- 100 requests per minute per client
- Burst size of 10 requests

### Error Handling

The API uses standard HTTP status codes and Connect RPC error codes.
Error responses include detailed error messages and error codes.
`
		info["version"] = e.version
		info["contact"] = map[string]interface{}{
			"name":  "V1 Consortium API Support",
			"email": "api-support@v1consortium.com",
		}
		info["license"] = map[string]interface{}{
			"name": "Proprietary",
		}
	}

	// Add security definitions if not present
	if _, ok := swagger["components"]; !ok {
		swagger["components"] = make(map[string]interface{})
	}

	if components, ok := swagger["components"].(map[string]interface{}); ok {
		if _, ok := components["securitySchemes"]; !ok {
			components["securitySchemes"] = map[string]interface{}{
				"bearerAuth": map[string]interface{}{
					"type":         "http",
					"scheme":       "bearer",
					"bearerFormat": "JWT",
					"description":  "JWT token for authentication",
				},
			}
		}
	}

	// Add global security requirement
	swagger["security"] = []map[string]interface{}{
		{"bearerAuth": []string{}},
	}

	// Add tags for better organization
	swagger["tags"] = []map[string]interface{}{
		{
			"name":        "Authentication",
			"description": "User authentication and authorization",
		},
		{
			"name":        "Organizations",
			"description": "Organization management and settings",
		},
		{
			"name":        "Compliance",
			"description": "Compliance monitoring and reporting",
		},
		{
			"name":        "Drug Testing",
			"description": "Drug and alcohol testing management",
		},
		{
			"name":        "Background Checks",
			"description": "Background check services",
		},
		{
			"name":        "MVR Reports",
			"description": "Motor vehicle record reporting",
		},
		{
			"name":        "DOT Physicals",
			"description": "DOT physical examination management",
		},
		{
			"name":        "Documents",
			"description": "Document storage and retrieval",
		},
		{
			"name":        "Notifications",
			"description": "Notification and messaging services",
		},
		{
			"name":        "Gateway",
			"description": "API Gateway configuration and monitoring",
		},
	}
}
