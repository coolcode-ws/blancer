package balance

import (
	"errors"
	"math/rand"
)

type RandomBalance struct {
}

func init() {
	RegisterBalancer("random", &RandomBalance{})
}

func (p *RandomBalance) DoBalance(insts []*Instance, key ...string) (inst *Instance, err error) {
	lens := len(insts)
	if lens == 0 {
		err = errors.New("No Instance")
		return
	}
	index := rand.Intn(lens)
	inst = insts[index]
	return
}
