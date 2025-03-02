package handlers

import (
	"fmt"
	"log"

	"github.com/flowkater/stopwatch-server/internal/application/services"
	"github.com/flowkater/stopwatch-server/internal/infrastructure/websocket"
	"github.com/flowkater/stopwatch-server/internal/interfaces/dto"
	"github.com/gofiber/fiber/v2"
	ws "github.com/gofiber/websocket/v2"
)

// WebSocketHandler 웹소켓 핸들러
type WebSocketHandler struct {
	userStateService *services.UserStateService
	clientManager    *websocket.ClientManager
}

// NewWebSocketHandler 웹소켓 핸들러 생성 함수
func NewWebSocketHandler(
	userStateService *services.UserStateService,
	clientManager *websocket.ClientManager,
) *WebSocketHandler {
	return &WebSocketHandler{
		userStateService: userStateService,
		clientManager:    clientManager,
	}
}

// Register 라우터에 핸들러 등록
func (h *WebSocketHandler) Register(app *fiber.App) {
	app.Use("/ws", func(c *fiber.Ctx) error {
		if ws.IsWebSocketUpgrade(c) {
			c.Set("Content-Type", "application/json; charset=utf-8")
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws", ws.New(h.handleWebSocket))
}

// handleWebSocket 웹소켓 핸들링
func (h *WebSocketHandler) handleWebSocket(c *ws.Conn) {
	// 연결 종료 시 실행될 defer 함수
	defer func() {
		userID := h.clientManager.Unregister(c)
		if userID != "" {
			// 사용자 오프라인 처리
			state := h.userStateService.SetUserOnline(userID, false)

			// 오프라인 상태 브로드캐스트
			h.userStateService.BroadcastUserState(state)
		}
		c.Close()
	}()

	// 메시지 루프
	for {
		var msg dto.Message
		if err := c.ReadJSON(&msg); err != nil {
			log.Println("read error:", err)
			break
		}

		fmt.Println("msg", msg)

		// 클라이언트 등록
		h.clientManager.Register(c, msg.UserID)

		// 사용자 상태 업데이트
		state := h.userStateService.UpdateUserState(
			msg.UserID,
			msg.Online, // 온라인
			msg.StopwatchRunning,
			msg.ElapsedTime,
		)

		// 상태 브로드캐스트
		h.userStateService.BroadcastUserState(state)
	}
}
