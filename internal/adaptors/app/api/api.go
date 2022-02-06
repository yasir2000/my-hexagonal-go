package api

import (
	"github.com/yasir2000/my-hexagonal-go/internal/ports"
)

type Adaptor struct {
	arith ports.ArithmaticPort
}
