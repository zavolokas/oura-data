package ouraring

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/zavolokas/oura-data/ouraring/activity"
)

type I interface {
	GetDailyActivity() (*activity.DailyResponse, error)
}

type client struct {
	httpClient http.Client
	token      string
}

func NewClient(token string) I {

	return &client{
		httpClient: http.Client{Timeout: 5 * time.Second},
		token:      token,
	}
}

func (cli *client) GetDailyActivity() (*activity.DailyResponse, error) {

	url := "https://api.ouraring.com/v2/usercollection/daily_activity"
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Printf("error %s", err)
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", cli.token))
	resp, err := cli.httpClient.Do(req)
	if err != nil {
		fmt.Printf("error %s", err)
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Body : %s", body)

	daresp := activity.DailyResponse{}
	err = json.Unmarshal(body, &daresp)
	if err != nil {
		fmt.Printf("error %s", err)
		return nil, err
	}

	return &daresp, nil
}
