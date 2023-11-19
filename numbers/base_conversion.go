/* Perform base conversion on an integer */

package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println(big.NewInt(1000000000000).Text(2))
	fmt.Println(big.NewInt(3777776).Text(10))
	fmt.Println(big.NewInt(123456789).Text(8))
}
