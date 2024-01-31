package main

import "fmt"

// Вы участвуете в разработке подсистемы проверки поля для игры <<Морской бой>>.
// Вам требуется написать проверку корректности количества кораблей на поле, учитывая их размеры.
// Напомним, что на поле должны быть: четыре однопалубных корабля, три двухпалубных корабля, два трёхпалубных корабля,
// один четырёхпалубный корабль. Вам заданы 10 целых чисел от 1 до 4.
// Проверьте, что заданные размеры соответствуют требованиям выше.

//5
//2 1 3 1 2 3 1 1 4 2
//1 1 1 2 2 2 3 3 3 4
//1 1 1 1 2 2 2 3 3 4
//4 3 3 2 2 2 1 1 1 1
//4 4 4 4 4 4 4 4 4 4

//YES
//NO
//YES
//YES
//NO

func main() {
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
