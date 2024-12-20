package recordcase

import (
	"context"
	"timeline/internal/entity/dto/recordto"
	"timeline/internal/repository"
	"timeline/internal/repository/mail"
	"timeline/internal/repository/mail/entity"
	"timeline/internal/repository/mapper/recordmap"

	"go.uber.org/zap"
)

type RecordUseCase struct {
	users   repository.UserRepository
	orgs    repository.OrgRepository
	records repository.RecordRepository
	mail    mail.Post
	Logger  *zap.Logger
}

func New(userRepo repository.UserRepository, orgRepo repository.OrgRepository, recordRepo repository.RecordRepository, logger *zap.Logger) *RecordUseCase {
	return &RecordUseCase{
		users:   userRepo,
		orgs:    orgRepo,
		records: recordRepo,
		Logger:  logger,
	}
}

func (r *RecordUseCase) Record(ctx context.Context, recordID int) (*recordto.RecordScrap, error) {
	data, err := r.records.Record(ctx, recordID)
	if err != nil {
		r.Logger.Error(
			"failed to get record",
			zap.Error(err),
		)
		return nil, err
	}
	return recordmap.RecordScrapToDTO(data), nil
}

func (r *RecordUseCase) RecordList(ctx context.Context, params *recordto.RecordListParams) (*recordto.RecordList, error) {
	data, found, err := r.records.RecordList(ctx, recordmap.RecordParamsToModel(params))
	if err != nil {
		r.Logger.Error(
			"failed to get record list",
			zap.Error(err),
		)
		return nil, err
	}
	resp := &recordto.RecordList{
		List:  recordmap.RecordListToDTO(data),
		Found: found,
	}
	return resp, nil
}

func (r *RecordUseCase) RecordAdd(ctx context.Context, rec *recordto.Record) error {
	record, err := r.records.RecordAdd(ctx, recordmap.RecordToModel(rec))
	if err != nil {
		r.Logger.Error(
			"failed to add record",
			zap.Error(err),
		)
		return err
	}
	r.mail.SendMsg(&entity.Message{
		Email:    record.UserEmail,
		Type:     mail.ReminderType,
		Value:    recordmap.RecordToReminder(record),
		IsAttach: true,
	})
	return nil
}

func (r *RecordUseCase) RecordPatch(ctx context.Context, rec *recordto.Record) error {
	if err := r.records.RecordPatch(ctx, recordmap.RecordToModel(rec)); err != nil {
		r.Logger.Error(
			"failed to add record",
			zap.Error(err),
		)
		return err
	}
	return nil
}

func (r *RecordUseCase) RecordDelete(ctx context.Context, rec *recordto.Record) error {
	if err := r.records.RecordDelete(ctx, recordmap.RecordToModel(rec)); err != nil {
		r.Logger.Error(
			"failed to add record",
			zap.Error(err),
		)
		return err
	}
	return nil
}
