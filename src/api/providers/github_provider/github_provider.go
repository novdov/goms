package github_provider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/novdov/goms/src/api/clients/rest_client"
	"github.com/novdov/goms/src/api/domain/github"
)

const (
	headerAuthorization       = "Authorization"
	headerAuthorizationFormat = "token %s"
	urlCreateRepo             = "https://api.github.com/user/repos"
)

func getAuthorizationHeader(accessToken string) string {
	return fmt.Sprintf(headerAuthorizationFormat, accessToken)
}

func CreateRepo(request github.CreateRepoRequest, accessToken string) (
	*github.CreateRepoResponse, *github.GhErrorResponse) {
	headers := http.Header{}
	headers.Set(headerAuthorization, getAuthorizationHeader(accessToken))
	response, err := rest_client.Post(urlCreateRepo, request, headers)
	if err != nil {
		log.Printf("error when trying to create new repo in github: %s", err.Error())
		// Cannot return nil when not returning pointer to struct
		return nil, &github.GhErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	jsonBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, &github.GhErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "invalid response body",
		}
	}
	defer response.Body.Close()

	if response.StatusCode >= 300 {
		var errResponse github.GhErrorResponse
		if err := json.Unmarshal(jsonBytes, &errResponse); err != nil {
			return nil, &github.GhErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Message:    "invalid json response body",
			}
		}
		errResponse.StatusCode = response.StatusCode
		return nil, &errResponse
	}

	var result github.CreateRepoResponse
	if err := json.Unmarshal(jsonBytes, &result); err != nil {
		log.Printf("error when trying to unmarshal github create repo respones: %s", err.Error())
		return nil, &github.GhErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return &result, nil
}
