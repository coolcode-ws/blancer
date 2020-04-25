package balance

import (
	"errors"
	"math/rand"
	"time"
)

//洗牌负载均衡
type Shuffle2Balance struct {
}

func init() {
	RegisterBalancer("shuffle2", &Shuffle2Balance{})
}

func (p *Shuffle2Balance) DoBalance(insts []*Instance, key ...string) (*Instance, error) {
	lens := len(insts)
	if lens == 0 {
		return nil, errors.New("No Instance found")
	}

	rand.Seed(time.Now().UnixNano())

	//shuffle2
	for i := lens; i > 0; i-- {
		lastIdx := i - 1
		idx := rand.Intn(i)
		insts[lastIdx], insts[idx] = insts[idx], insts[lastIdx]
	}

	inst := insts[0]
	inst.CallTimes++

	return inst, nil
}
