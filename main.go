package main

import (
	"balancer/balance"
	"fmt"
	"math/rand"
	"os"
)

func main() {

	var insts []*balance.Instance
	for i := 0; i < 5; i++ {
		host := fmt.Sprintf("192.168.%d.%d", rand.Intn(255), rand.Intn(255))
		wc := rand.Intn(10)
		one := balance.NewInstance(host, 8080, int64(wc))
		insts = append(insts, one)
	}

	var BalancerName = "weight_roundrobin"
	if len(os.Args) > 1 {
		BalancerName = os.Args[1]
	}

	for i := 0; i < 10000; i++ {
		_, err := balance.DoBalance(BalancerName, insts)
		if err != nil {
			fmt.Println("Do balance err:", err)
			continue
		}
	}

	for _, inst := range insts {
		fmt.Println(inst.GetResult())
	}

}
