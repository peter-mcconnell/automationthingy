package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/peter-mcconnell/automationthingy/config"
)

type Github struct {
	Logger       config.Logger
	ClientID     string
	clientSecret string
	RedirectUri  string
}

func (g *Github) LoginHandler(w http.ResponseWriter, r *http.Request) {
	redirect := fmt.Sprintf(
		"https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=%s",
		g.ClientID,
		g.RedirectUri,
	)
	http.Redirect(w, r, redirect, http.StatusTemporaryRedirect)
}

func (g *Github) CallbackHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	accessToken, err := g.getGithubAccessToken(code)
	if err != nil {
		g.Logger.Fatal(err)
		// TODO: handle properly
		panic(err)
	}
	data, err := g.getData(accessToken)
	if err != nil {
		g.Logger.Fatal(err)
		// TODO: handle properly
		panic(err)
	}

	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, []byte(data), "", "\t")
	if err != nil {
		g.Logger.Fatal(err)
		// TODO: handle properly
		panic(err)
	}
	fmt.Fprintf(w, string(prettyJSON.Bytes()))
}

func (g *Github) getGithubAccessToken(code string) (string, error) {
	requestBodyMap := map[string]string{
		"client_id":     g.ClientID,
		"client_secret": g.clientSecret,
		"code":          code,
	}
	requestJson, err := json.Marshal(requestBodyMap)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest(
		"POST",
		"https://github.com/login/oauth/access_token",
		bytes.NewBuffer(requestJson),
	)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	type githubAccessTokenResponse struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		Scope       string `json:"scope"`
	}

	var parsedResp githubAccessTokenResponse
	err = json.Unmarshal(body, &parsedResp)
	if err != nil {
		return "", err
	}

	return parsedResp.AccessToken, nil
}

func (g *Github) getData(accessToken string) (string, error) {
	req, err := http.NewRequest(
		"GET",
		"https://api.github.com/user",
		nil,
	)
	if err != nil {
		return "", err
	}
	authorizationHeaderValue := fmt.Sprintf("token %s", accessToken)
	req.Header.Set("Authorization", authorizationHeaderValue)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(respBody), nil
}
