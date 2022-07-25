package auth

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/peter-mcconnell/automationthingy/config"
	"github.com/peter-mcconnell/automationthingy/secretmgr"
)

type Github struct {
	Logger          config.Logger
	GithubConfig    *config.GithubAuth
	SecretmgrConfig *secretmgr.ConfigSecretMgr
}

func (g *Github) LoginHandler(w http.ResponseWriter, r *http.Request) {
	redirect := fmt.Sprintf(
		"https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=%s",
		g.GithubConfig.Clientid,
		g.GithubConfig.Redirecturi,
	)
	http.Redirect(w, r, redirect, http.StatusTemporaryRedirect)
}

func (g *Github) CallbackHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	accessToken, err := g.getGithubAccessToken(code)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	data, err := g.getData(accessToken)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, []byte(data), "", "\t")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintf(w, string(prettyJSON.Bytes()))
}

func (g *Github) loadSecret() (string, error) {
	secretMgr, err := secretmgr.GetSecretMgr(g.GithubConfig.Secretref, g.SecretmgrConfig)
	if err != nil {
		return "", err
	}
	secret, err := secretMgr.Get(g.GithubConfig.Secretref, "secret")
	if err != nil {
		return "", err
	}
	return string(secret), nil
}

func (g *Github) getGithubAccessToken(code string) (string, error) {
	clientSecret, err := g.loadSecret()
	if err != nil {
		return "", err
	}
	requestBodyMap := map[string]string{
		"client_id":     g.GithubConfig.Clientid,
		"client_secret": clientSecret,
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
		Error       string `json:"error,omitempty"`
		ErrorDesc   string `json:"error_description,omitempty"`
		ErrorUri    string `json:"error_uri,omitempty"`
		AccessToken string `json:"access_token,omitempty"`
		TokenType   string `json:"token_type,omitempty"`
		Scope       string `json:"scope,omitempty"`
	}

	var parsedResp githubAccessTokenResponse
	err = json.Unmarshal(body, &parsedResp)
	if err != nil {
		return "", err
	}

	if parsedResp.Error != "" {
		return "", errors.New(parsedResp.ErrorDesc)
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
