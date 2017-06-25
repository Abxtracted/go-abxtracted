// Package abxtracted provides a client to access the abxtracted.com API.
package abxtracted

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const apiURL = "https://api.abxtracted.com/public"

// ErrParseCerts occurs when parse of certificates fail
var ErrParseCerts = errors.New("failed to parse certs")

// Project represents the key of project
type Project string

// API for the abxtracted API
type API struct {
	client http.Client
}

// New abxtracted API client for a given project
func New() (*API, error) {
	var pool = x509.NewCertPool()
	if ok := pool.AppendCertsFromPEM([]byte(pemCerts)); !ok {
		return nil, ErrParseCerts
	}
	return &API{
		client: http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					RootCAs: pool,
				},
			},
		},
	}, nil
}

// Experiment from abxtracted API
type Experiment struct {
	CustomerIdentity string `json:"customerIdentity"`
	Experiment       string `json:"experiment"`
	Scenario         string `json:"scenario"`
}

// Get returns the experiment for a given customer identity and
// experiment key.
func (api *API) Get(
	project Project, customerIdentity, experimentKey string,
) (e Experiment, err error) {
	var url = fmt.Sprintf(
		apiURL+"/project/%s/customer/%s/experiment/%s",
		project,
		customerIdentity,
		experimentKey,
	)
	resp, err := api.client.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&e)
	return
}

// Complete completes a given experiment
func (api *API) Complete(project Project, experiment Experiment) error {
	var url = fmt.Sprintf(
		apiURL+"/project/%s/customer/%s/experiment/%s/check/complete",
		project,
		experiment.CustomerIdentity,
		experiment.Experiment,
	)
	_, err := api.client.Get(url)
	return err
}
