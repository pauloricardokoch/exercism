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

// genLetter returns a random letter from A-Z
func genLetter() string {
	return string(rune(65 + rand.Intn(90-65+1)))
}

// genNum returns a random number from 0-9
func genNum() string {
	return string(rune(48 + rand.Intn(57-48+1)))
}

func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}

	name := genLetter() + genLetter() + genNum() + genNum() + genNum()
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
