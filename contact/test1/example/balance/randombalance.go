package balance

import (
	"time"
	"fmt"
	"math/rand"
)

type RandomBalance struct{

}

func init() {
	rand.Seed(time.Now().Unix())
	randombalance := &RandomBalance{}
	RegisteBalance("random", randombalance)
}

func (r *RandomBalance)Dobalance(n int) {
	fmt.Println("random is working")
	fmt.Println(rand.Intn(n))
}