package activity

import "time"

var lassMapping = map[string]string{
	"0": "NoRing",
	"1": "Rest",
	"2": "Inactive",
	"3": "LowActivity",
	"4": "MediumActivity",
	"5": "HighActivity",
}

// daily updted
type Contributors struct {
	Prev7DailyTargetScore        int `json:"meet_daily_targets,omitempty"` //: 43,
	Prev24HourInactivAlertsScore int `json:"move_every_hour,omitempty"`    //: 100,
	Prev7DayRecoveryScore        int `json:"recovery_time,omitempty"`      //: 100,
	Prev24HourActivityScore      int `json:"stay_active,omitempty"`        //: 98,
	Prev7DaySportFreqScore       int `json:"training_frequency,omitempty"` //: 71,
	Prev7DaySprotVolumeScore     int `json:"training_volume,omitempty"`    //: 98
}

// dynamic
type MET struct {
	IntervalSec    float32   `json:"interval,omitempty"`  // 60,
	Values         []float32 `json:"items,omitempty"`     // [],
	StartTimestamp time.Time `json:"timestamp,omitempty"` // "2021-11-26T04:00:00.000-08:00"
}

type Dynamic struct {
	// dynamic
	ClassificationEvery5Min         string  `json:"class_5_min,omitempty"`                 // "000000000000000000000000000000000000000000000000000000000000000000000000003444544444445545455443454554454443334333322330000000000232232222222232222222322223222000000022332233422333222232233333222222222222222332223212233222122221111111111111121111111111111111111111111111111111111111111111",
	ActiveKCal                      int     `json:"active_calories,omitempty"`             // 1222,
	AvgMETMin                       float32 `json:"average_met_minutes,omitempty"`         // 1.90625,
	EnergyExpendInMetersWalked      int     `json:"equivalent_walking_distance,omitempty"` // 20122,
	HighActivityMETMin              int     `json:"high_activity_met_minutes,omitempty"`   // 444,
	HighActivityTimeSec             int     `json:"high_activity_time,omitempty"`          // 3000,
	AmountInactivityAlerts          int     `json:"inactivity_alerts,omitempty"`           // 0,
	LowActivityMETMin               int     `json:"low_activity_met_minutes,omitempty"`    // 117,
	LowActivityTimeSec              int     `json:"low_activity_time,omitempty"`           // 10020,
	MediumActivityMETMin            int     `json:"medium_activity_met_minutes,omitempty"` // 391,
	MediumActivityTimeSec           int     `json:"medium_activity_time,omitempty"`        // 6060,
	MET                             MET     `json:"met,omitempty"`                         // {},
	RemainingActivityToTargetMeters int     `json:"meters_to_target,omitempty"`            // -16200,
	TimeWithoutRingSec              int     `json:"non_wear_time,omitempty"`               // 27480,
	RestingTimeSec                  int     `json:"resting_time,omitempty"`                // 18840,
	SedentaryMETMin                 int     `json:"sedentary_met_minutes,omitempty"`       // 10,
	SedentaryTimeSec                int     `json:"sedentary_time,omitempty"`              // 21000,
	Steps                           int     `json:"steps,omitempty"`                       // 18430,
	TotalKCal                       int     `json:"total_calories,omitempty"`              // 3446,

	// ??
	Score int `json:"score,omitempty"` // 82,

}

type Daily struct {
	// updated daily
	Contributors              Contributors `json:"contributors,omitempty"`    // {},
	DailyActivityTargetKCal   int          `json:"target_calories,omitempty"` // 350,
	DailyActivityTargetMeters int          `json:"target_meters,omitempty"`   // 7000,
}

type Data struct {
	Dynamic
	Daily

	//Date                       time.Time           `json:"day,omitempty"`                         // "2021-11-26",
	Timestamp time.Time `json:"timestamp,omitempty"` // "2021-11-26T04:00:00-08:00"
}

type DailyResponse struct {
	Data      []Data
	NextToken string
}
