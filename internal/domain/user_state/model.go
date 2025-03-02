package user_state

// UserState 사용자 상태 도메인 모델
type UserState struct {
	UserID           string
	Online           bool
	StopwatchRunning bool
	ElapsedTime      float64
}

// NewUserState 사용자 상태 생성 함수
func NewUserState(userID string) *UserState {
	return &UserState{
		UserID:           userID,
		Online:           true,
		StopwatchRunning: false,
		ElapsedTime:      0,
	}
}

// SetOnline 사용자 온라인 상태 설정
func (u *UserState) SetOnline(online bool) {
	u.Online = online
}

// StartStopwatch 스탑워치 시작
func (u *UserState) StartStopwatch() {
	u.StopwatchRunning = true
}

// StopStopwatch 스탑워치 정지
func (u *UserState) StopStopwatch() {
	u.StopwatchRunning = false
}

// UpdateElapsedTime 경과 시간 업데이트
func (u *UserState) UpdateElapsedTime(time float64) {
	u.ElapsedTime = time
}

// UpdateStopwatch 스탑워치 상태 업데이트
func (u *UserState) UpdateStopwatch(running bool, elapsedTime float64) {
	if running {
		u.StartStopwatch()
	} else {
		u.StopStopwatch()
	}
	u.UpdateElapsedTime(elapsedTime)
}
