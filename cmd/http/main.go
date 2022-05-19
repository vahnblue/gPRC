// @title Example API
// @version 1.0
// @description PHAROS Example API
// @BasePath /example
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
package main

import (
	"go-skeleton-auth/internal/boot"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	if err := boot.HTTP(); err != nil {
		log.Println("[HTTP] failed to boot http server due to " + err.Error())
	}
}
