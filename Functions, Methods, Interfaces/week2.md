# M2.1.1 - First-Class Values

- Have at least the main function
- First-class values are associated with functional programming
    - Functional programming, way of programming
        - "a programming paradigm"
    - Examples of functional programming languages: Scheme, ML
    - Different way of thinking
- Treat a function as a first-class value
    - Functions can be treated like other types
        - integer, string, float...
    - Variable can be declared with a function type
        - Can then set variable equal to a function
    - Can be created dynamically, created while the code is running
    - Can create a new function inside a function
    - Can be passes as arguments
    - Can be returned as values
    - Can be stored in data structures

## Variables as functions
- Declare a variable as a function
- Variable acts as an alias, another name for the function

```
var funcVar func(int) int

// Variable declared to be a function that will take and int as a parameter and return and int, function signature

func incFn (x int) int {
    return x + 1
}

// Funtion incFn created

func main() {
    funcVar = incFn

// Assigning funcVar to be incFn function
// Note: Not () after incFn
    fmt.Print(funcVar(1))
}
```
- Use funcVar the same way you would use incFn

- Function signature: two type list. One is the input parameter and the other is the output result type

## Function as an arguments
- Funtion can be passed to another function as an argument

### Example 1
```
func applyIt (afunct func (int) int,

// function applyIt takes two parameters. One is a function called afunct, afunct function takes a int as a parameter and returns an int

    val int) int {

// The second parameter is an int valled val
// The applyIt function returns an int

    return afunct(val)

// applyIt takes two arguments: afunct and val
// applyIt takes afunct and applies val to afunct and return an int
}
```

### Example 2
```
func applyIt (afunct func (int) int,
   val int) int {
   return afunct(val)
}

func incFn(x int) int (return x + 1)
func decFn(x int) int (return x - 1)

func main() {

   fmt.Print(applyIt(incFn, 2)

// Will print out 3

   fmt.Print(applyIt(decFn, 2)

// Will print out 1
}
```



## Anonymous functions
- Anonymous function. Functio without a name.
- Use might be: when passing a function as an argument to another function
- Lambda calculator, not naming the function, just entering it into the argument.

### Example 3

```
func applyIt (afunct func (int) int,
    val int) int {
    return afunct(val)
 }

 func main() {
    v := applyIt(func (x int) int
                {return x+1}, 2)

// Define function without a name

    fmt.Print(v)

// Will print out 3
 }
```


# M2.1.2 - Returning Functions
- Functions can also return other functions as their return value.
- Why? Create a new function for a special purpose. Parameterizable function
- Return a new function with different properties

### Example: Distance to Origin function
- Takes a point (x, y coordinates)
- Returns distance to origin
- Case: Change the origin
- Option 1: Pass origin as argument
- Option 2: Define function with new origin

- Function defines a function
```
func MakeDistOrigin (o_x, o_y float64) func (float64, float64) float64 {

//Origin location is passed as an argument

fn:= func (x,y float64) float64 {
    return math.Sqrt(math.Pow(x - o_x, 2) +
    math.Pow(y - o_y, 2))

// Origin is built into the returned function

    return fn
}

func main() {
    Dist1 := MakeDistOrigin(0,0)
    Dist2 := MakeDistOrigin(2,2)
    fmt.Println(Dist1(2,2))
    fmt.Println(Dist2(2,2))
}
```



- Environment/Scope of a Function: set of all names that are valid inside a function
    - Names defined locally, in the function
    - Lexical Scoping: environment includes names defined in block where the function is defined
- Clousure: a function plus its environment
    - When you pass a function, you also pass the functions environment

# M2.2.2 - Variadic and Deferred
- Variable Argument Number
    - Function can take a variable number of arguments
    - ... to specify a variable number of arguments
    - ... is treated as a slice in the function
```
func getMax(vals ...int) int {

// ... means it can take as many integers as you want

    maxV := -1
    for _,v := range vals {
        if v > maxV {
            maxV = v
        }
    }
    return maxV
}
```


- Variadic Slice Argument
```
func main() {
    fmt.Println(getMax(1, 3, 6, 4))
    vslice := []int{1, 3, 6, 4}
    fmt.Println(getMax(vslice...))
}
```

- Can pass a slice to a variadic function
- Needs the ... suffix

- Deferred Function Call: not executed where it's called
- When the surrounding function completes, then it will run
- used for cleanup activies
```
func main() {
    defer fmt.Println("Bye!")
    fmt.Println("Hello!")
}
```
- Deferred Call Arguments are not evaluated in a deferred way
```
func main() {
    i := 1
    defer fmt.Println(i+1)
    i++
    fmt.Println("Hello!")
}
```
- The defer function would print "2" not "3"
- if the i++ is moved to be before the defer, then it would print "3"
- The value of the variable when defer is called, is what will be used in the defer function, but the function will only run at the deferred time.
