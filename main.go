package main

import (
	"fmt"
	"time"

	"github.com/wagslane/go-tinytime"
)

func main() {
	tt := tinytime.New(1585758374)
	tt = tt.Add(time.Hour * 48)
	fmt.Println(tt)
}
