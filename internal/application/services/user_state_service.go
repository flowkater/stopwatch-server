package services

import (
	"log"

	"github.com/flowkater/stopwatch-server/internal/domain/user_state"
	"github.com/flowkater/stopwatch-server/internal/infrastructure/websocket"
	"github.com/flowkater/stopwatch-server/internal/interfaces/dto"
)

// UserStateService 사용자 상태 서비스
type UserStateService struct {
	userStateRepo user_state.Repository
	clientManager *websocket.ClientManager
}

// NewUserStateService 사용자 상태 서비스 생성 함수
func NewUserStateService(
	userStateRepo user_state.Repository,
	clientManager *websocket.ClientManager,
) *UserStateService {
	return &UserStateService{
		userStateRepo: userStateRepo,
		clientManager: clientManager,
	}
}

// RegisterUser 사용자 등록
func (s *UserStateService) RegisterUser(userID string) *user_state.UserState {
	state := user_state.NewUserState(userID)
	s.userStateRepo.Save(state)
	return state
}

// GetUserState 사용자 상태 조회
func (s *UserStateService) GetUserState(userID string) *user_state.UserState {
	state, exists := s.userStateRepo.FindByID(userID)
	if !exists {
		return nil
	}
	return state
}

// SetUserOnline 사용자 온라인 상태 설정
func (s *UserStateService) SetUserOnline(userID string, online bool) *user_state.UserState {
	state := s.GetUserState(userID)
	state.SetOnline(online)
	s.userStateRepo.Save(state)
	return state
}

// UpdateUserState 사용자 상태 업데이트
func (s *UserStateService) UpdateUserState(userID string, online bool, stopwatchRunning bool, elapsedTime float64) *user_state.UserState {
	state, exists := s.userStateRepo.FindByID(userID)
	if !exists {
		state = s.RegisterUser(userID)
	}

	state.SetOnline(online)
	state.UpdateStopwatch(stopwatchRunning, elapsedTime)

	s.userStateRepo.Save(state)
	return state
}

// DeleteUser 사용자 삭제
func (s *UserStateService) DeleteUser(userID string) {
	s.userStateRepo.Delete(userID)
}

// BroadcastUserState 사용자 상태 브로드캐스트
func (s *UserStateService) BroadcastUserState(state *user_state.UserState) {
	msg := dto.FromUserState(state)
	log.Printf("Broadcasting state for user: %s, online: %v, running: %v, time: %v",
		state.UserID, state.Online, state.StopwatchRunning, state.ElapsedTime)
	s.clientManager.Broadcast(msg)
}

// GetUserStateList 사용자 상태 목록 조회
func (s *UserStateService) GetUserStateList() []*user_state.UserState {
	return s.userStateRepo.FindAll()
}
