package main

import (
	"context"
	"fmt"

	_mongoDriver "disspace/drivers/mongoDB"
)

func main() {
	var ctx = context.Background()
	db, err := _mongoDriver.ConnectDB(ctx)
	fmt.Println(db, err)

}
