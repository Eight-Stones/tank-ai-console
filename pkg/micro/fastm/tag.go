package fastm

import (
	"errors"
	"strings"
)

type variableLocation int

const (
	query variableLocation = iota
	path
)

const (
	tagLen = 2
)

var (
	ErrInvalidTag = errors.New("unable to parse a tag")
)

type tag struct {
	name     string
	location variableLocation
}

func newTag(tagValue string) (*tag, error) {
	this := &tag{}

	vals := strings.Split(tagValue, ",")
	if len(vals) != tagLen {
		return nil, ErrInvalidTag
	}

	this.name = vals[0]
	this.location = map[string]variableLocation{
		"query": query,
		"path":  path,
	}[vals[1]]

	return this, nil
}

func (t tag) Name() string {
	return t.name
}

func (t tag) Location() variableLocation {
	return t.location
}
