# M4.1.1 - Polymorphism

## Polymorphism
- Polymorphism: ability for an object to have different forms depending on the context.
    - a function or method can be different depending on object type
    - Ex. Area function. For rectange or triangle it could have the same name but it would do two different operations depending on the object type
    - High level of abstraction. Although the way of getting area is differnt in formula, it is still returning the area of the shape.
    - High level the return is identical. Low level (what the method does to get the return value) is different depending on type

Traditional OOP: inheritance through parent and child classes

- Subclass inherits the methods/data of the superclass
- Example: Speaker superclass
    - Speak() method, prints "<noise>"
- Subclass Cat and Dog
    - Also have the Speak() method
- Cat and Dog are different forms of speaker

- Go does not have inheritance
## Overriding
- Subclass redefines a method inherited from the superclass
- Example: Speaker, Cat, Dog
    - Speaker Speak() prints "<noise>"
    - Cat Speak() prints "meow"
    - Dog Speak() prints "woof"
- Subclass will redefine the speak value -> Speak will be polymorphic with this inheritance and overriding
    - Different implementations for each class
    - Same signature (name, params, return)


# M4.1.2 - Interfaces

## Interfaces
- Interface: Set of method signatures
    - Signatures: Name, parameters, return values and types
    - Helps with polymorphism
    - Used to express conceptual similarity between types
- Example Shape2D interface
    - All 2D shapes must have Area() and Perimeter()
    - Any type with these two methods would be considered a 2DShape
    - Does not define the method, but gives the signature for the methods
### Satisfying an Interface
    - Type satisfies an interface if type defines all methods specified in the interface
        - Same method signatures
    - Rectangle and Triangle types satisfy the Shape2D interface
        - Must have Area() and Perimeter() methods
        - Rectangle and Triangle can have other methods and data, but as long as they have the Area and Perimeter, they will be considered 2DShapes, they satisfy the interface

### Defining an Interface Type
```
type Shape2D interface {
    Area() float64
    Perimeter() float64
}
type Triangle {...}
func (t Triangle) Area() float64{...}
func (t Triangle) Perimeter() float64{...}
```
- Triangle type satisfies the Shape2D interface
- No need to state it explicitly
- In Go you do not need to explicitly say that Triangle satisfied the Shape2D interface

# M4.1.3 - Interface vs. Concrete Types

## Concrete Types
    - Specifies the exact representation of the data and methods
    - Complete method implementation is included
    - Also specifies the type of receiver type
    - Concrete type will have data that is associated with the type

## Interface Types
    - Specifies some method signatures
    - Implementations are abstracted
    - No data is specified, only methods
    - Interface will eventually be mapped to a concrete type

## Interface Value
    - When creating an interface you declare an interface type and make a value for that type
    - That interface type value can be treated like any other value (float, ints...)
        - Assigned to variables
        - Passed, returned

    - Interface value has two components
        - Dynamic Type
            - a concrete type that is assigned
            - Type that it is assigned to
        - Dynamic Value
            - The value of the dynamic type
### Example Shape2D Interface
    - Concrete types that satisfy interface
        - Rectangle
        - Triangle
        - Circle
    - Make an interface variable
        - give it a value, the value has to be mapped to a concrete type
        - Dynamic Type: The type of concrete type that the interface value is assigned to. Ex. Rectangle, Triangle, Circle
        - Dynamic Value: The value of the Dynamic Type

Interface value is a pair of dynamic type and dynamic value

## Defining an Interface Type

```
//Define Speaker interface
type Speaker interface {Speak()}

//Create dog type
type Dog struct {name string}

//Make dog type satisfy the Speaker interface by creating a function and having the receiver type be dog
func (d Dog) Speak() {

    fmt.Println(d.name)
}
func main() {
    //Declare Speaker
    //s1 is a variable and is the speaker interface value
    var s1 Speaker

    //Declare Dog
    //d1 is a dog type
    var d1 Dog("Brian")

    //speaker interface value can equal dog type because dog type satisfies the speaker interface
    s1=d1
    s1.Speak()

    //d1 is the concrete type assigned to speaker
}
```
- s1 Speaker type
- s1 is assigned to concrete type d1
- concrete type is the dynamic type, dog type
- dynamic value is the dog type value, the name

### Interface with Nil Dynamic value
    - An interface can have a dynamic type but not dynamic value
```
var s1 Speaker
//d1 is not a concrete object, it points to a dog
var d1 *Dog

//s1 does not have a value, but will be dog type and have a pointer to Dog
s1=d1
```
s1 doesnt have data. s1 has a dynamic type, of dog or a pointer to a dog. but no dynamic value
- interface here has a nil dynamic value
- Nil: nothing, empty end goal
- Can still call methods of s1
- Compiler knows that when Speak is called

```
func (d *Dog) Speak(){
    if d == nil {
        fmt.Println("<noise>")
    } else {
        fmt.Println(d.name)
    }
}
var s1 Speaker
var d1 *Dog
s1 = d1
s1.Speak()
```

    - if d == nil then it prints ("<noise>") some generic thing
    - else it prints (d.name), it prints the dog's name

### Nil Interface Value
- Interface with nil dynamic type
- No dynamic type, cannot call the methods on that interface

Nil dynamic value and valid dynamic type
```
var s1 Speaker
var d1 *Dog
s1 = d1
```
- Can call a method since type is known

Nil Dynamic Type
```
var s1 Speaker
```
Cannot call a method, runtime error

# M4.2.1 - Using Interfaces

## Ways to Use an Interface
- Need a function which takes multiple types of parameters
- Function foo() parameter
    - Type X or Type
- Define interface Z
- foo() parameter is interface Z
- Types X and Y satisy Z
- Interface methods much be those needed by foo()

### Example
- Pool in a yard
- Pool needs to fit in the yard
- Pool needs a fence around it
- Total area of the pool needs to be limited. Less than the area of the backyard
- Limited amount of fence. There is a limited perimeter depending on affordability
- Pool that is small enough to fit in the yard but also has a small enough perimeter to fit the budget.
- Function FitInYard() will return a boolean true if an argument satisfies the criteria
    - Arguments function takes
        - Shape: Need to calculate area and perimeter of any shape entered
        - Valid shape has to have an area and perimeter method

- Define interface for shapes that have area and perimeter. Shape2D Interface, because a 3D shape would not give an area.

```
type Shape2D interface {
    Area() float64
    Perimeter() float64
}

type Triangle {...}
func (t Triangle) Area() float64{...}
func (t Triangle) Perimeter() float64{...}

type Rectangle {...}
func (r Rectangle) Area() float64{...}
func (r Rectangle) Perimeter() float64{...}

func FitInYard(s Shape2D) bool {
    if (s.Area() > 100 &&
        s.Perimeter() > 100) {
            return true
    }
    return false
}
```

## Empty Interface
- Empty interface specifies no methods
- All types satisfy the empty interface
- Use it to have a function accept any type as a parameter
```
func PrintMe(val interface{}) {
    fmt.PrintIn(val)
}
```
# M4.2.2 - Type Assertions

## Concealing Type Differences
- Interfaces hide the differences between types
```
func FitInYard(s Shape2D) bool {
    if (s.Area() > 100 &&
        s.Perimeter() > 100) {
            return true
    }
    return false
```
- Sometimes you need to treat different types in different ways

## Exposing Type Differences
- In a situation where type difference matters, you may need to expose the type differences in an interface.
- Need to peel it apart
- Example: Graphics Program
- DrawShape() will draw any shape
```
func DrawShape (s Shape2D){...}
```
- Why: Underlying API has different drawing functions for each shape
    - func DrawRect (r Rectangle) {...}
    - func DrawTriangle (t Triangle) {...}
    - You would need to specify what type in the interface needs to call what func from the API, the concrete type os s Shape2D would need to be determined in this case

## Type Assertions for Disambiguation
- Used to determine and exact underlying concrete types
```
func DrawShape(s Shape2D) bool {
    rect, ok := s.(Rectangle)

    if ok {
        DrawRect(rect)
    }
    tri, ok := s.(Triangle)
    if ok {
        DrawRect(tri)
    }
}
```
- Type assertion extracts Rectangle from Shape2D
    - Concrete type in parentheses
- If interface contains concrete type
    - rect == concrete type, ok == true
- If interface does not contain concrete type
    - rect == zero, ok == false

## Type Switch
- Switch statement used with a type assertion
```
func DrawShape(s Shape2D) bool {
    switch := sh := s.(type) {
        case Rectangle:
            DrawRect(sh)
        case Triangle:
            DrawTri(sh)
    }
}
```
# M4.2.3 - Error Handling
## Error Interface
- Many Go programs return error interface objects to indicate errors
```
type error interface {
    Error() string
}
```
- Correct operation: error == nil
- Incorrect operation: Error() prints error message

## Handling Errors
```
f, err := os.Open("/harris/test.txt")
if err != nil {
    fmt.Println(err)
    return
}
```
- Check wether the error is nil
- If it is not nil, handle it
- fmt package calls the error() methd to generate string to print

## Questions
- What is polymorphism?
    When one thing can have multiple forms.
- Which of the following statements is true?
    Inheritance and overriding enable polymorphism.
- If a type satisfies an interface...
    The type defines all methods specified in the interface.
- An interface always has a dynamic type.
- Interfaces can support abstraction by concealing differences between types.
- Type assertions can reveal differences between type satisfying an interface.
- Type assertions return two values.
- What is a use for an empty interface?
    It can be used to allow a function to accept any type as a parameter.
- After executing the expression below, what is the value of err if there is no error?
```
f, err := os.Open(“/harris/test.txt”)
```
    nil

