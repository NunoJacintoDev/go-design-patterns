[Design Patterns](../../README.md) > [Behavioral Patterns](../README.md)

**Visitor pattern is very useful when you want to abstract away some functionality from a set of objects.It visits our object to perform operations on it**

# Visitor Pattern
In the Visitor design pattern, we are trying to separate the logic needed to work with a specific object outside of the object itself. So we could have many different visitors that do some things to specific types.
For example, imagine that we have a log writer that writes to console. We could make the logger "visitable" so that you can prepend any text to each log. We could write a Visitor pattern that prepends the date, the time, and the hostname to a field stored in the object.

## Objectives
- To separate the algorithm of some type from its implementation within some other type
- To improve the flexibility of some types by using them with little or no logic at all so all new functionality can be added without altering the object structure
- To fix a structure or behavior that would break the open/closed principle in a type

### **Open/close Principle**
In computer science, the open/closed principle states that: `Entities should be open for extension but closed for modification`. This simple state has lots of implications that allows building more maintainable software and less prone to errors. The Visitor pattern helps us to delegate some commonly changing algorithm from a type that we need it to be "stable" to an external type that can change often without affecting our original one.


# Example - A log appender
We will create a Visitor that appends different information to the types it "visits".

## Acceptance criteria
To effectively use the Visitor design pattern, we must have two roles - a `Visitor` and a `Visitable`. The `Visitor` is the type that `will act within a Visitable` type. So a Visitable interface implementation has an algorithm `detached` to the Visitor type:
1. We need two message loggers: `MessageA` and `MessageB` that will print a message with an `A:` or a `B:` respectively before the message.
2. We need a Visitor able to modify the message to be printed. It will append the text `Visited A` or `Visited B` to them, respectively.



# Example - Online Shop
Emulate an online shop with a few products. The products will have plain types, with just fields and we will make a couple of visitors to deal with them in the group.
Using Visitor Pattern we are able to create two different visitors that implement the example `Visitor` interface - `Price Visitor` that visits `Product` instances  and sum all the prices; and a `Name Printer Visitor` that visits `Product` instances and makes a list with each name of visited product.


# Attention!
We have seen a powerful abstraction to add new algorithms to some types. However, because of the lack of overloading in Go, this pattern could be limiting in some aspects (we have seen it in the first example, where we had to create the VisitA and VisitB implementations). In the second example, we haven't dealt with this limitation because we have used an interface to the Visit method of the Visitor struct, but we just used one type of visitor (ProductInfoRetriever) and we would have the same problem if we implemented a Visit method for a second type, which is one of the objectives of the original Gang of Four design patterns.