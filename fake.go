package idgen

import (
	"errors"
)

var ErrFakeIdGenOutOfIds = errors.New("all ids given out by fake id generator")

func NewFakeIdGen(ids []string) UUIDGenerator {
	return &fakeIdGen{ids, 0}
}

type fakeIdGen struct {
	ids []string
	x   int
}

func (f *fakeIdGen) NewId() (Id, error) {
	if f.x >= len(f.ids) {
		return "", ErrFakeIdGenOutOfIds
	}
	id := f.ids[f.x]
	f.x += 1
	return Id(id), nil
}
