package repository

import (
	"delyaneAPI/models"
)

// GetPaintingById return a demo model of painting DEMO
func GetPaintingById(id string) models.Painting {
	return models.Painting{Title: "Painting 1"}
}
