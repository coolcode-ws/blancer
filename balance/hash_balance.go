package balance

import (
	"fmt"
	"hash/crc32"
	"math/rand"
)

// 一致性hash负载均衡
type HashBalance struct {
}

func init() {
	RegisterBalancer("hash", &HashBalance{})
}

func (p *HashBalance) DoBalance(insts []*Instance, key ...string) (*Instance, error) {
	var defKey string = fmt.Sprintf("%d", rand.Int())
	if len(key) > 0 {
		defKey = key[0]
	}

	lens := len(insts)
	if lens == 0 {
		return nil, fmt.Errorf("No backend instance")
	}

	crcTable := crc32.MakeTable(crc32.IEEE)
	hashVal := crc32.Checksum([]byte(defKey), crcTable)
	index := int(hashVal) % lens
	inst := insts[index]
	inst.CallTimes++

	return inst, nil
}
