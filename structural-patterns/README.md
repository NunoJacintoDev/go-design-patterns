[Design Patterns](../README.md)

# Structural Patterns
- [Composite](./composite/)
- [Adapter](./adapter/)
- [Bridge](./bridge/)
- [Proxy](./proxy/)
- [Decorator](./decorator/)
- [Facade](./facade/)
- [Flyweight](./flyweight/)


# Decorator vs Proxy
[ TLDR: Decorator (runtime) vs (compile-time) Proxy ]

In the Decorator pattern, we decorate a type dynamically. This means that the decoration may or may not be there, or it may be composed of one or many types. If you remember, the Proxy pattern wraps a type in a similar fashion, but it does so at compile time and it's more like a way to access some type.

# Singleton vs Flyweight
With the Singleton pattern, we ensure that the same type is created only once. Also, the Singleton pattern is a Creational pattern. With Flyweight, which is a Structural pattern, we aren't worried about how the objects are created, but about how to structure a type to contain heavy information in a light way.