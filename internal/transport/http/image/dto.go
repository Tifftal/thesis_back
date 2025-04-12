package image

import (
	"thesis_back/internal/domain"
	layer_dto "thesis_back/internal/transport/http/layer"
)

type ImageResponse struct {
	ID        uint                      `json:"id"`
	Name      string                    `json:"name"`
	ProjectID uint                      `json:"projectID"`
	Layers    []layer_dto.LayerResponse `json:"layers"`
}

func ToImageResponse(image *domain.Image) ImageResponse {
	layers := make([]layer_dto.LayerResponse, len(image.Layers))
	for i, layer := range image.Layers {
		layers[i] = layer_dto.ToLayerResponse(&layer)
	}

	return ImageResponse{
		ID:        image.ID,
		Name:      image.Name,
		ProjectID: image.ProjectID,
		Layers:    layers,
	}
}
