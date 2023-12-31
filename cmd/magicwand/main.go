package main

import (
	"fmt"
	"github.com/meriy100/magicwand/controllers"
	"os"
)

func main() {
	if err := controllers.NewController().Run(); err != nil {
		fmt.Printf("%+v", err.Error())
		os.Exit(1)
	}
}
