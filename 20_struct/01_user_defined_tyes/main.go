package main

import "fmt"

type seconds float64

func main() {
	var time_to_run_100 seconds
	time_to_run_100 = 9.8
	fmt.Printf("My time to 100m is %0.2f s", time_to_run_100)
}
/*
user defined types -> aliased types
In the example above, we declared a type seconds whose underlying type is float64

Generally, it is a bad practice to alias types
But sometimes, programmers need to attach methods to a self-defined type
for example, Go has:
type Duration float64 // A Duration represents the elapsed time between two instants as an int64 nanoseconds count

has attached methods

func (d Duration) Hours() float64 // returns the duration as a floating point number of hours
func (d Duration) Minutes() float64 // returns the duration as a floating point number of minutes
func (d Duration) Seconds() float64 // returns the duration as a floating point number of seconds
 */
