package form

import (
	"net/url"

	"github.com/bcowtech/httparg/internal"
	"github.com/bcowtech/structproto"
)

const (
	TagName = "form"
	True    = "true"
)

var _ internal.ContentProcessor = Process

func Process(content []byte, target interface{}) error {
	values, err := url.ParseQuery(string(content))
	if err != nil {
		return err
	}

	provider := internal.NewQueryArgsBinder(values)

	prototype, err := structproto.Prototypify(target,
		&structproto.StructProtoOption{
			TagName: TagName,
		})
	if err != nil {
		return err
	}

	return prototype.Bind(provider)
}
