package main

import (
	json2 "encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func fetch(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	html := string(body)
	return html
}

func fetchProblems() string {
	resp, err := http.Post("https://leetcode-cn.com/graphql", "application/json", strings.NewReader(`{"operationName":"getQuestionTranslation","variables":{},"query":"query getQuestionTranslation($lang: String) {\n  translations: allAppliedQuestionTranslations(lang: $lang) {\n    title\n    questionId\n    __typename\n  }\n}\n"}`))

	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	html := string(body)
	//fmt.Println(html)
	return html
}

func main() {
	jsonStr := fetch("https://leetcode-cn.com/api/problems/all/")
	json := map[string]interface{}{}
	err := json2.Unmarshal([]byte(jsonStr), &json)
	if err != nil {
		panic(err)
	}
	list := json["stat_status_pairs"].([]interface{})
	pid := os.Args[1]
	println("pid:"+pid)
	//pid := "13"
	for i := range list {
		stat := list[i].(map[string]interface{})["stat"].(map[string]interface{})
		if stat["frontend_question_id"] == pid {
			questionTitleSlug := stat["question__title_slug"].(string)
			println("create " + questionTitleSlug)
			filename := strings.ReplaceAll(stat["question__title_slug"].(string), "-", "")
			err := os.Mkdir(pid, os.FileMode(0755))
			if err != nil {
				panic(err)
			}
			file, err := os.Create(pid + "/" + filename + ".go")
			if err != nil {
				panic(err)
			}
			_, err = file.WriteString("package _" + pid + "\n\n//https://leetcode-cn.com/problems/" + questionTitleSlug + "/")
			if err != nil {
				panic(err)
			}
			err = file.Close()
			if err != nil {
				panic(err)
			}
		}
	}
}
