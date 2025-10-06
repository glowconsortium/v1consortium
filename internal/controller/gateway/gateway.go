package gateway

import (
	"context"
	v1 "v1consortium/api/gateway/v1"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type Controller struct {
	v1.UnimplementedConfigurationServiceServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterConfigurationServiceServer(s.Server, &Controller{})
}

func (*Controller) GetGatewayConfig(ctx context.Context, req *v1.GetGatewayConfigRequest) (res *v1.GetGatewayConfigResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) UpdateGatewayConfig(ctx context.Context, req *v1.UpdateGatewayConfigRequest) (res *v1.UpdateGatewayConfigResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) AddRoute(ctx context.Context, req *v1.AddRouteRequest) (res *v1.AddRouteResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) RemoveRoute(ctx context.Context, req *v1.RemoveRouteRequest) (res *v1.RemoveRouteResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) UpdateRoute(ctx context.Context, req *v1.UpdateRouteRequest) (res *v1.UpdateRouteResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) ListRoutes(ctx context.Context, req *v1.ListRoutesRequest) (res *v1.ListRoutesResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) Authenticate(ctx context.Context, req *v1.AuthenticateRequest) (res *v1.AuthenticateResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) CheckRateLimit(ctx context.Context, req *v1.RateLimitRequest) (res *v1.RateLimitResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) ValidateRequest(ctx context.Context, req *v1.ValidateRequestRequest) (res *v1.ValidateRequestResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GetServiceHealth(ctx context.Context, req *v1.ServiceHealthRequest) (res *v1.ServiceHealthResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) ListServices(ctx context.Context, req *v1.ListServicesRequest) (res *v1.ListServicesResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) LogApiRequest(ctx context.Context, req *v1.LogApiRequestRequest) (res *v1.LogApiRequestResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GetCircuitBreakerStatus(ctx context.Context, req *v1.CircuitBreakerStatusRequest) (res *v1.CircuitBreakerStatusResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) RecordMetric(ctx context.Context, req *v1.RecordMetricRequest) (res *v1.RecordMetricResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GetMetrics(ctx context.Context, req *v1.GetMetricsRequest) (res *v1.GetMetricsResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GetGatewayStatus(ctx context.Context, req *v1.GetGatewayStatusRequest) (res *v1.GetGatewayStatusResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) RecordError(ctx context.Context, req *v1.RecordErrorRequest) (res *v1.RecordErrorResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GetErrorStats(ctx context.Context, req *v1.GetErrorStatsRequest) (res *v1.GetErrorStatsResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GetPerformanceStats(ctx context.Context, req *v1.GetPerformanceStatsRequest) (res *v1.GetPerformanceStatsResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GetRateLimitStats(ctx context.Context, req *v1.GetRateLimitStatsRequest) (res *v1.GetRateLimitStatsResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
