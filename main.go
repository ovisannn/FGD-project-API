package main

import (
	"context"
	"log"
	"time"

	_routes "disspace/app/routes"
	_mongoDriver "disspace/drivers/mongoDB"

	_threadUseCase "disspace/business/threads"
	_threadController "disspace/controllers/threads"
	_threadRepository "disspace/drivers/databases/threads"

	"github.com/labstack/echo/v4"
	// "go.mongodb.org/mongo-driver/bson"
)

func main() {
	// Start Try
	var ctx = context.Background()
	db, _ := _mongoDriver.ConnectToDB(ctx)
	// if err != nil {
	// 	panic(err)
	// }

	// cursor, err := db.Collection("threads").Find(ctx, bson.D{})

	// var results []bson.M
	// if err = cursor.All(ctx, &results); err != nil {
	// 	panic(err)
	// }
	// for _, result := range results {
	// 	output, err := json.MarshalIndent(result, "", "    ")
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Printf("%s\n", output)
	// }
	// End Try

	e := echo.New()
	timeoutContext := time.Duration(20 * time.Second)

	threadRepository := _threadRepository.NewMongoDBThreadRepository(db)
	threadUseCase := _threadUseCase.NewThreadUseCase(threadRepository, timeoutContext)
	threadController := _threadController.NewThreadController(threadUseCase)

	routesInit := _routes.ControllerList{
		ThreadController: *threadController,
	}

	routesInit.RouteRegister(e)
	log.Fatal(e.Start(":8080"))
}
