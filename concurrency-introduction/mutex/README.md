# Mutex


If you are working with concurrent applications, you have to deal with more than one resource potentially accessing some memory location. This is usually called **race condition**.

In simpler terms, a race condition is similar to that moment where two people try to get the last piece of pizza at exactly the same time and their hands collide. Replace the pizza with a variable and their hands with Goroutines and we'll have a perfect analogy.

There is one character at the dinner table to solve this issues, a father or mother. They have kept the pizza on a different table and we have to ask for permission to stand up before getting our slice of pizza. It doesn't matter if all the kids ask at the same time, they will only allow one kid to stand.
Well, a mutex is like our parents. They'll control who can access the pizza - In our case, a variable - and they won't allow anyone else to access it.

To use a mutex, we have to actively lock it; if it's already locked (another Goroutine is using it), we'll have to wait until it's unlocked again. Once we get access to the mutex, we can lock it again, do whatever modifications are needed, and unlock it again. 

## Run example
``make mutex``


# Race detector flag

We already know what is a race condition - it used when  two processes try to access the same resource at the same time with one or more writing operations (both processes writing or one process writing while the other reads) involved at that precise moment.

Go has a very handy tool to help diagnose race conditions, that you can run in your tests or your main application directly

````
go run -race ./concurrency-introduction/mutex/cli/main.go
````

The flag ``-race`` will detect if there's a potential race condition. Try to comment `Lock()` and `Unlock()` and re-run `make mutex` and see the errors


## Note
Notice that if we reduce the number of processors to just one, we will not have any `race condition error`, because its impossible to occur a race condition.