package main

import "fmt"

type job struct {
	title string
	id int
	location string
}

type employee struct {
	job job // job field for employee is a job type
	name string
	age int
	gender string
}

func main() {
	person := employee{
		job: job{
			title: "Manager",
			id: 100,
			location: "Pittsburgh, PA",
		},
		name: "Tom",
		age: 24,
		gender: "Male",
	}

	fmt.Printf("%s is a %d-year-old %s who worked in %s as a %s",
				person.name, person.age, person.gender, person.job.location, person.job.title)
}