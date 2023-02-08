package main

type college struct {
	departments []department
}

func (c *college) addDepartment(departmentName string, numOfProfessors int) {
	if departmentName == "computerscience" {
		computerScienceDepartment := &computerscience{numOfProfessors: numOfProfessors}
		c.departments = append(c.departments, computerScienceDepartment)
	}
	if departmentName == "literature" {
		literatureDepartment := &literature{numOfProfessors: numOfProfessors}
		c.departments = append(c.departments, literatureDepartment)
	}
}

func (c *college) getDepartment(departmentName string) department {
	for _, department := range c.departments {
		if department.getName() == departmentName {
			return department
		}
	}
	//Return null department if the department doesn't exist
	return &nullDepartment{}
}
