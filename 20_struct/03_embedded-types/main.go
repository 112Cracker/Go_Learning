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

	fmt.Printf("%s is a %d-year-old %s who worked in %s as a %s.\n",
				person.name, person.age, person.gender, person.job.location, person.job.title)
}
/*
(1) overriding fields in Go -> promotion!
for example, the above employee rewrite as follows:

type employee struct {
	job job
	name string
	age int
	gender string
	location string
}

person := employee {
	job: job {
		title: "Manager"
		id: 100
		location: "Pittsburgh, PA"
	},
	name: "Tom"
	age: 24,
	gender: "Male"
	location: "Cleveland"
}
location value in the outer scope override the inner scope location value

(2) overriding methods in Go -> promotion!

job could have a method called A
employee could have a method called A as a promotion(overriding A belongs to job) to A belongs to job

 */