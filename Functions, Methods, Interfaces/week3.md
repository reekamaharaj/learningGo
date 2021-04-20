# M3.1.1 - Classes and Encapsulation
## Classes - object-oriented programming paradigm
- Collection of data fields and functions (methods) that share a well defined responsibility. All related to the same concept
- Class is a template. Contains data fields, but not the data
- Example for a geometry program
    - Data: x and y Coordinates
    - Function(Methods): DistToOrigin(), Quadrant(), AddXOffset(), AddYOffset(), SetX(), SetY()

## Object
- an instance of the class
- contains data
- Example for a geometry program
    - Point class with points (0,0) (5,5) (6,0)


## Encapsulation
- concept associated with OOP
- associated with the use of abstraction
- Data can be protected from the programmer (not creator)
- Data can be accessed only using methods
- Why: Maybe don't trust the programmer to keep data consistent
- Helps debugg and avoid mistakes from user error
- Benefit for internal consistency of data
- Data can't be accessed directly from the outside
- Hard abstraction barrier: Methods to access the data
- Example for a geometry program
    - Option 1: Make method DoubleDist()

# M3.1.2 - Support for Classes (1)
- Go doesn't have a class keyword
- Most OOP languages have a class keyword
    - Data fiels and methods are defined inside a class block

## Associating methods with data in Go
# Receiver type
- Type that the method is associated with
- To associate the data with the method the function is given the reciever
- When you call the method, you used standard dot notation

- When associating a method with a receiver the type needs to be defined in the same package as the associated method
- Function definition is different than normal function definitions, to define the receiver

Example
```
// created type

type MyInt int

// before the name of the function is the receiver type (mi MyInt)

func (mi MyInt) Double () int {
    return int(mi*2)
}
func main() {
    v := MyInt (3)
    fmt.Println(v.Double())
    // left of dot, v is the object, the data
    // right of dot, function/method
}
```

- Type MyInt, is a type like string or int but the programmer creates it
- The Method Double() is defined and will only work with a MyInt typed variable
- All variables defined with a MyInt type of data will be an int, and will work with the Double() Method
- variable v is set to be a MyInt type, which is an int and that int is 3
- The Method Double() is called with variable v. Since variable v is MyInt type, it can be used with the method Double()
- The Method Double will use the value of the MyInt type as an argument defined as mi. v.Double() makes this process happen

## Flow
- Create a type called MyInt which will be an int
- Create a method that only works with MyInt types
- Method will take one argument of a MyInt type and use it as the data for mi
- Method will return a value as an int
- MyInt is the receiver type for method Double
- To call Double, it needs to be prefixed with dot notation
- The prefix needs to be a MyInt type

## Implicit Method Argument
- Double does not take an argument in the definition
- There is an implicit argumet for Double
- For a Method, there is a receiver type
- A receiver type will have an object assigned, the thing to the left of the dot notation. That is where the data is coming from. Passed implicitly
- The receiver type is an implicit argment to the method
- Double looks like it doesn't have an argument. But it has one implicit argument
- When the Double method is called the object used in the call will be passed to the Double method
- It is passed with call by value
- a copy of v will be made and passed to Double

# M3.1.3 - Support for Classes (2)

- Class defined, data associated with methods
- Receiver type in go, not classes
- Struct is commonly used as a receiver type
    - Struct can compose all kinds of data fields

## Struct with methods
- Like a class in other programs
- Associated data and methods

```
type Point struct {
    x float64
    y float 64
}

func (p Point) DistToOrig(){
    t := math.Pow(p.x, 2) + math.Pow(p.y, 2)
    return math.Sqrt(t)
}

func main(){
    p1 := Point(3,4)
    fmt.Println(p1.DistToOrg())
}
```
- Point is a stuct, type is Point and it is composed of two floating point numbers
- Point, called p is a receiver type. It is implicitly  passed to DistToOrig
- DistToOrig does not take explicit arguments
- p1 is used to make a point with given values and p1.DistToOrig will call the function with the defined point

# M3.2.1 - Encapsulation

- keeping data private
- controlled access to data

## Controlling Access
- Can define public functions to allow access to hidden data
- People can use data in predefined ways

### Example
```
package data
var x int=1
func PrintX() {
    fmt.Println(x)
}
```
```
package main
import "data"
func main(){
    data.PrintX()
}
```

- x is a hidden variable, x is exported
    - capital letter -> exported
- if the data package is imported, then the main method has access to x as a variable
- calling the function PrintX in main, after importing will allow access to print x

## Controlling Access to Structs
- Hide fields of struct by starting field name with a lower-case letter
- Define public methods which access hidden data
- case where you don't want the outside user to directly modify x and y
- give them lower-case names
- define a set of functions inside the pacakge
- functions are public
- Need InitMe() to assign hidden data fields

### InitMe
```
package data
type Point struct {
    x float64
    y float64
}
func (p *Point) InitMe (xn, yn float64 {
    p.x = xn
    p.y = yn
})
```
- function is associated with the receiver type, p *Point
- InitMe will allow you to initialize x and y
- it will set p.x and p.y equal to the arguments provided
- Can directly change x and y but can use the function to change x and y

### Scale

```
func (p *Point) Scale (v float64) {
    p.x = p.x * v
    p.y = p.y * v
}
```
- Scales x and y by an argument
- Function is already written, programmer just has to give the scale factor

### PrintMe
```
func (p *Point) PrintMe(){
    fmt.Println(p.x, p.y)
}
```
- Print the x and y values

- Defined set functions
- Defined set methods
- Associated them with a type, Point
- Methods are public, with a capital letter
    - PrintMe, Scale, InitMe
- Can access outside of package

### Main
```
package main
func main (){
    var p data.Point
    p.InitMe(3,4)
    p.Scale(2)
    p.PrintMe()
}
```

- Initializes point with x=3 and y=4
- Scale x and y by 2. So x=6 and y=8
- Prints x and y values.

# M3.2.2 - Point Receivers
## Limitations of Methods
- Receiver is passed implicitly as an argument to the method
- Method cannot modify the data inside the receiver
- Example: OffsetX() should increase x coordinate
```
function main(){
    p1 := Point(3,4)
    p1.OffsetX(5)
}
```
- OffsetX cannot change the x value
- OffsetX is being passed a copy of p1
- OffsetX can change the copy of p1
- The copy of p1 is changed but when the function is done executing the environment goes away, and the copy of p1 goes with it
- Since it is passed by value, it can not change p1
#### Large Receivers
- if the receiver is large then there is a lot to be copied

```
type Image [100] [100]int
func main() {
    i1 := GrabImage()
    i;.BlurImage()
}
```
- 10,000 ints copied to BlurImage
- Instead of calling by value, you call by reference

### Pointer Receivers
```
func (p *Point) OffsetX(v float64){
    p.x = p.x + v
}
```
- explicitly pass the pointer to the object and not the object
- p is a pointer to a Point
- a pointer is being passed to the function
- Receiver can be a pointer to a type
- Call by reference, pointer is passed to the method

# M3.2.3 - Point Receivers, Referencing, Dereferencing

Using a pointer as a receiver means you don't need to dereference the pointer

```
func (p *Point) OffsetX(v int) {
    p.x = p.x + v
}
```

- Want OffsetX to actually modify the x coordinate, so pass the pointer to p
- Using a receiver means you don't have to use `*p.x=*p.x+v`, which is dereferencing
- Point is references as p. not *p

No Need to Reference
```
func main(){
    p := Point{3,4}
    p.OffsetX(5)
    fmt.Println(p.x)
}
```
- Do not need to reference when calling the method

## Using Pointer Receivers
- Good practice
    - All methods for a type have pointer receivers, or
    - All methods for a type have non-pointer receivers

What is the difference between an object and a class?
    A class is a template and an object is an instance of that template.
Which of the following refers to data hiding?
    Encapsulation
Assume that that the type t is the receiver type for a method called Foo(). Under what conditions would it be better to make the receiver type of Foo() a pointer to t, rather than itself?
    I. When the receiver type t uses a large amount of memory.
    II. When the method Foo() must modify the data in the object of the receiver type.
        Both I and II
What is the difference between a struct in Go and a class in an object-oriented language?
    A struct contains only data while a class can also contain methods.
How do you associate a method with an arbitrary data type on Go?
    Define the method so that its receiver type is the data type of interest.
Say that you have defined a type t and you have declared an object of that type called t1. Assume that the type t is the receiver type for a method called Foo(). Which expression shows a proper invocation of the the method Foo()?
    t1.Foo()


In Go, how do you hide variables or functions in a package, so that functions outside of the package cannot access them?
    NOT: Define the variable/function inside the package. Or Use the package keyword
    Note: Puts the variable/function inside a package but does not necessarily hide them inside the package.Defines a package, but does not necessarily hide everything inside the package.

    ?Use the private keyword.
    ->?Give the variable/function a name which starts with a lower-case letter

