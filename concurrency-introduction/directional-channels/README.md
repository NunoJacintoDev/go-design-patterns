# Directional channels

One cool feature about Go channels is that, when we use them as parameters, we can restrict their directionality so that they can be used only to send or to receive. The compiler will complain if a channel is used in the restricted direction. This feature applies a new level of static typing to Go apps and makes code more understandable and more readable.

## Run example
``make directional-channel``

In the example, the line where we launch the new Goroutine `go func(ch chan<- string)` states that the channel passed to this function can only be used as an input channel, and you can't listen to it.

As you can see, the arrow is on the opposite side of the keyword `chan`, indicating an extracting operation from the channel. Keep in mind that the **channel arrow always points left**:
- To indicate a **receiving channel**, it must go on the left: `<-chan`
- To indicate a **inserting channel**, it must go on the right: `chan<-`