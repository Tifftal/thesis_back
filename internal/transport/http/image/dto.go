package image

import (
	"mime/multipart"
	"thesis_back/internal/domain"
	layer_dto "thesis_back/internal/transport/http/layer"
)

type AddImageDTO struct {
	ProjectID uint                  `form:"projectID" binding:"required"`
	Name      string                `form:"name" binding:"required"`
	ImageFile *multipart.FileHeader `form:"image" binding:"required"`
}

type UpdateImageDTO struct {
	Name string `json:"name" binding:"required"`
}

type ImageResponse struct {
	ID        uint                      `json:"id"`
	Name      string                    `json:"name"`
	FileName  string                    `json:"fileName"`
	ProjectID uint                      `json:"projectID"`
	Layers    []layer_dto.LayerResponse `json:"layers"`
	URL       string                    `json:"url"`
}

func ToImageResponse(image *domain.Image) ImageResponse {
	layers := make([]layer_dto.LayerResponse, len(image.Layers))
	for i, layer := range image.Layers {
		layers[i] = layer_dto.ToLayerResponse(&layer)
	}

	return ImageResponse{
		ID:        image.ID,
		Name:      image.Name,
		FileName:  image.FileName,
		ProjectID: image.ProjectID,
		URL:       image.URL,
		Layers:    layers,
	}
}

type ErrorResponse struct {
	Message string `json:"message"`
}
