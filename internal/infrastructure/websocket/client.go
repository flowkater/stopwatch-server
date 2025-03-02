package websocket

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/flowkater/stopwatch-server/internal/interfaces/dto"
	"github.com/gofiber/websocket/v2"
)

// Client 웹소켓 클라이언트
type Client struct {
	Conn   *websocket.Conn
	UserID string
}

// ClientManager 웹소켓 클라이언트 관리자
type ClientManager struct {
	clients map[*websocket.Conn]*Client
	mu      sync.RWMutex
}

// NewClientManager 클라이언트 관리자 생성 함수
func NewClientManager() *ClientManager {
	return &ClientManager{
		clients: make(map[*websocket.Conn]*Client),
	}
}

// Register 클라이언트 등록
func (m *ClientManager) Register(conn *websocket.Conn, userID string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.clients[conn] = &Client{
		Conn:   conn,
		UserID: userID,
	}
}

// Unregister 클라이언트 등록 해제
func (m *ClientManager) Unregister(conn *websocket.Conn) string {
	m.mu.Lock()
	defer m.mu.Unlock()

	if client, ok := m.clients[conn]; ok {
		userID := client.UserID
		delete(m.clients, conn)
		return userID
	}
	return ""
}

// GetUserID 연결로부터 사용자 ID 조회
func (m *ClientManager) GetUserID(conn *websocket.Conn) (string, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	client, ok := m.clients[conn]
	if !ok {
		return "", false
	}
	return client.UserID, true
}

// Broadcast 모든 클라이언트에게 메시지 전송
func (m *ClientManager) Broadcast(msg *dto.Message) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	jsonData, err := json.Marshal(msg)
	if err != nil {
		log.Println("JSON marshaling error:", err)
		return
	}

	log.Printf("Broadcasting message: %s", string(jsonData))

	for conn := range m.clients {
		if err := conn.WriteJSON(msg); err != nil {
			log.Println("write error:", err)
		}
	}
}
