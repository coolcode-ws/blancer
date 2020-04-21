package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	//初始化主机
	//insts := make([]*balance.Instance, 16) //append会自动扩容
	var insts []*balance.Instance
	for i := 0; i < 16; i++ {
		host := fmt.Sprintf("192.168.%d.%d", rand.Intn(255), rand.Intn(255))
		one := balance.NewInstance(host, 8080)
		insts = append(insts, one)
	}
	//var balancer balance.Balancer
	var BalancerName = "random"
	if len(os.Args) > 1 {
		BalancerName = os.Args[1]
	}
	for {
		inst, err := balance.DoBalance(BalancerName, insts)
		if err != nil {
			fmt.Println("Do balance err:", err)
			continue
		}
		fmt.Println(inst)
		time.Sleep(time.Second)
	}

}
