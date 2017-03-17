package module

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id int64
}

func (this *User) Get() (err error) {
	if err := this.Check(); err != nil {
		return err
	}
	return
}

func (this *User) Check() (err error) {
	if this.Id < 1 {
		return errors.New(`not init`)
	}
	return
}

func checkLogin(name, password string) error {
	if len(name) < 1 {
		return errors.New(`empty name`)
	}

	if len(password) < 1 {
		return errors.New(`empty password`)
	}
	return nil
}

func UserCreate(name, password string) (user *User, err error) {

	err = checkLogin(name, password)
	if err != nil {
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		err = errors.New(`can not hash password`)
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

		var errLogin error
		user, errLogin = UserLogin(name, password)
		if errLogin == nil {
			err = nil
		}
		return
	}

	if i, _ = result.LastInsertId(); i < 1 {
		err = errors.New(`unknown name`)
		return
	}

	user = &User{
		Id: i,
	}

	return
}

func UserLogin(name, password string) (user *User, err error) {
	err = checkLogin(name, password)
	if err != nil {
		return
	}

	query := `SELECT id, password FROM user WHERE name = ?`
	row, err := db.Query(query, name)
	if err != nil {
		err = errors.New(`db error`)
		return
	}

	if !row.Next() {
		err = errors.New(`no user`)
		return
	}

	var (
		id          int64
		passwordGet string
	)

	if err = row.Scan(&id, &passwordGet); err != nil {
		err = errors.New(`db error`)
		return
	}

	// fmt.Println(`password`, password, passwordGet)

	err = bcrypt.CompareHashAndPassword([]byte(passwordGet), []byte(password))
	if err != nil {
		err = errors.New(`password error`)
		return
	}

	user = &User{
		Id: id,
	}
	return
}
