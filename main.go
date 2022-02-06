package main

import (
	"github.com/yasir2000/my-hexagonal-go/internal/adaptors/core/arithmatic"
	"github.com/yasir2000/my-hexagonal-go/ports"
)

func main() {

	var core ports.ArithmaticPort

	core = arithmatic.NewAdaptor()
	// creation of type adaptor which has access to all methods
	//arithAdaptor := arithmatic.NewAdaptor()

}
