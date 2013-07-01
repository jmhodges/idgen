package idgen

import (
	"crypto/rand"
	"encoding/hex"
	"io"
)

const EmptyId = Id("")

var privateIdGen = NewIdGen()

type Id string

type UUIDGenerator interface {
	NewId() (Id, error)
}

func NewIdGen() UUIDGenerator {
	return &idGen{}
}

type idGen struct {}

// Generate an id suitable for inserting into a postgres uuid column. This
// means that it will not be simply 16 random bytes, but a hex-formatted
// version of the 16 byte slice with dashes, making it 36 bytes long.
func (ig idGen) NewId() (Id, error) {
	id := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, id)
	if err != nil {
		return EmptyId, err
	}
	dashedId := make([]byte, 36)
	hex.Encode(dashedId[0:8], id[0:4])
	dashedId[8] = '-'
	hex.Encode(dashedId[9:13], id[4:6])
	dashedId[13] = '-'
	hex.Encode(dashedId[14:18], id[6:8])
	dashedId[18] = '-'
	hex.Encode(dashedId[19:23], id[8:10])
	dashedId[23] = '-'
	hex.Encode(dashedId[24:36], id[10:16])
	return Id(string(dashedId)), nil
}

func NewId() Id {
	id, err := privateIdGen.NewId()
	if err != nil {
		panic("/dev/urandom is broken, process is untrustworthy, uuid gen failed")
	}
	return id
}
