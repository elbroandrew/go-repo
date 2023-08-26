package main

import (
	"bufio"
	"fmt"
	"os"
	// "math"
)

/*
купил три, заплатил за два
7⋅2 + 5⋅3 надо будет оплатить 5⋅2 + 4⋅3 = 22.

7 / 3 = 2 floor
7 - 2 = 5
5*2

5 / 3 = 1 floor
5 - 1 = 4
4*3

6
12
2 2 2 2 2 2 2 3 3 3 3 3
12
2 3 2 3 2 2 3 2 3 2 2 3
1
10000
9
1 2 3 1 2 3 1 2 3
6
10000 10000 10000 10000 10000 10000
6
300 100 200 300 200 300

output:
22
22
10000
12
40000
1100

*/

func payment(n map[int]int) int {
	offPriceNumber := 3

	var sum int = 0

	for price, quantity := range n {
		var k int = quantity / offPriceNumber
		sum += (quantity - k) * price

	}

	return sum

}

func main() {

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)


	var testCount int
	fmt.Fscan(reader, &testCount)

	m := make(map[int]int)


	for i := 0; i < testCount; i++ {
		

		var goodsCount, price int
		fmt.Fscan(reader, &goodsCount)
		
		for j := 0; j < goodsCount; j++ {

			fmt.Fscan(reader, &price)
			if _, ok := m[price]; ok {
				m[price] += 1
			}else{
				m[price] = 1
			}
		}

		fmt.Printf("%#v\n", m)

		fmt.Fprintln(writer, payment(m))
	}

	writer.Flush()

}
