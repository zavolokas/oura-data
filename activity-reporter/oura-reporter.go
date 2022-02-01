package activityreporter

import (
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
	return -1, NewNotImplementedError("Oura.GetSteps")
}
