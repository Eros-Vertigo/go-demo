package user

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

// User user model
type User struct {
	gorm.Model
	Name         string
	Email        *string
	Age          uint8
	Birthday     time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
}
type Params struct {
	Age  uint8
	Name string
}

func (u *User) Save(db *gorm.DB) error {
	res := db.Create(u)
	return res.Error
}

func (p *Params) Find(db *gorm.DB) (User, error) {
	var (
		u   User
		err error
	)
	db.Where("name = ?", p.Name).First(&u)
	return u, err
}
