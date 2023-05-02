package main

import "fmt"

//input
//siska stevani 12345
//2
//100 80
//
//output
//Name: stevani, siska
//ID: 12345
//Grade: O

type Person struct {
	firstname, lastname string
	identification      int
}

type Student struct {
	Person
	testScores []int
}

func (p *Person) printPerson() {
	fmt.Printf("Name: %v, %v\nID: %v\n", p.lastname, p.firstname, p.identification)
}

func (s *Student) calculate() string {
	total := 0
	avg := 0
	for i := 0; i < len(s.testScores); i++ {
		total += s.testScores[i]
	}
	avg = total / len(s.testScores)
	switch {
	case avg <= 100 && avg >= 90:
		return "0"
	case avg < 90 && avg >= 80:
		return "E"
	case avg < 80 && avg >= 70:
		return "A"
	case avg < 70 && avg >= 55:
		return "P"
	case avg < 55 && avg >= 40:
		return "D"
	case avg < 40:
		return "T"
	default:
		return ""
	}
}

func main() {
	var (
		firstname      string
		lastname       string
		identification int
		numScores      int
		testScores     []int
	)

	fmt.Scanln(&firstname, &lastname, &identification)
	fmt.Scanln(&numScores)

	for i := 0; i < numScores; i++ {
		var tempScores int
		fmt.Scanf("%d", &tempScores)
		testScores = append(testScores, tempScores)
	}

	student := Student{
		Person{firstname: firstname, lastname: lastname, identification: identification}, testScores,
	}
	fmt.Println(testScores)
	student.printPerson()
	fmt.Println("Grade:", student.calculate())

}
