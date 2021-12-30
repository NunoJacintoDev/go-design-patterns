[Design Patterns](../../README.md) > [Structural Patterns](../README.md)

# Bridge
The Bridge pattern decouples an abstraction (an object) from its implementation (what the object does) so that the two can vary independently.

## Objectives
- Bring flexibility to a struct that often change.
- By knowing the inputs and outputs of a method, it allows us to change code without knowing too much about it and leaving the freedom for both sides to be modified more easily


## Example Use Case
### Two printers and two ways of printing for each
- A ```PrinterAPI``` that accepts a message to print
- An implementation of the API that simply prints the message to the console
- An implementation of the API that prints to an ```io.Writer``` interface
- A ```Printer``` abstraction with a ```Print``` method to implement in printing types
- A ```normal``` printer object, which will implement the ```Printer``` and the ```PrinterAPI``` interface
- The ```normal``` printer will forward the message directly to the implementation
- A ```Packt``` printer, which will implement the ```Printer``` abstraction and the ```PrinterAPI``` interface
- The ```Packt``` printer will append the message ```Message from Packt:``` to all prints
