package multiBalance

import (
	"fmt"
	"strconv"
	"sync"
	"github.com/imroc/req"
	"github.com/button-tech/utils-node-tool/btc/handlers/storage"
)

type Data struct {
	sync.Mutex
	Result map[string]float64
}

func New() *Data {
	return &Data{
		Result: make(map[string]float64),
	}
}

func (ds *Data) set(key string, value float64) {
	ds.Result[key] = value
}

func (ds *Data) Set(key string, value float64) {
	ds.Lock()
	defer ds.Unlock()
	ds.set(key, value)
}

func Worker(wg *sync.WaitGroup, addr string, r *Data) {
	defer wg.Done()
	balance, err := req.Get(storage.BtcURL + "/insight-api/addr/" + addr + "/balance")
	if err != nil {
		fmt.Println(err)
		return
	}
	balanceFloat, _ := strconv.ParseFloat(balance.String(), 64)
	balanceFloat *= 0.00000001
	r.Set(addr, balanceFloat)
}