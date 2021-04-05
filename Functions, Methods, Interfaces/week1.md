# NOTE
There is a small amount of controversy surrounding the expressions "pass by pointer" and "pass by reference." Just search the web for "Pass by Pointer vs. Pass by Reference in Go Programming" and you'll see for yourself! Be aware that our instructor uses "pass by reference" where some Go users would say "pass by pointer" or "pass a pointer."

# M1.1.1 - Why Use Functions?
- Function: Set of instructions with a name, name is optional
- main function in Go, where the execution starts
- function contents are between {}
- Calling a function is to execute the function
- () after function name will have the parameters
- func nameOfFunc(argument string) returnValue string{functioDoesSomething}

- Why use a function
    - Reusability: don't have to rewrite the same code
    - Saves time for coder
    - Shrinks size of code
    - Can use the function across projects
    - Others could import code as a library and use the code
    - Abstraction: hiding details that are less important
        - Helps simplify things
        - Group things
        - Hide details
        - Only need to understand the input-output behavior
    - Improve understandability of code
    - Organization
        - Functions organized in a hierarchy
    - Ability to glance at code and understand how it works
- Naming a function well helps with understandability

# M1.1.2 - Function Parameters and Return Values
- Parameters: function inputs - function definition
- Argument: Data applied as parameters - function call
- A function doesn't need parameters/arguments
- A function can accept multiple parameters
- Parameter can come from different places
- Don't have to give a parameter a type
- Comma separate multiple arguments of the same type

- Return values: function output - function definition
- A function can have multiple outputs
- Return values have to have a type
- Return value from function can be assigned to a variable

# M1.1.3 - Call by Value, Reference
- Call by value: how arguments are passed to parameters
    - Arguments are passed as parameters
    - The data used in a function is a copy of the original, not the original value
- Important: function being called can't interfere with the original variable
    - Modifying a parameter in a function doesn't have an effect outside of the function
- Advantage: Data encapsulation: functions can not alter the variable
    - Limits the propagation of erros
    - Localizes errors
- Disadvantage: Copying time
    - Time to copy the argument into the parameters

- Call by reference
    - Not built into Go
    - Manually do it by passing a pointer instead of an argument
        - Pointer is a reference
    - *y *int, pointer to y
    - &x when calling the function, passes the pointer to x, not a copy of x.
    - A modification to the pointer, is changing the value at the location of storage,
    - Not passing the integer or data to the function, the pointer/reference to the data is being passed. Allowing the ability to alter the variable
- Advantage: copying time, could copy only a portion of an argument instead of having to copy the whole argument by copying the pointer
- Disadvantage: Data encapsulation, a bug inside a function could alter variables outside of the function and create bugs, errors

- pass by reference, if you definitely want the function to modify the variables in the calling function
# M1.1.4 - Passing Arrays and Slices
- Pass array as an argument, arguments are all copied (call by value)
- Instead of copying the whole array, save time and memory space by
    - Employ call by reference, use a pointer to the array
        - Using a pointer to the array means the function can modify the array
        - Dereferencing and referencing operations to pass array pointers and use array pointers. * and &
    - A slice allows you to only copy a portion of the array, it is a window to an array
        - Using a slice copies the pointer
        - A slice isn't an array its a structure that contains:
            - A pointer to the array or to the start of the slice, in the array
            - The length of the slice
            - The capacity of the slice
        - A slice helps do the call by reference stuff, which out having to do the call by reference stuff... the deferencing and referencing. Because the slice has a pointer to the array and allows modifying the slice and the underlying array the slice is representing
# M1.2.1 - Well-Written Functions
- Goal for good code. Organized and easy to understand
- Understandable code is good for debugging and maintenance
- Function definition impacts understandability
- Understandability: If you want to find a feature, can you find it quickly?
- Organization makes it easier to find things

- With no global variables, parameters from straight from functions, easier to debug, you know where the parameter came from

# M1.2.2 - Guidelines for Functions
- Function naming: Good descriptive name
- Parameter naming: Good descriptive name

# M1.2.3 - Function Guidelines+
- Functions aren't too complex
- Function length: shorter will be less complex, sometimes
- Function call hierarchy to simplify functions. Function that calls a function that calls a function...
- Control-flow complexity, paths from top of the function to the bottom, way to judge complexity of function
- separate the conditionals into different functions and reduce the max complexity,the max number of control-flow paths. Reduce that overall, making the code easier to debug.