package main

import (
	"github.com/bount-ing/bount.ing/api/routes"
)

func main() {
	r := routes.SetupRouter()
	r.Run() // listen and serve on 0.0.0.0:8080
}
