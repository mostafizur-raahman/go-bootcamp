package main

import (
	"fmt"

	"github.com/mostafizur-raahman/go-bootcamp/auth"
)

func main() {
	auth.LoginWithCredientials("Mostafiz", "1234567")
	session := auth.GetSession()

	fmt.Println(session)
}
