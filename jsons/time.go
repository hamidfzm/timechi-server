package jsons

import (
	"github.com/hamidfzm/timechi-server/models"
	"time"
	"github.com/hamidfzm/timechi-server/helpers"
)

type TimeV1 struct {
	ID        uint             `json:"id"`
	Title     string           `json:"title"`
	StartedAt helpers.JSONTime `json:"started_at"`
	StoppedAt helpers.JSONTime `json:"stopped_at"`
	Duration  time.Duration    `json:"duration"`
	UserID    uint             `json:"user_id"`
}

func (j *TimeV1) From(m *models.Time) {
	j.ID = m.ID
	j.Title = m.Title
	j.StartedAt = helpers.JSONTime{Time: m.StartedAt}
	j.StoppedAt = helpers.JSONTime{Time: m.StoppedAt}
	j.Duration = m.Duration
	j.UserID = m.UserID
}
