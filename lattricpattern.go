package main

/*THIS IS SOLVED BY BIONOMIAL THEORM. THIS IMPLEMATION IS WRONG GO GET RIGHT VALUE SET K TO INPUT AND N TO *2 OR IN MY CASE IT WAS 40,20*/
import (
	"fmt"
	"math/big"
)

func lattricPath(n int) {

	k := n / 2
	r := new(big.Int).Binomial(int64(n*2), int64(k))
	fmt.Print("No =", r)

}
