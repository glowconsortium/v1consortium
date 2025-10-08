package organization

import (
	"context"
	"fmt"
	"v1consortium/internal/dao"
	"v1consortium/internal/model/do"
	"v1consortium/internal/model/entity"
	"v1consortium/internal/service"
)

func new() service.IOrganizationService {
	return &sOrganizationService{}
}

func init() {
	service.RegisterOrganizationService(new())
}

type sOrganizationService struct{}

func (s *sOrganizationService) GetOrganization(ctx context.Context, orgnanizationid string) (*entity.Organizations, error) {
	var organization entity.Organizations
	err := dao.Organizations.Ctx(ctx).Where(dao.Organizations.Columns().Id, orgnanizationid).Scan(&organization)
	if err != nil {
		return nil, err
	}
	return &organization, nil

}

func (s *sOrganizationService) CreateOrganization(ctx context.Context, createdata *do.Organizations) (*entity.Organizations, error) {

	_, err := dao.Organizations.Ctx(ctx).Data(createdata).Insert()
	if err != nil {
		return nil, err
	}

	// createdata.Id is interface{}; assert or convert to string
	var id string
	switch v := createdata.Id.(type) {
	case string:
		id = v
	case []byte:
		id = string(v)
	default:
		return nil, fmt.Errorf("unsupported id type: %T", createdata.Id)
	}

	return s.GetOrganization(ctx, id)
}
func (s *sOrganizationService) UpdateOrganization(ctx context.Context, id string, updatedata *do.Organizations) (*entity.Organizations, error) {

	_, err := dao.Organizations.Ctx(ctx).Where(dao.Organizations.Columns().Id, id).Data(updatedata).Update()
	if err != nil {
		return nil, err
	}
	return s.GetOrganization(ctx, id)
}
func (s *sOrganizationService) DeleteOrganization(ctx context.Context, id string) error {
	_, err := dao.Organizations.Ctx(ctx).Where(dao.Organizations.Columns().Id, id).Delete()
	return err
}

func (s *sOrganizationService) ListOrganizations(ctx context.Context, offset, limit int) ([]*entity.Organizations, error) {
	var organizations []*entity.Organizations
	err := dao.Organizations.Ctx(ctx).Offset(offset).Limit(limit).Scan(&organizations)
	if err != nil {
		return nil, err
	}
	return organizations, nil
}

func (s *sOrganizationService) DeactivateOrganization(ctx context.Context, id string) error {
	_, err := dao.Organizations.Ctx(ctx).Where(dao.Organizations.Columns().Id, id).Data(do.Organizations{
		IsActive: false,
	}).Update()
	return err
}

// func (s *sOrganizationService) CreateOrganizationUser(ctx context.Context, createdata *do.OrganizationUsers) (*entity.OrganizationUsers, error) {

// 	_, err := dao.OrganizationUsers.Ctx(ctx).Data(createdata).Insert()
// 	if err != nil {
// 		return nil, err
// 	}

// 	// createdata.Id is interface{}; assert or convert to string
// 	var id string
// 	switch v := createdata.Id.(type) {
// 	case string:
// 		id = v
// 	case []byte:
// 		id = string(v)
// 	default:
// 		return nil, fmt.Errorf("unsupported id type: %T", createdata.Id)
// 	}

// 	return s.GetOrganizationUser(ctx, id)
// }
// func (s *sOrganizationService) GetOrganizationUser(ctx context.Context, id string) (*entity.OrganizationUsers, error) {
// 	var organizationuser entity.OrganizationUsers
// 	err := dao.OrganizationUsers.Ctx(ctx).Where(dao.OrganizationUsers.Columns().Id, id).Scan(&organizationuser)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &organizationuser, nil

// }
