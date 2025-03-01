package repository

import (
	"sync"

	"github.com/flowkater/stopwatch-server/internal/domain/stopwatch"
)

// StopwatchMemoryRepository 스탑워치 인메모리 레포지토리
type StopwatchMemoryRepository struct {
	stopwatches map[string]*stopwatch.Stopwatch
	mu          sync.RWMutex
}

// NewStopwatchMemoryRepository 스탑워치 인메모리 레포지토리 생성 함수
func NewStopwatchMemoryRepository() *StopwatchMemoryRepository {
	return &StopwatchMemoryRepository{
		stopwatches: make(map[string]*stopwatch.Stopwatch),
	}
}

// Save 스탑워치 저장
func (r *StopwatchMemoryRepository) Save(userID string, sw *stopwatch.Stopwatch) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.stopwatches[userID] = sw
	return nil
}

// FindByUserID 사용자 ID로 스탑워치 조회
func (r *StopwatchMemoryRepository) FindByUserID(userID string) (*stopwatch.Stopwatch, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	sw, ok := r.stopwatches[userID]
	return sw, ok
}

// Delete 스탑워치 삭제
func (r *StopwatchMemoryRepository) Delete(userID string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.stopwatches, userID)
}
