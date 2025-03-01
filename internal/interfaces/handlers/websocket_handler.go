package handlers

import (
	"log"

	"github.com/flowkater/stopwatch-server/internal/application/services"
	"github.com/flowkater/stopwatch-server/internal/domain/message"
	"github.com/flowkater/stopwatch-server/internal/infrastructure/websocket"
	"github.com/gofiber/fiber/v2"
	ws "github.com/gofiber/websocket/v2"
)

// WebSocketHandler 웹소켓 핸들러
type WebSocketHandler struct {
	userService      *services.UserService
	stopwatchService *services.StopwatchService
	clientManager    *websocket.ClientManager
}

// NewWebSocketHandler 웹소켓 핸들러 생성 함수
func NewWebSocketHandler(
	userService *services.UserService,
	stopwatchService *services.StopwatchService,
	clientManager *websocket.ClientManager,
) *WebSocketHandler {
	return &WebSocketHandler{
		userService:      userService,
		stopwatchService: stopwatchService,
		clientManager:    clientManager,
	}
}

// Register 라우터에 핸들러 등록
func (h *WebSocketHandler) Register(app *fiber.App) {
	app.Use("/ws", func(c *fiber.Ctx) error {
		if ws.IsWebSocketUpgrade(c) {
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
			h.userService.SetUserOnline(userID, false)

			// 오프라인 상태 브로드캐스트
			sw := h.stopwatchService.GetStopwatch(userID)
			msg := message.NewMessage(userID, false, sw.Running, sw.ElapsedTime)
			h.clientManager.Broadcast(msg)
		}
		c.Close()
	}()

	// 메시지 루프
	for {
		var msg message.Message
		if err := c.ReadJSON(&msg); err != nil {
			log.Println("read error:", err)
			break
		}

		// 클라이언트 등록
		h.clientManager.Register(c, msg.UserID)

		// 사용자 온라인 처리
		h.userService.SetUserOnline(msg.UserID, true)

		// 스탑워치 상태 업데이트
		h.stopwatchService.UpdateStopwatch(msg.UserID, msg.StopwatchRunning, msg.ElapsedTime)

		// 상태 브로드캐스트
		h.clientManager.Broadcast(&msg)
	}
}
