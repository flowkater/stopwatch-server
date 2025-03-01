package services

import (
	"github.com/flowkater/stopwatch-server/internal/domain/message"
	"github.com/flowkater/stopwatch-server/internal/domain/user"
	"github.com/flowkater/stopwatch-server/internal/infrastructure/websocket"
)

// UserService 사용자 서비스
type UserService struct {
	userRepo      user.Repository
	clientManager *websocket.ClientManager
}

// NewUserService 사용자 서비스 생성 함수
func NewUserService(userRepo user.Repository, clientManager *websocket.ClientManager) *UserService {
	return &UserService{
		userRepo:      userRepo,
		clientManager: clientManager,
	}
}

// RegisterUser 사용자 등록
func (s *UserService) RegisterUser(userID string) *user.User {
	user := user.NewUser(userID)
	s.userRepo.Save(user)
	return user
}

// SetUserOnline 사용자 온라인 상태 설정
func (s *UserService) SetUserOnline(userID string, online bool) {
	user, exists := s.userRepo.FindByID(userID)
	if !exists {
		user = s.RegisterUser(userID)
	}

	user.SetOnline(online)
	s.userRepo.Save(user)
}

// DeleteUser 사용자 삭제
func (s *UserService) DeleteUser(userID string) {
	s.userRepo.Delete(userID)
}

// BroadcastUserState 사용자 상태 브로드캐스트
func (s *UserService) BroadcastUserState(msg *message.Message) {
	s.clientManager.Broadcast(msg)
}
