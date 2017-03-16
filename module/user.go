package module

import (
	"errors"
	// "fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	id int64
}

func (this *User) Get() (err error) {
	if err := this.Check(); err != nil {
		return err
	}
	return
}

func (this *User) Check() (err error) {
	if this.id < int64(1) {
		return errors.New(`not init`)
	}
	return
}

func UserCreate(name, password string) (id int64, err error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		err = errors.New(`can not hash password`)
		return
	}

	if len(name) < 1 {
		err = errors.New(`empty name`)
		return
	}

	if len(password) < 1 {
		err = errors.New(`empty password`)
		return
	}

	query := `INSERT IGNORE INTO user SET name = ?, ts_create = ?, password = ?`
	result, err := db.Exec(query, name, time.Now().Unix(), hashedPassword)
	if err != nil {
		err = errors.New(`db error`)
		return
	}

	var i int64
	if i, _ = result.RowsAffected(); i < 1 {
		err = errors.New(`duplicate name`)
		return
	}

	if i, _ = result.LastInsertId(); i < 1 {
		err = errors.New(`unknown name`)
		return
	}

	id = i

	return
}
