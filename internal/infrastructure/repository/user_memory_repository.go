package repository

import (
	"sync"

	"github.com/flowkater/stopwatch-server/internal/domain/user"
)

// UserMemoryRepository 사용자 인메모리 레포지토리
type UserMemoryRepository struct {
	users map[string]*user.User
	mu    sync.RWMutex
}

// NewUserMemoryRepository 사용자 인메모리 레포지토리 생성 함수
func NewUserMemoryRepository() *UserMemoryRepository {
	return &UserMemoryRepository{
		users: make(map[string]*user.User),
	}
}

// Save 사용자 저장
func (r *UserMemoryRepository) Save(user *user.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.users[user.ID] = user
	return nil
}

// FindByID ID로 사용자 조회
func (r *UserMemoryRepository) FindByID(id string) (*user.User, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	u, ok := r.users[id]
	return u, ok
}

// Delete 사용자 삭제
func (r *UserMemoryRepository) Delete(id string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.users, id)
}

// FindAll 모든 사용자 조회
func (r *UserMemoryRepository) FindAll() []*user.User {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]*user.User, 0, len(r.users))
	for _, u := range r.users {
		result = append(result, u)
	}
	return result
}
