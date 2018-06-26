package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type (
	// Config for the plugin.
	Config struct {
		ProjectID   string
		ProjectKey  string
		Environment string
		BuildAuthor string
		BuildCommit string
		RepoLink    string
	}

	// Plugin structure
	Plugin struct {
		Config Config
	}

	airbrakePayload struct {
		Environment string `json:"environment"`
		Username    string `json:"username"`
		Repository  string `json:"repository"`
		Revision    string `json:"revision"`
	}
)

// Exec executes the plugin.
func (p Plugin) Exec() error {
	fmt.Println("=====================================")
	fmt.Println("= Here is drone-airbrake-deployment =")
	fmt.Println("=====================================")

	airbrakeURL := fmt.Sprintf("https://airbrake.io/api/v4/projects/%s/deploys?key=%s", p.Config.ProjectID, p.Config.ProjectKey)
	payload := &airbrakePayload{
		Environment: p.Config.Environment,
		Username:    p.Config.BuildAuthor,
		Repository:  p.Config.RepoLink,
		Revision:    p.Config.BuildCommit,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", airbrakeURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	fmt.Println("Sending deploy track to airbrake...")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		fmt.Println("!!! Error happened")

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		fmt.Printf("Response Status: %s\n", resp.Status)
		fmt.Printf("Response body: %s\n", string(bodyBytes))

		return errors.New("Airbrake deploy track fail")
	}
	fmt.Println("Success!")

	return nil
}
