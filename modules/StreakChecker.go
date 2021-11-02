package StreakChecker

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"log"
	"time"
	"net/http"
)

type Submission struct {
	// Id				int           `json:"id"`
	Epoch_second  	int           `json:"epoch_second"`
	// Problem_id    	string        `json:"problem_id"`
	// Contest_id    	string        `json:"contest_id"`
	// User_id			string        `json:"user_id"`
	// Language      	string        `json:"language"`
	// Point         	float64       `json:"point"`
	// Length        	int           `json:"length"`
	Result        	string        `json:"result"` // must "AC"
	// Exection_time 	int           `json:"execution_time"`
}


type StreakRank struct {
    Count  int `json:"count"`
    Rank   int `json:"rank"`
}

func IsAcceptedToday () bool {
	unixTime := getTimeUnix()
	submission := getSubmissions(unixTime)
	for _, data := range(submission) {
		fmt.Println(data.Epoch_second, data.Result)
		if data.Result == "AC" {
			return true
		}
	}
	return false
}

func getTimeUnix() int {
	nowUTC := time.Now().UTC()
	jst := time.FixedZone("JST", +9*60*60)
	nowJST := nowUTC.In(jst)
	today := time.Date(nowJST.Year(), nowJST.Month(), nowJST.Day(), 0, 0, 0, 0, jst)
	unix := today.Unix()
	return int(unix)
}

func getSubmissions(unixTime int) []Submission {
	url := "https://kenkoooo.com/atcoder/atcoder-api/v3/user/submissions?user=totori0908&from_second=" + strconv.Itoa(unixTime)
	jsonStr := httpGetStr(url)
	submissions := formatSubmission(jsonStr)
	return submissions
}

func getStreakRank() *StreakRank {
	jsonStr := httpGetStr("https://kenkoooo.com/atcoder/atcoder-api/v3/user/streak_rank?user=totori0908")
	streak := formatStreakRank(jsonStr)
	return streak
}

func httpGetStr(url string) string {
	// HTTPリクエストを発行しレスポンスを取得する
	response, err := http.Get(url)
	if err != nil {
		log.Fatal("Get Http Error:", err)
	}
	// レスポンスボディを読み込む
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal("IO Read Error:", err)
	}
	// 読み込み終わったらレスポンスボディを閉じる
	defer response.Body.Close()
	return string(body)
}


func formatSubmission(jsonStr string) []Submission {
	var sub []Submission
	// fmt.Println(jsonStr)
	if err := json.Unmarshal([]byte(jsonStr), &sub); err != nil {
		log.Fatal(err)
	}
	return sub
}

func formatStreakRank(jsonStr string) *StreakRank {
	sub := new(StreakRank)
	// fmt.Println(jsonStr)
	if err := json.Unmarshal([]byte(jsonStr), sub); err != nil {
		log.Fatal(err)
	}
	return sub
}