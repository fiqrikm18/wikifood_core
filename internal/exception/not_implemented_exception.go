package exception

type NotImplementedException struct{}

func (e *NotImplementedException) Error() string {
	return "Not Implemented"
}
