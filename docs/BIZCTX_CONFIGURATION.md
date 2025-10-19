# V1 Consortium Configuration Example

This file shows examples of how to configure the BizCtx cookie interceptors.

## Environment Variables

You can configure the BizCtx cookie settings using environment variables:

```bash
# Basic cookie configuration
export BIZCTX_SESSIONCOOKIENAME="v1_session_id"
export BIZCTX_ORGANIZATIONCOOKIENAME="v1_org_id"
export BIZCTX_COOKIEDOMAIN=".yourdomain.com"
export BIZCTX_COOKIEPATH="/"
export BIZCTX_COOKIEMAXAGEHOURS="24"

# Security settings
export BIZCTX_COOKIESECURE="true"        # Set to false for development
export BIZCTX_COOKIEHTTPONLY="true"
export BIZCTX_COOKIESAMESITE="strict"    # Options: strict, lax, none

# Enable/disable the interceptor
export INTERCEPTORS_BIZCTXENABLED="true"
```

## Configuration File (config.yaml)

```yaml
# Server configuration
server:
  address: ":8000"
  readTimeout: "15s"
  writeTimeout: "15s"
  idleTimeout: "60s"

# Environment
environment: "development"  # or "production"

# Interceptors configuration
interceptors:
  recoveryEnabled: true
  loggingEnabled: true
  metricsEnabled: true
  rateLimitEnabled: true
  validationEnabled: true
  authEnabled: false        # Disabled for development
  corsEnabled: true
  bizCtxEnabled: true       # Enable BizCtx cookie interceptor

# BizCtx cookie configuration
bizCtx:
  sessionCookieName: "v1_session_id"
  organizationCookieName: "v1_org_id"
  cookieDomain: ""          # Empty for development, set to ".yourdomain.com" for production
  cookiePath: "/"
  cookieMaxAgeHours: 24     # 24 hours
  cookieSecure: false       # Set to true for production (HTTPS)
  cookieHttpOnly: true
  cookieSameSite: "lax"     # Options: strict, lax, none

# JWT configuration
jwt:
  secretKey: "your-super-secret-key-change-in-production"
  publicEndpoints:
    - "/v1consortium.auth.AuthService/Login"
    - "/v1consortium.auth.AuthService/RegisterUser"
    - "/health"
    - "/swagger.json"
    - "/enhanced-swagger.json"
    - "/docs"
    - "/docs/"
  requiredRole: ""

# CORS configuration
cors:
  allowedOrigins: ["*"]     # Restrict in production
  allowedMethods: ["GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD", "PATCH"]
  allowedHeaders:
    - "Accept"
    - "Accept-Encoding"
    - "Accept-Post"
    - "Content-Encoding"
    - "Content-Length"
    - "Content-Type"
    - "Connect-Protocol-Version"
    - "Connect-Timeout-Ms"
    - "Grpc-Timeout"
    - "Grpc-Encoding"
    - "Grpc-Accept-Encoding"
    - "User-Agent"
    - "X-User-Agent"
    - "X-Grpc-Web"
    - "Authorization"
    - "Cookie"              # Important for cookie support
  allowCredentials: true    # Important for cookie support

# Rate limiting
rateLimit:
  enabled: true
  requestsPerMinute: 100
  burstSize: 10
```

## Development vs Production Settings

### Development Configuration

```yaml
environment: "development"

bizCtx:
  cookieSecure: false       # Allow HTTP
  cookieSameSite: "lax"     # More permissive

cors:
  allowedOrigins: ["http://localhost:3000", "http://localhost:5173"]
  allowCredentials: true

interceptors:
  authEnabled: false        # Disable for easier development
```

### Production Configuration

```yaml
environment: "production"

bizCtx:
  cookieDomain: ".yourdomain.com"
  cookieSecure: true        # HTTPS only
  cookieSameSite: "strict"  # Strict security

cors:
  allowedOrigins: ["https://yourdomain.com", "https://app.yourdomain.com"]
  allowCredentials: true

interceptors:
  authEnabled: true         # Enable for production
```

## Usage in Your Application

Once configured, the interceptors will automatically:

1. **Read cookies** from incoming requests and set them in the business context
2. **Write values** from the business context back to cookies in responses

### In Your Connect Services

```go
func (s *YourService) SomeMethod(ctx context.Context, req *connect.Request[SomeRequest]) (*connect.Response[SomeResponse], error) {
    // Session ID and Organization ID are automatically available in context
    sessionID, err := service.BizCtx().GetCurrentSessionID(ctx)
    if err != nil {
        return nil, connect.NewError(connect.CodeUnauthenticated, err)
    }
    
    orgID, err := service.BizCtx().GetCurrentOrganizationID(ctx)
    if err != nil {
        log.Printf("No organization ID: %v", err)
    }
    
    // Your business logic here
    
    // If you update the context values, they'll be written back to cookies
    ctx = service.BizCtx().SetCurrentSessionID(ctx, newSessionID)
    ctx = service.BizCtx().SetCurrentOrganizationID(ctx, newOrgID)
    
    return connect.NewResponse(&SomeResponse{}), nil
}
```

### Client-Side Cookie Handling

Ensure your client sends cookies with requests:

```javascript
// Fetch API
fetch('/v1consortium.auth.AuthService/SomeMethod', {
  method: 'POST',
  credentials: 'include',  // Important: include cookies
  headers: {
    'Content-Type': 'application/json',
  },
  body: JSON.stringify(requestData)
});

// Axios
axios.defaults.withCredentials = true;
```

## Security Considerations

1. **Always use HTTPS in production** (`cookieSecure: true`)
2. **Set appropriate domain** for cookie scope
3. **Use HttpOnly cookies** to prevent XSS attacks
4. **Choose appropriate SameSite policy** based on your needs
5. **Enable CORS credentials** if using cookies across origins
6. **Validate and sanitize** cookie values in your business logic

## Troubleshooting

### Cookies Not Being Set

1. Check that `allowCredentials: true` in CORS config
2. Ensure client sends `credentials: 'include'`
3. Verify cookie domain matches request domain
4. Check browser developer tools for cookie errors

### Cookies Not Being Read

1. Verify `Cookie` header is in CORS `allowedHeaders`
2. Check that cookies aren't expired
3. Ensure cookie path matches request path
4. Verify SameSite policy allows the request

### Development Issues

1. Set `cookieSecure: false` for HTTP development
2. Use `cookieSameSite: "lax"` for cross-origin development
3. Check browser security restrictions for localhost