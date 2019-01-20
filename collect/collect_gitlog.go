package collect

import (
	"encoding/json"
	"fmt"
	"github.com/Ryu955/shintyoku/setting"
	"io/ioutil"
	"net/http"
	"sort"
	"time"
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

func GetCommitLog(url string) GitLog {
	api_key := setting.GetApiKey()

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
	}
	return gl
}

func JsonToStruct(commit_map map[string]int, gl GitLog) {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	for _, log := range gl {
		commit_date := log.Commit.Author.Date.In(jst).Format("2006-01-02")
		_, is_exist := commit_map[commit_date]
		if is_exist {
			commit_map[commit_date] = commit_map[commit_date] + 1
		} else {
			commit_map[commit_date] = 1
		}
	}
}

func GetGitLog(repo_name string) List {

	//url := "https://api.github.com/repos/hillive/" + repo_name + "/commits?per_page=100&page=2"
	url := "https://api.github.com/repos/hillive/" + repo_name + "/commits?per_page=100"
	// https://developer.github.com/v3/#pagination
	// url := "https://api.github.com/repos/hillive/201801_GraduationThesis_ryutai/commits"
	//url := "https://api.github.com/repos/ryu955/dotfiles/commits"

	gl := GetCommitLog(url)

	commit_map := make(map[string]int)
	JsonToStruct(commit_map, gl)

	//fmt.Println(len(commit_map))
	// fmt.Println(commit_map)

	if len(gl) == 100 {
		url = "https://api.github.com/repos/hillive/" + repo_name + "/commits?per_page=100&page=2"
		gl2 := GetCommitLog(url)
		//	fmt.Println(len(gl2))
		JsonToStruct(commit_map, gl2)
	}
	//fmt.Println(commit_map)
	//fmt.Println(len(commit_map))

	var count int
	for _, v := range commit_map {
		count = count + v
	}

	fmt.Print("総コミット回数：")
	fmt.Println(count)

	sorted_log := List{}
	for k, v := range commit_map {
		e := Log{k, v}
		sorted_log = append(sorted_log, e)
	}

	sort.Sort(sorted_log)
	return sorted_log
}
