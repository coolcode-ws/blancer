package balance

import "fmt"

//初始化一个实例
var (
	mgr = BalanceMgr{
		allBalance: make(map[string]Balancer),
	}
)

//管理负载均衡算法，在random roundrobin 算法中可以注册
type BalanceMgr struct {
	allBalance map[string]Balancer
}

func (p *BalanceMgr) registerBalancer(name string, b Balancer) {
	p.allBalance[name] = b
}

func RegisterBalancer(name string, b Balancer) {
	mgr.registerBalancer(name, b)
}
func DoBalance(name string, insts []*Instance) (inst *Instance, err error) {
	//从map里面找有没有这个调度算法
	Balancer, ok := mgr.allBalance[name]
	if !ok {
		fmt.Errorf("Not found %s balancer", name)
		return
	}
	fmt.Printf("use %s balancer\n", name)
	inst, err = Balancer.DoBalance(insts)
	return
}
