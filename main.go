package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/hootsuite/healthchecks"
	"github.com/hootsuite/healthchecks/checks/sqlsc"
	_ "github.com/lib/pq"
)

func main() {

	DBname := os.Getenv("DB_NAME")
	DBuser := os.Getenv("DB_USER")
	DBpass := os.Getenv("DB_PASS")
	DBhost := os.Getenv("DB_HOST")
	DBport := os.Getenv("DB_PORT")

	// Create a db object by opening the connection
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", DBuser, DBpass, DBhost, DBport, DBname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// Define a StatusEndpoint at '/status/db' for a database dependency
	dbs := healthchecks.StatusEndpoint{
		Name:          "Postgres DB",
		Slug:          "db",
		Type:          "internal",
		IsTraversable: false,
		StatusCheck: sqlsc.SQLDBStatusChecker{
			DB: db,
		},
		TraverseCheck: nil,
	}

	// Add all your StatusEndpoints to a slice for your service.
	// This will be used to initialize the framework in the next step.
	statusEndpoints := []healthchecks.StatusEndpoint{dbs}

	// aboutFilePath = path for the about file for metadata about service.
	// Ussually hosted in project for developers / ops to edit.
	// Sample at https://github.com/hootsuite/healthchecks/blob/master/test/about.json
	aboutFilePath := "conf/about.json"

	// versionFilePath = path for the version file for current version of running service.
	// Ussually created at build / deploy time.
	// Sample at https://github.com/hootsuite/healthchecks/blob/master/test/version.txt
	versionFilePath := "conf/version.txt"

	// OPTIONAL - Set up any service injected customData for /status/about response.
	// Values can be any valid JSON conversion and will override values set in about.json.
	customData := make(map[string]interface{})

	// Register all "/status/..." requests to use our health checking framework
	http.Handle("/status/", healthchecks.Handler(statusEndpoints, aboutFilePath, versionFilePath, customData))
	log.Fatal(http.ListenAndServe(":8000", nil))
}
