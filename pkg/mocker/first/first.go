package first

import (
	"github.com/ctx42/tst-a/pkg/mocker/second"
)

type Medium interface {
	Method0() error
	Method1(a *second.Second) error
}
