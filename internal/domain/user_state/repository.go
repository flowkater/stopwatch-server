package user_state

// Repository 사용자 상태 레포지토리 인터페이스
type Repository interface {
	Save(userState *UserState) error
	FindByID(id string) (*UserState, bool)
	Delete(id string)
	FindAll() []*UserState
}
