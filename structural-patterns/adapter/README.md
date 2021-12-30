[Design Patterns](../../) > [Structural Patterns](../)

# Adapter
The Adapter pattern is very useful when, for example, an interface gets outdated and it's not possible to replace it easily or fast. Instead, you create a new interface to deal with the current needs of your application, which, under the hood, uses implementations of the old interface.
Adapter also helps us to maintain the open/closed principle in our apps, making them more predictable too. They also allow us to write code which uses some base that we can't modify.

## Objectives
The Adapter design pattern will help you fit the needs of two parts of the code that are incompatible at first. This is the key to being kept in mind when deciding if the Adapter pattern is a good design for your problem-two interfaces that are incompatible, but which must work together, are good candidates for an Adapter pattern (but they could also use the facade pattern, for example).


## Go most common use of Adapter pattern
- http.Handler
- io.Writer
- io.Reader
- io.Pipe()