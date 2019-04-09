package main

import (
	"time"
	"fmt"
)

type Employee struct {
	ID			int
	Name		string
	Address		string
	DoB			time.Time
	Position	string
	Salary		int
	ManagerID	int
}

// 这个方法仅用于测试，没什么意义
func EmployeeByID(id int) *Employee  {
	return &Employee{
		ID:id,
	}
}

func main()  {
	var William Employee
	William.Salary -= 5000

	position := &William.Position
	*position = "Senior " + *position

	var employeeOfTheMonth *Employee = &William
	employeeOfTheMonth.Position += " (proactive team player)"
	// 以上两行代码相当于
	//(*employeeOfTheMonth).Position += " (proactive team player)"

	fmt.Println(William)

	fmt.Println(EmployeeByID(1).ID)
	id := William.ID
	EmployeeByID(id).Salary = 0
}
