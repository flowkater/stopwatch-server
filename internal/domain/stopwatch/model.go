package stopwatch

// Stopwatch 도메인 모델
type Stopwatch struct {
	Running     bool
	ElapsedTime float64
}

// NewStopwatch 스탑워치 생성 함수
func NewStopwatch() *Stopwatch {
	return &Stopwatch{
		Running:     false,
		ElapsedTime: 0,
	}
}

// Start 스탑워치 시작
func (s *Stopwatch) Start() {
	s.Running = true
}

// Stop 스탑워치 정지
func (s *Stopwatch) Stop() {
	s.Running = false
}

// UpdateElapsedTime 경과 시간 업데이트
func (s *Stopwatch) UpdateElapsedTime(time float64) {
	s.ElapsedTime = time
}
