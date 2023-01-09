package user

import (
	"database/sql"
	log "github.com/sirupsen/logrus"
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
	Age uint8
}

func (u *User) Save(db *gorm.DB) error {
	res := db.Create(u)
	return res.Error
}

func (p *Params) Find(db *gorm.DB) (*User, error) {
	var (
		u   = &User{}
		err error
	)
	err = db.Find(u).Error
	if err != nil {
		log.Error(err)
		return u, err
	}
	return u, err
}
