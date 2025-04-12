package layer

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"thesis_back/internal/usecase/layer"
)

type LayerHandler struct {
	lu     *layer.ILayerUseCase
	logger *zap.Logger
}

func NewLayerHandler(lu *layer.ILayerUseCase, logger *zap.Logger) *LayerHandler {
	return &LayerHandler{
		lu:     lu,
		logger: logger.Named("LayerHandler"),
	}
}

func (h *LayerHandler) Create(c *gin.Context) {
	userID := c.GetInt("userID")

	if userID == 0 {

	}
}

func (h *LayerHandler) Delete(c *gin.Context) {

}

func (h *LayerHandler) Get(c *gin.Context) {

}

func (h *LayerHandler) GetByID(c *gin.Context) {

}

func (h *LayerHandler) Update(c *gin.Context) {

}
