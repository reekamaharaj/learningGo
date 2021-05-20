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

    - Overriding
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
Thank you. Thank you. So interfaces, we've talked about them. And one thing we want to talk about now is a little bit of how to use them. So what are they used for, language-wise? And in language, why would you need an interface? And I already said, interfaces, they express some sort of a conceptual similarity between different types. So the idea is that if two types satisfy an interface, then they must be similar in some way that is important to the application. So one common, sort of, practical thing that you would use an interface for, is when you need a function, you want to write a function which takes multiple types as a parameter. So specifically, normally, a function, it takes, say, it takes an integer as its arguments, right? It can only take an integer. But what if you want to take an integer or a float? Something like this, maybe a integer, a float or a string, right? You want it to take multiple types. In different types, it'll do different things, but you want it to be able to take different types. You can use an interface for that. So giving an example, say I've got a function, foo ( ), and it's gotta take a parameter. And this parameter, it could be either type X or type Y. And I'm talking very generically now. I'll give you a concrete example in the next slide, but so I got this foo. It's going to take a parameter which is type X or type Y. I want it to take either type. So the way I can do that is I can define an interface called Z.
Play video starting at :1:23 and follow transcript1:23
And X as a parameter can be, the type it can be, in interface type Z.
Play video starting at :1:29 and follow transcript1:29
Then I define X and Y to satisfy Z. Right, so, X and Y are both satisfying the interface Z. Then foo(), since it can take anything that satisfies interface Z, it can take X and Y as its arguments. So this is a common way to use an interface. And basically, an interface in this way, it sort of generalizes, right? It hides the details of the differences between the types. It's like look, these two types are similar in the way that's important to me. And so your function could just take the interface. It needs to take anything that is similar in that way.
Play video starting at :2:4 and follow transcript2:04
So, to be a little bit more specific, I made up a problem about a pool in a yard, right? So, I have a backyard, and I want to put a pool in my yard. But the pool, before I can put the pool in the yard, it needs to fit in my yard. And it needs to be fenced, because I don't want my kids to fall into the pool, so I need fencing around the pool. And it needs to fit in the yard. So to fit in the yard, the total area this pool needs to be limited. Less than the area of the backyard.
Play video starting at :2:30 and follow transcript2:30
Also to fence it, I only have so much fence, because fence cost money. I only have a limited amount of fence. So the perimeter of this thing has to be limited to within some limit, depending on how much I can afford. So I need to determine if a particular pool shape satisfies this criteria, because I'm trying to go through a bunch of different pool shapes. And I want to pick one that satisfies these criterias. It's sort of small enough that it can fit in the yard. And also the perimeter's small enough so that I can afford to fence it. So I'm going to write this function called FitInYard (). And this is going to return a boolean. So it takes a shape as an argument. Some shape, maybe I got some triangular shape, I was pass that to FitInYard, and it returns true, if the shape satisfies the criteria. So if the area is small enough, and the pool is small enough, then it says true. If I pass it a shape like rectangle and it's a big rectangle and it doesn't fit in my yard then it will return false, okay? So that's what FitInYard is. Now the thing about fit in yard is it's got to take a shape as an argument. But I want it to take any shape as an argument. I don't care if it's a triangle, circle, square, rectangle, whatever, it can be any shape. I should take it as an argument. But I have to be able to compute the area and compute the perimeter, okay? So not any shape, its got to be a shape whose area and perimeter I can compute. So let's say the idea is should take rectangles, triangles, whatever. But a valid shape has to have an area method and a perimeter method, right? So if the shape, if I can't compute the area, then I won't be able to tell if it fits in my yard. Say it's a sphere or something, right? There's no area, it's a 3D shape, that's a 3D object, I can't compute area of a thing like that, right? So, that is not a valid shape that I want to try to fit in my yard.
Play video starting at :4:15 and follow transcript4:15
So, any shape that has area and perimeter, that's okay with me.
Play video starting at :4:20 and follow transcript4:20
So what I can do is I can define this interface for shapes that have the area and perimeter. So I make my Shape2D interface. We already talked about this. But I make my Shape2D interface. It specifies area and perimeter, which return float64. Then I can make my types, triangle type, rectangle type, whatever types I want. And as long as these types, I don't care how they define what data's inside them, as long as they have methods that use them as receiver types. That it has area method and perimeter method. So, a triangle, you got area that has triangle as a receiver method and also perimeter. Same thing for rectangles, it's got an area and a perimeter. As long as they have area and parameter, I should be able to take this as an argument. So they satisfy this interface Shape2D. So in my FitInYard implementation, you can see that the argument that it takes is called s, and its type is the interface type, Shape2D. So what that means is that this argument could be any type that satisfies that Shape2D interface. Like rectangle, triangle, whatever the types are. And it returns a boolean and all the function does is very simple. Just says if s.Area is less than 100, let's say 100 is the size of my backyard, right? And as our perimeter is less than 100, because that's all the fence I could afford, then return true, else return false. So a valid argument to this is any type that satisfies the shape to the interface.
Play video starting at :5:57 and follow transcript5:57
Now the empty interface is standard interfaces predefined and it just specifies no methods. So that means that any type can actually satisfy that interface. And what you use it for is when you want to have a function argument be any type. You don't want to restrict it at all in terms of the type that this function can accept. Then you just make it's type the empty interface. So as an example, we got this function PrintMe. And its val argument is the empty interface. That's how you specify the empty interface that I haven't read. So that means that val can just be any type. And all this does is just print it. So it will print any type you give it. You give it an int, float, string, whatever, it will just print that to the screen.
Play video starting at :6:43 and follow transcript6:43
Thank you.

# M4.2.2 - Type Assertions
Module four, Interfaces for Abstraction. Topic 2.2, Type Assertions.
Play video starting at ::5 and follow transcript0:05
So a lot of the point of an interface is to conceal differences between types. So if you think about it, an interface can hide the differences between two types. It basically highlights the two similarities. So while in this fit and yard implementation, there are rectangles and there are triangles. But from the inside fit in yard, they're all treated the same, right? As long as they both satisfy Shape2D, I can call s.Area, s.Perimeter, right? So what you're doing is, what interface allow you to do is to treat different types that have some similarities, some similar methods, treat them the same. All right, so you are hiding differences by using interfaces. But sometimes you need to disambiguate, sometimes you need to treat different types in different ways, okay? So sometimes, like in this function, we don't, right? We can just say s.Area, s.Perimeter. We treat them exactly the same because they have the same methods. But sometimes you do need to differentiate based on the type. You need to be able to figure out what is the concrete type, right? So in this example, s, since you're just calling area and perimeter, it doesn't matter exactly what the concrete type is, right? The concrete type could be rectangle, it could be triangle. It doesn't matter, either way area and perimeter do what you think. So in this case, the concrete type that underlies the interface value, that doesn't matter, okay? But there are definitely cases where the concrete type matters. In those cases, you're going to have to expose those type differences. So you're going to have to take this interface, which is hiding the differences between the types, and peel it apart again, and say okay, actually, this is really a rectangle, this is really a triangle. So situations like that might be a graphics program, okay? So I got my graphics program which I've used many times, but in my graphics program, this time I want to write a function called DrawShape. And it should draw any shape. So I want to be able to pass it as an argument any two dimensional shape. So I declare it sort of the top line I'm showing right there, func DrawShape. It takes Shape2D, that's the type of it's argument, Shape2D. So it can take any two dimensional shape as an argument. So that's good, right? I've used my interface to generalize and to hide the differences between the types, rectangle, triangle, circle, doesn't matter, for passing it as an argument to draw shape anyway. Now, inside DrawShape, though, in this case, I'm going to have to disambiguate. I'm going to have to determine this s, is it a rectangle, is it a triangle, what is it? Because maybe in the underlying API there's some kind of drawing functions that I'm using in this API, right? And the underlying drawing functions, they actually are specific to the type of shaping drawing. So for instance, maybe the underlying API gives me a draw rectangle, DrawRect, right? And then another draw triangle, and a draw circle, and so on, which is not uncommon in these drawing APIs, right? So you got draw rectangle, draw triangle, draw circle, all those. Now, these API functions, my DrawShape is going to have call these, right? When it wants to draw a rectangle, it's going to have to call DrawRectangle. DrawTriangle to draw triangle. And these underlying API functions, they don't take just any shape. They won't take Shape2D. DrawRect only takes rectangles. DrawTriangle only takes triangles. And so on. So this is the case where I want to use my interface so my DrawShape can take any argument, any type of reasonable shape. But inside my DrawShape, I'm going to have to differentiate. I'm going to say, look, if you're a rectangle, call DrawRect. If you're a triangle, call DrawTriangle, and so on. So in this case, inside DrawShape, I'm going to have to determine the concrete type that s is based on, that the shape is based on.
Play video starting at :3:47 and follow transcript3:47
So for that, I use what's called a type assertion. So type assertions can be used to disambiguate between the different concrete types that actually satisfy a particular interface. And you can see that here with DrawShape. It needs to actually disambiguate. So if it's a rectangle that's being passed, it needs to call DrawRect. If it's a triangle, it needs to call it DrawTraingle. So you can see us doing that here. At the top, you've got that first type assertion, where it says rect, ok:=, so that will return a rectangle, if the s is actually a rectangle. So if ok is true, and it found a rectangle, it will call DrawRect with that rectangle. Otherwise, the next type assertion actually checks to see if the type of the interface of s is a triangle. So it says, tri, ok:= s.(Triangle) this time. And so, ok will be true if s is actually a triangle. And in that case, tri is going to equal that triangle. And so, you call DrawTriangle with a triangle. So either way, we use this type assertion to disambiguate, to determine the actual underlying concrete type for this Shape2D interface.
Play video starting at :5:4 and follow transcript5:04
Now, another way to do this sort of a common thing that you need to do is what we just did in the last slide, we went down a list of possible types. So rectangle and triangle in this case. But note that an interface can actually be satisfied by many different types. There's an interface that's satisfied by 10 different types, you might need to disambiguate all 10. So to run down the list. If you're this type, then do this. If you're that type, then do that, and so on. And so, there's a switch construct type switch which is just for that purpose. So you got one case for every different type that you need to deal with. So in this case, you got two cases, case Rectangle, case Triangle. And in each case, case Rectangle draws the rectangle, case Triangle draws the triangle. But right before that, you start with the switch. So notice the type Type assertion says s.type. In parentheses you just say type, the generic word type. And so, what happens is sh will be whatever the, it'll be the concrete type that s represents. So if s is actually a rectangle, then sh will be that rectangle, right? If s is a triangle, then sh will be that triangle. And you'll hit the appropriate case. So if sh is a rectangle, then you'll execute the case rectangle. If sh is a triangle, then you'll execute the case triangle. So this is just a more convenient way to sort of run down a list to disambiguate a whole set of types that all satisfy a particular interface.
Play video starting at :6:36 and follow transcript6:36
Thank you.

# M4.2.3 - Error Handling
Module 4, Interfaces for Abstraction, Topic 2.3, Error Handling. So I just want to show a common use of interfaces in Go.
Play video starting at ::10 and follow transcript0:10
The Error Interface. So there are a lot of different Go functions that are built into packages, which return errors. And when I say return errors, what they do is they return whatever they're supposed to return, and then their second return value is an error. An error interface. So and we see it defined over here. The error interface just is any type that satisfies this interface, and error interface just specifies that you have to have a method called, error, which prints the error message essentially which prints something, some text that's useful. So under correct operation, the error return might be nil, so for instance, let's say, I want to open a file. If it opens the file correctly, it'll return nil for the error, and there's no problem. But if the error actually has a value, then you'll probably print the error, and it'll call its error method, which will successfully print the error. So show you an example of that. So basically the idea is, when you, this happens, there a lot of different Go language function like this, which return error as the second argument, okay? And so, when that happens, you should check that error, after the call, and handle it if you need to. So you can see on the top line, we're opening a file. So os.open, open the file by that name, and it returns two things. One is the file F, and the second thing is an error, if an error exist. So then, right after that, for safety's sake, you should check the error. So if error not equal to nil, so if it's equal to nil, you're fine, you go on. If it's not equal to nil, then, do a print that sort of the most obvious thing to do to handle the error, is just to print it. So you println(err) and return.
Play video starting at :1:57 and follow transcript1:57
So printing the error, the format package, the fmt package, which println is a part of. That package will call the error method of the error to generate the string, and print that string. And so, this is sort of the generic way of handling errors in Go. It's a very common way to handle errors in Go.
Play video starting at :2:18 and follow transcript2:18
Thank you.