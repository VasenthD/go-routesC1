package main

import (
	"context"
	"fmt"
	"log"
	"mongodb1/config"
	"mongodb1/constants"
	"mongodb1/controller"
	"mongodb1/routes"
	services "mongodb1/transactionservice"

	//	"rest-api/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	mongoclient *mongo.Client
	ctx         context.Context
	server      *gin.Engine
)

func initRoutes() {
	routes.Default(server)
}

func initApp(mongoClient *mongo.Client) {
	ctx = context.TODO()
	// creating a collection

	profileCollection := mongoClient.Database(constants.DatabaseName).Collection("profiles")

	//pass the collection and context to instances the service

	profileService := services.InitProfileServiceInit(profileCollection, ctx)
	//pass the service to instantieate the controller
	profileController := controller.InitProfileController(*profileService)

	//pass the controller to route
	routes.ProfileRoute(server, profileController)
}

func main() {
	server = gin.Default()
	mongoclient, err := config.ConnectDataBase()
	defer mongoclient.Disconnect(ctx)
	if err != nil {
		panic(err)
	}
	initRoutes()
	initApp(mongoclient)
	fmt.Println("server running on port", constants.Port)
	log.Fatal(server.Run(constants.Port))
}

//miscr

/*






 */
