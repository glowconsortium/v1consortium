package bizctx

import (
	"context"
	"fmt"
	"v1consortium/internal/service"
)

const (
	SessionIDKey      bizctxKey = "session_id"
	OrganizationIDKey bizctxKey = "organization_id"
)

type bizctxKey string

type sBizCtx struct {
	sessionID      string
	OrganizationID string
}

func new() service.IBizCtx {
	return &sBizCtx{}
}

func init() {
	service.RegisterBizCtx(new())
}

func (s *sBizCtx) Init(ctx context.Context) error {
	return nil
}

func (s *sBizCtx) SetCurrentSessionID(ctx context.Context, sessionID string) context.Context {
	return context.WithValue(ctx, SessionIDKey, sessionID)
}

func (s *sBizCtx) SetCurrentOrganizationID(ctx context.Context, organizationID string) context.Context {
	return context.WithValue(ctx, OrganizationIDKey, organizationID)
}

func (s *sBizCtx) GetCurrentOrganizationID(ctx context.Context) (string, error) {
	if value := ctx.Value(OrganizationIDKey); value != nil {
		if orgID, ok := value.(string); ok {
			return orgID, nil
		}
	}
	return "", fmt.Errorf("organization ID not found in context")
}

func (s *sBizCtx) GetCurrentSessionID(ctx context.Context) (string, error) {
	if value := ctx.Value(SessionIDKey); value != nil {
		if sessionID, ok := value.(string); ok {
			return sessionID, nil
		}
	}
	return "", fmt.Errorf("session ID not found in context")
}
