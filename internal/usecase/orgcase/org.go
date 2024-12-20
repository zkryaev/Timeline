package orgcase

import (
	"context"
	"errors"
	"fmt"
	"timeline/internal/entity/dto/orgdto"
	"timeline/internal/repository"
	"timeline/internal/repository/database/postgres"
	"timeline/internal/repository/mapper/orgmap"

	"go.uber.org/zap"
)

type OrgUseCase struct {
	user   repository.UserRepository
	org    repository.OrgRepository
	Logger *zap.Logger
}

func New(userRepo repository.UserRepository, orgRepo repository.OrgRepository, logger *zap.Logger) *OrgUseCase {
	return &OrgUseCase{
		user:   userRepo,
		org:    orgRepo,
		Logger: logger,
	}
}

func (o *OrgUseCase) Organization(ctx context.Context, id int) (*orgdto.Organization, error) {
	if id <= 0 {
		return nil, fmt.Errorf("Id must be > 0")
	}
	data, err := o.org.OrgByID(ctx, id)
	if err != nil {
		o.Logger.Error(
			"failed to get org",
			zap.Error(err),
		)
		return nil, err
	}
	return orgmap.OrganizationToDTO(data), nil
}

func (o *OrgUseCase) OrgUpdate(ctx context.Context, newOrg *orgdto.OrgUpdateReq) error {
	if err := o.org.OrgUpdate(ctx, orgmap.OrgUpdateToModel(newOrg)); err != nil {
		if errors.Is(err, postgres.ErrOrgNotFound) {
			return err
		}
		o.Logger.Error(
			"failed to update org",
			zap.Error(err),
		)
		return err
	}
	return nil
}
