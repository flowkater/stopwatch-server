package user

// Repository 사용자 레포지토리 인터페이스
type Repository interface {
	Save(user *User) error
	FindByID(id string) (*User, bool)
	Delete(id string)
	FindAll() []*User
}
