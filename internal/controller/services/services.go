package services

import (
	"context"
	v1 "v1consortium/api/services/v1"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type Controller struct {
	v1.UnimplementedBackgroundCheckServiceServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterBackgroundCheckServiceServer(s.Server, &Controller{})
}

func (*Controller) OrderBackgroundCheck(ctx context.Context, req *v1.OrderBackgroundCheckRequest) (res *v1.OrderBackgroundCheckResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GetBackgroundCheck(ctx context.Context, req *v1.GetBackgroundCheckRequest) (res *v1.GetBackgroundCheckResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) UpdateBackgroundCheck(ctx context.Context, req *v1.UpdateBackgroundCheckRequest) (res *v1.UpdateBackgroundCheckResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) ListBackgroundChecks(ctx context.Context, req *v1.ListBackgroundChecksRequest) (res *v1.ListBackgroundChecksResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) AddFinding(ctx context.Context, req *v1.AddFindingRequest) (res *v1.AddFindingResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) UpdateFinding(ctx context.Context, req *v1.UpdateFindingRequest) (res *v1.UpdateFindingResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) ListFindings(ctx context.Context, req *v1.ListFindingsRequest) (res *v1.ListFindingsResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) InitiateAdverseAction(ctx context.Context, req *v1.InitiateAdverseActionRequest) (res *v1.InitiateAdverseActionResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) HandleDispute(ctx context.Context, req *v1.HandleDisputeRequest) (res *v1.HandleDisputeResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GetAdverseActionStatus(ctx context.Context, req *v1.GetAdverseActionStatusRequest) (res *v1.GetAdverseActionStatusResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GetAvailablePackages(ctx context.Context, req *v1.GetAvailablePackagesRequest) (res *v1.GetAvailablePackagesResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GetProviderStatus(ctx context.Context, req *v1.GetProviderStatusRequest) (res *v1.GetProviderStatusResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GetBackgroundCheckAnalytics(ctx context.Context, req *v1.GetBackgroundCheckAnalyticsRequest) (res *v1.GetBackgroundCheckAnalyticsResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GetComplianceStatus(ctx context.Context, req *v1.GetComplianceStatusRequest) (res *v1.GetComplianceStatusResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) UpdateComplianceStatus(ctx context.Context, req *v1.UpdateComplianceStatusRequest) (res *v1.UpdateComplianceStatusResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) ListComplianceStatus(ctx context.Context, req *v1.ListComplianceStatusRequest) (res *v1.ListComplianceStatusResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GenerateComplianceCertificate(ctx context.Context, req *v1.GenerateComplianceCertificateRequest) (res *v1.GenerateComplianceCertificateResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GetCertificate(ctx context.Context, req *v1.GetCertificateRequest) (res *v1.GetCertificateResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) ListCertificates(ctx context.Context, req *v1.ListCertificatesRequest) (res *v1.ListCertificatesResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) RevokeCertificate(ctx context.Context, req *v1.RevokeCertificateRequest) (res *v1.RevokeCertificateResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GenerateComplianceReport(ctx context.Context, req *v1.GenerateComplianceReportRequest) (res *v1.GenerateComplianceReportResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GetSavedReport(ctx context.Context, req *v1.GetSavedReportRequest) (res *v1.GetSavedReportResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) ListSavedReports(ctx context.Context, req *v1.ListSavedReportsRequest) (res *v1.ListSavedReportsResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) DeleteSavedReport(ctx context.Context, req *v1.DeleteSavedReportRequest) (res *v1.DeleteSavedReportResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GetComplianceAlerts(ctx context.Context, req *v1.GetComplianceAlertsRequest) (res *v1.GetComplianceAlertsResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GetComplianceAnalytics(ctx context.Context, req *v1.GetComplianceAnalyticsRequest) (res *v1.GetComplianceAnalyticsResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GetComplianceAuditTrail(ctx context.Context, req *v1.GetComplianceAuditTrailRequest) (res *v1.GetComplianceAuditTrailResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) UploadDocument(ctx context.Context, req *v1.UploadDocumentRequest) (res *v1.UploadDocumentResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GetDocument(ctx context.Context, req *v1.GetDocumentRequest) (res *v1.GetDocumentResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) ListDocuments(ctx context.Context, req *v1.ListDocumentsRequest) (res *v1.ListDocumentsResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) UpdateDocument(ctx context.Context, req *v1.UpdateDocumentRequest) (res *v1.UpdateDocumentResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) DeleteDocument(ctx context.Context, req *v1.DeleteDocumentRequest) (res *v1.DeleteDocumentResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GetDocumentVersions(ctx context.Context, req *v1.GetDocumentVersionsRequest) (res *v1.GetDocumentVersionsResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) ShareDocument(ctx context.Context, req *v1.ShareDocumentRequest) (res *v1.ShareDocumentResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GetSharedDocuments(ctx context.Context, req *v1.GetSharedDocumentsRequest) (res *v1.GetSharedDocumentsResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) RevokeDocumentShare(ctx context.Context, req *v1.RevokeDocumentShareRequest) (res *v1.RevokeDocumentShareResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) SearchDocuments(ctx context.Context, req *v1.SearchDocumentsRequest) (res *v1.SearchDocumentsResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GetDocumentAnalytics(ctx context.Context, req *v1.GetDocumentAnalyticsRequest) (res *v1.GetDocumentAnalyticsResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GetDocumentRetentionStatus(ctx context.Context, req *v1.GetDocumentRetentionStatusRequest) (res *v1.GetDocumentRetentionStatusResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) ApplyRetentionPolicy(ctx context.Context, req *v1.ApplyRetentionPolicyRequest) (res *v1.ApplyRetentionPolicyResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) ScheduleDOTPhysical(ctx context.Context, req *v1.ScheduleDOTPhysicalRequest) (res *v1.ScheduleDOTPhysicalResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GetDOTPhysical(ctx context.Context, req *v1.GetDOTPhysicalRequest) (res *v1.GetDOTPhysicalResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) UpdateDOTPhysical(ctx context.Context, req *v1.UpdateDOTPhysicalRequest) (res *v1.UpdateDOTPhysicalResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) ListDOTPhysicals(ctx context.Context, req *v1.ListDOTPhysicalsRequest) (res *v1.ListDOTPhysicalsResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) RegisterMedicalExaminer(ctx context.Context, req *v1.RegisterMedicalExaminerRequest) (res *v1.RegisterMedicalExaminerResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GetMedicalExaminer(ctx context.Context, req *v1.GetMedicalExaminerRequest) (res *v1.GetMedicalExaminerResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) ListMedicalExaminers(ctx context.Context, req *v1.ListMedicalExaminersRequest) (res *v1.ListMedicalExaminersResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GenerateCertificate(ctx context.Context, req *v1.GenerateCertificateRequest) (res *v1.GenerateCertificateResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) ValidateCertificate(ctx context.Context, req *v1.ValidateCertificateRequest) (res *v1.ValidateCertificateResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GetExpiringCertificates(ctx context.Context, req *v1.GetExpiringCertificatesRequest) (res *v1.GetExpiringCertificatesResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) SetExpirationReminder(ctx context.Context, req *v1.SetExpirationReminderRequest) (res *v1.SetExpirationReminderResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) CreateTestingProgram(ctx context.Context, req *v1.CreateTestingProgramRequest) (res *v1.CreateTestingProgramResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GetTestingProgram(ctx context.Context, req *v1.GetTestingProgramRequest) (res *v1.GetTestingProgramResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) ListTestingPrograms(ctx context.Context, req *v1.ListTestingProgramsRequest) (res *v1.ListTestingProgramsResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) OrderDrugTest(ctx context.Context, req *v1.OrderDrugTestRequest) (res *v1.OrderDrugTestResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GetDrugTest(ctx context.Context, req *v1.GetDrugTestRequest) (res *v1.GetDrugTestResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) UpdateDrugTest(ctx context.Context, req *v1.UpdateDrugTestRequest) (res *v1.UpdateDrugTestResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) ListDrugTests(ctx context.Context, req *v1.ListDrugTestsRequest) (res *v1.ListDrugTestsResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) CreateRandomPool(ctx context.Context, req *v1.CreateRandomPoolRequest) (res *v1.CreateRandomPoolResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) AddUsersToPool(ctx context.Context, req *v1.AddUsersToPoolRequest) (res *v1.AddUsersToPoolResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) RemoveUsersFromPool(ctx context.Context, req *v1.RemoveUsersFromPoolRequest) (res *v1.RemoveUsersFromPoolResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GetRandomPool(ctx context.Context, req *v1.GetRandomPoolRequest) (res *v1.GetRandomPoolResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) ListRandomPools(ctx context.Context, req *v1.ListRandomPoolsRequest) (res *v1.ListRandomPoolsResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) ConductRandomSelection(ctx context.Context, req *v1.ConductRandomSelectionRequest) (res *v1.ConductRandomSelectionResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GetRandomSelection(ctx context.Context, req *v1.GetRandomSelectionRequest) (res *v1.GetRandomSelectionResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) ListRandomSelections(ctx context.Context, req *v1.ListRandomSelectionsRequest) (res *v1.ListRandomSelectionsResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) ValidateRandomSelection(ctx context.Context, req *v1.ValidateRandomSelectionRequest) (res *v1.ValidateRandomSelectionResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) OrderMVR(ctx context.Context, req *v1.OrderMVRRequest) (res *v1.OrderMVRResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GetMVRReport(ctx context.Context, req *v1.GetMVRReportRequest) (res *v1.GetMVRReportResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) UpdateMVRReport(ctx context.Context, req *v1.UpdateMVRReportRequest) (res *v1.UpdateMVRReportResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) ListMVRReports(ctx context.Context, req *v1.ListMVRReportsRequest) (res *v1.ListMVRReportsResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) AddMVRViolation(ctx context.Context, req *v1.AddMVRViolationRequest) (res *v1.AddMVRViolationResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) UpdateMVRViolation(ctx context.Context, req *v1.UpdateMVRViolationRequest) (res *v1.UpdateMVRViolationResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) ListMVRViolations(ctx context.Context, req *v1.ListMVRViolationsRequest) (res *v1.ListMVRViolationsResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) EnableContinuousMonitoring(ctx context.Context, req *v1.EnableContinuousMonitoringRequest) (res *v1.EnableContinuousMonitoringResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GetMonitoringStatus(ctx context.Context, req *v1.GetMonitoringStatusRequest) (res *v1.GetMonitoringStatusResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GetMVRAnalytics(ctx context.Context, req *v1.GetMVRAnalyticsRequest) (res *v1.GetMVRAnalyticsResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) SyncProviderData(ctx context.Context, req *v1.SyncProviderDataRequest) (res *v1.SyncProviderDataResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) SendNotification(ctx context.Context, req *v1.SendNotificationRequest) (res *v1.SendNotificationResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GetNotification(ctx context.Context, req *v1.GetNotificationRequest) (res *v1.GetNotificationResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) ListNotifications(ctx context.Context, req *v1.ListNotificationsRequest) (res *v1.ListNotificationsResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) MarkNotificationRead(ctx context.Context, req *v1.MarkNotificationReadRequest) (res *v1.MarkNotificationReadResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) MarkAllNotificationsRead(ctx context.Context, req *v1.MarkAllNotificationsReadRequest) (res *v1.MarkAllNotificationsReadResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) CreateNotificationTemplate(ctx context.Context, req *v1.CreateNotificationTemplateRequest) (res *v1.CreateNotificationTemplateResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GetNotificationTemplate(ctx context.Context, req *v1.GetNotificationTemplateRequest) (res *v1.GetNotificationTemplateResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) ListNotificationTemplates(ctx context.Context, req *v1.ListNotificationTemplatesRequest) (res *v1.ListNotificationTemplatesResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) UpdateNotificationTemplate(ctx context.Context, req *v1.UpdateNotificationTemplateRequest) (res *v1.UpdateNotificationTemplateResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GetNotificationPreferences(ctx context.Context, req *v1.GetNotificationPreferencesRequest) (res *v1.GetNotificationPreferencesResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) UpdateNotificationPreferences(ctx context.Context, req *v1.UpdateNotificationPreferencesRequest) (res *v1.UpdateNotificationPreferencesResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) SendBulkNotification(ctx context.Context, req *v1.SendBulkNotificationRequest) (res *v1.SendBulkNotificationResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) ScheduleNotification(ctx context.Context, req *v1.ScheduleNotificationRequest) (res *v1.ScheduleNotificationResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) ListScheduledNotifications(ctx context.Context, req *v1.ListScheduledNotificationsRequest) (res *v1.ListScheduledNotificationsResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) CancelScheduledNotification(ctx context.Context, req *v1.CancelScheduledNotificationRequest) (res *v1.CancelScheduledNotificationResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GetNotificationAnalytics(ctx context.Context, req *v1.GetNotificationAnalyticsRequest) (res *v1.GetNotificationAnalyticsResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) CreateOrganization(ctx context.Context, req *v1.CreateOrganizationRequest) (res *v1.CreateOrganizationResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GetOrganization(ctx context.Context, req *v1.GetOrganizationRequest) (res *v1.GetOrganizationResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) UpdateOrganization(ctx context.Context, req *v1.UpdateOrganizationRequest) (res *v1.UpdateOrganizationResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) ListOrganizations(ctx context.Context, req *v1.ListOrganizationsRequest) (res *v1.ListOrganizationsResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) DeactivateOrganization(ctx context.Context, req *v1.DeactivateOrganizationRequest) (res *v1.DeactivateOrganizationResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) CreateUser(ctx context.Context, req *v1.CreateUserRequest) (res *v1.CreateUserResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GetUser(ctx context.Context, req *v1.GetUserRequest) (res *v1.GetUserResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (res *v1.UpdateUserResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) ListUsers(ctx context.Context, req *v1.ListUsersRequest) (res *v1.ListUsersResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) DeactivateUser(ctx context.Context, req *v1.DeactivateUserRequest) (res *v1.DeactivateUserResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) CreateSubscription(ctx context.Context, req *v1.CreateSubscriptionRequest) (res *v1.CreateSubscriptionResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GetSubscription(ctx context.Context, req *v1.GetSubscriptionRequest) (res *v1.GetSubscriptionResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) UpdateSubscription(ctx context.Context, req *v1.UpdateSubscriptionRequest) (res *v1.UpdateSubscriptionResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) ListSubscriptionPlans(ctx context.Context, req *v1.ListSubscriptionPlansRequest) (res *v1.ListSubscriptionPlansResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) StartWorkflow(ctx context.Context, req *v1.StartWorkflowRequest) (res *v1.StartWorkflowResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GetWorkflow(ctx context.Context, req *v1.GetWorkflowRequest) (res *v1.GetWorkflowResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) ListWorkflows(ctx context.Context, req *v1.ListWorkflowsRequest) (res *v1.ListWorkflowsResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) UpdateWorkflow(ctx context.Context, req *v1.UpdateWorkflowRequest) (res *v1.UpdateWorkflowResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) CancelWorkflow(ctx context.Context, req *v1.CancelWorkflowRequest) (res *v1.CancelWorkflowResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) RetryWorkflow(ctx context.Context, req *v1.RetryWorkflowRequest) (res *v1.RetryWorkflowResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) StartDrugTestOrderWorkflow(ctx context.Context, req *v1.StartDrugTestOrderWorkflowRequest) (res *v1.StartDrugTestOrderWorkflowResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) StartMVRMonitoringWorkflow(ctx context.Context, req *v1.StartMVRMonitoringWorkflowRequest) (res *v1.StartMVRMonitoringWorkflowResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) StartRandomSelectionWorkflow(ctx context.Context, req *v1.StartRandomSelectionWorkflowRequest) (res *v1.StartRandomSelectionWorkflowResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) StartBackgroundCheckWorkflow(ctx context.Context, req *v1.StartBackgroundCheckWorkflowRequest) (res *v1.StartBackgroundCheckWorkflowResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) StartNotificationWorkflow(ctx context.Context, req *v1.StartNotificationWorkflowRequest) (res *v1.StartNotificationWorkflowResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GetWorkflowAnalytics(ctx context.Context, req *v1.GetWorkflowAnalyticsRequest) (res *v1.GetWorkflowAnalyticsResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GetFailedWorkflows(ctx context.Context, req *v1.GetFailedWorkflowsRequest) (res *v1.GetFailedWorkflowsResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GetRunningWorkflows(ctx context.Context, req *v1.GetRunningWorkflowsRequest) (res *v1.GetRunningWorkflowsResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
