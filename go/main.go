package main

import (
	"kew/cryptopal/test2"
)

func main() {
	// str := "email=foo@bar.com&uid=10&role=user"
	// enc, err := test2.ProfileFor(str)
	// if err != nil {
	// 	fmt.Print(err.Error())
	// }
	// fmt.Println(enc)
	test2.PerformAttack()
}
