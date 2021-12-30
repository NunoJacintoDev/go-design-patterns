# Callbacks

If you have ever worked with languages like JavaScript that use them extensively, this section will be familiar to you. A callback is an anonymous function that will be executed within the context of a different function.


# Run example
``make callback``


## Callback Hell
The term callback hell is commonly used to refer to when many callbacks have been stacked within each other. This makes them difficult to reason with and handle when they grow too much. For example, using the same code as before, we could stack another asynchronous call with the contents that we previously printed to the console:

````
 func main() {
     wait.Add(1)

     toUpperAsync("Hello Callbacks!", func(v string) {
        toUpperAsync(fmt.Sprintf("Callback: %s\n", v), func(v string) {
            fmt.Printf("Callback within %s", v)
            wait.Done()
       })
    })

    println("Waiting async response...")
    wait.Wait()
}
````


In this case, we can assume that the outer callback will be executed before the inner one. That's why we don't need to add one more to the WaitGroup.

The point here is that we must be careful when using callbacks. In very complex systems, too many callbacks are hard to reason with and hard to deal with. But with care and rationality, they are powerful tools.