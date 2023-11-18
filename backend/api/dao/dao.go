package dao




type DaoInterFace interface {
	GetUserCounites(userID int)
	AddCounitesForUser(counties []string, userId int)
	DeleteCountiesforUser(counties []string, userId int)
}


