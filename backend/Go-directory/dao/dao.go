package dao




type DaoInterFace interface {
	GetUserCounites(userID int)
	AddCounitesForUser(counties []int, userId int)
	DeleteCountiesforUser(counties []int, userId int)
}