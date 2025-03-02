package handlers

import (
	"log"
	"net/url"

	"github.com/flowkater/stopwatch-server/internal/application/services"
	"github.com/flowkater/stopwatch-server/internal/interfaces/dto"
	"github.com/gofiber/fiber/v2"
)

type UserStateHandler struct {
	userStateService *services.UserStateService
}

func NewUserStateHandler(userStateService *services.UserStateService) *UserStateHandler {
	return &UserStateHandler{
		userStateService: userStateService,
	}
}

func (h *UserStateHandler) Register(app *fiber.App) {
	api := app.Group("/api/user-states")

	api.Get("/:userID", h.GetUserState)
	api.Get("/", h.GetUserStateList)
}

// GetUserState 사용자 상태 조회
func (h *UserStateHandler) GetUserState(c *fiber.Ctx) error {
	userID := c.Params("userID")
	if userID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "userID is required",
		})
	}
	decodedUserID, err := url.QueryUnescape(userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID encoding",
		})
	}

	userState := h.userStateService.GetUserState(decodedUserID)
	if userState == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "사용자 상태를 찾을 수 없습니다",
		})
	}
	response := dto.FromUserState(userState)
	log.Println("response", response)

	return c.JSON(response)
}

func (h *UserStateHandler) GetUserStateList(c *fiber.Ctx) error {
	userStateList := h.userStateService.GetUserStateList()
	response := dto.FromUserStateList(userStateList)
	log.Println("response", response)

	return c.JSON(response)
}
