package factories

import (
	"github.com/hamidfzm/timechi-server/helpers"
	"github.com/hamidfzm/timechi-server/models"
	"github.com/hamidfzm/timechi-server/entities"
)

func GetTimeV1(m *models.Time) *entities.TimeV1 {
	return &entities.TimeV1{
		ID:        m.ID,
		Title:     m.Title,
		StartedAt: helpers.JSONTime{Time: m.StartedAt},
		StoppedAt: helpers.JSONTime{Time: m.StoppedAt},
		Duration:  m.Duration,
		UserID:    m.UserID,
	}
}

func GetTimesV1(ms []*models.Time, page int, perPage int, total int) *entities.TimesV1 {
	items := make([]*entities.TimeV1, len(ms))
	for i, m := range ms {
		items[i] = GetTimeV1(m)
	}
	
	return &entities.TimesV1{
		Page:    page,
		PerPage: perPage,
		Total:   total,
		Items:   items,
	}
}
