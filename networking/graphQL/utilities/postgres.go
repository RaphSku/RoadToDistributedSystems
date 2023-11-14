package utilities

import (
	"os"
	"strconv"

	"github.com/jackc/pgx"
	"github.com/labstack/gommon/log"
)

// A connection to the database is being established and returned
func GetPSQLConnection(database string) *pgx.Conn {
	// Get port number from an environment variable
	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Error("You have to provide a valid port number as an environment variable! Error happened due to:", err)
		os.Exit(1)
	}
	dbPortNumber := uint16(dbPort)

	// Create a Postgres config
	psqlConfig := pgx.ConnConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     dbPortNumber,
		Database: database,
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PW"),
	}

	// Establish a connection with the Postgres database
	conn, err := pgx.Connect(psqlConfig)
	if err != nil {
		log.Error("Unable to connect to database due to:", err)
		os.Exit(1)
	}

	return conn
}
