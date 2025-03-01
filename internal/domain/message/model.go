package message

// Message: 클라이언트와 주고받는 데이터 구조
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
