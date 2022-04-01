package mongo

import (
	"context"
	"fmt"
	"log"
	"test3/hariprathap-hp/system_design/tinyURL/utils/errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	URI = "mongodb://localhost:27017"
)

var (
	clientOptions *options.ClientOptions
	client        *mongo.Client
	dbConnerr     error
	Collection    *mongo.Collection
)

const (
	DbName   = "oauth"
	CollName = "oauth_keys"
)

func init() {
	clientOptions = options.Client().ApplyURI(URI)
	clientOptions.SetConnectTimeout(time.Second * 10)
	client, dbConnerr = mongo.Connect(context.TODO(), clientOptions)
	if dbConnerr != nil {
		fmt.Println("Connection to DB Failed")
		//fmt.Fprintln(w, "Connect to DB Failed", dbConnerr)
	}

	err := client.Ping(context.TODO(), nil)

	if err != nil {
		fmt.Println("The ping error")
		log.Fatal(err)
	}
}

func GetClient() *mongo.Client {
	return client
}

func ListDBS() *errors.RestErr {
	fmt.Println("Listing DBS")
	dbs, err := client.ListDatabaseNames(context.TODO(), bson.M{})
	if err != nil {
		return errors.NewInternalServerError("Error while listing databases")
	}
	fmt.Println("Printing DBS", dbs)
	return nil
}

func ListColls() *errors.RestErr {
	fmt.Println("Listing Collections")
	//colls, err := client.Database(dbName).ListCollectionNames(context.TODO(), bson.M{})
	Collection := client.Database(DbName).Collection(CollName)
	var at bson.M
	if err := Collection.FindOne(context.TODO(), bson.M{}).Decode(&at); err != nil {
		log.Fatal(err)
	}
	fmt.Println(at)
	return nil
}
