package activityreporter

import (
	"time"

	"github.com/zavolokas/oura-data/ouraring"
)

type I interface {
	GetSteps(start time.Time, end time.Time) (int, error)
}

type reporter struct {
	OuraClient ouraring.I
}

type Option func(r *reporter)

func New(options ...Option) I {
	r := &reporter{}

	for _, opt := range options {
		opt(r)
	}

	return r
}

func (r *reporter) GetSteps(start time.Time, end time.Time) (int, error) {
	if r.OuraClient != nil {

	}

	return -1, NewNoStepsProviderError()
}

func WithOuraRingData(ouraToken string) Option {
	return func(r *reporter) {
		r.OuraClient = ouraring.NewClient(ouraToken)
	}
}
