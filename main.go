package main

import (
	"context"
	"fmt"

	mongoDriver "capstone-project-API/drivers/mongoDB"
)

func main() {
	var ctx = context.Background()
	db, err := mongoDriver.Connect_to_db(ctx)
	fmt.Println(db, err)

}
