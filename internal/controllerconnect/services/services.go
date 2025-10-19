package servicesconnect

import (
	"context"
	v1 "v1consortium/api/services/v1"
	"v1consortium/internal/controller/services"

	"connectrpc.com/connect"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type ServicesConnectService struct {
	servicesController *services.Controller
}

func NewServicesConnectService(ctx context.Context) *ServicesConnectService {
	return &ServicesConnectService{
		servicesController: &services.Controller{},
	}
}

func (s *ServicesConnectService) OrderBackgroundCheck(ctx context.Context, req *connect.Request[v1.OrderBackgroundCheckRequest]) (res *connect.Response[v1.OrderBackgroundCheckResponse], err error) {
	resp, err := s.servicesController.OrderBackgroundCheck(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (s *ServicesConnectService) GetBackgroundCheck(ctx context.Context, req *connect.Request[v1.GetBackgroundCheckRequest]) (res *connect.Response[v1.GetBackgroundCheckResponse], err error) {
	resp, err := s.servicesController.GetBackgroundCheck(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (s *ServicesConnectService) UpdateBackgroundCheck(ctx context.Context, req *connect.Request[v1.UpdateBackgroundCheckRequest]) (res *connect.Response[v1.UpdateBackgroundCheckResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) ListBackgroundChecks(ctx context.Context, req *connect.Request[v1.ListBackgroundChecksRequest]) (res *connect.Response[v1.ListBackgroundChecksResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) AddFinding(ctx context.Context, req *connect.Request[v1.AddFindingRequest]) (res *connect.Response[v1.AddFindingResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) UpdateFinding(ctx context.Context, req *connect.Request[v1.UpdateFindingRequest]) (res *connect.Response[v1.UpdateFindingResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) ListFindings(ctx context.Context, req *connect.Request[v1.ListFindingsRequest]) (res *connect.Response[v1.ListFindingsResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) InitiateAdverseAction(ctx context.Context, req *connect.Request[v1.InitiateAdverseActionRequest]) (res *connect.Response[v1.InitiateAdverseActionResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) HandleDispute(ctx context.Context, req *connect.Request[v1.HandleDisputeRequest]) (res *connect.Response[v1.HandleDisputeResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) GetAdverseActionStatus(ctx context.Context, req *connect.Request[v1.GetAdverseActionStatusRequest]) (res *connect.Response[v1.GetAdverseActionStatusResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) GetAvailablePackages(ctx context.Context, req *connect.Request[v1.GetAvailablePackagesRequest]) (res *connect.Response[v1.GetAvailablePackagesResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) GetProviderStatus(ctx context.Context, req *connect.Request[v1.GetProviderStatusRequest]) (res *connect.Response[v1.GetProviderStatusResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) GetBackgroundCheckAnalytics(ctx context.Context, req *connect.Request[v1.GetBackgroundCheckAnalyticsRequest]) (res *connect.Response[v1.GetBackgroundCheckAnalyticsResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) GetComplianceStatus(ctx context.Context, req *connect.Request[v1.GetComplianceStatusRequest]) (res *connect.Response[v1.GetComplianceStatusResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) UpdateComplianceStatus(ctx context.Context, req *connect.Request[v1.UpdateComplianceStatusRequest]) (res *connect.Response[v1.UpdateComplianceStatusResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) ListComplianceStatus(ctx context.Context, req *connect.Request[v1.ListComplianceStatusRequest]) (res *connect.Response[v1.ListComplianceStatusResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) GenerateComplianceCertificate(ctx context.Context, req *connect.Request[v1.GenerateComplianceCertificateRequest]) (res *connect.Response[v1.GenerateComplianceCertificateResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) GetCertificate(ctx context.Context, req *connect.Request[v1.GetCertificateRequest]) (res *connect.Response[v1.GetCertificateResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) ListCertificates(ctx context.Context, req *connect.Request[v1.ListCertificatesRequest]) (res *connect.Response[v1.ListCertificatesResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) RevokeCertificate(ctx context.Context, req *connect.Request[v1.RevokeCertificateRequest]) (res *connect.Response[v1.RevokeCertificateResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) GenerateComplianceReport(ctx context.Context, req *connect.Request[v1.GenerateComplianceReportRequest]) (res *connect.Response[v1.GenerateComplianceReportResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) GetSavedReport(ctx context.Context, req *connect.Request[v1.GetSavedReportRequest]) (res *connect.Response[v1.GetSavedReportResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) ListSavedReports(ctx context.Context, req *connect.Request[v1.ListSavedReportsRequest]) (res *connect.Response[v1.ListSavedReportsResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) DeleteSavedReport(ctx context.Context, req *connect.Request[v1.DeleteSavedReportRequest]) (res *connect.Response[v1.DeleteSavedReportResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) GetComplianceAlerts(ctx context.Context, req *connect.Request[v1.GetComplianceAlertsRequest]) (res *connect.Response[v1.GetComplianceAlertsResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) GetComplianceAnalytics(ctx context.Context, req *connect.Request[v1.GetComplianceAnalyticsRequest]) (res *connect.Response[v1.GetComplianceAnalyticsResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) GetComplianceAuditTrail(ctx context.Context, req *connect.Request[v1.GetComplianceAuditTrailRequest]) (res *connect.Response[v1.GetComplianceAuditTrailResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) UploadDocument(ctx context.Context, req *connect.Request[v1.UploadDocumentRequest]) (res *connect.Response[v1.UploadDocumentResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) GetDocument(ctx context.Context, req *connect.Request[v1.GetDocumentRequest]) (res *connect.Response[v1.GetDocumentResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) ListDocuments(ctx context.Context, req *connect.Request[v1.ListDocumentsRequest]) (res *connect.Response[v1.ListDocumentsResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) UpdateDocument(ctx context.Context, req *connect.Request[v1.UpdateDocumentRequest]) (res *connect.Response[v1.UpdateDocumentResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) DeleteDocument(ctx context.Context, req *connect.Request[v1.DeleteDocumentRequest]) (res *connect.Response[v1.DeleteDocumentResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) GetDocumentVersions(ctx context.Context, req *connect.Request[v1.GetDocumentVersionsRequest]) (res *connect.Response[v1.GetDocumentVersionsResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) ShareDocument(ctx context.Context, req *connect.Request[v1.ShareDocumentRequest]) (res *connect.Response[v1.ShareDocumentResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) GetSharedDocuments(ctx context.Context, req *connect.Request[v1.GetSharedDocumentsRequest]) (res *connect.Response[v1.GetSharedDocumentsResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) RevokeDocumentShare(ctx context.Context, req *connect.Request[v1.RevokeDocumentShareRequest]) (res *connect.Response[v1.RevokeDocumentShareResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) SearchDocuments(ctx context.Context, req *connect.Request[v1.SearchDocumentsRequest]) (res *connect.Response[v1.SearchDocumentsResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) GetDocumentAnalytics(ctx context.Context, req *connect.Request[v1.GetDocumentAnalyticsRequest]) (res *connect.Response[v1.GetDocumentAnalyticsResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) GetDocumentRetentionStatus(ctx context.Context, req *connect.Request[v1.GetDocumentRetentionStatusRequest]) (res *connect.Response[v1.GetDocumentRetentionStatusResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) ApplyRetentionPolicy(ctx context.Context, req *connect.Request[v1.ApplyRetentionPolicyRequest]) (res *connect.Response[v1.ApplyRetentionPolicyResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) ScheduleDOTPhysical(ctx context.Context, req *connect.Request[v1.ScheduleDOTPhysicalRequest]) (res *connect.Response[v1.ScheduleDOTPhysicalResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) GetDOTPhysical(ctx context.Context, req *connect.Request[v1.GetDOTPhysicalRequest]) (res *connect.Response[v1.GetDOTPhysicalResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) UpdateDOTPhysical(ctx context.Context, req *connect.Request[v1.UpdateDOTPhysicalRequest]) (res *connect.Response[v1.UpdateDOTPhysicalResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) ListDOTPhysicals(ctx context.Context, req *connect.Request[v1.ListDOTPhysicalsRequest]) (res *connect.Response[v1.ListDOTPhysicalsResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) RegisterMedicalExaminer(ctx context.Context, req *connect.Request[v1.RegisterMedicalExaminerRequest]) (res *connect.Response[v1.RegisterMedicalExaminerResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) GetMedicalExaminer(ctx context.Context, req *connect.Request[v1.GetMedicalExaminerRequest]) (res *connect.Response[v1.GetMedicalExaminerResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) ListMedicalExaminers(ctx context.Context, req *connect.Request[v1.ListMedicalExaminersRequest]) (res *connect.Response[v1.ListMedicalExaminersResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) GenerateCertificate(ctx context.Context, req *connect.Request[v1.GenerateCertificateRequest]) (res *connect.Response[v1.GenerateCertificateResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) ValidateCertificate(ctx context.Context, req *connect.Request[v1.ValidateCertificateRequest]) (res *connect.Response[v1.ValidateCertificateResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) GetExpiringCertificates(ctx context.Context, req *connect.Request[v1.GetExpiringCertificatesRequest]) (res *connect.Response[v1.GetExpiringCertificatesResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) SetExpirationReminder(ctx context.Context, req *connect.Request[v1.SetExpirationReminderRequest]) (res *connect.Response[v1.SetExpirationReminderResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) CreateTestingProgram(ctx context.Context, req *connect.Request[v1.CreateTestingProgramRequest]) (res *connect.Response[v1.CreateTestingProgramResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) GetTestingProgram(ctx context.Context, req *connect.Request[v1.GetTestingProgramRequest]) (res *connect.Response[v1.GetTestingProgramResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) ListTestingPrograms(ctx context.Context, req *connect.Request[v1.ListTestingProgramsRequest]) (res *connect.Response[v1.ListTestingProgramsResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) OrderDrugTest(ctx context.Context, req *connect.Request[v1.OrderDrugTestRequest]) (res *connect.Response[v1.OrderDrugTestResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) GetDrugTest(ctx context.Context, req *connect.Request[v1.GetDrugTestRequest]) (res *connect.Response[v1.GetDrugTestResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) UpdateDrugTest(ctx context.Context, req *connect.Request[v1.UpdateDrugTestRequest]) (res *connect.Response[v1.UpdateDrugTestResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) ListDrugTests(ctx context.Context, req *connect.Request[v1.ListDrugTestsRequest]) (res *connect.Response[v1.ListDrugTestsResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) CreateRandomPool(ctx context.Context, req *connect.Request[v1.CreateRandomPoolRequest]) (res *connect.Response[v1.CreateRandomPoolResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) AddUsersToPool(ctx context.Context, req *connect.Request[v1.AddUsersToPoolRequest]) (res *connect.Response[v1.AddUsersToPoolResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) RemoveUsersFromPool(ctx context.Context, req *connect.Request[v1.RemoveUsersFromPoolRequest]) (res *connect.Response[v1.RemoveUsersFromPoolResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) GetRandomPool(ctx context.Context, req *connect.Request[v1.GetRandomPoolRequest]) (res *connect.Response[v1.GetRandomPoolResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) ListRandomPools(ctx context.Context, req *connect.Request[v1.ListRandomPoolsRequest]) (res *connect.Response[v1.ListRandomPoolsResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) ConductRandomSelection(ctx context.Context, req *connect.Request[v1.ConductRandomSelectionRequest]) (res *connect.Response[v1.ConductRandomSelectionResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) GetRandomSelection(ctx context.Context, req *connect.Request[v1.GetRandomSelectionRequest]) (res *connect.Response[v1.GetRandomSelectionResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) ListRandomSelections(ctx context.Context, req *connect.Request[v1.ListRandomSelectionsRequest]) (res *connect.Response[v1.ListRandomSelectionsResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) ValidateRandomSelection(ctx context.Context, req *connect.Request[v1.ValidateRandomSelectionRequest]) (res *connect.Response[v1.ValidateRandomSelectionResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) OrderMVR(ctx context.Context, req *connect.Request[v1.OrderMVRRequest]) (res *connect.Response[v1.OrderMVRResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) GetMVRReport(ctx context.Context, req *connect.Request[v1.GetMVRReportRequest]) (res *connect.Response[v1.GetMVRReportResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) UpdateMVRReport(ctx context.Context, req *connect.Request[v1.UpdateMVRReportRequest]) (res *connect.Response[v1.UpdateMVRReportResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) ListMVRReports(ctx context.Context, req *connect.Request[v1.ListMVRReportsRequest]) (res *connect.Response[v1.ListMVRReportsResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) AddMVRViolation(ctx context.Context, req *connect.Request[v1.AddMVRViolationRequest]) (res *connect.Response[v1.AddMVRViolationResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) UpdateMVRViolation(ctx context.Context, req *connect.Request[v1.UpdateMVRViolationRequest]) (res *connect.Response[v1.UpdateMVRViolationResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) ListMVRViolations(ctx context.Context, req *connect.Request[v1.ListMVRViolationsRequest]) (res *connect.Response[v1.ListMVRViolationsResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) EnableContinuousMonitoring(ctx context.Context, req *connect.Request[v1.EnableContinuousMonitoringRequest]) (res *connect.Response[v1.EnableContinuousMonitoringResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) GetMonitoringStatus(ctx context.Context, req *connect.Request[v1.GetMonitoringStatusRequest]) (res *connect.Response[v1.GetMonitoringStatusResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) GetMVRAnalytics(ctx context.Context, req *connect.Request[v1.GetMVRAnalyticsRequest]) (res *connect.Response[v1.GetMVRAnalyticsResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) SyncProviderData(ctx context.Context, req *connect.Request[v1.SyncProviderDataRequest]) (res *connect.Response[v1.SyncProviderDataResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) SendNotification(ctx context.Context, req *connect.Request[v1.SendNotificationRequest]) (res *connect.Response[v1.SendNotificationResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) GetNotification(ctx context.Context, req *connect.Request[v1.GetNotificationRequest]) (res *connect.Response[v1.GetNotificationResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) ListNotifications(ctx context.Context, req *connect.Request[v1.ListNotificationsRequest]) (res *connect.Response[v1.ListNotificationsResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) MarkNotificationRead(ctx context.Context, req *connect.Request[v1.MarkNotificationReadRequest]) (res *connect.Response[v1.MarkNotificationReadResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) MarkAllNotificationsRead(ctx context.Context, req *connect.Request[v1.MarkAllNotificationsReadRequest]) (res *connect.Response[v1.MarkAllNotificationsReadResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) CreateNotificationTemplate(ctx context.Context, req *connect.Request[v1.CreateNotificationTemplateRequest]) (res *connect.Response[v1.CreateNotificationTemplateResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) GetNotificationTemplate(ctx context.Context, req *connect.Request[v1.GetNotificationTemplateRequest]) (res *connect.Response[v1.GetNotificationTemplateResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) ListNotificationTemplates(ctx context.Context, req *connect.Request[v1.ListNotificationTemplatesRequest]) (res *connect.Response[v1.ListNotificationTemplatesResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) UpdateNotificationTemplate(ctx context.Context, req *connect.Request[v1.UpdateNotificationTemplateRequest]) (res *connect.Response[v1.UpdateNotificationTemplateResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) GetNotificationPreferences(ctx context.Context, req *connect.Request[v1.GetNotificationPreferencesRequest]) (res *connect.Response[v1.GetNotificationPreferencesResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) UpdateNotificationPreferences(ctx context.Context, req *connect.Request[v1.UpdateNotificationPreferencesRequest]) (res *connect.Response[v1.UpdateNotificationPreferencesResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) SendBulkNotification(ctx context.Context, req *connect.Request[v1.SendBulkNotificationRequest]) (res *connect.Response[v1.SendBulkNotificationResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) ScheduleNotification(ctx context.Context, req *connect.Request[v1.ScheduleNotificationRequest]) (res *connect.Response[v1.ScheduleNotificationResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) ListScheduledNotifications(ctx context.Context, req *connect.Request[v1.ListScheduledNotificationsRequest]) (res *connect.Response[v1.ListScheduledNotificationsResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) CancelScheduledNotification(ctx context.Context, req *connect.Request[v1.CancelScheduledNotificationRequest]) (res *connect.Response[v1.CancelScheduledNotificationResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) GetNotificationAnalytics(ctx context.Context, req *connect.Request[v1.GetNotificationAnalyticsRequest]) (res *connect.Response[v1.GetNotificationAnalyticsResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) CreateOrganization(ctx context.Context, req *connect.Request[v1.CreateOrganizationRequest]) (res *connect.Response[v1.CreateOrganizationResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) GetOrganization(ctx context.Context, req *connect.Request[v1.GetOrganizationRequest]) (res *connect.Response[v1.GetOrganizationResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) UpdateOrganization(ctx context.Context, req *connect.Request[v1.UpdateOrganizationRequest]) (res *connect.Response[v1.UpdateOrganizationResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) ListOrganizations(ctx context.Context, req *connect.Request[v1.ListOrganizationsRequest]) (res *connect.Response[v1.ListOrganizationsResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) DeactivateOrganization(ctx context.Context, req *connect.Request[v1.DeactivateOrganizationRequest]) (res *connect.Response[v1.DeactivateOrganizationResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) CreateUser(ctx context.Context, req *connect.Request[v1.CreateUserRequest]) (res *connect.Response[v1.CreateUserResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) GetUser(ctx context.Context, req *connect.Request[v1.GetUserRequest]) (res *connect.Response[v1.GetUserResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) UpdateUser(ctx context.Context, req *connect.Request[v1.UpdateUserRequest]) (res *connect.Response[v1.UpdateUserResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) ListUsers(ctx context.Context, req *connect.Request[v1.ListUsersRequest]) (res *connect.Response[v1.ListUsersResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) DeactivateUser(ctx context.Context, req *connect.Request[v1.DeactivateUserRequest]) (res *connect.Response[v1.DeactivateUserResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) CreateSubscription(ctx context.Context, req *connect.Request[v1.CreateSubscriptionRequest]) (res *connect.Response[v1.CreateSubscriptionResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) GetSubscription(ctx context.Context, req *connect.Request[v1.GetSubscriptionRequest]) (res *connect.Response[v1.GetSubscriptionResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) UpdateSubscription(ctx context.Context, req *connect.Request[v1.UpdateSubscriptionRequest]) (res *connect.Response[v1.UpdateSubscriptionResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) ListSubscriptionPlans(ctx context.Context, req *connect.Request[v1.ListSubscriptionPlansRequest]) (res *connect.Response[v1.ListSubscriptionPlansResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) StartWorkflow(ctx context.Context, req *connect.Request[v1.StartWorkflowRequest]) (res *connect.Response[v1.StartWorkflowResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) GetWorkflow(ctx context.Context, req *connect.Request[v1.GetWorkflowRequest]) (res *connect.Response[v1.GetWorkflowResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) ListWorkflows(ctx context.Context, req *connect.Request[v1.ListWorkflowsRequest]) (res *connect.Response[v1.ListWorkflowsResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) UpdateWorkflow(ctx context.Context, req *connect.Request[v1.UpdateWorkflowRequest]) (res *connect.Response[v1.UpdateWorkflowResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) CancelWorkflow(ctx context.Context, req *connect.Request[v1.CancelWorkflowRequest]) (res *connect.Response[v1.CancelWorkflowResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) RetryWorkflow(ctx context.Context, req *connect.Request[v1.RetryWorkflowRequest]) (res *connect.Response[v1.RetryWorkflowResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) StartDrugTestOrderWorkflow(ctx context.Context, req *connect.Request[v1.StartDrugTestOrderWorkflowRequest]) (res *connect.Response[v1.StartDrugTestOrderWorkflowResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) StartMVRMonitoringWorkflow(ctx context.Context, req *connect.Request[v1.StartMVRMonitoringWorkflowRequest]) (res *connect.Response[v1.StartMVRMonitoringWorkflowResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) StartRandomSelectionWorkflow(ctx context.Context, req *connect.Request[v1.StartRandomSelectionWorkflowRequest]) (res *connect.Response[v1.StartRandomSelectionWorkflowResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) StartBackgroundCheckWorkflow(ctx context.Context, req *connect.Request[v1.StartBackgroundCheckWorkflowRequest]) (res *connect.Response[v1.StartBackgroundCheckWorkflowResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) StartNotificationWorkflow(ctx context.Context, req *connect.Request[v1.StartNotificationWorkflowRequest]) (res *connect.Response[v1.StartNotificationWorkflowResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) GetWorkflowAnalytics(ctx context.Context, req *connect.Request[v1.GetWorkflowAnalyticsRequest]) (res *connect.Response[v1.GetWorkflowAnalyticsResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) GetFailedWorkflows(ctx context.Context, req *connect.Request[v1.GetFailedWorkflowsRequest]) (res *connect.Response[v1.GetFailedWorkflowsResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *ServicesConnectService) GetRunningWorkflows(ctx context.Context, req *connect.Request[v1.GetRunningWorkflowsRequest]) (res *connect.Response[v1.GetRunningWorkflowsResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
