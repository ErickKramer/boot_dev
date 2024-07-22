package main

import "fmt"

type car struct {
	make  string
	model string
}

type truck struct {
	car
	bedsize int
}

type user struct {
	name   string
	number int
}

type sender struct {
	user
	rateLimit int
}

func main() {
	lanesTruck := truck{
		bedsize: 10,
		car: car{
			make:  "toyota",
			model: "camry",
		},
	}

	fmt.Println(lanesTruck.bedsize)
	fmt.Println(lanesTruck.make)
	fmt.Println(lanesTruck.model)

	dummySender := sender{
		rateLimit: 10,
		user:      user{name: "Erick", number: 20},
	}
	fmt.Println(dummySender.name)
	fmt.Println(dummySender.rateLimit)
}
