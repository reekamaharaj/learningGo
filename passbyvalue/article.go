//code from pass by value article, link in notes

// int
func modifyInt(n int) int {
	return n + 5
}age := 30 fmt.Println("Before function call: ", age)   // 30 fmt.Println("Function call:", modifyInt(age))           // 35
fmt.Println("After function call: ", age)               // 30

//# float
func modifyFloat(n float64) float64 {
	return n + 5.0
}
cash := 10.50
fmt.Println("Before function call: ", cash)             // 10.5
fmt.Println("Function call:", modifyFloat(cash))        // 15.5 fmt.Println("After function call: ", cash)              // 10.5

//# bool
func modifyBool(n bool) bool {
	return !n
}
old := false
f mt.Println("Before function call: ", old)             // false
fmt.Println("Function call:",  modifyBool(old))         // true fmt.Println("After function call: ", old)               // false


//# string
func modifyString(n string) string {
	n = "Golang"
	return n
}
message := "Go"
fmt.Println("Before function call: ", message)       // Go fmt.Println("Function call:", modifyString(message)) // Golang
fmt.Println("After function call: ", message)        // Go

//# array
func modifyArray(coffee [3]string) [3]string {
    coffee[2] = "germany"
    return coffee
}
country := [3]string{"nigeria", "egypt", "sweden"}
fmt.Println("Before function call: ", country)
// [nigeria egypt sweden]
fmt.Println("Function call:", modifyArray(country))
// [nigeria egypt germany]
fmt.Println("After function call: ", country)
// [nigeria egypt sweden]

//# struct
func modifyStruct(p profile) profile {
	p.Age = 85
	p.Name = "Balqees"
	p.Salary = 500.45
	p.TechInterest = true

	return p
}
myProfile := profile{
	Age:  15,
	Name: "Adeshina",
	Salary: 300,
	TechInterest: false,
}
fmt.Println("Before function call: ", myProfile)
// {15 Adeshina 300 false}
fmt.Println("Function call:", modifyStruct(myProfile))
// {85 Balqees 500.45 true}
fmt.Println("After function call: ", myProfile)
// {15 Adeshina 300 false}