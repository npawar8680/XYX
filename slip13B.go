package main
import (
	"testing"
	)
func Square(x int) int{
return x * x
}
func BenchmarkSquare(b*testing.B){
for i:=0;i<b.N;i++{
Square(10) 
}
}

