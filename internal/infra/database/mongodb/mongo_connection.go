package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnection(driver, dbName, user, password, host, port string) (*mongo.Database, error) {
	connString := fmt.Sprintf("%s://%s:%s@%s:%s", driver, user, password, host, port)
	dbConn, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connString))
	if err != nil {
		return nil, err
	}

	return dbConn.Database(dbName), nil
}
