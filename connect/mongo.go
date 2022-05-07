package connect

import (
    "context"
	"fmt"
	"strings"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
	"github.com/spf13/viper"
)
var (
	MongoClient string 
	mongoUri = []string{
		"mongodb://",
		viper.GetString("mongodb.server.user"),":",
		viper.GetString("mongodb.server.pass"),"@",
		viper.GetString("mongodb.server.host"),
		"/?maxPollSize=20&w=majority"}
)


func init() {
	fmt.Println(mongoUri)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(strings.Join(mongoUri,"")))
	if err != nil {
		panic(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
}
