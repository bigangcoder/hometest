package main

import (
	"Testprj03/hometest.globalgroup.com/backend"
	"fmt"
	"time"
)

func init() {
}
func main() {
	fmt.Println("started", time.Now().Format("2006-01-02 15:04:05.000000"))
	backend.DoMainCtl()
}
