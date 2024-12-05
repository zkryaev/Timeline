package orgcase

import (
	"context"
	"timeline/internal/entity/dto/orgdto"
	"timeline/internal/repository/mapper/orgmap"

	"go.uber.org/zap"
)

func (o *OrgUseCase) Service(ctx context.Context, ServiceID, OrgID int) (*orgdto.ServiceResp, error) {
	service, err := o.org.Service(ctx, ServiceID, OrgID)
	if err != nil {
		o.Logger.Error(
			"failed to get service",
			zap.Error(err),
		)
		return nil, err
	}
	return orgmap.ServiceToDTO(service), nil
}

func (o *OrgUseCase) ServiceWorkerList(ctx context.Context, ServiceID, OrgID int) ([]*orgdto.WorkerResp, error) {
	data, err := o.org.ServiceWorkerList(ctx, ServiceID, OrgID)
	if err != nil {
		o.Logger.Error(
			"failed to get service",
			zap.Error(err),
		)
		return nil, err
	}
	workers := make([]*orgdto.WorkerResp, 0, len(data))
	for _, worker := range data {
		workers = append(workers, orgmap.WorkerToDTO(worker))
	}

	return workers, nil
}
func (o *OrgUseCase) ServiceAdd(ctx context.Context, Service *orgdto.AddServiceReq) (*orgdto.ServiceResp, error) {
	serviceID, err := o.org.ServiceAdd(ctx, orgmap.AddServiceToModel(Service))
	if err != nil {
		o.Logger.Error(
			"failed to add service",
			zap.Error(err),
		)
		return nil, nil
	}
	return &orgdto.ServiceResp{
		ServiceID: serviceID,
	}, nil
}

func (o *OrgUseCase) ServiceUpdate(ctx context.Context, Service *orgdto.UpdateServiceReq) (*orgdto.UpdateServiceReq, error) {
	if err := o.org.ServiceUpdate(ctx, orgmap.UpdateService(Service)); err != nil {
		o.Logger.Error(
			"failed to update service",
			zap.Error(err),
		)
		return nil, nil
	}
	return Service, nil
}

func (o *OrgUseCase) ServiceList(ctx context.Context, OrgID int) ([]*orgdto.ServiceResp, error) {
	data, err := o.org.ServiceList(ctx, OrgID)
	if err != nil {
		o.Logger.Error(
			"failed to retrieve list of services",
			zap.Error(err),
		)
		return nil, nil
	}
	serviceList := make([]*orgdto.ServiceResp, 0, len(data))
	for _, v := range data {
		serviceList = append(serviceList, orgmap.ServiceToDTO(v))
	}
	return serviceList, nil
}

func (o *OrgUseCase) ServiceDelete(ctx context.Context, ServiceID, OrgID int) error {
	if err := o.org.ServiceDelete(ctx, ServiceID, OrgID); err != nil {
		o.Logger.Error(
			"failed to delete service",
			zap.Error(err),
		)
	}
	return nil
}