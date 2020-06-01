package task2

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type manager struct {
	id        int
	title     string
	firstName string
	lastName  string
	salary    int
}

func Work() {
	db, err := sql.Open("mysql", "root:@/employees")

	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT employees.titles.title, employees.employees.first_name, employees.employees.last_name, employees.salaries.salary " +
		"FROM ((((employees.departments " +
		"INNER JOIN employees.dept_manager ON employees.dept_manager.dept_no = employees.departments.dept_no)" +
		"INNER JOIN employees.employees ON employees.employees.emp_no = employees.dept_manager.emp_no)" +
		"INNER JOIN employees.salaries ON employees.salaries.emp_no = employees.employees.emp_no)" +
		"INNER JOIN employees.titles ON employees.titles.emp_no = employees.salaries.emp_no)" +
		"WHERE employees.salaries.to_date > CURDATE() AND employees.titles.to_date > curdate();")
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
	fmt.Println("All current managers of each department and  his/her title\n" +
		"(I think it was ment to display current title because otherwise it has no sense),\n" +
		"first name, last name, current salary.")
	for _, m := range managers {
		fmt.Println(m.title, m.firstName, m.lastName, m.salary)
	}

}
