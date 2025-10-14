package gatewayconnect

import (
	"context"
	v1 "v1consortium/api/gateway/v1"
	"v1consortium/internal/controller/gateway"

	"connectrpc.com/connect"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type GatewayConnectService struct {
	gatewayController *gateway.Controller
}

func NewGatewayConnectService(ctx context.Context) *GatewayConnectService {
	return &GatewayConnectService{
		gatewayController: &gateway.Controller{},
	}
}

func (s *GatewayConnectService) GetGatewayConfig(ctx context.Context, req *connect.Request[v1.GetGatewayConfigRequest]) (res *connect.Response[v1.GetGatewayConfigResponse], err error) {
	resp, err := s.gatewayController.GetGatewayConfig(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (*GatewayConnectService) UpdateGatewayConfig(ctx context.Context, req *connect.Request[v1.UpdateGatewayConfigRequest]) (res *connect.Response[v1.UpdateGatewayConfigResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*GatewayConnectService) AddRoute(ctx context.Context, req *connect.Request[v1.AddRouteRequest]) (res *connect.Response[v1.AddRouteResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*GatewayConnectService) RemoveRoute(ctx context.Context, req *connect.Request[v1.RemoveRouteRequest]) (res *connect.Response[v1.RemoveRouteResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*GatewayConnectService) UpdateRoute(ctx context.Context, req *connect.Request[v1.UpdateRouteRequest]) (res *connect.Response[v1.UpdateRouteResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*GatewayConnectService) ListRoutes(ctx context.Context, req *connect.Request[v1.ListRoutesRequest]) (res *connect.Response[v1.ListRoutesResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*GatewayConnectService) Authenticate(ctx context.Context, req *connect.Request[v1.AuthenticateRequest]) (res *connect.Response[v1.AuthenticateResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*GatewayConnectService) CheckRateLimit(ctx context.Context, req *connect.Request[v1.RateLimitRequest]) (res *connect.Response[v1.RateLimitResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*GatewayConnectService) ValidateRequest(ctx context.Context, req *connect.Request[v1.ValidateRequestRequest]) (res *connect.Response[v1.ValidateRequestResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*GatewayConnectService) GetServiceHealth(ctx context.Context, req *connect.Request[v1.ServiceHealthRequest]) (res *connect.Response[v1.ServiceHealthResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*GatewayConnectService) ListServices(ctx context.Context, req *connect.Request[v1.ListServicesRequest]) (res *connect.Response[v1.ListServicesResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*GatewayConnectService) LogApiRequest(ctx context.Context, req *connect.Request[v1.LogApiRequestRequest]) (res *connect.Response[v1.LogApiRequestResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*GatewayConnectService) GetCircuitBreakerStatus(ctx context.Context, req *connect.Request[v1.CircuitBreakerStatusRequest]) (res *connect.Response[v1.CircuitBreakerStatusResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*GatewayConnectService) RecordMetric(ctx context.Context, req *connect.Request[v1.RecordMetricRequest]) (res *connect.Response[v1.RecordMetricResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*GatewayConnectService) GetMetrics(ctx context.Context, req *connect.Request[v1.GetMetricsRequest]) (res *connect.Response[v1.GetMetricsResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*GatewayConnectService) GetGatewayStatus(ctx context.Context, req *connect.Request[v1.GetGatewayStatusRequest]) (res *connect.Response[v1.GetGatewayStatusResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*GatewayConnectService) RecordError(ctx context.Context, req *connect.Request[v1.RecordErrorRequest]) (res *connect.Response[v1.RecordErrorResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*GatewayConnectService) GetErrorStats(ctx context.Context, req *connect.Request[v1.GetErrorStatsRequest]) (res *connect.Response[v1.GetErrorStatsResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*GatewayConnectService) GetPerformanceStats(ctx context.Context, req *connect.Request[v1.GetPerformanceStatsRequest]) (res *connect.Response[v1.GetPerformanceStatsResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*GatewayConnectService) GetRateLimitStats(ctx context.Context, req *connect.Request[v1.GetRateLimitStatsRequest]) (res *connect.Response[v1.GetRateLimitStatsResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
