package main

import (
	_ "tronbooth/routers"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/astaxie/beego"
)

const (
	retryDelay       = 60 * time.Second
	maxRetryAttempts = 3
)

func main() {
	log.Println("Env $PORT :", os.Getenv("PORT"))
	if os.Getenv("PORT") != "" {
		port, err := strconv.Atoi(os.Getenv("PORT"))
		if err != nil {
			log.Fatal(err)
			log.Fatal("$PORT must be set")
		}
		log.Println("port : ", port)
		beego.BConfig.Listen.HTTPPort = port
		beego.BConfig.Listen.HTTPSPort = port
	}
	if os.Getenv("BEEGO_ENV") != "" {
		log.Println("Env $BEEGO_ENV :", os.Getenv("BEEGO_ENV"))
		beego.BConfig.RunMode = os.Getenv("BEEGO_ENV")
	}

	// gotenv.Load()
	// db, err := postgres.NewPgDb(false)

	// if err != nil {
	// 	log.Fatalf("error in establishing database connection: %s", err.Error())
	// }
	// defer db.Close()

	// go startPullingProfitEventLog()

	beego.Run()
}
