[Design Patterns](../../README.md) > [Behavioral Patterns](../README.md)

**It's a pattern that will be in between two types to exchange information**

# Mediator Pattern
One of the key objectives of any design pattern is to avoid tight coupling between objects, a particularly effective method to achieve it, when the application grows a lot, is the Mediator pattern.

The Mediator pattern is the perfect example of a pattern that is commonly used by every programmer without thinking very much about it.

Mediator pattern will act as the `type in charge of exchanging communication between two objects`. This way, the communicating `objects don't need to know each other` and can change more freely. The pattern that maintains which objects give what information is the `Mediator`.

## Objectives
- To provide loose coupling between two objects that must communicate between them
- To reduce the amount of dependencies of a particular type to the minimum by passing these needs to the Mediator pattern

# Example - A calculator
For the Mediator pattern, we are going to develop an extremely simple arithmetic calculator.
Our calculator will only do two very simple operations: sum and subtract.

## Acceptance criteria
1. Define an operation called `Sum` that takes a number and adds it to another number. (Only for 1 and 2)

# Uncoupling two types with the Mediator
We have carried out a disruptive example to try to think outside the box and reason deeply about the Mediator pattern. Tight coupling between entities in an app can become really complex to deal with in the future and allow more difficult refactoring if needed.
Just remember that the Mediator pattern is there to act as a managing type between two types that don't know about each other so that you can take one of the types without affecting the other and replace a type in a more easy and convenient way.