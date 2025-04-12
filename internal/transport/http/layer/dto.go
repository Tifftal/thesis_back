package layer

import (
	"encoding/json"
	"thesis_back/internal/domain"
)

type LayerResponse struct {
	ID           uint                   `json:"id"`
	ImageID      uint                   `json:"image_id"`
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
		Measurements: measurements,
	}
}
