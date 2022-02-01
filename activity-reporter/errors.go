package activityreporter

import "fmt"

type baseError struct {
	err error
	msg string
}

type NoStepsProviderError baseError

func NewNoStepsProviderError() NoStepsProviderError {
	return NoStepsProviderError{}
}

func (me NoStepsProviderError) Error() string {
	return "ActivityReporter has no configured steps provider"
}

func (me NoStepsProviderError) Unwrap() error {
	return me.err
}

func (me NoStepsProviderError) Is(target error) bool {
	_, ok := target.(NoStepsProviderError)
	return ok
}

type NotImplementedError struct {
	featureName string
	baseError
}

func NewNotImplementedError(featureName string) NotImplementedError {
	return NotImplementedError{featureName: featureName}
}

func (me NotImplementedError) Error() string {
	return fmt.Sprintf("feature %s is not implemented", me.featureName)
}

func (me NotImplementedError) Unwrap() error {
	return me.err
}

func (me NotImplementedError) Is(target error) bool {
	_, ok := target.(NotImplementedError)
	return ok
}
