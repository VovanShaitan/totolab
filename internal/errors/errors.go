package errors

import "fmt"

type CustomErrorExample struct {
	ErrorInfo string
}

func (e *CustomErrorExample) Error() string {
	return fmt.Sprintf("Oops! Something went wrong. %v", e.ErrorInfo)
}

// infoError := errors.New("Bla-bla-bla...")
