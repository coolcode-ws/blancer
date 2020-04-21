package balance

import "strconv"

type Instance struct {
	Host      string
	Port      int64
	CallTimes int64
}

func NewInstance(host string, port int64) *Instance {
	return &Instance{
		Host: host,
		Port: port,
	}
}

func (i *Instance) GetPort() int64 {
	return i.Port
}

func (i *Instance) GetHost() string {
	return i.Host
}

func (i *Instance) GetCallTimes() string {
	return i.CallTimes
}

func (i *Instance) GetUrl() string {
	return i.Port + ":" + strconv.FormatInt(i.Port, 10)
}
