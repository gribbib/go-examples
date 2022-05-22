package main

import (
	"fmt"
	"math"
)
import "log"
import "example.com/greetings"
import "golang.org/x/tour/pic"

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Jess")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)

	messages, err := greetings.Hellos([]string {"Matt", "Teddy", "Eule"})
	if err != nil {
		log.Fatal(err)
	}

	for _, message := range messages {
		fmt.Println(message)
	}

	fmt.Println(Sqrt(123456789))
	fmt.Println(math.Sqrt(123456789))

	pic.Show(Pic)
}

func Sqrt(x float64) float64 {
	z := x/2
	i := 0
	for {
		fmt.Println("z: ", z)
		fmt.Println("iteration: ", i+1)
		oldZ := z
		z -= (z*z - x) / (2 * z)
		if oldZ == z || roundTo(oldZ-z, 10) == 0 {
			break
		}
		i++
	}
	return z
}

func roundTo(n float64, decimals int) float64 {
	return math.Round(n*math.Pow10(decimals)) / math.Pow10(decimals)
}

func Pic(dx, dy int) [][]uint8 {
	var pic [][]uint8
	for i := 0; i < dy; i++ {
		var s []uint8
		/*if i%2 == 0 {
			for j := 0; j < dx; j++ {
				s = append(s, uint8(j))
			}
		} else {
			for j := dx - 1; j >= 0; j-- {
				s = append(s, uint8(j))
			}
		}*/
		
		for j := 0; j < dx; j++ {
			//s = append(s, uint8(j^i))
			s = append(s, uint8(j*i))
		}
		pic = append(pic, s)
	}
	return pic
}

func fibonacci() func() int {
	fibA, fibB := 0, -1
	return func() int {
		if fibB < 1 {
			fibB++
		} else {
			old := fibB
			fibB += fibA
			fibA = old
		}
		return fibB
	}
}
