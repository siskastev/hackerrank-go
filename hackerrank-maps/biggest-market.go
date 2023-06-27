package main

import "fmt"

// find country the most

func main() {
	user := []User{
		{
			"siska",
			"malang",
		},
		{
			"stevani",
			"malang",
		},
		{
			"ika",
			"jakarta",
		},
		{
			"kaka",
			"surabaya",
		},
	}

	country, count := calculate(user)
	fmt.Printf("the biggest user in country %s with user %d", country, count)

}

type User struct {
	name, country string
}

func calculate(users []User) (string, int) {
	counts := make(map[string]int)
	for _, v := range users {
		counts[v.country]++
	}
	maxCountry := ""
	maxCount := 0
	for country, count := range counts {
		if count > maxCount {
			maxCount = count
			maxCountry = country
		}
	}
	return maxCountry, maxCount
}
