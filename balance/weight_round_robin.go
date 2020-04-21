package balance

import (
	"fmt"
)

//WeightRoundRobinf负载均衡
type WeightRoundRobinBalance struct {
	Index  int64
	Weight int64
}

func init() {
	RegisterBalancer("weight_roundrobin", &WeightRoundRobinBalance{})
}

func (p *WeightRoundRobinBalance) DoBalance(insts []*Instance, key ...string) (*Instance, error) {
	lens := len(insts)
	if lens == 0 {
		return nil, fmt.Errorf("No Instance found")
	}

	inst := p.GetInst(insts)
	inst.CallTimes++

	return inst, nil
}

func (p *WeightRoundRobinBalance) GetInst(insts []*Instance) *Instance {
	gcd := getGCD(insts)
	for {
		p.Index = (p.Index + 1) % int64(len(insts))
		if p.Index == 0 {
			p.Weight = p.Weight - gcd
			if p.Weight <= 0 {
				p.Weight = getMaxWeight(insts)
				if p.Weight == 0 {
					return &Instance{}
				}
			}
		}

		if insts[p.Index].Weight >= p.Weight {
			return insts[p.Index]
		}
	}
}

//计算两个数的最大公约数
func gcd(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

//计算多个数的最大公约数
func getGCD(insts []*Instance) int64 {
	var weights []int64

	for _, inst := range insts {
		weights = append(weights, inst.Weight)
	}

	g := weights[0]
	for i := 1; i < len(weights)-1; i++ {
		oldGcd := g
		g = gcd(oldGcd, weights[i])
	}
	return g
}

//获取最大权重
func getMaxWeight(insts []*Instance) int64 {
	var max int64 = 0
	for _, inst := range insts {
		if inst.Weight >= int64(max) {
			max = inst.Weight
		}
	}

	return max
}
