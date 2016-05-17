//Package betaseries provide structs and functions for accessing version 2.4
// of the Betaseries API.
package betaseries

import (
	"log"
	"os"
)

var betaseriesKey string

type BetaseriesAPI struct {
	Key   string
	Token string
}

func init() {
	if betaseriesKey = os.Getenv("BETASERIES_KEY"); betaseriesKey == "" {
		log.Fatal("BETASERIES_KEY must be set in env")
	}
}
