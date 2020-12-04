package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("d1.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	r := bufio.NewReader(f)

	var sum, sum2 int
	var instructions = make(map[int]int)
	var line string

	i := 0
	for {
		line, err = r.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		if err != nil {
			break
		}
		instructions[i], _ = strconv.Atoi(line)
		i++
	}

	i = 0
	var n1, n2, n3 int
	for i = 0; i < len(instructions)-2; i++ {
		for j := i + 1; j < len(instructions)-1; j++ {
			for k := j + 1; k < len(instructions); k++ {
				n1 = instructions[i]
				n2 = instructions[j]
				n3 = instructions[k]

				if n1+n2 == 2020 {
					sum = n1 * n2
				}

				if n1+n2+n3 == 2020 {
					sum2 = n1 * n2 * n3
				}
			}
		}
	}
	fmt.Println("Answer 1:", sum)
	fmt.Println("Answer 2:", sum2)
}
