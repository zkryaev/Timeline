package recordmodel

import (
	"timeline/internal/repository/models/orgmodel"
	"timeline/internal/repository/models/usermodel"
)

type Record struct {
	RecordID  int  `db:"record_id"`
	OrgID     int  `db:"org_id"`
	UserID    int  `db:"user_id"`
	SlotID    int  `db:"slot_id"`
	ServiceID int  `db:"service_id"`
	WorkerID  int  `db:"worker_id"`
	Reviewed  bool `db:"reviewed"`
}

type RecordListParams struct {
	OrgID    int  `db:"org_id"`
	UserID   int  `db:"user_id"`
	Reviewed bool `db:"reviewed"`
}

type RecordScrap struct {
	RecordID int  `db:"record_id"`
	Reviewed bool `db:"reviewed"`
	Org      *orgmodel.OrgInfo
	User     *usermodel.UserInfo
	Slot     *orgmodel.Slot
	Service  *orgmodel.Service
	Worker   *orgmodel.Worker
	Feedback *Feedback
}
