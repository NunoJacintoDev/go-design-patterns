[Design Patterns](../../README.md) > [Creational Patterns](../README.md)

# Prototype design pattern
The aim of the Prototype pattern is to have an object or a set of objects that is already created at compilation time, but which you can clone as many times as you want at runtime. This is useful, for example, as a default template for a user who has just registered with your webpage or a default pricing plan in some service. The key difference between this and a Builder pattern is that objects are cloned for the user instead of building them at runtime. You can also build a cache-like solution, storing information using a prototype.

# Objective
- Maintain a set of objects that will be cloned to create new instances
- Provide a default value of some type to start working on top of it
- Free CPU of complex object initialization to take more memory resources

# Wrap up
The Prototype pattern is a powerful tool to build caches and default objects. You have probably realized too that some patterns can overlap a bit, but they have small differences that make them more appropriate in some cases and not so much in others.