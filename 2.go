package main

import "fmt"

func main2() {
	var N, d, m, y int
	_, err := fmt.Scan(&N)
	if err != nil {
		return
	}

	monthValues := make(map[int]int)

	monthValues[1] = 31
	monthValues[3] = 31
	monthValues[4] = 30
	monthValues[5] = 31
	monthValues[6] = 30
	monthValues[7] = 31
	monthValues[8] = 31
	monthValues[9] = 30
	monthValues[10] = 31
	monthValues[11] = 30
	monthValues[12] = 31

	for i := 0; i < N; i++ {
		_, err2 := fmt.Scan(&d, &m, &y)
		if err2 != nil {
			return
		}
		if (1 <= d && d <= 31) && (1 <= m && m <= 12) && (1950 <= y && y <= 2300) {
			if y%400 == 0 || (y%4 == 0 && y%100 != 0) {
				monthValues[2] = 29
				if d <= monthValues[m] {
					fmt.Println("YES")
				} else {
					fmt.Println("NO")
				}
			} else {
				monthValues[2] = 28
				if d <= monthValues[m] {
					fmt.Println("YES")
				} else {
					fmt.Println("NO")
				}
			}
		} else {
			fmt.Println("NO")
		}

	}
}
