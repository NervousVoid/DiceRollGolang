package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var (
	seed  = flag.Int("seed", 0, "seed for random generator. unixNano(now) be default")
	start = flag.Int("start", 1, "begining of dice sides amount")
	end   = flag.Int("end", 6, "ending of dice sides amount")
	n     = flag.Int("n", 1, "amount of rolls")
	norepeat = flag.Bool("norepeat", false, "no repeating numbers in ouput")
)

func randInterval(l, r int) int {
	return rand.Intn(r+1) + l
}

func main() {
	flag.Parse()

	if *start > *end {
		fmt.Println("Staring amount of sides can`t be more than ending size!")
	} else {
		if *seed == 0 {
			rand.Seed(time.Now().UnixNano())
		} else {
			rand.Seed(int64(*seed))
		}
		for i := 1; i <= *n; i++ {
			fmt.Println(randInterval(*start, *end))
		}
	}
}
