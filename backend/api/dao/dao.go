package dao

//"go.mongodb.org/mongo-driver/mongo"
//"context"

/********************************************************************************
* Description: dao.go is an interface that contains all database functions to   *
* be implemented by daoPostgres                                                 *
*********************************************************************************/
type DaoInterFace interface {
	GetUserCounites(userID string)
	AddCounitesForUser(counties []string, userId int)
	DeleteCountiesforUser(counties []string, userId int)
}
