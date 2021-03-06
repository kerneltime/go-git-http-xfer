package githttpxfer

import "fmt"

type URLNotFoundError struct {
	Method string
	Path   string
}

func (e *URLNotFoundError) Error() string {
	return fmt.Sprintf("Url Not Found: Method %s, Path %s", e.Method, e.Path)
}

type MethodNotAllowedError struct {
	Method string
	Path   string
}

func (e *MethodNotAllowedError) Error() string {
	return fmt.Sprintf("Method Not Allowed: Method %s, Path %s", e.Method, e.Path)
}
