# ReaddleGoTask
Implemented Go application for Readdle Internship

## Statement

Statement of tasks:

**Test task 1 (HTTP, APIs, time)**

Use 3rd-party JSON API: https://date.nager.at/PublicHoliday/Country/UA
Write a console application that prints if it’s a holiday today (and the name of it). If today isn’t a holiday, the application should print the next closest holiday. 

Additionally, if the holiday is adjacent to a weekend (so that amount of non-working days is extended), the application should print this information. I.e. the next holiday is May 1, Friday, and it’s adjacent to Saturday (May 2) and Sunday (May 3), so the application should print something like: “The next holiday is International Workers' Day, May 1, and the weekend will last 3 days: May 1 - May 3”.

P.S. A candidate is expected to calculate long weekends manually, without using any other 3rd-party API, except the one with national holidays.

**Test task 2 (MySQL)**

Download and install the Employee sample database (https://dev.mysql.com/doc/employee/en/employees-installation.html).

Structure: https://dev.mysql.com/doc/employee/en/sakila-st ructure.html.

Create queries:
Find all current managers of each department and display his/her title, first name, last name, current salary.
Find all employees (department, title, first name, last name, hire date, how many years they have been working) to congratulate them on their hire anniversary this month.
Find all departments, their current employee count, their current sum salary.

## Pre-Installations
1. Add GO MySQL Driver driver to $GOPATH
``` bash
go get github.com/go-sql-driver/mysql
```
2. Install sample database https://github.com/datacharmer/test_db
``` bash
mysql -t < employees.sql
```

## Important details of implementation
1. Default user and password for db: root ''
2. Second query in task2 has to big output so this query was **limitted to 50 ROWS**. 
3. Third query in task2 takes in count only current employees because due to the statement we must display their title.

## Usage
``` go
import (
	"InternshipTask/task1"
	"InternshipTask/task2"
)

func main() {
	task1.Work()
	task2.Work()
}
  
```


