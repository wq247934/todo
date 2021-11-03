package main

import (
	"fmt"
	"todo/routers"
)

func main() {
	fmt.Println("Hello World")
	app := routers.SetRouters()
	app.Run(":8080")

}
