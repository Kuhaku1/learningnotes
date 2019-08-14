package main

import "fmt"

func main() {

	/*var skipDeploy bool

	for a := 0;a < 10; a++{
		fmt.Println(a)
		for b := 0; b < 10;b++{
			fmt.Println(b)
			if b == 2{
				skipDeploy = true
				break
			}
		}
		if skipDeploy {
			break
		}
	}
	fmt.Println("Done")*/

	for a := 0; a < 10; a++ {
		fmt.Println(a)
		for b := 0; b < 10; b++ {
			fmt.Println(b)
			if b == 2 {
				goto skipDeploy
			}
		}
	}
	return

skipDeploy:
	fmt.Println("Done")
}
