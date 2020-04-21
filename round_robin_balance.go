package balance

import "errors"

type RoundRobinBalance struct {
	curIndex int
}

func init() {
	RegisterBalancer("roundrobin", &RoundRobinBalance{})
}
func (p *RoundRobinBalance) DoBalance(insts []*Instance, key ...string) (inst *Instance, err error) {
	lens := len(insts)
	if lens == 0 {
		err = errors.New("No Instance")
		return
	}
	if p.curIndex >= lens {
		p.curIndex = 0
	}
	inst = insts[p.curIndex]
	//p.curIndex++
	p.curIndex = (p.curIndex + 1) % lens
	return
}
