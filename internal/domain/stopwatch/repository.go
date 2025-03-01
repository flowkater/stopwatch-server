package stopwatch

// Repository 스탑워치 레포지토리 인터페이스
type Repository interface {
	Save(userID string, stopwatch *Stopwatch) error
	FindByUserID(userID string) (*Stopwatch, bool)
	Delete(userID string)
}
