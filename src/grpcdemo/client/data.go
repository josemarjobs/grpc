package main

import (
	"grpcdemo/pb"
)

var newEmployees = []pb.Employee{
	pb.Employee{
		BadgeNumber:         123,
		FirstName:           "Peter",
		LastName:            "griffin",
		VacationAccrualRate: 2,
		VacationAccrued:     30,
	},
	pb.Employee{
		BadgeNumber:         234,
		FirstName:           "Joe",
		LastName:            "Swanson",
		VacationAccrualRate: 2.3,
		VacationAccrued:     23.4,
	},
	pb.Employee{
		BadgeNumber:         524,
		FirstName:           "Glen",
		LastName:            "Quagmire",
		VacationAccrualRate: 3,
		VacationAccrued:     31.7,
	},
	pb.Employee{
		BadgeNumber:         5134,
		FirstName:           "Cleveland",
		LastName:            "Brown",
		VacationAccrualRate: 3,
		VacationAccrued:     31.7,
	},
}
