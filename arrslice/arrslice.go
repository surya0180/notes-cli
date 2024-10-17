package arrslice

import "fmt"

func Execute() {
	// hobbies := [3]string{"Dancing", "Singing", "Swimming"}

	// fmt.Println(hobbies)
	// fmt.Println(hobbies[0])

	// slice1 := hobbies[1:] // Or 1:3
	// fmt.Println(slice1)

	// slice2 := hobbies[0:2]
	// fmt.Println(slice2)

	// slice3 := hobbies[:2]
	// fmt.Println(slice3)

	// slice4 := slice2[1:3]
	// fmt.Println(slice4)

	// goals := []string{"Goal1", "Goal2", "Goal3", "Goal4", "Goal5"}
	// goals[1] = "Goal Different"
	// goals = append(goals, "Goal6")
	// fmt.Println(goals)

	// products := []product{{title: "title1", id: 1, price: 5.6}, {title: "title2", id: 2, price: 8.2}}
	// products = append(products, product{title: "title3", id: 3, price: 4.89})
	// fmt.Println(products)

	exercises := map[string]string{
		"name1": "value1",
		"name2": "value2",
	}
	fmt.Println(exercises)

	for key, value := range exercises {
		fmt.Println(key)
		fmt.Println(value)
	}
}
