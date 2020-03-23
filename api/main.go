package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ATechnoHazard/apollo-backend/internal/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"github.com/lib/pq"
	negronilogrus "github.com/meatballhat/negroni-logrus"
	log "github.com/sirupsen/logrus"
	"github.com/thoas/stats"
	"github.com/urfave/negroni"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	log.SetFormatter(&log.JSONFormatter{PrettyPrint: true})
	log.SetOutput(os.Stdout)
	if os.Getenv("DEBUG") == "true" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Running in a debug environment")
	} else {
		log.Printf("Running in a production environment")
	}
}

func connectDb() *gorm.DB {
	conn, err := pq.ParseURL(os.Getenv("DB_URI"))
	if err != nil {
		log.Fatal(err)
	}

	db, err := gorm.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}

	if os.Getenv("DEBUG") == "true" {
		db = db.Debug()
	}
	return db
}

func main() {
	// stats middleware
	s := stats.New()

	// http router
	router := httprouter.New()
	router.HandlerFunc("GET", "/stats", func(w http.ResponseWriter, r *http.Request) {
		data := s.Data()
		msg := utils.Message(http.StatusOK, "Statistics")
		msg["stats"] = data
		utils.Respond(w, msg)
	})

	// init negroni
	n := negroni.New()
	n.Use(negroni.NewRecovery())

	// log requests if debug
	if _, ok := os.LookupEnv("DEBUG"); ok {
		n.Use(negronilogrus.NewCustomMiddleware(log.DebugLevel, &log.JSONFormatter{PrettyPrint: true}, "API requests")) // request logger middleware
	}
	n.Use(s)

	n.UseHandler(router) // wrap router with middleware
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	log.WithField("event", "START").Info("Listening on port " + port)

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), n)
	if err != nil {
		log.Panic(err)
	}
}
