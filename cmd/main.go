package main

import (
	"fmt"

	"github.com/zavolokas/oura-data/ouraring"
)

const OuraToken = ""

func main() {

	client := ouraring.NewClient(OuraToken)
	da, _ := client.GetDailyActivity()
	fmt.Print(da)
}
