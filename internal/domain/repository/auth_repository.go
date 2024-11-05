package repository

type AuthRepository interface {
	GenerateToken(userID uint, role string) (string, error)
}
