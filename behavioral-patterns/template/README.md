[Design Patterns](../../README.md) > [Behavioral Patterns](../README.md)

The Template pattern is one of those widely used patterns that are incredibly useful, especially when writing libraries and frameworks. The idea is to provide a user some way to execute code within an algorithm.
In this section, we will see how to write idiomatic Go Template patterns and see some Go source code where it's wisely used. We will write an algorithm of three steps where the second step is delegated to the user while the first and third aren't. The first and third steps on the algorithm represent the template.

# Template Pattern
While with the Strategy pattern we were encapsulating algorithm implementation in different strategies, with the `Template pattern we will try to achieve something similar but with just part of the algorithm`.

`The Template design pattern lets the user write a part of an algorithm while the rest is executed by the abstraction`. This is common when creating libraries to ease in some complex task or when reusability of some algorithm is compromised by only a part of it.
Imagine, for example, that we have a long transaction of HTTP requests. We have to perform the following steps:

1. Authenticate user.
2. Authorize him.
3. Retrieve some details from a database.
4. *Make some modification*
5. Send the details back in a new request.

It wouldn't make sense to repeat steps 1 to 5 in the user's code every time he needs to modify something on the database. Instead, steps 1, 2, 3, and 5 will be abstracted in the same algorithm that receives an interface with whatever the fifth step needs to finish the transaction. It doesn't need to be a interface either, it could be a callback.

## Objectives
The Template design pattern is all about `reusability` and giving `responsibilities to the user`. So the objectives for this pattern are following:
- Defer a part of an algorithm of the library to a the user
- Improve reusability by abstracting the parts of the code that are not common between executions

# Example - algorithm with a deferred step
In our first example, we are going to write an algorithm that is composed of three steps and each of them returns a message. The first and third steps are controlled by the Template and just the second step is deferred to the user.

1. Each step in the algorithm must return a string.
2. The first step is a method called first() and returns the string hello.
3. The third step is a method called third() and returns the string template.
4. The second step is whatever string the user wants to return but it's defined by the MessageRetriever interface that has a Message() string method.
5. The algorithm is executed sequentially by a method called ExecuteAlgorithm and returns the strings returned by each step joined in a single string by a space.


# Summary
We wanted to put a lot of focus on this pattern because it is very important when `developing libraries and frameworks and allows a lot of flexibility and control to users of our library`.

We have also seen again that it's very common to mix patterns to `provide flexibility to the users, not only in a behavioral way but also structural`. This will come very handy when working with concurrent apps where we need to restrict access to parts of our code to avoid races.