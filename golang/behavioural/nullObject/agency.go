package main

import "fmt"

func main() {
	college1 := createCollege1()
	college2 := createCollege2()
	totalProfessors := 0
	departmentArray := []string{"computerscience", "painting", "electricalengineering", "literature", "physics"}

	for _, departmentName := range departmentArray {
		d := college1.getDepartment(departmentName)
		totalProfessors += d.getNumOfProfessors()
	}

	fmt.Printf("Total number of professors in college1 is %d\n", totalProfessors)

	//reset prof count

	totalProfessors = 0
	for _, deparmentName := range departmentArray {
		d := college2.getDepartment(deparmentName)
		totalProfessors += d.getNumOfProfessors()
	}
	fmt.Printf("Total number of professors in college2 is %d\n", totalProfessors)
}

func createCollege1() *college {
	college := &college{}
	college.addDepartment("computerscience", 4)
	college.addDepartment("literature", 5)
	return college
}

func createCollege2() *college {
	college := &college{}
	college.addDepartment("painting", 2)
	return college
}
