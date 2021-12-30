# Channels


Channels are the way we communicate between processes. We could be sharing a memory location and using mutexes to control the processes' access. But channels provide us with a more natural way to handle concurrent applications that also produces better concurrent designs in our programs.

## Run example
``make channel``