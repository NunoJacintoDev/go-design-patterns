[Design Patterns](../../README.md) > [Behavioral Patterns](../README.md)

**"An object kept as a reminder of a person or event."**

# Memento Pattern
The meaning of memento is very similar to the functionality it provides in design patterns. Basically, we'll have `a type with some state and we want to be able to save milestones of its state`. Having a finite amount of states saved, we can recover them if necessary for a variety of tasks-undo operations, historic, and so on.

The Memento design pattern `usually has three actors`:
- `Memento`: A type that stores the type we want to save. Usually, we won't store the business type directly and we provide an extra layer of abstraction through this type.
- `Originator`: A type that is in charge of creating mementos and storing the current active state. We said that the Memento type wraps states of the business type and we use originator as the creator of mementos.
- `Care Taker`: A type that stores the list of mementos that can have the logic to store them in a database or to not store more than a specified number of them.

## Objectives
Memento is all about a `sequence of actions over time`, say to `undo one or two operations` or to provide some kind of transactionality to some application.

Memento provides the foundations for many tasks, but its main objectives could be defined as follows:
- `Capture an object state without modifying the object itself`
- `Save a limited amount of states so we can retrieve them later`

# Example - Simple example with strings
We will develop a simple example using a string as the state we want to save. This way, we will focus on the common Memento pattern implementations before making it a bit more complex with a new example.
The string, stored in a field of a State instance, will be modified and we will be able to undo the operations done in this state.

## Requirements and acceptance criteria
We are constantly talking about state; all in all, the Memento pattern is about storing and retrieving states. Our acceptance criteria must be all about states:
1. We need to store a finite amount of states of type string.
2. We need a way to restore the current stored state to one of the state list.
With these two simple requirements, we can already start writing some tests for this example.

# Summary
With the Memento pattern, we have learned a powerful way to create `undoable operations` that are very useful when `writing UI applications` but also when you have to develop `transactional operations`. In any case, the situation is the same: you need a `Memento`, an `Originator`, and a `caretaker` actor.