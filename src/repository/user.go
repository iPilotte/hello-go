package repository

import (
	"helloworld/model"
	"log"

	"github.com/jmoiron/sqlx"
)

// UserRepoMYSQL ที่เกี่ยวกับดาต้าเบส
type UserRepoMYSQL struct {
	DBConnection *sqlx.DB
}

//UserRepo interface
type UserRepo interface {
	CreateUser(user model.User) error
}

//CreateUser repo
func (userRepo UserRepoMYSQL) CreateUser(user model.User) error {
	statement := `INSERT INTO user (username, password) VALUES (?,?)`
	tx := userRepo.DBConnection.MustBegin()
	tx.MustExec(statement, user.Username, user.Password)
	if err := tx.Commit(); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
