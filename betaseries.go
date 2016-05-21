//Package betaseries provide structs and functions for accessing version 2.4
// of the Betaseries API.
package betaseries

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var betaseriesKey string

type BetaAPI struct {
	Key   string
	Token string
}

type betaseriesAuthResp struct {
	Token string `json:"token"`
}

func (api *BetaAPI) SetToken(token string) {
	api.Token = token
}

func NewBetaseriesAPI(apiKey string) *BetaAPI {
	api := &BetaAPI{
		Key:   apiKey,
		Token: "",
	}
	return api
}

func toMD5(s string) string {
	h := md5.New()
	io.WriteString(h, s)
	checksum := fmt.Sprintf("%x", h.Sum(nil))
	return checksum
}

func (api *BetaAPI) Auth(login, password string) {
	md5 := toMD5(password)
	resp, err := http.PostForm("https://api.betaseries.com/members/auth", url.Values{"key": {api.Key}, "login": {login}, "password": {md5}})
	defer resp.Body.Close()
	if err != nil {
		log.Fatalf("error when connecting to betaseries : %s", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error when reding the response from betaseries : %s", err)
	}
	var betaResp betaseriesAuthResp
	err = json.Unmarshal(body, &betaResp)
	if err != nil {
		log.Fatalf("error when unmarshalling betaseries json response : %s", err)
	}
	token := betaResp.Token
	api.SetToken(token)
}
