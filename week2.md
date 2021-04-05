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

- `var funcVar func(int) int`
- Variable declared to be a function that will take and int as a parameter and return and int, function signature
- `func incFn (x int) int {`
- `   return x + 1`
- `}`
- Funtion incFn created
- `func main() {`
- `   funcVar = incFn`
- Assigning funcVar to be incFn function
- Note: Not () after incFn
- `   fmt.Print(funcVar(1))`
- `}`
- Use funcVar the same way you would use incFn

- Function signature: two type list. One is the input parameter and the other is the output result type

## Function as an arguments
- Funtion can be passed to another function as an argument

### Example 1
- `func applyIt (afunct func (int) int,`
- function applyIt takes two parameters. One is a function called afunct, afunct function takes a int as a parameter and returns an int
-    `val int) int {`
- The second parameter is an int valled val
- The applyIt function returns an int
-    `return afunct(val)`
- applyIt takes two arguments: afunct and val
- applyIt takes afunct and applies val to afunct and return an int
- `}`

### Example 2
- `func applyIt (afunct func (int) int,`
-    `val int) int {`
-    `return afunct(val)`
- `}`
- `func incFn(x int) int (return x + 1)`
- `func decFn(x int) int (return x - 1)`
- `func main() {`
-    `fmt.Print(applyIt(incFn, 2)`
- Will print out 3
-    `fmt.Print(applyIt(decFn, 2)`
- Will print out 1
- `}`

## Anonymous functions
- Anonymous function. Functio without a name.
- Use might be: when passing a function as an argument to another function
- Lambda calculator, not naming the function, just entering it into the argument.

### Example 3
- `func applyIt (afunct func (int) int,`
-    `val int) int {`
-    `return afunct(val)`
- `}`
- `func main() {`
-    `v := applyIt(func (x int) int `
-    `            {return x+1}, 2)`
- Define function without a name
-    `fmt.Print(v)`
- Will print out 3
- `}`

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
- `func MakeDistOrigin (o_x, o_y float64)`
- Origin location is passed as an argument
- `            func (float64, float64) float64 {`
- `    fn:= func (x,y float64) float64 {`
- ` return math.Sqrt(math.Pow(x - o_x, 2) +`
- ` math.Pow(y - o_y, 2))`
- Origin is built into the returned function
- ` return fn`
- `}`


So a way we do it, is we can make a function that defines a function and it returns a function. So this function is called MakeDistOrigin, right? And the point of this function is to create the function that I want, this distance to origin function. So if we look down at the bottom of this function, you could see it says a return fn, right? So what its going to do is its going to define a function inside it, and then it's going to return that function as its return value. So let's look at the, go back to the top, make this into origin. It takes several arguments, okay.

The arguments right there are o_x, o_y, right? These are two floats. Now, these two floats, they're supposed to be the origin. The o stands for origin in this case. So the o_x and o_y are the origin location, right? So if I want my origin to be 0,0, those first two arguments would be 0,0. So wherever I want my origin, I pass that as arguments to this MakeDistOrigin function. Now, the return value of this MakeDistOrigin function is a function. And you can see that in the second line, line 2, func at two arguments float64,float64 returns one float64. So the two arguments of this return value of the two floats are going to be the x and y coordinates whose distance we want to compute to the origin. And then, the last float64 is a return value which is going to be a float which is the actual distance, okay? So just get this straight, this make distance to origin function, it has two arguments which are the x, y coordinates of the origin. And it has a return value which is a function, which is going to compute the distance to that origin that we passed as arguments. Now, if we look into the function, line three. We have a variable called fn, and that's going to be the function that we're creating. So fn: equals, and you notice we're defining a new function, func x, y to floats and returns a float. And this new function that we're creating, it returns math.square root, so it computes square root. It's doing Pythagorean theorem. Math,pow x minus o, x, right? So I take the difference between the x coordinate and the origin. And the origin's x. And I square that, and then I do the same thing, I add that to y squared. So it's y minus it's origin with a y component of the origin, squared. And then, I take the square root of that. So that's just Pythagorean theorem. It does its Pythagorean theorem, and notice that it take the x, y coordinates that you want to find, whose distant you want to find from the origin. Subtracts x from the origins x, y from the origins y. Does the Pythagorean theorem. And then, it returns that function. So this MakeDistOrigin function doesn't actually compute the Pythagorean theorem, it returns a function whose job it is to do Pythagorean theorem, to compute Pythagorean theorem to figure out the distance from an origin, o, x comma o, y. So this is a function that is made special purpose, this actually creates special purpose functions. So notice, I could use this function, this MakeDistOrigin function to make many functions. I could make one distance to origin function that computes a distance to one origin. I could make another one that computes the distance to a different origin. So I can make as many as I want with different origins. And so, the origin is now built in to the return function.

Okay, so now, let's look at how we might use this. In my main, I'm going to create two functions. I say, look, I want to have origins. I want to have origin 0,0 and another origin at 2,2, for one reason or another in my problem, I need two different origins. So Dist1 is going to be the function that computes the distance to origins 0,0. Dist2 is going to be the function that computes the distance to origin 2,2. And you can see me define those in the first two lines. I just call MakeDistOrigin with 0,0, and then 2,2. Then the last two lines, the print lines, they just print, they compute the distance from 2,2 to the origins. So the first one uses Dist1. So it computes the distance from 2,2 to the origin 0,0, which is about 2.1 something. So that's what we'd print out. Where the second one computes the distance from 2,2 to origin 2,2, which it should be zero. So you print that out and you get zero for that. So what we've done is we've made two special-purpose functions. And we made them special-purpose by giving them parameters, specifically the origins, right? The origins are the parameters that we use to make the function. So this is somethings that's kind of a cool use of returning a function, right? A function can create a new function that serves sort of a catered, special purpose.

Now, every function has an environment.

An environment Is a set of all the names that are valid inside the function, okay? All the variables and other things the you define that you can refer to inside that function. So this includes all the names that are defined locally in the function, any variables that you create in the function. But also, it uses Lexical Scoping. Is lexically scoped. So what that means is the function can access variables that are defined in the block where the function is defined. So in the example code down there, we got this function foo. Now, the variables that I have highlighted in red, those are all inside its environment. So let's take z. z is defined inside foo, so clearly, it's within foo's environment. Actually, notice how I use the words scope to refer to environment? People do that all the time. It's very common. It's not technically correct. I think you're supposed to use the term environment rather than scope. So anyway, the variable z is defined inside foo. So it's inside foo's environment. The variable y is a local variable to foo. It's one of it's parameters, right? So that is also inside the environment of foo. So foo can access y and z. Now, the variable x is defined in the same block as the function foo is defined, so foo can see that variable x, too. So let's say this whole, this piece of code is all defined inside another function, right? This variable x is defined in the same place, the same block, as foo is defined, and so foo can have access to that, too. So all those variables highlighted in red are within the environment of foo. All the variables that foo has access to and can use when it's executed.

So that's what environment is, and that's important. When you start passing around functions as arguments, the environment goes along with the function.

So, this term Closure. A Closure is a function plus its environment, right? Together. In fact, in Go, I think it's implemented actually directly as a struct. You have a pointer to the function and pointed to the environment, and they're put together. So when you pass a function as an argument to another function, you're also passing its environment with it. So what that means is when you eventually execute this function that you just passed, it still has access to its environment, the environment where it was defined. So what implications does this have? Sometimes it makes figuring out the variable values kind of confusing. But just remember that the closure, the environment to that function, goes with the function when you pass it as an argument. So let's take an example. This, again, is a make distant to origin function, right? Now, the function that it defines fn equals func, that function, that function, it has and environment. And notice that o underscore x, and o underscore y are part of its environment, okay? That's important, right? They are a part of its environment. So when you execute that function later, so this make distance to origin, it's going to return this function. And later on, when that function that you return gets executed, it still has access to the same environment. So o_x and o_y, that it had when it was defined, they were paths to make distance to origin, those variables are still accessible to this function when you call it later, okay? So that's why, so basically what I'm saying is it remembers these origin values. The o_x, o_y, that gets carried with the function. And so, when you execute the function, it's still using the same origin values, o_x, o_y when it gets executed. So that's called a closure. And when you pass a function as an argument, you pass this closure, the function plus its environment together, they go together. So you have to remember that when you're trying to figure out how these functions get evaluated. Where are the variables are coming from, they're coming from this closure, from where it was defined, because that's how Golang is scoped. It's lexically scoped, so it gets its environment from where it was defined.



# M2.2.2 - Variadic and Deferred
Module two, functions and organization. Topic 2.1 variadic and deferred. So, we've been talking about functions generally, we're going to talk about a few more variations on functions about how you can pass the arguments and how you can get them to execute different times. One useful tool is to be able to pass a variable number of arguments to the function. So, normally when you define a function, you have to hard code the arguments that it takes. So, if it takes three arguments, you list them, comma-separated inside parentheses. But sometimes you want to make a function that takes a variable number of arguments. So, there are a lot of functions like this, maybe you want to take a number of integers and you don't know how many integers. If you take two, you take 10, you can still work with them and do the same thing with the whole set of integers regardless of how many is taken. So, in that case, you'd like to be able to pass it a variable number of arguments, you can do that using this ellipsis character, not character really but ellipsis it's just three dots. Three period dots in a row and you put that there inside the argument list to specify that you want to have a variable number of arguments. Inside the function when you get this argument it looks like a slice. So, if we look at the function there, it's called getMax and its supposed to get the maximum integer out of a set of integers that you pass it as an argument. So, if you pass it two integers to 10 integers or whatever it is, it should go through all those integers, find the greatest one and return that. So, we want to be able to take a variable number of arguments, so you can see highlighted in red. I say, vals...int. So, it takes an integer but that "..." before integer, means it can take as many integers as you want to take. So, then inside the function, this vals argument is treated like a slice of integers. So, the function just basically you can see what it does it, just goes through this whole, you can see the for loop. It goes through the range of vals so it just iterate through all the integers inside vals, finds the biggest one, sets max v to whichever one is the biggest. Then in the end it returns max v. So, this is just a useful tool you can take a variable number of arguments just use this ellipsis is... inside the argument list and you can treat the parameter just like a slice. Now, another variation on that is let's say you got one of these variadic functions, it takes a variable number of arguments, you can pass it a comma-separated list of arguments. So, say for my getMax, I want to pass it five integers. I could pass it one comma, two comma, three comma, four comma, five, as many as I want. Which is what I do actually in this example right here, you can see getMax, one, three, six , four. I can make that list as long as I want. But another way to pass a variable number of arguments is to just pass it a slice. So, that one, three, six, four, that could already be pre-packaged in a slice and then you could pass the slice to this getMax function. So, that's what I'm doing below, vslice myslice is equal to slice of one, three, six, four. Then I pass that in the last line where I do the print line, I call getMax and I pass it vslice which is my slice. Then notice when I do that, that right after the word vslice, I have the ellipses again. So "..." you have to put that there so that it knows it instead of passing a comma separated sequence of arguments, this vslice is meant to be a slice of all the arguments put together. But once you do that, you can just pass the entire slice to the function and it works fine. So, another thing that is sometimes useful with functions is to have a deferred function call. Deferred functions mean that they don't get executed right where they're called, they get executed later. So, when the surrounding function completes, they get executed. So, typically use this for cleanup activities so, say you're doing something opening files or doing whatever you're doing, maybe you'll have a deferred function which closes all the files at the end or something like that. So, this function doesn't actually get called until the surrounding function is done and say you're done with all the files that you're interested in, then it gets called as you're exiting and closes all the files. So, it does some kinda clean up activity, so this is a common thing to use it for, for this type of cleanups afterwards. So, as an example we've got our main function right here. First thing we do, all you do to do the differ, is just put the keyword differ in front of the function call. So, here we got differ print line. So, defer fmt print line bye. Now and then the next line is just fmt println Hello. Now if they were executed in the order that they're written, you'd print bye and then hello. But of course since we deferred it, what will happen is hello will get executed first. Then defer will not be executed until the main function, the surrounding function actually completes. So, what will actually get printed is hello and then bye. So, one thing to remember about these deferred function calls, is that the arguments are not evaluated in a deferred way. The arguments are evaluated immediately but the call is deferred. So, what does that mean? Sometimes it doesn't mean anything. If you just pass the function some kind of a fixed argument that can't change that doesn't meet evaluation, that doesn't mean anything. But if you pass it an argument that needs to be evaluated, you have to note that argument is evaluated right there where the first statement is. Not later when the call actually happens. So, as to show this, you got a main. This main, you can see in there there's a defer print line and then there's the fmt.println at the end. But there's also a variable called i, i set it equal to one. Now when I do the differ, I say print line i +1. Now at the point where that the first statement is, i =1, so i +1=2, so two should get printed. But then noticed that the line after that, I say, i++, then I'd say I print hello. Remember that differ will not be executed until later after the whole main is complete. So, by the time that differed print line executes, the value of i will actually be two. Because i starts off at one as one, it gets incremented plus plus and so it should be two by the time you actually execute that differed statement. Then if you would evaluate the argument at that time, you would say i+1, 2+1, three, a three would get printed. What actually gets printed as a two, because that i +1 is evaluated right there when the first statement, when it hits the differed statement, the i +1 is evaluated and at that time, i is a one, so i +1 is a two. So, later when the differed statement actually gets executed it's still prints a two.