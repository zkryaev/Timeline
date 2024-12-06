package recordmap

import (
	"timeline/internal/entity/dto/recordto"
	"timeline/internal/repository/mapper/orgmap"
	"timeline/internal/repository/mapper/usermap"
	"timeline/internal/repository/models/recordmodel"
)

func RecordToModel(dto *recordto.Record) *recordmodel.Record {
	return &recordmodel.Record{
		RecordID:  dto.RecordID,
		OrgID:     dto.OrgID,
		UserID:    dto.UserID,
		SlotID:    dto.SlotID,
		ServiceID: dto.ServiceID,
		WorkerID:  dto.WorkerID,
		Reviewed:  dto.Reviewed,
	}
}

func RecordToDTO(model *recordmodel.Record) *recordto.Record {
	return &recordto.Record{
		RecordID:  model.RecordID,
		OrgID:     model.OrgID,
		UserID:    model.UserID,
		SlotID:    model.SlotID,
		ServiceID: model.ServiceID,
		WorkerID:  model.WorkerID,
		Reviewed:  model.Reviewed,
	}
}

func RecordParamsToModel(dto *recordto.RecordListParams) *recordmodel.RecordListParams {
	return &recordmodel.RecordListParams{
		OrgID:    dto.OrgID,
		UserID:   dto.UserID,
		Reviewed: dto.Reviewed,
	}
}

func RecordScrapToDTO(model *recordmodel.RecordScrap) *recordto.RecordScrap {
	return &recordto.RecordScrap{
		RecordID: model.RecordID,
		Reviewed: model.Reviewed,
		Org:      orgmap.OrgInfoToEntity(model.Org),
		User:     usermap.UserInfoToDTO(model.User),
		Slot:     orgmap.SlotInfoToDTO(model.Slot),
		Service:  orgmap.ServiceToEntity(model.Service),
		Worker:   orgmap.WorkerToEntity(model.Worker),
		Feedback: FeedbackToDTO(model.Feedback),
	}
}

func RecordListToDTO(model []*recordmodel.RecordScrap) *recordto.RecordList {
	resp := &recordto.RecordList{
		List: make([]*recordto.RecordScrap, 0, len(model)),
	}
	for _, v := range model {
		resp.List = append(resp.List, RecordScrapToDTO(v))
	}
	return resp
}
