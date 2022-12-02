package main

import (
	"fmt"
)

func get_elves_and_calories() ([]int, int) {
	var calories int
	elves := []int{}
	total_calories := 0
	largest := -1
	largest_calories := -1

	n, _ := fmt.Scanf("%d\n", &calories);
	for n >= 0 {
		if n == 0 {
			elves = append(elves, total_calories)
			if total_calories > largest_calories {
				largest = len(elves)-1
				largest_calories = total_calories
			}
			total_calories = 0
		} else if n > 0 {
			total_calories += calories
		}
		n, _ = fmt.Scanf("%d\n", &calories);
		if n == 0 && total_calories == 0 {
			break
		}
	}
	return elves, largest
}

func find_elf_with_most_calories(elves []int) (int, int) {
	largest_calories := -1
	largest_elf := -1
	for elf, cal := range elves {
		if cal > largest_calories {
			largest_calories = cal
			largest_elf = elf
		}
	}
	return largest_calories, largest_elf
}

func find_top_three_calories(elves []int) (int) {
	top3 := [3]int{-1,-1,-1}
	var cal int
	total_cal := 0

	cal, top3[0] = find_elf_with_most_calories(elves)
	total_cal += cal
	elves = append(elves[:top3[0]], elves[top3[0]+1:]...)
	cal, top3[1] = find_elf_with_most_calories(elves)
	total_cal += cal
	elves = append(elves[:top3[1]], elves[top3[1]+1:]...)
	cal, top3[2] = find_elf_with_most_calories(elves)
	total_cal += cal

	return total_cal
}

func main() {
	elves := []int{}
	elves, largest := get_elves_and_calories()

	fmt.Println(elves[largest], largest)
	// part 2
	total_calories := find_top_three_calories(elves[:])
	fmt.Println(total_calories)
}
