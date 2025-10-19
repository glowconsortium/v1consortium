# BizCtx Cookie Interceptors

This package provides Connect RPC interceptors that can read values from HTTP cookies into the business context (`bizctx`) and write values from `bizctx` back to cookies.

## Overview

The interceptors bridge the gap between HTTP cookies and your application's business context, allowing you to:

1. **Read from cookies**: Extract session IDs and organization IDs from incoming HTTP cookies and set them in the business context
2. **Write to cookies**: Take values from the business context and set them as HTTP cookies in the response

## Interceptors

### 1. BizCtxCookieReaderInterceptor

Reads cookies from incoming requests and sets values in the business context.

```go
config := BizCtxCookieConfig{
    SessionCookieName:      "session_id",
    OrganizationCookieName: "org_id",
    // ... other config
}

reader := BizCtxCookieReaderInterceptor(config)
```

### 2. BizCtxCookieWriterInterceptor

Reads values from the business context and sets them as cookies in the response.

```go
writer := BizCtxCookieWriterInterceptor(config)
```

### 3. BizCtxCookieInterceptor (Combined)

Combines both reading and writing functionality in a single interceptor.

```go
combined := BizCtxCookieInterceptor(config)
```

## Configuration

Use `BizCtxCookieConfig` to configure cookie behavior:

```go
config := BizCtxCookieConfig{
    SessionCookieName:      "v1_session_id",     // Cookie name for session ID
    OrganizationCookieName: "v1_org_id",        // Cookie name for organization ID
    CookieDomain:           ".yourdomain.com",   // Cookie domain
    CookiePath:             "/",                 // Cookie path
    CookieMaxAge:           24 * time.Hour,      // Cookie lifetime
    CookieSecure:           true,                // Secure flag (HTTPS only)
    CookieHttpOnly:         true,                // HttpOnly flag (no JS access)
    CookieSameSite:         http.SameSiteStrictMode, // SameSite policy
}
```

Or use defaults:

```go
config := DefaultBizCtxCookieConfig()
```

## Usage Examples

### Basic Setup

```go
import (
    "v1consortium/internal/pkg/interceptors"
    "connectrpc.com/connect"
)

func setupInterceptors() []connect.Interceptor {
    cookieConfig := interceptors.DefaultBizCtxCookieConfig()
    
    // Customize for your environment
    cookieConfig.SessionCookieName = "my_session"
    cookieConfig.OrganizationCookieName = "my_org"
    cookieConfig.CookieSecure = false // For development
    
    return []connect.Interceptor{
        interceptors.RecoveryInterceptor(),
        interceptors.LoggingInterceptor(),
        interceptors.BizCtxCookieInterceptor(cookieConfig),
        interceptors.MetricsInterceptor(),
    }
}
```

### Separate Reader and Writer

```go
func setupSeparateInterceptors() []connect.Interceptor {
    cookieConfig := interceptors.DefaultBizCtxCookieConfig()
    
    return []connect.Interceptor{
        interceptors.RecoveryInterceptor(),
        
        // Read cookies early in the chain
        interceptors.BizCtxCookieReaderInterceptor(cookieConfig),
        
        // Your business logic interceptors here
        interceptors.AuthInterceptor(authConfig),
        
        // Write cookies late in the chain
        interceptors.BizCtxCookieWriterInterceptor(cookieConfig),
        
        interceptors.ValidationInterceptor(),
    }
}
```

### In Your Connect Service

```go
func (s *MyService) SomeMethod(ctx context.Context, req *connect.Request[SomeRequest]) (*connect.Response[SomeResponse], error) {
    // The session ID and organization ID are now available in the context
    sessionID, err := service.BizCtx().GetCurrentSessionID(ctx)
    if err != nil {
        // Handle case where session ID is not present
        return nil, connect.NewError(connect.CodeUnauthenticated, err)
    }
    
    orgID, err := service.BizCtx().GetCurrentOrganizationID(ctx)
    if err != nil {
        // Handle case where organization ID is not present
        log.Printf("No organization ID in context: %v", err)
    }
    
    // Your business logic here
    
    // If you modify the session or organization in your business logic,
    // the writer interceptor will automatically set the updated values as cookies
    ctx = service.BizCtx().SetCurrentSessionID(ctx, newSessionID)
    ctx = service.BizCtx().SetCurrentOrganizationID(ctx, newOrgID)
    
    return connect.NewResponse(&SomeResponse{}), nil
}
```

## Important Notes

### Connect RPC and Cookies

Connect RPC is primarily designed for API communication, and cookies are more commonly used in browser-based applications. However, these interceptors can be useful when:

1. You have a web application that needs to maintain session state across Connect RPC calls
2. You're building a hybrid application that uses both traditional web forms and Connect RPC
3. You need to bridge between cookie-based authentication and your Connect RPC services

### HTTP Middleware Alternative

For more robust cookie handling, consider implementing cookie management at the HTTP middleware level rather than in Connect interceptors. The interceptors provided here are useful for simple cases but may not handle all cookie edge cases.

### Security Considerations

- Always use `CookieSecure: true` in production (HTTPS)
- Set `CookieHttpOnly: true` to prevent XSS attacks
- Choose appropriate `SameSite` policies based on your application's needs
- Consider cookie encryption for sensitive data
- Validate and sanitize cookie values

## Development vs Production

```go
// Development configuration
devConfig := DefaultBizCtxCookieConfig()
devConfig.CookieSecure = false  // Allow HTTP
devConfig.CookieDomain = ""     // Don't restrict domain

// Production configuration  
prodConfig := DefaultBizCtxCookieConfig()
prodConfig.CookieSecure = true
prodConfig.CookieDomain = ".yourdomain.com"
prodConfig.CookieSameSite = http.SameSiteStrictMode
```