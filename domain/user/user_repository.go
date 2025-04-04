package user

type Repository interface {
	GetIdByToken(token string) int64
}
