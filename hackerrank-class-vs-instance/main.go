package main
import "fmt"


/*
Task
Write a Person class with an instance variable,age,
and a constructor that takes an integer, initialAge, as a parameter.
The constructor must assign  to  after confirming the argument passed as initialAge is not negative;
if a negative argument is passed as initialAge, the constructor should set age to 0
and print Age is not valid, setting age to 0.. In addition, you must write the following instance methods:

yearPasses() should increase the age instance variable by 1.
amIOld() should perform the following conditional actions:
If age < 13, print You are young..
If age >= 13 and < 18, print You are a teenager..
Otherwise, print You are old.

Sample input
4
-1
10
16
18

output
Age is not valid, setting age to 0.
You are young.
You are young.

You are young.
You are a teenager.

You are a teenager.
You are old.

You are old.
You are old.
 */

type person struct {
	age int
}

func (p person) NewPerson(initialAge int) person {
	//Add some more code to run some checks on initialAge
	if initialAge < 0{
		fmt.Println("Age is not valid, setting age to 0.")
		p.age = 0
	}
	return p
}

func (p person) amIOld() {
	//Do some computation in here and print out the correct statement to the console
	if p.age < 13{
		fmt.Println("You are young.")
	}else if p.age >= 13 && p.age < 18 {
		fmt.Println("You are a teenager.")
	}else{
		fmt.Println("You are old.")
	}

}

func (p person) yearPasses() person {
	//Increment the age of the person in here
	p.age++
	return p
}

func main() {
	var T,age int

	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		fmt.Scan(&age)
		p := person{age: age}
		p = p.NewPerson(age)
		p.amIOld()
		for j := 0; j < 3; j++ {
			p = p.yearPasses()
		}
		p.amIOld()
		fmt.Println()
	}
}