package mongodb

import "time"

type demon struct {
	Name     string    `bson:"name"`
	Age      int       `bson:"age"`
	CreateAt time.Time `bson:"createAt"`
}
