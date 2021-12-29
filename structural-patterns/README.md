[Design Patterns](../README.md)

# Structural Patterns
- [Composite](./composite/README.md)
- [Adapter](./adapter/README.md)
- [Bridge](./bridge/README.md)
- [Proxy](./proxy/README.md)
- [Decorator](./decorator/README.md)
- [Facade](./facade/README.md)
- [Flyweight](./flyweight/README.md)


## Decorator vs Proxy
 - TLDR: Decorator (runtime) vs (compile-time) Proxy

    In the Decorator pattern, we decorate a type dynamically. This means that the decoration may or may not be there, or it may be composed of one or many types. If you remember, the Proxy pattern wraps a type in a similar fashion, but it does so at compile time and it's more like a way to access some type.