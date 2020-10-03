package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"

	"github.com/gilcrest/cspc"
	"github.com/gilcrest/cspc/app"
	"github.com/gilcrest/cspc/datastore"
	"github.com/gilcrest/cspc/datastore/countrystore"
	"github.com/gilcrest/cspc/datastore/countystore"
	"github.com/gilcrest/cspc/datastore/statestore"
	"github.com/gilcrest/errs"
	"github.com/rs/zerolog"
)

// cliFlags are the command line flags parsed at startup
type cliFlags struct {
	countries bool
	states    bool
	counties  bool
	all       bool
}

type countyInit struct {
	CountyCode string `json:"county_cd"`
	CountyName string `json:"county_name"`
}

func main() {
	// Initialize cliFlags and return a pointer to it
	cf := new(cliFlags)

	ctx := context.Background()

	// countries flag allows for loading all countries into lookup.country_lkup
	// If not set, defaults to false and countries are not loaded
	flag.BoolVar(&cf.countries, "countries", false, "loads all countries into lookup.country_lkup")

	// states flag allows for loading all states into lookup.state_prov_lkup
	// If not set, defaults to false and states are not loaded
	flag.BoolVar(&cf.states, "states", false, "loads all states into lookup.state_prov_lkup")

	// counties flag allows for loading all counties into lookup.county_lkup
	// If not set, defaults to false and counties are not loaded
	flag.BoolVar(&cf.counties, "counties", false, "loads all counties into lookup.county_lkup")

	// all flag allows for loading all countries, states and counties
	// If not set, defaults to false and individual flags are considered
	flag.BoolVar(&cf.counties, "counties", false, "all flag allows for loading all countries, states and counties. If not set, individual flags are considered")

	if cf.all {
		cf.countries = true
		cf.states = true
		cf.counties = true
	}

	if cf.countries {
		err := loadCountries(ctx)
		if err != nil {
			fmt.Println(err)
		}
	}

	if cf.states {
		err := loadUSStates(ctx)
		if err != nil {
			fmt.Println(err)
		}
	}

	if cf.counties {
		s, err := mapCounties2States()
		if err != nil {
			fmt.Println(err)
		}
		ctx := context.Background()
		err = loadCounties2db(ctx, s)
		if err != nil {
			fmt.Println(err)
		}
	}
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

func loadCountries(ctx context.Context) error {
	const op errs.Op = "main/loadCountries"

	var countries []cspc.Country

	logger := app.NewLogger(zerolog.DebugLevel)

	db, cleanup, err := datastore.NewDB(datastore.LocalDatastore, logger)
	if err != nil {
		return errs.E(op, err)
	}
	defer cleanup()

	// declare variable as the Transactor interface
	var datastorer datastore.Datastorer

	datastorer = datastore.NewDatastore(db)

	a := app.NewApplication(app.Local, datastorer, logger)

	tx, err := a.Datastorer.BeginTx(ctx)
	if err != nil {
		return errs.E(op, err)
	}

	// declare variable as the Transactor interface
	var transactor countrystore.Transactor

	transactor, err = countrystore.NewTx(tx)
	if err != nil {
		return errs.E(op, err)
	}

	err = json.Unmarshal([]byte(cspc.CountryFullJSON), &countries)
	if err != nil {
		return errs.E(op, err)
	}

	for _, c := range countries {
		fmt.Printf("Country Name = %s, Alpha 2 Code = %s\n", c.Name, c.Alpha2Code)
		err = transactor.CreateCountry(ctx, c)
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

func loadUSStates(ctx context.Context) error {
	const op errs.Op = "main/loadStates"

	var states []cspc.StateProvince

	logger := app.NewLogger(zerolog.DebugLevel)

	db, cleanup, err := datastore.NewDB(datastore.LocalDatastore, logger)
	if err != nil {
		return errs.E(op, err)
	}
	defer cleanup()

	// declare variable as the Transactor interface
	var datastorer datastore.Datastorer

	datastorer = datastore.NewDatastore(db)

	a := app.NewApplication(app.Local, datastorer, logger)

	tx, err := a.Datastorer.BeginTx(ctx)
	if err != nil {
		return errs.E(op, err)
	}

	// declare variable as the Transactor interface
	var transactor statestore.Transactor

	transactor, err = statestore.NewTx(tx)
	if err != nil {
		return errs.E(op, err)
	}

	err = json.Unmarshal([]byte(cspc.USStatesJSON), &states)
	if err != nil {
		return errs.E(op, err)
	}

	for _, s := range states {
		fmt.Printf("State Name = %s, Alpha 2 Code = %s\n", s.Name, s.Code)
		err = transactor.CreateStateProvince(ctx, "US", s)
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

func mapCounties2States() ([]*cspc.StateProvince, error) {
	const op errs.Op = "main/initLoadCounty"

	var (
		counties []countyInit
		states   []*cspc.StateProvince
	)

	err := json.Unmarshal([]byte(cspc.USCountyJSON), &counties)
	if err != nil {
		return nil, errs.E(op, err)
	}

	err = json.Unmarshal([]byte(cspc.USStatesJSON), &states)
	if err != nil {
		return nil, errs.E(op, err)
	}

	for _, state := range states {
		for _, ic := range counties {
			statefip := ic.CountyCode[:2]

			if state.FIPSCode == statefip {

				c := cspc.County{
					Code:             ic.CountyCode,
					Name:             ic.CountyName,
					LatitudeAverage:  "",
					LongitudeAverage: "",
				}

				state.Counties = append(state.Counties, c)
			}
		}
	}

	return states, nil
}

func loadCounties2db(ctx context.Context, states []*cspc.StateProvince) error {
	const op errs.Op = "main/loadCounties2db"

	logger := app.NewLogger(zerolog.DebugLevel)

	db, cleanup, err := datastore.NewDB(datastore.LocalDatastore, logger)
	if err != nil {
		return errs.E(op, err)
	}
	defer cleanup()

	// declare variable as the Transactor interface
	var datastorer datastore.Datastorer

	datastorer = datastore.NewDatastore(db)

	a := app.NewApplication(app.Local, datastorer, logger)

	tx, err := a.Datastorer.BeginTx(ctx)
	if err != nil {
		return errs.E(op, err)
	}

	// declare variable as the Transactor interface
	var transactor countystore.Transactor

	transactor, err = countystore.NewTx(tx)
	if err != nil {
		return errs.E(op, err)
	}

	for _, s := range states {
		for _, c := range s.Counties {
			err = transactor.CreateCounty(ctx, "US", s.Code, c)
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
