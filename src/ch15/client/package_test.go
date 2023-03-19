package client

import (
	"ch15/series"
	"testing"
)

func TestPackage(t *testing.T) {
	fib := series.GetFibonaicSerie(10)
	t.Log(fib)

	squ := series.Square(100)
	t.Log(squ)
}
