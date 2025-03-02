package dto

import "github.com/flowkater/stopwatch-server/internal/domain/user_state"

// Message: 클라이언트와 주고받는 데이터 구조 (DTO)
type Message struct {
	UserID           string  `json:"userId"`
	Online           bool    `json:"online"`
	StopwatchRunning bool    `json:"stopwatchRunning"`
	ElapsedTime      float64 `json:"elapsedTime"`
}

// NewMessage 메시지 생성 함수
func NewMessage(userID string, online bool, stopwatchRunning bool, elapsedTime float64) *Message {
	return &Message{
		UserID:           userID,
		Online:           online,
		StopwatchRunning: stopwatchRunning,
		ElapsedTime:      elapsedTime,
	}
}

// FromUserState 사용자 상태로부터 메시지 생성
func FromUserState(state *user_state.UserState) *Message {
	return &Message{
		UserID:           state.UserID,
		Online:           state.Online,
		StopwatchRunning: state.StopwatchRunning,
		ElapsedTime:      state.ElapsedTime,
	}
}

// FromUserStateList 사용자 상태 목록으로부터 메시지 생성
func FromUserStateList(states []*user_state.UserState) []*Message {
	messages := make([]*Message, len(states))
	for i, state := range states {
		messages[i] = FromUserState(state)
	}
	return messages
}

// ToUserState 메시지로부터 사용자 상태 생성
func (m *Message) ToUserState() *user_state.UserState {
	state := user_state.NewUserState(m.UserID)
	state.SetOnline(m.Online)
	state.UpdateStopwatch(m.StopwatchRunning, m.ElapsedTime)
	return state
}
