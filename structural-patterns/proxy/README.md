[Design Patterns](../../) > [Structural Patterns](../)

# Proxy
The Proxy pattern usually wraps an object to hide some of its characteristics. These characteristics could be the fact that it is a remote object (remote proxy), a very heavy object such as a very big image or the dump of a terabyte database (virtual proxy), or a restricted access object (protection proxy).


## Objectives
- Hide an object behind the proxy so the features can be hidden, restricted, and so on
- Provide a new abstraction layer that is easy to work with, and can be changed easily

# Usage
Wrap proxies around types that need some intermediate action, like giving authorization to the user or providing access to a database, like in our example.
For example to separate application needs from database needs. If our application accesses the database too much, a solution for this is not in your database. Remember that the Proxy uses the same interface as the type it wraps, and, for the user, there shouldn't be any difference between the two.