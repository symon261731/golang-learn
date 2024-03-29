package notdoneexercise

import "fmt"

func CalcMoney(input []int) {
	sum := 0

	for _, value := range input {

		switch value {
		case 5:
			sum += value
		case 10:
			sum += 10
			sum -= 5
			if sum < 0 {
				fmt.Print(false)
				return
			}
			sum += 5
		case 20:
			sum += 20
			sum -= 15
			if sum < 0 {
				fmt.Print(false)
				return
			}
			sum += 5
		}

	}

	fmt.Print(true)
}

func CalcGroupsOfStudents() {
	var n, k, studentIndex int

	fmt.Println("Введите пожалуйста сколько студентов, групп, порядковый номер студента: ")
	fmt.Scan(&n, &k, &studentIndex)

	groupNum := studentIndex % k
	studentsInOneGroup := n / k

	for i := 1; i <= k; i++ {
		if groupNum == 0 {
			groupNum = k
		}
		if studentIndex <= studentsInOneGroup*i {
			fmt.Printf("студент находиться в %d группе", i)
		}
	}

	return
}
