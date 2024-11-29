package orgdto

import "timeline/internal/entity"

type Organization struct {
	OrgID int             `json:"id"`
	Info  *entity.OrgInfo `json:"info"`
}

type OrgUpdateReq struct {
	OrgID     int                 `json:"org_id" validate:"required"`
	Name      string              `json:"name" validate:"min=3,max=100"`
	Address   string              `json:"address"`
	Long      float64             `json:"long" validate:"longitude"`
	Lat       float64             `json:"lat" validate:"latitude"`
	Type      string              `json:"type"`
	Telephone string              `json:"telephone" validate:"e164"`
	City      string              `json:"city"`
	About     string              `json:"about,omitempty" validate:"max=1500"`
	Timetable []*entity.OpenHours `json:"timetable,omitempty"`
}

type OrgUpdateResp struct {
	OrgID     int                 `json:"org_id"`
	Name      string              `json:"name"`
	Address   string              `json:"address"`
	Long      float64             `json:"long"`
	Lat       float64             `json:"lat"`
	Type      string              `json:"type"`
	Telephone string              `json:"telephone"`
	City      string              `json:"city"`
	About     string              `json:"about,omitempty"`
	Timetable []*entity.OpenHours `json:"timetable,omitempty"`
}
