package balance

import (
	"fmt"
)

type Rotatebalance struct {
	index int
}

func init() {
	rotatebalance := &Rotatebalance{}
	rotatebalance.index = 0
	RegisteBalance("rotate", rotatebalance)
}

func (r *Rotatebalance) Dobalance(n int) {
	fmt.Println("rotate is working")
	fmt.Println(r.index)
	r.index = (r.index + 1) % n
}
