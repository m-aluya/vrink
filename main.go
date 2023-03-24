package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	var router = gin.Default()
	var address = ":3000"

	log.Fatal(router.Run(address))
}
