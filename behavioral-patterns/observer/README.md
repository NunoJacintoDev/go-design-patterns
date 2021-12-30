[Design Patterns](../../README.md) > [Behavioral Patterns](../README.md)


# Observer Pattern
Also known as **publish/subscriber or publish/listener pattern**

The idea behind the Observer pattern is simple to subscribe to some event that will trigger some behavior on many subscribed types. Why is this so interesting? Because we uncouple an event from its possible handlers.

For example, imagine a login button. We could code that when the user clicks the button, the button color changes, an action is executed, and a form check is performed in the background. But with the Observer pattern, the type that changes the color will subscribe to the event of the clicking of the button. The type that checks the form and the type that performs an action will subscribe to this event too.

## Objectives
- Provide an event-driven architecture where one event can trigger one or more actions
- Uncouple the actions that are performed from the event that triggers them
- Provide more than one event that triggers the same action

<br>

# Example - The notifier
`Publisher` struct, which is the one that triggers an event so it must **accept new observers and remove them if necessary**. When the `Publisher` struct is `triggered`, it must `notify all its observers` of the new event with the data associated.


# Use Cases
The Observer pattern is commonly used in UI's. Android programming is filled with Observer patterns so that the Android SDK can delegate the actions to be performed by the programmers creating an app.

