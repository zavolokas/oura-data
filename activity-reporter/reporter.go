package activityreporter

import (
	"time"
)

type StepsProvider interface {
	GetSteps(start time.Time, end time.Time) (int, error)
}

type I interface {
	StepsProvider
}

type reporter struct {
	StepsProviders []StepsProvider
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
	if len(r.StepsProviders) == 0 {
		return -1, NewNoStepsProviderError()
	}

	// TODO: should return aggregated and provoder tagged data
	return r.StepsProviders[0].GetSteps(start, end)
}

// TODO: move the option to Ring package
func WithOuraRingData(ouraToken string) Option {
	return func(r *reporter) {
		ouraReporter := newOuraReporter(ouraToken)

		r.StepsProviders = append(r.StepsProviders, ouraReporter)
	}
}
