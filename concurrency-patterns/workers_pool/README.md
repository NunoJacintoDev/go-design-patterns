[Design Patterns](../../README.md) > [Concurrency Patterns](../README.md)

One problem we may face with some of the approaches to concurrency is their unbounded context. We cannot let an app create an unlimited amount of Goroutines.
Goroutines are light, but the work they perform could be very heavy. A workers pool helps us to solve this problem.

#  Workers Pool
With a pool of workers, we want to bound the amount of Goroutines available so that we have a deeper control of the pool of resources. This is easy to achieve by creating a channel for each worker and having workers with either an idle or busy status.

## Objectives
Creating a Worker pool is all about resource control: CPU, RAM, time, connections, and so on. The workers pool design pattern helps us to do the following:
- Control access to shared resources using quotas
- Create a limited amount of Goroutines per app
- Provide more parallelism capabilities to other concurrent structures

# Example - Pool of pipelines

We want to launch a bounded number of pipelines so that the Go scheduler can try to process each request in parallel.
The idea is to control the number of Goroutines, stop them gracefully when the app has finished, and maximize parallelism using a concurrent structure without race conditions.

The pipeline we will use is similar to the one we used in **Concurrency Patters > Pipeline**

## Acceptance Criteria
1. When making a request with a string value (any), it must be uppercase.
2. Once the string is uppercase, a predefined text must be appended to it. This text should not be uppercase.
3. With the previous result, the worker ID must be prefixed to the final string.
4. The resulting string must be passed to a predefined handler.


# Wrapping up the Worker pool
With the workers pool, we have our first complex concurrent application that can be used in real-world production systems. It also has room to improve, but it is a very good design pattern to build concurrent bounded apps.

It is key that we always have the number of Goroutines that are being launched under control. While it's easy to launch thousands to achieve more parallelism in an app, we must be very careful that they don't have code that can hang them in an infinite loop, too.

With the workers pool, we can now fragment a simple operation in many parallel tasks. Think about it; this could achieve the same result with one simple call to fmt.Printf, but we have done a pipeline with it; then, we launched few instances of this pipeline and finally, distributed the workload between all those pipes.