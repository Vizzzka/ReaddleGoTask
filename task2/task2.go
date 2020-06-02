package task2

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type manager struct {
	title     string
	firstName string
	lastName  string
	salary    int
}

type employee struct {
	departmentName string
	title          string
	firstName      string
	lastName       string
	hireDate       string
	yearsOfWork    int
}

type department struct {
	departmentName string
	employeeCount  int
	sumSalary      int
}

func getManagers() []manager {
	rows, err := db.Query("SELECT employees.titles.title, employees.employees.first_name, employees.employees.last_name, employees.salaries.salary " +
		"FROM ((((employees.departments " +
		"INNER JOIN employees.dept_manager ON employees.dept_manager.dept_no = employees.departments.dept_no)" +
		"INNER JOIN employees.employees ON employees.employees.emp_no = employees.dept_manager.emp_no)" +
		"INNER JOIN employees.salaries ON employees.salaries.emp_no = employees.employees.emp_no)" +
		"INNER JOIN employees.titles ON employees.titles.emp_no = employees.salaries.emp_no)" +
		"WHERE employees.salaries.to_date > CURDATE() AND employees.titles.to_date > curdate() AND employees.dept_manager.to_date > curdate();")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	managers := []manager{}

	for rows.Next() {
		m := manager{}
		err := rows.Scan(&m.title, &m.firstName, &m.lastName, &m.salary)
		if err != nil {
			fmt.Println(err)
			continue
		}
		managers = append(managers, m)
	}
	return managers
}

func getEmployees() []employee {
	rows, err := db.Query("SELECT employees.departments.dept_name, employees.titles.title," +
		" employees.employees.first_name, employees.employees.last_name, employees.employees.hire_date," +
		" YEAR(curdate()) - YEAR(employees.employees.hire_date) " +
		"FROM (((employees.departments " +
		"INNER JOIN employees.dept_emp ON employees.dept_emp.dept_no = employees.departments.dept_no) " +
		"INNER JOIN employees.employees ON employees.employees.emp_no = employees.dept_emp.emp_no) " +
		"INNER JOIN employees.titles ON employees.titles.emp_no = employees.employees.emp_no) " +
		"WHERE employees.dept_emp.to_date > CURDATE() AND employees.titles.to_date > curdate() AND MONTH(curdate()) = MONTH(employees.employees.hire_date)" +
		"LIMIT 50;")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	employees := []employee{}

	for rows.Next() {
		emp := employee{}
		err := rows.Scan(&emp.departmentName, &emp.title, &emp.firstName, &emp.lastName,
			&emp.hireDate, &emp.yearsOfWork)
		if err != nil {
			fmt.Println(err)
			continue
		}
		employees = append(employees, emp)
	}
	return employees
}

func getDepartments() []department {
	rows, err := db.Query("SELECT employees.departments.dept_name, COUNT(employees.employees.emp_no), SUM(employees.salaries.salary) " +
		"FROM (((employees.departments " +
		"INNER JOIN employees.dept_emp ON employees.departments.dept_no = employees.dept_emp.dept_no) " +
		"INNER JOIN employees.employees ON employees.employees.emp_no = employees.dept_emp.emp_no) " +
		"INNER JOIN employees.salaries ON employees.employees.emp_no = employees.salaries.emp_no) " +
		"WHERE employees.dept_emp.to_date > CURDATE() AND employees.salaries.to_date > curdate() " +
		"GROUP BY employees.departments.dept_no;")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	departments := []department{}

	for rows.Next() {
		d := department{}
		err := rows.Scan(&d.departmentName, &d.employeeCount, &d.sumSalary)
		if err != nil {
			fmt.Println(err)
			continue
		}
		departments = append(departments, d)
	}
	return departments
}

var db *sql.DB

func Work() {
	var err error
	db, err = sql.Open("mysql", "root:@/employees")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	managers := getManagers()
	employees := getEmployees()
	departments := getDepartments()

	fmt.Println("All current managers of each department " +
		"( his/her title, first name, last name, current salary.")
	for _, m := range managers {
		fmt.Println(m.title, m.firstName, m.lastName, m.salary)
	}
	fmt.Println("\nAll employees (department, title, first name, last name," +
		" hire date, how many years they have been working)\n" +
		" to congratulate them on their hire anniversary this month " +
		"-- which means that this month is a month of their hire date")
	for _, emp := range employees {
		fmt.Println(emp.departmentName, emp.title, emp.firstName, emp.lastName, emp.hireDate, emp.yearsOfWork)
	}
	fmt.Println("\nAll departments, their current employee count, their current sum salary")
	for _, d := range departments {
		fmt.Println(d.departmentName, d.employeeCount, d.sumSalary)
	}

}
