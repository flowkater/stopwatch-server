package repository

import (
	"sync"

	"github.com/flowkater/stopwatch-server/internal/domain/user_state"
)

// UserStateMemoryRepository 사용자 상태 인메모리 레포지토리
type UserStateMemoryRepository struct {
	userStates map[string]*user_state.UserState
	mu         sync.RWMutex
}

// NewUserStateMemoryRepository 사용자 상태 인메모리 레포지토리 생성 함수
func NewUserStateMemoryRepository() *UserStateMemoryRepository {
	return &UserStateMemoryRepository{
		userStates: make(map[string]*user_state.UserState),
	}
}

// Save 사용자 상태 저장
func (r *UserStateMemoryRepository) Save(userState *user_state.UserState) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.userStates[userState.UserID] = userState
	return nil
}

// FindByID ID로 사용자 상태 조회
func (r *UserStateMemoryRepository) FindByID(id string) (*user_state.UserState, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	u, ok := r.userStates[id]
	return u, ok
}

// Delete 사용자 상태 삭제
func (r *UserStateMemoryRepository) Delete(id string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.userStates, id)
}

// FindAll 모든 사용자 상태 조회
func (r *UserStateMemoryRepository) FindAll() []*user_state.UserState {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]*user_state.UserState, 0, len(r.userStates))
	for _, u := range r.userStates {
		result = append(result, u)
	}
	return result
}
