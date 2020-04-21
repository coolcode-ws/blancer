package balance

import (
	"fmt"
)

//RoundRobinf负载均衡
type RoundRobinBalance struct {
	curIndex int
}

func init() {
	RegisterBalancer("roundrobin", &RoundRobinBalance{})
}

func (p *RoundRobinBalance) DoBalance(insts []*Instance, key ...string) (*Instance, error) {
	lens := len(insts)
	if lens == 0 {
		return nil, fmt.Errorf("No Instance found")
	}

	if p.curIndex >= lens {
		p.curIndex = 0
	}
	inst := insts[p.curIndex]

	p.curIndex = (p.curIndex + 1) % lens

	inst.CallTimes++

	return inst, nil
}
