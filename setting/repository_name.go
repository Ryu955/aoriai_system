package setting

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetRepoName() []string {
	// https://api.github.com/search/repositories?q=ryutai+org:hillive&type=Repositories
	data, err := ioutil.ReadFile("./repos.csv")
	if err != nil {
		panic(err)
	}
	str_data := strings.TrimRight(string(data), "\n")
	split_data := strings.Split(str_data, ",")
	fmt.Println(split_data)

	// url を検索

	//q :="201801_GraduationThesis"
	q := split_data[0] + "_" + split_data[1]

	// url := "https://api.github.com/search/repositories?q=ryutai+org:hillive&type=Repositories"
	url := "https://api.github.com/search/repositories?q=" + q + "+org:hillive&type=Repositories"

	api_key := GetApiKey()

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", api_key)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	bytes := ([]byte)(body)
	// fmt.Println(string(bytes))

	var rs RepoSearch
	err = json.Unmarshal(bytes, &rs)
	if err != nil {
		fmt.Println("error:", err)
	}

	var repos []string
	for _, item := range rs.Items {
		//fmt.Println(item.Name)
		repos = append(repos, item.Name)
	}
	return repos
}
