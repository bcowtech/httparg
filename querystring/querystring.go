package querystring

import (
	"net/url"

	"github.com/bcowtech/httparg/internal"
	"github.com/bcowtech/structproto"
)

const (
	TagName = "query"
)

var _ internal.StringContentProcessor = Process

func Process(content string, target interface{}) error {
	values, err := url.ParseQuery(content)
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
