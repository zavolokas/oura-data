package ouraring

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type ActivityContributor struct {
	MeetDailyTargets  int `json:"meet_daily_targets"` //: 43,
	MoveEveryHour     int `json:"move_every_hour"`    //: 100,
	RecoveryTime      int `json:"recovery_time"`      //: 100,
	StayActive        int `json:"stay_active"`        //: 98,
	TrainingFrequency int `json:"training_frequency"` //: 71,
	TrainingVolume    int `json:"training_volume"`    //: 98
}

type MET struct {
	Interval  float32   `json:"interval"`  // 60,
	Items     []float32 `json:"items"`     // [],
	Timestamp time.Time `json:"timestamp"` // "2021-11-26T04:00:00.000-08:00"
}

type ActivityData struct {
	Class5Min                 string              `json:"class_5_min"`                 // "000000000000000000000000000000000000000000000000000000000000000000000000003444544444445545455443454554454443334333322330000000000232232222222232222222322223222000000022332233422333222232233333222222222222222332223212233222122221111111111111121111111111111111111111111111111111111111111111",
	Score                     int                 `json:"score"`                       // 82,
	ActiveCalories            int                 `json:"active_calories"`             // 1222,
	AvgMETMinutes             float32             `json:"average_met_minutes"`         // 1.90625,
	Contributors              ActivityContributor `json:"contributors"`                // {},
	EquivalentWalkingDistance int                 `json:"equivalent_walking_distance"` // 20122,
	HighActivityMETMinutes    int                 `json:"high_activity_met_minutes"`   // 444,
	HighActivityTime          int                 `json:"high_activity_time"`          // 3000,
	InactivityAlerts          int                 `json:"inactivity_alerts"`           // 0,
	LowActivityMETMinutes     int                 `json:"low_activity_met_minutes"`    // 117,
	LowActivityTime           int                 `json:"low_activity_time"`           // 10020,
	MediumActivityMETMinutes  int                 `json:"medium_activity_met_minutes"` // 391,
	MediumActivityTime        int                 `json:"medium_activity_time"`        // 6060,
	MET                       MET                 `json:"met"`                         // {},
	MetersToTarget            int                 `json:"meters_to_target"`            // -16200,
	NonWearTime               int                 `json:"non_wear_time"`               // 27480,
	RestingTime               int                 `json:"resting_time"`                // 18840,
	SedentaryMETMinutes       int                 `json:"sedentary_met_minutes"`       // 10,
	SedentaryTime             int                 `json:"sedentary_time"`              // 21000,
	Steps                     int                 `json:"steps"`                       // 18430,
	TargetCalories            int                 `json:"target_calories"`             // 350,
	TargetMeters              int                 `json:"target_meters"`               // 7000,
	TotalCalories             int                 `json:"total_calories"`              // 3446,
	//Day                       time.Time           `json:"day"`                         // "2021-11-26",
	Timestamp time.Time `json:"timestamp"` // "2021-11-26T04:00:00-08:00"
}

type DailyActivityResponse struct {
	Data      []ActivityData
	NextToken string
}

type I interface {
	GetDailyActivity() (*DailyActivityResponse, error)
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

func (cli *client) GetDailyActivity() (*DailyActivityResponse, error) {

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

	daresp := DailyActivityResponse{}
	err = json.Unmarshal(body, &daresp)
	if err != nil {
		fmt.Printf("error %s", err)
		return nil, err
	}

	return &daresp, nil
}
