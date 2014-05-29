package gojsonld

import (
	"fmt"
)

type UniqueNamer struct {
	prefix   string
	counter  int
	existing map[string]string
}

func NewUniqueNamer(prefix string) *UniqueNamer {
	return &UniqueNamer{
		prefix:   prefix,
		existing: make(map[string]string),
	}
}

//TODO: sync/atomic counter
func (un *UniqueNamer) get(old string) string {
	if len(old) > 0 {
		if name := un.existing[old]; len(name) > 0 {
			return name
		}
	}

	name := fmt.Sprintf("%s%d", un.prefix, un.counter)
	un.counter += 1
	if len(old) > 0 {
		un.existing[old] = name
	}

	return name
}

func (un *UniqueNamer) contains(key string) bool {
	return len(un.existing[key]) > 0
}
