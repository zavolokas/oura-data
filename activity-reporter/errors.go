package activityreporter

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
