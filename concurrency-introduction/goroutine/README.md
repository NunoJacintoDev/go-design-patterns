# Goroutines

In Go, we achieve concurrency by working with **Goroutines**. They are like processes that run applications in a computer concurrently; in fact, the main loop of Go could be considered a Goroutine, too. Goroutines are used in places where we would use **actors**. They execute some logic and die (or keep looping if necessary).

But **Goroutines are not threads**. We can launch thousands of concurrent Goroutines, even millions. They are incredibly cheap, with a small growth stack. We will use Goroutines to execute code that we want to work concurrently.

For example, three calls to three services to compose a response can be designed concurrently with three Goroutines to do the service calls potentially in parallel and a fourth Goroutine to receive them and compose the response. What's the point here? That if we have a computer with four cores, we could potentially run this service call in parallel, but if we use a one-core computer, the design will still be correct and the calls will be executed concurrently in only one core. By designing concurrent applications, we don't need to worry about parallel execution.

Returning to the bike analogy, we were pushing the pedals of the bike with our two legs. That's two Goroutines concurrently pushing the pedals. When we use the tandem, we had a total of four Goroutines, possibly working in parallel. But we also have two hands to handle the front and rear brakes. That's a total of eight Goroutines for our two threads bike. Actually, we don't pedal when we brake and we don't brake when we pedal; that's a correct concurrent design. Our nervous system transports the information about when to stop pedaling and when to start braking.
**In Go, our nervous system is composed of channels**.

# Run example
``make waitgroup``