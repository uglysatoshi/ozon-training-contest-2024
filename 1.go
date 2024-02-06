package main

import "fmt"

func main1() {
	var N int
	_, err := fmt.Scan(&N)
	if err != nil {
		return
	}
	for i := 0; i < N; i++ {
		answers := make([]int, 4)
		boats := make([]int, 10)

		for j := 0; j < len(boats); j++ {
			_, err2 := fmt.Scan(&boats[j])
			if err2 != nil {
				return
			}
		}

		for k := 0; k < len(boats); k++ {
			switch boats[k] {
			case 1:
				{
					answers[0]++
				}
			case 2:
				{
					answers[1]++
				}
			case 3:
				{
					answers[2]++
				}
			case 4:
				{
					answers[3]++
				}
			}
		}

		if answers[0] == 4 && answers[1] == 3 && answers[2] == 2 && answers[3] == 1 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

	}
}
