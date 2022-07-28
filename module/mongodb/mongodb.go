package mongodb

// demon mongodb

import (
	"demon/components/mongo"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	mongo2 "go.mongodb.org/mongo-driver/mongo"
	"time"
)

var (
	col *mongo2.Collection
)

func init() {
	log.Info("demon mongodb module")
	col = mongo.Mdb.Database("demon").Collection("demon")
	insertOne()
	//findOne()
	//findAll()
}

// 插入
func insertOne() {
	res, err := col.InsertOne(
		mongo.Ctx,
		bson.D{
			{"item", "canvas"},
			{"qty", 100},
			{"tags", bson.A{"cotton"}},
			{"size", bson.D{
				{"h", 28},
				{"w", 35.5},
				{"uom", "cm"},
			}},
			{"time", time.Now().Add(8 * time.Hour)},
		})
	if err != nil {
		log.Fatalf("mongodb insertOne err : [%v]", err)
	}
	log.Infof("mongodb insertOne with docId = [%v]", res.InsertedID)
}

// 查询单条
func findOne() {
	var result bson.M
	err := col.FindOne(mongo.Ctx, bson.M{
		"name": "YuanTong",
	}).Decode(&result)
	if err != nil {
		log.Fatalf("mongodb findOne err : [%v]", err)
	}
	log.Infof("mongodb findOne with doc = [%v]", result)
	log.Infof("createAt is [%v]", result["createAt"])
}

// 查询多条
func findAll() {
	temp := make([]*demon, 0)
	cursor, err := col.Find(mongo.Ctx, bson.M{})
	if err != nil {
		log.Fatalf("mongodb findAll err : [%v]", err)
	}
	err = cursor.All(mongo.Ctx, &temp)
	if err != nil {
		log.Fatalf("mongodb findAll err : [%v]", err)
	}
	log.Infof("mongodb findAll with doc = [%v]", temp[0].CreateAt)
}
