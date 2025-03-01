package services

import (
	"github.com/flowkater/stopwatch-server/internal/domain/message"
	"github.com/flowkater/stopwatch-server/internal/domain/stopwatch"
	"github.com/flowkater/stopwatch-server/internal/infrastructure/websocket"
)

// StopwatchService 스탑워치 서비스
type StopwatchService struct {
	stopwatchRepo stopwatch.Repository
	clientManager *websocket.ClientManager
}

// NewStopwatchService 스탑워치 서비스 생성 함수
func NewStopwatchService(
	stopwatchRepo stopwatch.Repository,
	clientManager *websocket.ClientManager,
) *StopwatchService {
	return &StopwatchService{
		stopwatchRepo: stopwatchRepo,
		clientManager: clientManager,
	}
}

// GetStopwatch 사용자의 스탑워치 조회
func (s *StopwatchService) GetStopwatch(userID string) *stopwatch.Stopwatch {
	sw, exists := s.stopwatchRepo.FindByUserID(userID)
	if !exists {
		sw = stopwatch.NewStopwatch()
		s.stopwatchRepo.Save(userID, sw)
	}
	return sw
}

// UpdateStopwatch 스탑워치 상태 업데이트
func (s *StopwatchService) UpdateStopwatch(userID string, running bool, elapsedTime float64) *stopwatch.Stopwatch {
	sw := s.GetStopwatch(userID)

	if running {
		sw.Start()
	} else {
		sw.Stop()
	}

	sw.UpdateElapsedTime(elapsedTime)
	s.stopwatchRepo.Save(userID, sw)

	return sw
}

// DeleteStopwatch 스탑워치 삭제
func (s *StopwatchService) DeleteStopwatch(userID string) {
	s.stopwatchRepo.Delete(userID)
}

// BroadcastStopwatchState 스탑워치 상태 브로드캐스트
func (s *StopwatchService) BroadcastStopwatchState(msg *message.Message) {
	s.clientManager.Broadcast(msg)
}
