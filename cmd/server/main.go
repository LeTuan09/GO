package main

import (
	"GO-API-SERVER/internal/routers"
)

func main() {
	r := routers.NewRouter()
	r.Run(":14447") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
