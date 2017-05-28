package jobs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"sync"

	"github.com/Yu-taro/taue/taue/models"
)

// getContributes from API
func getContributes(users []models.User) (completeHandler []models.User) {

	ch := make(chan models.User, len(users))
	var wg sync.WaitGroup

	for _, user := range users {
		wg.Add(1)
		go getGitActivity(&wg, user, ch)
	}

	wg.Wait()
	close(ch)

	for data := range ch {
		completeHandler = append(completeHandler, data)
	}

	return completeHandler

}

func getGitActivity(wg *sync.WaitGroup, user models.User, ch chan models.User) {
	user.GitHubEvents = getGitHubContributes(user)
	user.GitLabEvents = getGitLabContributes(user)
	ch <- user
	wg.Done()
}

func getGitHubContributes(user models.User) (githubEvents []models.GitHubEvent) {
	if !user.GitHubName.Valid && user.GitHubName.String == "" {
		return
	}

	value := url.Values{}
	value.Add("per_page", "100")
	if user.GitHubToken.Valid && user.GitHubToken.String != "" {
		value.Add("access_token", user.GitHubToken.String)
	}

	const baseURL = "https://api.github.com"
	urlString := baseURL + "/users/" + user.GitHubName.String + "/events"

	resp, err := http.Get(urlString + "?" + value.Encode())
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("%s", body)
		return
	}

	//var githubEvents []models.GitHubEvent
	err = json.Unmarshal(body, &githubEvents)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return githubEvents

}

func getGitLabContributes(user models.User) (gitlabEvents []models.GitLabEvent) {
	if !user.GitLabID.Valid {
		return
	}

	value := url.Values{}
	value.Add("per_page", "100")
	if user.GitHubToken.Valid && user.GitLabToken.String != "" {
		value.Add("private_token", user.GitLabToken.String)
	}

	const baseURL = "https://gitlab.com/api/v3"
	urlString := baseURL + "/users/" + strconv.FormatInt(user.GitLabID.Int64, 10) + "/events"

	resp, err := http.Get(urlString + "?" + value.Encode())
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("%s", body)
		return
	}

	err = json.Unmarshal(body, &gitlabEvents)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return gitlabEvents

}
