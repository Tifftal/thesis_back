package user

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"thesis_back/internal/domain"
	"thesis_back/internal/usecase/user"
)

type UserHandler struct {
	uc     user.IUserUseCase
	logger *zap.Logger
}

func NewUserHandler(uc user.IUserUseCase, logger *zap.Logger) *UserHandler {

	return &UserHandler{
		uc:     uc,
		logger: logger.Named("UserHandler"),
	}
}

// Register godoc
// @Summary Зарегистрировать новго пользователя
// @tags auth
// @Accept json
// @Produce json
// @Param input body CreateUserDTO true "Данные пользователя"
// @Success 201 {object} AuthResponse
// @Success 400 {object} ErrorResponse
// @Router /auth/register [post]
func (h *UserHandler) Register(c *gin.Context) {
	var req CreateUserDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Warn("Validation error", zap.Error(err))
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: domain.ErrInvalidRequestBody.Error()})
		return
	}

	user, err := h.uc.Register(c.Request.Context(), &domain.User{
		Username:   req.Username,
		Password:   req.Password,
		FirstName:  req.FirstName,
		LastName:   req.LastName,
		Patronymic: req.Patronymic,
	})
	if err != nil {
		h.logger.Error("Register error", zap.Error(err))
		c.JSON(errorStatusCode(err), ErrorResponse{Message: err.Error()})
		return
	}

	response := ToUserResponse(user.ID, user.Username, user.FirstName, user.LastName, user.Patronymic, user.CreatedAt, user.UpdatedAt)

	c.JSON(http.StatusCreated, response)
}

// Login godoc
// @Summary Аутентификация пользователя
// @Tags auth
// @Accept json
// @Produce json
// @Param input body LoginUserDTO true "Данные для авторизации"
// @Success 200 {object} AuthResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /auth/login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var req LoginUserDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Warn("Validation error", zap.Error(err))
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: domain.ErrInvalidRequestBody.Error()})
		return
	}

	user, err := h.uc.Authenticate(c.Request.Context(), req.Username, req.Password)
	if err != nil {
		h.logger.Warn("Authenticate error", zap.Error(err))
		c.JSON(errorStatusCode(err), ErrorResponse{Message: err.Error()})
		return
	}

	tokens, err := h.uc.GenerateTokens(user)
	if err != nil {
		h.logger.Warn("GenerateTokens error", zap.Error(err))
		c.JSON(errorStatusCode(err), ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, ToAuthResponse(tokens, user))
}

func (h *UserHandler) Refresh(c *gin.Context) {}

// Me godoc
// @Summary Получение информации о текущем пользователе
// @Tags user
// @Security BearerAuth
// @Produce json
// @Success 200 {object} UserResponse
// @Failure 401 {object} ErrorResponse
// @Router /user/me [get]
func (h *UserHandler) Me(c *gin.Context) {}

func errorStatusCode(err error) int {
	switch {
	case errors.Is(err, domain.ErrUserNotFound):
		return http.StatusNotFound
	case errors.Is(err, domain.ErrInvalidCredentials):
		return http.StatusUnauthorized
	case errors.Is(err, domain.ErrUserExists):
		return http.StatusConflict
	case errors.Is(err, domain.ErrUnauthorized):
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}
