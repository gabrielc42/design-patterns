// facade:
// encapsulates complex subsystem behind a simple interface
// can make changes to existing subsystem and not affect client
// car starting, startEngine() and stopEngine()

package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println()
	walletFacade := newWalletFacade("walit", 1337)
	fmt.Println()

	err := walletFacade.addMoneyToWallet("walit", 1337, 2)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	fmt.Println()
	err = walletFacade.deductMoneyFromWallet("walit", 1337, 1)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}
