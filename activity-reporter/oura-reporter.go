package activityreporter

import (
	"log"
	"time"

	"github.com/zavolokas/oura-data/ouraring"
)

type ouraReporter struct {
	OuraClient ouraring.I
}

func newOuraReporter(ouraToken string) I {
	return &ouraReporter{
		OuraClient: ouraring.NewClient(ouraToken),
	}
}

func (r *ouraReporter) GetSteps(start time.Time, end time.Time) (int, error) {
	response, err := r.OuraClient.GetDailyActivity()
	if err != nil {
		log.Printf("error getting daily activity from oura")
		return -1, err
	}
	return response.Data[0].Steps, nil

	//return -1, NewNotImplementedError("Oura.GetSteps")
}
