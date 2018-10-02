package entities

import (
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

type TimesV1 struct {
	Items   []*TimeV1 `json:"items"`
	Page    int       `json:"page,omitempty"`
	PerPage int       `json:"per_page,omitempty"`
	Total   int       `json:"total,omitempty"`
}
