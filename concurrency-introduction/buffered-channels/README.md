# Buffered channels

A buffered channel works in a similar way to default unbuffered channels. You also pass and take values from them by using the arrows, but, unlike unbuffered channels, senders don't need to wait until some Goroutine picks the data that they are sending:

## Run example
``make buffered-channel``
Notice that it doesn't print `Hello World! 2` - That's because the channel only has a size of one, the second message will block the goroutine execution