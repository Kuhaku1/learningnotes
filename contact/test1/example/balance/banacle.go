package balance

import (
	"fmt"
)

type Balancein interface {
	Dobalance(int)
}

type Balance struct {
	name map[string]Balancein
}
var balancestu *Balance

func init() {
	balancestu = &Balance{}
	balancestu.name = make(map[string]Balancein)
}

func RegisteBalance(name string, a Balancein) {
	// balancestu
	balancestu.name[name] = a
}

func Dobalance(name string, n int) {
	balance, err:= balancestu.name[name]
	if err != true {
		fmt.Println("未注册", err)
		return
	}
	balance.Dobalance(n)
}
