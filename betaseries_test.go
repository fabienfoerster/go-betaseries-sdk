package betaseries_test

import (
	"testing"

	"github.com/fabienfoerster/go-betaseries-sdk"
)

var expectedToken string = "baafb5f7d64c"

func TestAuth(t *testing.T) {
	api := betaseries.NewBetaseriesAPI("70b5e42ba85a")
	if api.Token != "" {
		t.Errorf("Before auth the token should be empty")
	}
	api.Auth("binou42", "courgette")
	if api.Token != expectedToken {
		t.Errorf("Expected token : %s , actual : %s", expectedToken, api.Token)
	}
}
