package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var (
	seed     = flag.Int("seed", 0, "seed for random generator. unixNano(now) be default")
	start    = flag.Int("start", 1, "begining of dice sides amount")
	end      = flag.Int("end", 6, "ending of dice sides amount")
	n        = flag.Int("n", 1, "amount of rolls")
	norepeat = flag.Bool("norepeat", false, "no repeating numbers in output")
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

		//Создаётся срез со всеми целыми числами от L до R,
		//и с каждой итерацией из него достаётся элемент по случайному индексу
		//и удаляется, чтобы больше не попасться
		if *norepeat {
			len := *end - *start
			var x []int

			if len < *n {
				fmt.Println()
				fmt.Printf("You can`t generate %d unrepeated numbers in range from %d to %d!", *n, *start, *end)
				os.Exit(0)
			}

			for i := 0; i <= len; i++ {
				x = append(x, *start)
				*start++
			}

			for i := 0; i < *n; i++ {
				randInd := randInterval(0, len)
				fmt.Println(x[randInd])
				x = append(x[:randInd], x[randInd+1:]...)
				len--
			}
		} else {
			for i := 1; i <= *n; i++ {
				fmt.Println(randInterval(*start, *end))
			}
		}
	}
}
