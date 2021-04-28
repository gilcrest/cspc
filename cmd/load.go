package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/gilcrest/cspc/logger"

	"github.com/google/uuid"
	"github.com/peterbourgon/ff"

	"github.com/gilcrest/cspc"
	"github.com/gilcrest/cspc/app"
	"github.com/gilcrest/cspc/datastore"
	"github.com/gilcrest/cspc/datastore/countrystore"
	"github.com/gilcrest/cspc/datastore/countystore"
	"github.com/gilcrest/cspc/datastore/statestore"
	"github.com/gilcrest/errs"
	"github.com/rs/zerolog"
)

const (
	// exitFail is the exit code if the program
	// fails.
	exitFail = 1
	// log level environment variable name
	loglevelEnv string = "LOG_LEVEL"
	// minimum accepted log level environment variable name
	logLevelMinEnv string = "LOG_LEVEL_MIN"
	// log error stack environment variable name
	logErrorStackEnv string = "LOG_ERROR_STACK"
	// server port environment variable name
	portEnv string = "PORT"
	// database host environment variable name
	dbHostEnv string = "DB_HOST"
	// database port environment variable name
	dbPortEnv string = "DB_PORT"
	// database name environment variable name
	dbNameEnv string = "DB_NAME"
	// database user environment variable name
	dbUserEnv string = "DB_USER"
	// database user password environment variable name
	dbPasswordEnv string = "DB_PASSWORD"
)

// cliFlags are the command line flags parsed at startup
type cliFlags struct {
}

type countyInit struct {
	CountyCode string `json:"county_cd"`
	CountyName string `json:"county_name"`
}

type stateInit struct {
	Code             string `json:"state_prov_cd"`
	Name             string `json:"state_name"`
	FIPSCode         string `json:"state_fips_cd"`
	LatitudeAverage  string `json:"latitude_average"`
	LongitudeAverage string `json:"longitude_average"`
}

func main() {
	if err := run(os.Args); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitFail)
	}
}

func run(args []string) error {

	flgs, err := newFlags(args)
	if err != nil {
		return err
	}

	// determine minimum logging level based on flag input
	minlvl, err := zerolog.ParseLevel(flgs.logLvlMin)
	if err != nil {
		return err
	}

	// determine logging level based on flag input
	lvl, err := zerolog.ParseLevel(flgs.loglvl)
	if err != nil {
		return err
	}

	// setup logger with appropriate defaults
	lgr := logger.NewLogger(os.Stdout, minlvl, true)

	// logs will be written at the level set in NewLogger (which is
	// also the minimum level). If the logs are to be written at a
	// different level than the minimum, use SetGlobalLevel to set
	// the global logging level to that. Minimum rules will still
	// apply.
	if minlvl != lvl {
		zerolog.SetGlobalLevel(lvl)
	}

	// set global logging time field format to Unix timestamp
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	lgr.Info().Msgf("minimum accepted logging level set to %s", minlvl)
	lgr.Info().Msgf("logging level set to %s", lvl)

	// set global to log errors with stack (or not) based on flag
	logger.WriteErrorStackGlobal(flgs.logErrorStack)
	lgr.Info().Msgf("log error stack global set to %t", flgs.logErrorStack)

	// validate port in acceptable range
	err = portRange(flgs.port)
	if err != nil {
		lgr.Fatal().Err(err).Msg("portRange() error")
	}

	//get struct holding PostgreSQL datasource name details
	dsn := datastore.NewPGDatasourceName(flgs.dbhost, flgs.dbname, flgs.dbuser, flgs.dbpassword, flgs.dbport)

	// initialize a non-nil, empty context
	ctx := context.Background()

	logger := app.NewLogger(zerolog.DebugLevel)

	db, cleanup, err := datastore.NewDB(dsn, logger)
	if err != nil {
		fmt.Println(err)
	}
	defer cleanup()
	defaultDatastore := datastore.NewDefaultDatastore(db)

	a := app.NewApplication(app.Local, defaultDatastore, logger)

	if flgs.all {
		flgs.countries = true
		flgs.states = true
		flgs.counties = true
	}

	if flgs.countries {
		err := loadCountries(ctx, a)
		if err != nil {
			fmt.Println(err)
		}
	}

	if flgs.states {
		err := loadUSStates(ctx, a)
		if err != nil {
			fmt.Println(err)
		}
	}

	if flgs.counties {
		s, err := mapUSCounties2States(ctx, a)
		if err != nil {
			fmt.Println(err)
		}
		ctx := context.Background()
		err = loadCounties2db(ctx, a, s)
		if err != nil {
			fmt.Println(err)
		}
	}
	return nil
}

type flags struct {
	// log-level flag allows for setting logging level, e.g. to run the server
	// with level set to debug, it'd be: ./server -log-level=debug
	// If not set, defaults to error
	loglvl string

	// log-level-min flag sets the minimum accepted logging level
	// - e.g. in production, you may have a policy to never allow logs at
	// trace level. You could set the minimum log level to Debug. Even
	// if the Global log level is set to Trace, only logs at Debug
	// and above would be logged. Default level is trace.
	logLvlMin string

	// logErrorStack flag determines whether or not a full error stack
	// should be logged. If true, error stacks are logged, if false,
	// just the error is logged
	logErrorStack bool

	// port flag is what http.ListenAndServe will listen on. default is 8080 if not set
	port int

	// dbhost is the database host
	dbhost string

	// dbport is the database port
	dbport int

	// dbname is the database name
	dbname string

	// dbuser is the database user
	dbuser string

	// dbpassword is the database user's password
	dbpassword string

	countries bool
	states    bool
	counties  bool
	all       bool
}

// newFlags parses the command line flags using ff and returns
// a flags struct or an error
func newFlags(args []string) (flgs flags, err error) {
	// create new FlagSet using the program name being executed (args[0])
	// as the name of the FlagSet
	fs := flag.NewFlagSet(args[0], flag.ContinueOnError)

	var (
		logLvlMin     = fs.String("log-level-min", "trace", fmt.Sprintf("sets minimum log level (trace, debug, info, warn, error, fatal, panic, disabled), (also via %s)", logLevelMinEnv))
		loglvl        = fs.String("log-level", "info", fmt.Sprintf("sets log level (trace, debug, info, warn, error, fatal, panic, disabled), (also via %s)", loglevelEnv))
		logErrorStack = fs.Bool("log-error-stack", true, fmt.Sprintf("if true, log full error stacktrace, else just log error, (also via %s)", logErrorStackEnv))
		port          = fs.Int("port", 8080, fmt.Sprintf("listen port for server (also via %s)", portEnv))
		dbhost        = fs.String("db-host", "", fmt.Sprintf("postgresql database host (also via %s)", dbHostEnv))
		dbport        = fs.Int("db-port", 5432, fmt.Sprintf("postgresql database port (also via %s)", dbPortEnv))
		dbname        = fs.String("db-name", "", fmt.Sprintf("postgresql database name (also via %s)", dbNameEnv))
		dbuser        = fs.String("db-user", "", fmt.Sprintf("postgresql database user (also via %s)", dbUserEnv))
		dbpassword    = fs.String("db-password", "", fmt.Sprintf("postgresql database password (also via %s)", dbPasswordEnv))
		// countries flag allows for loading all countries into lookup.country_lkup
		// If not set, defaults to false and countries are not loaded
		countries = fs.Bool("countries", false, "loads all countries into lookup.country_lkup")
		// states flag allows for loading all states into lookup.state_prov_lkup
		// If not set, defaults to false and states are not loaded
		states = fs.Bool("states", false, "loads all states into lookup.state_prov_lkup")
		// counties flag allows for loading all counties into lookup.county_lkup
		// If not set, defaults to false and counties are not loaded
		counties = fs.Bool("counties", false, "loads all counties into lookup.county_lkup")
		// all flag allows for loading all countries, states and counties
		// If not set, defaults to false and individual flags are considered
		all = fs.Bool("all", false, "all flag allows for loading all countries, states and counties. If not set, individual flags are considered")
	)

	// Parse the command line flags from above
	err = ff.Parse(fs, args[1:], ff.WithEnvVarNoPrefix())
	if err != nil {
		return flgs, err
	}

	return flags{
		loglvl:        *loglvl,
		logLvlMin:     *logLvlMin,
		logErrorStack: *logErrorStack,
		port:          *port,
		dbhost:        *dbhost,
		dbport:        *dbport,
		dbname:        *dbname,
		dbuser:        *dbuser,
		dbpassword:    *dbpassword,
		countries:     *countries,
		states:        *states,
		counties:      *counties,
		all:           *all,
	}, nil
}

// createNameJSON was used to create the CountryNamesJSON constant
func createNameJSON() ([]string, error) {
	const op errs.Op = "main/createNameJSON"

	var (
		countries []cspc.Country
		names     []string
	)

	err := json.Unmarshal([]byte(cspc.CountryFullJSON), &countries)
	if err != nil {
		return nil, errs.E(op, err)
	}

	for _, c := range countries {
		names = append(names, c.Name)
	}

	return names, nil
}

// createAlpha2JSON was used to create the CountryAlpha2CodeJSON constant
func createAlpha2JSON() ([]string, error) {
	const op errs.Op = "main/createNameJSON"

	var (
		countries []cspc.Country
		alpha2s   []string
	)

	err := json.Unmarshal([]byte(cspc.CountryFullJSON), &countries)
	if err != nil {
		return nil, errs.E(op, err)
	}

	for _, c := range countries {
		alpha2s = append(alpha2s, c.Alpha2Code)
	}

	return alpha2s, nil
}

func loadCountries(ctx context.Context, a *app.Application) error {
	const op errs.Op = "main/loadCountries"

	var countries []*cspc.Country

	tx, err := a.Datastorer.BeginTx(ctx)
	if err != nil {
		return errs.E(op, err)
	}

	// declare variable as the Transactor interface
	var transactor countrystore.Transactor

	transactor = countrystore.NewDefaultTransactor(a.Datastorer)

	err = json.Unmarshal([]byte(cspc.CountryFullJSON), &countries)
	if err != nil {
		return errs.E(op, err)
	}

	for _, c := range countries {
		fmt.Printf("Country Name = %s, Alpha 2 Code = %s\n", c.Name, c.Alpha2Code)
		now := time.Now()
		c.ID = uuid.New()
		c.CreateUsername = "gilcrest"
		c.CreateTimestamp = now
		c.UpdateUsername = "gilcrest"
		c.UpdateTimestamp = now
		err = transactor.CreateCountry(ctx, c, tx)
		if err != nil {
			return errs.E(op, a.Datastorer.RollbackTx(tx, err))
		}
	}

	// Commit the Transaction
	if err = a.Datastorer.CommitTx(tx); err != nil {
		return errs.E(op, err)
	}

	return nil
}

// one-time use to get JSON strings - not needed going forward
func countryInit() {
	names, err := createNameJSON()
	if err != nil {
		fmt.Println(err)
	}

	n, err := json.Marshal(names)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(n))

	a2j, err := createAlpha2JSON()
	if err != nil {
		fmt.Println(err)
	}
	a2, err := json.Marshal(a2j)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(a2))
}

// createUSStatesNameJSON was used to create the USStatesNameJSON constant
func createUSStatesNameJSON() ([]string, error) {
	const op errs.Op = "main/createStateProvinceNameJSON"

	var (
		states []cspc.StateProvince
		names  []string
	)

	err := json.Unmarshal([]byte(cspc.USStatesJSON), &states)
	if err != nil {
		return nil, errs.E(op, err)
	}

	for _, s := range states {
		names = append(names, s.Name)
	}

	return names, nil
}

// createUSStatesCodeJSON was used to create the USStatesCodeJSON constant
func createUSStatesCodeJSON() ([]string, error) {
	const op errs.Op = "main/createUSStatesCodeJSON"

	var (
		states []cspc.StateProvince
		codes  []string
	)

	err := json.Unmarshal([]byte(cspc.USStatesJSON), &states)
	if err != nil {
		return nil, errs.E(op, err)
	}

	for _, s := range states {
		codes = append(codes, s.Code)
	}

	return codes, nil
}

func loadUSStates(ctx context.Context, a *app.Application) error {
	const op errs.Op = "main/loadStates"

	var states []cspc.StateProvince

	tx, err := a.Datastorer.BeginTx(ctx)
	if err != nil {
		return errs.E(op, err)
	}

	var (
		countrySelector countrystore.Selector
		transactor      statestore.Transactor
	)

	countrySelector = countrystore.NewDefaultSelector(a.Datastorer)

	transactor = statestore.NewDefaultTransactor(a.Datastorer)

	err = json.Unmarshal([]byte(cspc.USStatesJSON), &states)
	if err != nil {
		return errs.E(op, err)
	}

	for _, s := range states {
		fmt.Printf("State Name = %s, Alpha 2 Code = %s\n", s.Name, s.Code)

		us, err := countrySelector.FindByAlpha2Code(ctx, "US")
		if err != nil {
			return errs.E(op, a.Datastorer.RollbackTx(tx, err))
		}

		s.ID = uuid.New()
		now := time.Now()
		s.CreateUsername = "gilcrest"
		s.CreateTimestamp = now
		s.UpdateUsername = "gilcrest"
		s.UpdateTimestamp = now

		err = transactor.CreateStateProvince(ctx, us, s, tx)
		if err != nil {
			return errs.E(op, a.Datastorer.RollbackTx(tx, err))
		}
	}

	// Commit the Transaction
	if err = a.Datastorer.CommitTx(tx); err != nil {
		return errs.E(op, err)
	}

	return nil
}

// one-time use - not needed going forward
func statesInit() {
	names, err := createUSStatesNameJSON()
	if err != nil {
		fmt.Println(err)
	}

	n, err := json.Marshal(names)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(n))

	usj, err := createUSStatesCodeJSON()
	if err != nil {
		fmt.Println(err)
	}
	u, err := json.Marshal(usj)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(u))
}

func mapUSCounties2States(ctx context.Context, a *app.Application) ([]cspc.StateProvince, error) {
	const op errs.Op = "main/mapUSCounties2States"

	var (
		ci              []countyInit
		si              []stateInit
		states          []*cspc.StateProvince
		finalStates     []cspc.StateProvince
		countrySelector countrystore.Selector
		stateSelector   statestore.Selector
	)

	err := json.Unmarshal([]byte(cspc.USCountyJSON), &ci)
	if err != nil {
		return nil, errs.E(op, err)
	}

	err = json.Unmarshal([]byte(cspc.USStatesJSON), &si)
	if err != nil {
		return nil, errs.E(op, err)
	}

	countrySelector = countrystore.NewDefaultSelector(a.Datastorer)
	stateSelector = statestore.NewDefaultSelector(a.Datastorer)

	us, err := countrySelector.FindByAlpha2Code(ctx, "US")
	if err != nil {
		return nil, errs.E(op, err)
	}

	// Initialize a true StateProvince struce for each state
	// initialized from JSON
	for _, s := range si {
		ts, err := stateSelector.FindByStateProvCode(ctx, us, s.Code)
		if err != nil {
			return nil, errs.E(op, err)
		}
		states = append(states, &ts)
	}

	for _, state := range states {
		for _, ic := range ci {
			statefip := ic.CountyCode[:2]

			if state.FIPSCode == statefip {
				now := time.Now()
				c := cspc.County{
					ID:               uuid.New(),
					Code:             ic.CountyCode,
					Name:             ic.CountyName,
					LatitudeAverage:  "",
					LongitudeAverage: "",
					CreateUsername:   "gilcrest",
					CreateTimestamp:  now,
					UpdateUsername:   "gilcrest",
					UpdateTimestamp:  now,
				}

				state.Counties = append(state.Counties, c)
			}
		}
	}

	for _, s := range states {
		fmt.Printf("%s has %d counties\n", s.Name, len(s.Counties))
		finalStates = append(finalStates, *s)
	}

	return finalStates, nil
}

func loadCounties2db(ctx context.Context, a *app.Application, states []cspc.StateProvince) error {
	const op errs.Op = "main/loadCounties2db"

	tx, err := a.Datastorer.BeginTx(ctx)
	if err != nil {
		return errs.E(op, err)
	}

	var transactor countystore.Transactor
	transactor = countystore.NewDefaultTransactor(a.Datastorer)

	for _, state := range states {
		for _, c := range state.Counties {
			err = transactor.CreateCounty(ctx, state, c, tx)
			if err != nil {
				return errs.E(op, a.Datastorer.RollbackTx(tx, err))
			}
		}
	}

	// Commit the Transaction
	if err = a.Datastorer.CommitTx(tx); err != nil {
		return errs.E(op, err)
	}

	return nil
}

// portRange validates the port be in an acceptable range
func portRange(port int) error {
	if port < 0 || port > 65535 {
		return errs.E(errors.New(fmt.Sprintf("port %d is not within valid port range (0 to 65535)", port)))
	}
	return nil
}
