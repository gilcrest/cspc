package datastore

import (
	"database/sql"
	"errors"

	"github.com/gilcrest/cspc/errs"
	"github.com/rs/zerolog"
)

// NewDB returns an open database handle of 0 or more underlying PostgreSQL connections
func NewDB(dsn PGDatasourceName, logger zerolog.Logger) (*sql.DB, func(), error) {

	f := func() {}

	// Open the postgres database using the postgres driver (pq)
	// func Open(driverName, dataSourceName string) (*DB, error)
	db, err := sql.Open("postgres", dsn.String())
	if err != nil {
		return nil, f, errs.E(errs.Database, err)
	}

	logger.Info().Msgf("sql database opened for %s on port %d", dsn.Host, dsn.Port)

	err = validateDB(db, logger)
	if err != nil {
		return nil, f, err
	}

	return db, func() { db.Close() }, nil
}

// validateDB pings the database and logs the current user and database
func validateDB(db *sql.DB, log zerolog.Logger) error {
	err := db.Ping()
	if err != nil {
		return errs.E(err)
	}
	log.Info().Msg("sql database Ping returned successfully")

	var (
		currentDatabase string
		currentUser     string
		dbVersion       string
	)
	sqlStatement := `select current_database(), current_user, version();`
	row := db.QueryRow(sqlStatement)
	switch err := row.Scan(&currentDatabase, &currentUser, &dbVersion); err {
	case sql.ErrNoRows:
		return errs.E(errors.New("no rows were returned"))
	case nil:
		log.Info().Msgf("database version: %s", dbVersion)
		log.Info().Msgf("current database user: %s", currentUser)
		log.Info().Msgf("current database: %s", currentDatabase)
	default:
		return errs.E(err)
	}
	return nil
}
