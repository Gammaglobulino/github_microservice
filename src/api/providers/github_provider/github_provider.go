package github_provider

import (
	"api/clients/restclient"
	"api/domain/github"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	headerAuthorization       = "Authorization"
	headerAuthorizationFormat = "token %s"
	urlCreateRepo             = "https://api.github.com/user/repos"
)

func getAuthorizationHeader(accessToken string) string {
	return fmt.Sprintf(headerAuthorizationFormat, accessToken)
}
func CreateRepo(accessToken string, request github.CreateRepoRequest) (*github.CreateRepoResponse, *github.GitHubErrorResponse) {
	headers := http.Header{}
	headers.Set(headerAuthorization, getAuthorizationHeader(accessToken))
	resp, err := restclient.Post(urlCreateRepo, request, headers)
	if err != nil {
		log.Println("Error when trying to create a repo in github: %s", err.Error())
		return nil, &github.GitHubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error when trying to create a repo in github: %s", err.Error())
		return nil, &github.GitHubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "invalid response body",
		}
	}
	if resp.StatusCode > 299 {
		var errResponse github.GitHubErrorResponse
		if err := json.Unmarshal(bytes, &errResponse); err != nil {
			return nil, &github.GitHubErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Message:    "invalid json response",
			}
		}
		errResponse.StatusCode = resp.StatusCode
		return nil, &errResponse
	}
	var result github.CreateRepoResponse
	if err := json.Unmarshal(bytes, &result); err != nil {
		log.Println("Error when trying to create a repo in github: %s", err.Error())
		return nil, &github.GitHubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "invalid json response",
		}
	}
	return &result, nil
}
