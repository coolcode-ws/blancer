package balance

import (
	"errors"
	"math/rand"
	"time"
)

//洗牌负载均衡
type ShuffleBalance struct {
}

func init() {
	RegisterBalancer("shuffle", &ShuffleBalance{})
}

func (p *ShuffleBalance) DoBalance(insts []*Instance, key ...string) (*Instance, error) {
	lens := len(insts)
	if lens == 0 {
		return nil, errors.New("No Instance found")
	}

	rand.Seed(time.Now().UnixNano())

	//shuffle
	for i := 0; i < lens/2; i++ {
		a := rand.Intn(lens)
		b := rand.Intn(lens)
		insts[a], insts[b] = insts[b], insts[a]
	}

	inst := insts[0]
	inst.CallTimes++

	return inst, nil
}
