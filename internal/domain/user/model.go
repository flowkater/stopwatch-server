package user

// User 도메인 모델
type User struct {
	ID     string
	Online bool
}

// NewUser 사용자 생성 함수
func NewUser(id string) *User {
	return &User{
		ID:     id,
		Online: true,
	}
}

// SetOnline 사용자 온라인 상태 설정
func (u *User) SetOnline(online bool) {
	u.Online = online
}
