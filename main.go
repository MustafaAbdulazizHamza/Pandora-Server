package main

import (
	"database/sql"
	"github.com/MustafaAbdulazizHamza/Pandora/APIs"
	"github.com/MustafaAbdulazizHamza/Pandora/Middleware"
	_ "github.com/mattn/go-sqlite3"

	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	var (
		address string = ":8080"
		db      *sql.DB
		err     error
	)
	if db, err = sql.Open("sqlite3", "./Pandora.db"); err != nil {
		os.Exit(404)
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		os.Exit(404)
	}
	router.Use(Middleware.AuthenticateUser(db))

	router.POST("/user", APIs.InsertUser(db))
	router.PATCH("/user", APIs.UpdateUserCredentials(db))
	router.DELETE("/user", APIs.DeleteUser(db))
	router.POST("/secret", APIs.PostSecret(db))
	router.GET("/secret", APIs.GetSecret(db))
	router.PATCH("/secret", APIs.UpdateSecret(db))
	router.DELETE("/secret", APIs.DeleteSecret(db))
	err = router.RunTLS(address, "server.crt", "server.key")
	log.Fatal(err)
}
