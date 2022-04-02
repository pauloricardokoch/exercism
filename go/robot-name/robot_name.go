package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Robot struct {
	name string
}

var robots = make(map[string]Robot)
var uniqueNames = 26 * 26 * 10 * 10 * 10

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}

	name := fmt.Sprintf(
		"%c%c%03d",
		rand.Intn(26)+'A',
		rand.Intn(26)+'A',
		rand.Intn(1000),
	)

	if _, ok := robots[name]; ok {
		if len(robots) < uniqueNames {
			return r.Name()
		}
		return "", errors.New(fmt.Sprintf("All %d available names have been taken.", len(robots)))
	}

	r.name = name
	robots[name] = *r
	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = ""
	_, err := r.Name()
	if err != nil {
		panic(err.Error())
	}
}
