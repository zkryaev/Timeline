package orgmodel

import "time"

type ScheduleList struct {
	WorkerID int `db:"worker_id"`
	OrgID    int `db:"org_id"`
	Schedule []*Schedule
}

type Schedule struct {
	WorkerScheduleID int       `db:"worker_schedule_id"`
	Weekday          int       `db:"weekday"`
	Start            time.Time `db:"start"`
	Over             time.Time `db:"over"`
}

type ScheduleParams struct {
	WorkerID int `db:"worker_id"`
	OrgID    int `db:"org_id"`
	Weekday  int `db:"weekday"`
}
