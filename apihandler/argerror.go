package apihandler

import "fmt"

type apiHandlerError struct {
	arg    string
	reason string
}

func (a *apiHandlerError) Error() string {
	return fmt.Sprintf("%s :%s", a.reason, a.arg)
}
