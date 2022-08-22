package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/ghodss/yaml"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

var db *sql.DB

var server = env("SERVER")
var port = env("PORT")
var user = env("USER")
var password = env("PASSWORD")
var database = env("DATABASE")

var config Config

const (
	collector = "query_exporter"
)

func env(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	// load env file
	godotenv.Load()
	var err error
	var configFile, bind string
	// =====================
	// Get OS parameter
	// =====================
	flag.StringVar(&configFile, "config", "config.yml", "configuration file")
	flag.StringVar(&bind, "bind", "0.0.0.0:9104", "bind")
	flag.Parse()

	// =====================
	// Load config & yaml
	// =====================
	var b []byte
	if b, err = ioutil.ReadFile(configFile); err != nil {
		log.Errorf("Failed to read config file: %s", err)
		os.Exit(1)
	}

	// Load yaml
	if err := yaml.Unmarshal(b, &config); err != nil {
		log.Errorf("Failed to load config: %s", err)
		os.Exit(1)
	}

	for {
		queryMetrics(err)
		// Calling Sleep method
		log.Info("Sleeping for 10 seconds")
		time.Sleep(10 * time.Second)

	}

}

func queryMetrics(err error) {
	// Execute each queries in metrics
	port, err := strconv.Atoi(port)

	for name, metric := range config.Metrics {
		// Build connection string
		connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
			server, user, password, port, database)

		// Create connection pool
		db, err = sql.Open("sqlserver", connString)
		if err != nil {
			log.Fatal("Error creating connection pool: ", err.Error())
		}
		ctx := context.Background()
		err = db.PingContext(ctx)
		if err != nil {
			log.Fatal(err.Error())
		}

		fmt.Printf("Connected!\n")

		tsql := fmt.Sprintf(metric.Query)
		log.Info(" Name!!!! " + name + metric.Description)
		//tsql := fmt.Sprintf(metric.Query)

		// Execute query
		rows, err := db.QueryContext(ctx, tsql)
		if err != nil {
			log.Info("Catch me")
			log.Error(err)
		}

		defer rows.Close()

		var count int

		// Iterate through the result set.
		for rows.Next() {
			var tablename string
			var totalspacemb string
			var unusedspacemb string

			// Get values from row.
			err := rows.Scan(&tablename, &totalspacemb, &unusedspacemb)
			if err != nil {
				log.Error(err)
			}

			fmt.Printf("Table Name: %s,\t Total Space: %s,\t Unused Space: %s \n", tablename, totalspacemb, unusedspacemb)
			count++
		}
	}

}

// =============================
// Config config structure
// =============================
type Config struct {
	DSN     string
	Metrics map[string]struct {
		Query       string
		Type        string
		Description string
		Labels      []string
		Value       string
		metricDesc  string
	}
}
