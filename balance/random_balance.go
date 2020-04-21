package balance

import (
	"errors"
	"math/rand"
)

//随机负载均衡
type RandomBalance struct {
}

func init() {
	RegisterBalancer("random", &RandomBalance{})
}

func (p *RandomBalance) DoBalance(insts []*Instance, key ...string) (*Instance, error) {
	lens := len(insts)
	if lens == 0 {
		return nil, errors.New("No Instance found")
	}

	index := rand.Intn(lens)
	inst := insts[index]
	inst.CallTimes++

	return inst, nil
}
