
# Dynatrace MYSQL Exporter

Exporter to query a database based on a frequency to obtain metric values for Dynatrace.


## Features

- Query SQL Server
- YAML Syntax to build queries and objects

## Installation

Install Go

Clone Repo

Create a .env file see the example below:

## .env file
```
    SERVER="database-1.asdf.us-east-1.rds.amazonaws.com"
    PORT=1433
    USER="sa"
    PASSWORD="..."
    DATABASE="xyz"
```

```run go main.go```
## Library Setup
Run the following Go commands to install the libraries
```
 go get github.com/prometheus/client_golang/prometheus
 go get github.com/prometheus/client_golang/prometheus/promhttp
 go get github.com/prometheus/common/version
 go get github.com/sirupsen/logrus
 go get github.com/ghodss/yaml
 go get github.com/go-sql-driver/mysql
 go get github.com/denisenkom/go-mssqldb
 go get github.com/joho/godotenv
 go mod vendor
 
 go run .

```


## Roadmap

- Dockerize the go installation files

- Build Better Ingestion features

- Redo configuration structure

- Add Error Handling

- Add Dynatrace API Push Feature

- Add Logging for Updates

- Considering Prometheus Agent to push/pull data
## Authors

- [@Aetrius](https://www.github.com/Aetrius)

