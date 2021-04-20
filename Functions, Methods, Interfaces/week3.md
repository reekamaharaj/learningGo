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
Module 3: Object-Orientation in Go, Topic 1.3: Support for Classes.
Play video starting at ::7 and follow transcript0:07
Now, in a normal object directed language, it's a class is defined, it's data associated with some kind of methods. And usually, you can associate lots of different data, you can roll up lots of different variables. Maybe an anti-float, whatever type of data you can put a lot of it, as much as you want together. And then associate that with any number of methods, asnd you can do the same thing in go. Of course, you're going to use a receiver type, just like we talked about, you don't have classes, you have receiver types. But you can just use a type that has lots of data in it. So, before we're using examples where the type was just an int myant, right. It was just an int, one piece of data. But it's very common to use a struct as a receiver type, a struct of some kind. So structs basically allow you to compose all kinds of different data fields. So in this case, my point struct, I'm just composing two numbers, an x and a y, both floats. So two floating point numbers, they're composed into one struct. But remember, with a struct you can compose an arbitrary amount of information, you can put together. So it's very common to see a receiver type be a struct of some kind with lots of different data. It's a traditional feature of classes, people just roll lots of different data together.
Play video starting at :1:27 and follow transcript1:27
Now, the structs with methods, you could take a struct and define it as a type, like we just did with that point type. And then you can associate methods with that type. And then you get what you would normally think of as a class in another language. You get the struct with lots of different data, associate it with as many different methods as you want to associate it with struct. So, we got an example of that right here, we're using the point that I defined just in the last slide. So, this point, I want to make a function called DistToOrigin, and I've defined it right there. Notice that to the left of the name of the function DistToOrigin, I pass it a point, p Point, right? When I say I pass it, it's an implicit pass, right. So it doesn't have any explicit arguments, but its receiver type is a Point called p, and that will be implicitly passed to DistToOrig. And then if you look at the function, the insides of it, the internals, it's just doing the Pythagorean theorem, right. It's squaring the x, squaring the y, adding together, then it returns the square root. So it just does Pythagorean theorem nothing sophisticated. Then in my main, I can make a point p1 is three comma four and then I can just call p1.DistToOrigin. And that p1 together with its x and y coordinates will be implicitly passed to dist to origin. Dist to origin will then do Pythagorean theorem and return the distance which is five in this case, thanks.

# M3.2.1 - Encapsulation
Module 3 or Object Orientation in Go. Topic 2.1, Encapsulation. So Go provides a lot of different support for encapsulation and keeping private data. But you want to be able to have controlled access of the data. So typically, even if you have private data in some package, you probably don't want to hide it completely. All right, or else why are you even importing it anyway? You hide it but you want to have controlled access to it. So what that means is you want people to be able to use that data but only in the way that you define, okay? Using your methods or functions. So what you can do is you can define a set of functions, public functions, that allow another package, an external package, to access the hidden data. So as an example, say I got my data package right there, package data, got my hidden variable x int= 1. And then I can define inside that same package a function called PrintX. And PrintX just prints x, okay? It does exactly what it says. Now PrintX, notice it starts with a capital letter, so that means it gets exported. So if my main package decides to import the data package, the main package will be able to access this PrintX method even though it can't directly access the x. Okay, and so now what happens is I can access the main method, the main function can access the x variable only through this PrintX function. So if I want to see the x value, I have to call PrintX. So if I look in my main code, I import the data and then in my main I can call data.PrintX and then I can see the value of x. Even though I couldn't directly access x from my main, I can indirectly access it through these public functions. So this is generally how we're going to control access to data that we want to hide. You want to give access but only in a controlled fashion. We let them see what we want them to see, is the idea. Also, to modify code, to modify x, right. I mean as it is, x cannot be modified externally. There's no way the main can directly see x or modify it. But if I wanted to allow the main to be able to modify x, I could make some kind of a function inside the package, start it with a capital letter that main could call to access the variable.
Play video starting at :2:21 and follow transcript2:21
So we can do this with structures too. So say we have some kind of type that's a structure. Like a point type. We put that in our data package again, right? And maybe the x and the y and coordinates, we don't want the outside user, the person who is using this type to be able to directly modify x and y. We want to be able to control their observation and their modifications to x and y. So we give them lower-case names, lower-case x, lower-case y. But we define a set of functions inside that package, the data package, that are public, and allow another package to use to actually access x and y in some way. So for instance, first one you might want to define is InitMe that I'm defining down here. And that, notice it is associated with the point type, the receiver type is Point, so p *Point. I call InitMe and InitMe just allows me to initialize x and y, right? That's something clearly you're going to want to do. You make a point, you want to initialize the x and y values. So I do it through this InitMe method that I'm defining. And it just sets p.x equal to the first argument, p.y equal to the second argument. So in this way, using this function, this InitMe function, I can modify x and y, even though I can't directly touch x and y I can do it through this function. Then a few more functions that you might want to add to allow access to the x and y, they're hidden. This is Scale. So Scale, again, it's associated with Point, its receiver type is Point. And Scale, you pass it a floating point number v, and it just scales x and y together, so it multiplies p.x times the scale factor, p.y times scale factor. Again, we're not trusting the programmer to do this, we're scaling both of them together. So if they want to scale, they have to call our scale function and they can scale them both. Also, print me, maybe I want to be able to print the x and y values. And since another package can't directly access x and y to call Println on it, we have to provide a function for that, PrintMe. And it just goes in there and it prints out the x and y, prints out p.x, p.y. So now we define a set of functions, a set of methods, really, because they're all associated with the type Point. And these methods are all public because we started them with capital letters, Print and Scale, they're all capital. So we can access them outside in, say, our main package.
Play video starting at :4:48 and follow transcript4:48
So in our main package we can use them, so for instance, in
Play video starting at :4:53 and follow transcript4:53
this main we make a point, data.Point call p. Then we call p.InitMe, it initializes x and y to 3 and 4. Then we call p.Scale to scale it, to multiply 3 and 4 times 2, so it should be 6 and 8. Then we call p.PrintMe, it prints 6 and 8. So if we ran this it would work. And in this way even though we can't from the main, we can't directly access x and y, we can't say p.x= bam p.y=, right? But we can access them through these methods that are provided to us in a controlled way. Thank you.

# M3.2.2 - Point Receivers
Module 3, Object-Orientation in Go, Topic 2.2, Point Receivers. So we've been talking about methods, defining methods for an association with different receiver types. So there are a few limitations to this process that we may need to overcome. So remember [COUGH] that this receiver type is implicitly passed. The receiver object is implicitly passed as an argument to the method. So even though it's not explicitly passed, it is implicitly passed. And remember that argument passing in Go is passed by value, called by value. So that means that the method can't modify the data inside the receiver object. So as an example, let's say we had some mythical OffsetX, and it should increase the x coordinate of a point, right? We wanted to add some constant to the x coordinate at some point. So in our function main, we say pl := Point (3, 4), we make a point. Now we say pl.OffsetX, and we pass it 5 is the value that we want to add to the X. That won't change the X. It can't change the X coordinate. The reason why is because this OffsetX is being passed a copy of p1, not appointed a p1, a copy of p1. And so as it gets a copy of p1, it can change it's copy, so it gets its own p1.x copy, it can change that from 3 to 8. But that goes away as soon as soon as the function is done executing. So as Offset as x is done executing, that goes away, because it's in its environment is gone. So what we want in this case is to be able to change the actual values inside p1, but you can't, because they get the p1. The p1, the object is actually passed by value. Another problem with that is if the receiver is large, a lot of copying happens when you make a call. So when you call by value and the receiver object is passed as an argument, the whole thing has to get copied onto the stack internally. And if it's a large receiver object, then there's a lot of copying. So in this case, I've got my type Image [COUGH] and this type is a 100 by 100 array events, which is actually small for an image, okay? So that's 10,000 int. So when I call this BlurImage [COUGH] which is some method that's associated with whose receiver type is this image. When I call BlurImage, the i1, this image, actually has to get passed to the BlurImage method. And that's 10,000 ints that you've got to copy out of the stack and that can take a long time. And images actually are a sort of worse case, because they can get gigantic, right? 100 by 100 is not even big, so that can waste a lot of time.
Play video starting at :2:49 and follow transcript2:49
So this is the problem.
Play video starting at :2:51 and follow transcript2:51
So what do you do? Well, we do what we did before with method argument, not even with method, with regular functions. We can just pass, instead of passing, calling by value, you can call by reference. So you explicitly pass the appointed to the object, rather than the object itself. So the way you manage this is when you declare the function. Like we're seeing here, we're declaring this OffsetX. See this receiver type to the left of the name of the function. So in parentheses is p *Point. I say *Point this time, *Point, right? Instead if I said p *Point, then I'm passing the point called by value. But if I say p *Point, then now p is a pointer to a point, right? So now, and when I implicitly pass this p value to offset x, it's going to pass a pointer to the point type that we're talking about. So now inside the function, I say p.x = p.x + v and that'll actually work, because p.x is now, since p is actually a pointer to this structure. p.x points to the actual x value in memory. So you can actually modify it now, because you didn't call it by reference.
Play video starting at :4:7 and follow transcript4:07
Thank you.

# M3.2.3 - Point Receivers, Referencing, Dereferencing
Module 3, object-orientation in go. Topic 2.3, point receivers, referencing, and dereferencing.
Play video starting at ::7 and follow transcript0:07
So one thing about using a pointer receiver is that there's no need to dereference the pointer inside the method. So what I mean by this is that, so let's take a look at this example OffsetX. I want it to have a pointer receiver, because I want OffsetX to actually be able to modify the X coordinate of the type. So I have to pass it a pointer to p. So now if you look to the left of the function name OffsetX, you see p star point. So the receiver is a pointer type. Now notice inside the function I say p.x=p.x+v.
Play video starting at ::43 and follow transcript0:43
I don't say *p.x=*p.x+v, right. That's a dereferencing that you would normally use with pointers. I don't have to do that because this is a common enough thing that going just allows you to get it with basically the ( ) recognizes it and says okay, I know what you mean. If you can just say p.x and it knows even though it's a pointer, it knows to get the x p, it knows to basically de-reference it automatically. So, it's just a handy shorthand, you don't have to do the dereferencing when you're doing this.
Play video starting at :1:16 and follow transcript1:16
And likewise there's no need to reference either. So, say I'm in my main and I want to I've defined my offset X, as I showed you, where it accepts a pointer,
Play video starting at :1:29 and follow transcript1:29
the receiver type is a pointer, right? Now in this case, in this main, I'm declaring this point p I'm making it 3,4. So p is actually a struct, right? It is the type, it is the actual struct. It is not appointed to the struct, it is the struct. But then when I call OffsetX, I say p.OffsetX, when really, since offset X is supposed to have a point to receiver you would think you would have to say anforsand P dot offset X. But you don't, you can just say P dot offset X and the goal compiler recognizes that just because this is a common thing to do. So it just makes it easier, it's a convenience. So when using pointer receivers it is good programming practice to either have all methods use pointer receivers, or have none of them use pointer receivers. It's just a good standard. It's easy to get confused. So if you have some methods, use pointer receivers and some not use pointer receivers It can get confusing. You'll send a pointer to the one that doesn't need a pointer and so on. So it's just more appropriate to use, it's good practice, you don't have to. You can mix and match if you want, but it's good practice to use all pointer references for a particular type or all non pointer references.
Play video starting at :2:47 and follow transcript2:47
Thank you.