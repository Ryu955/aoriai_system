package collect

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
		"time"
	"strings"
)

type GitLog []struct {
	Sha    string `json:"sha"`
	NodeID string `json:"node_id"`
	Commit struct {
		Author struct {
			Name  string    `json:"name"`
			Email string    `json:"email"`
			Date  time.Time `json:"date"`
		} `json:"author"`
		Committer struct {
			Name  string    `json:"name"`
			Email string    `json:"email"`
			Date  time.Time `json:"date"`
		} `json:"committer"`
		Message string `json:"message"`
		Tree    struct {
			Sha string `json:"sha"`
			URL string `json:"url"`
		} `json:"tree"`
		URL          string `json:"url"`
		CommentCount int    `json:"comment_count"`
		Verification struct {
			Verified  bool        `json:"verified"`
			Reason    string      `json:"reason"`
			Signature interface{} `json:"signature"`
			Payload   interface{} `json:"payload"`
		} `json:"verification"`
	} `json:"commit"`
	URL         string `json:"url"`
	HTMLURL     string `json:"html_url"`
	CommentsURL string `json:"comments_url"`
	Author      struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		NodeID            string `json:"node_id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"author"`
	Committer struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		NodeID            string `json:"node_id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"committer"`
	Parents []struct {
		Sha     string `json:"sha"`
		URL     string `json:"url"`
		HTMLURL string `json:"html_url"`
	} `json:"parents"`
}

func GetGitLog() {

	url := "https://api.github.com/repos/hillive/201801_GraduationThesis_ryutai/commits"
	//url := "https://api.github.com/repos/ryu955/dotfiles/commits"

	api_key := GetApiKey()
	fmt.Println(api_key)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", api_key)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	bytes := ([]byte)(body)

	var gl GitLog
	err := json.Unmarshal(bytes, &gl)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	//fmt.Println(gl[0].Commit.Author.Date)

	for count, log := range gl {
		fmt.Println(count)
		fmt.Println(log.Commit.Author.Date)
	}

	// fmt.Println(string(body))

}

func GetApiKey() string {

	data, err := ioutil.ReadFile("./.env")
	if err != nil {
		// エラー処理
		fmt.Println("don't read file")
		panic(err)
	}

	api_key := "token " +strings.TrimRight(string(data), "\n")
	return api_key
}
