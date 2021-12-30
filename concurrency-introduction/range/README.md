# Ranging over channels

We can use `range` keyword to `range` over channels, buffered or unbuffered.
Notice in the example that `range` keeps iterating until the channel is closed

Range is very useful in taking data from a channel, and it's commonly used in **fan-in patterns** where many different Goroutines send data to the same channel.

## Run example
```make channel-range```