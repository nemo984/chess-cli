package lichess

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func AnalysisURL(pgn string) (string, error) {
	url := "https://lichess.org/api/import"
	values := map[string]string{"pgn": strings.TrimSpace(pgn)}
	jsonValue, _ := json.Marshal(values)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("Non-OK HTTP status: %v", resp.StatusCode)
	}

	j := struct {
		ID  string `json:"id"`
		URL string `json:"url"`
	}{}
	err = json.NewDecoder(resp.Body).Decode(&j)
	if err != nil {
		return "", err
	}
	return j.URL, nil
}

func GetPuzzle() (DailyPuzzle,error) {
	j := DailyPuzzle{}
	url := "https://lichess.org/api/puzzle/daily"
	resp,err := http.Get(url)
	if err != nil {
		return j,err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return j,fmt.Errorf("Non-OK HTTP status: %v",resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&j)
	if err != nil {
		return j,err
	}
	fmt.Println(j)

	return j,nil

}