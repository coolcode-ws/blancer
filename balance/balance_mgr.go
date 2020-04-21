package balance

import "fmt"

var (
	mgr = BalanceMgr{
		allBalance: make(map[string]Balancer),
	}
)

type BalanceMgr struct {
	allBalance map[string]Balancer
}

func (p *BalanceMgr) registerBalancer(balanceType string, b Balancer) {
	p.allBalance[balanceType] = b
}

func RegisterBalancer(balanceType string, b Balancer) {
	mgr.registerBalancer(balanceType, b)
}

func DoBalance(balanceType string, insts []*Instance) (*Instance, error) {
	Balancer, ok := mgr.allBalance[balanceType]
	if !ok {
		fmt.Printf("Not found %s balancer", balanceType)
		return nil, fmt.Errorf("not fund balancer")
	}
	return Balancer.DoBalance(insts)
}
