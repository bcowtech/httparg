package httparg

import (
	"sync"

	"github.com/bcowtech/arg"
	"github.com/bcowtech/httparg/form"
	"github.com/bcowtech/httparg/internal"
	"github.com/bcowtech/httparg/json"
	"github.com/bcowtech/httparg/querystring"
)

type InvalidArgumentError = arg.InvalidArgumentError

// interface
type (
	Validatable interface {
		Validate() error
	}
)

// function
type (
	ContentProcessor = internal.ContentProcessor
	ErrorHandler     func(err error)
)

var (
	errorHandlerOnce sync.Once
	errorHandler     ErrorHandler

	stdErrorHandler = func(err error) { panic(err) }

	canonicalContentProcessors = map[string]ContentProcessor{
		"application/octet-stream":          nil,
		"text/plain":                        nil,
		"application/x-www-form-urlencoded": form.Process,
		"application/json":                  json.Process,
	}
)

func init() {
	internal.ContentProcessRegistryService.Setup(
		querystring.Process,
		canonicalContentProcessors,
	)
}

func initErrorHandler(fn ErrorHandler) {
	errorHandlerOnce.Do(func() {
		if fn != nil {
			errorHandler = fn
		}
	})
}
