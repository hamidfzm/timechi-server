package jsons

import (
	"github.com/hamidfzm/timechi-server/models"
	"time"
	"github.com/hamidfzm/timechi-server/helpers"
)

type StartTimerV1 struct {
	Title string `json:"title" validate:"required"`
}

type TimeV1 struct {
	ID        uint             `json:"id"`
	Title     string           `json:"title"`
	StartedAt helpers.JSONTime `json:"started_at"`
	StoppedAt helpers.JSONTime `json:"stopped_at,omitempty"`
	Duration  time.Duration    `json:"duration,omitempty"`
	UserID    uint             `json:"user_id,omitempty"`
}

func (j *TimeV1) From(m *models.Time) {
	j.ID = m.ID
	j.Title = m.Title
	j.StartedAt = helpers.JSONTime{Time: m.StartedAt}
	j.StoppedAt = helpers.JSONTime{Time: m.StoppedAt}
	j.Duration = m.Duration
	j.UserID = m.UserID
}

type TimesV1 struct {
	Items   []TimeV1 `json:"items"`
	Page    int      `json:"page,omitempty"`
	PerPage int      `json:"per_page,omitempty"`
	Total   int      `json:"total,omitempty"`
}

func (j *TimesV1) From(ms *[]models.Time, page int, perPage int, total int) {
	j.Items = make([]TimeV1, len(*ms))
	for i, m := range *ms {
		j.Items[i].From(&m)
	}
	j.Page = page
	j.PerPage = perPage
	j.Total = total
}
