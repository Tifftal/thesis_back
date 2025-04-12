package layer

import (
	"encoding/json"
	"thesis_back/internal/domain"
)

type LayerResponse struct {
	ID           uint                   `json:"id"`
	ImageID      uint                   `json:"imageID"`
	Name         string                 `json:"name"`
	Measurements map[string]interface{} `json:"measurements"`
}

type CreateLayerDTO struct {
	ImageID uint   `json:"imageID"`
	Name    string `json:"name"`
}

type UpdateLayerDTO struct {
	Name         string                 `json:"name"`
	Measurements map[string]interface{} `json:"measurements"`
}

func ToLayerResponse(layer *domain.Layer) LayerResponse {
	var measurements map[string]interface{}
	if err := json.Unmarshal(layer.Measurements, &measurements); err != nil {
		measurements = make(map[string]interface{})
	}

	return LayerResponse{
		ID:           layer.ID,
		ImageID:      layer.ImageID,
		Name:         layer.Name,
		Measurements: measurements,
	}
}

type ErrorResponse struct {
	Message string `json:"message"`
}
