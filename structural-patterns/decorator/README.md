[Design Patterns](../../) > [Structural Patterns](../)

# Decorator
The Decorator design pattern allows you to decorate an already existing type with more functional features without actually touching it. How is it possible? Well, it uses an approach similar to **matryoshka dolls**, where you have a small doll that you can put inside a doll of the same shape but bigger, and so on and so forth.
The Decorator type implements the same interface of the type it decorates, and stores an instance of that type in its members. This way, you can stack as many decorators (dolls) as you want by simply storing the old decorator in a field of the new one.

## Objectives
- When you need to add functionality to some code that you don't have access to, or you don't want to modify to avoid a negative effect on the code, and follow the open/close principle (like legacy code)
- When you want the functionality of an object to be created or altered dynamically, and the number of features is unknown and could grow fast


## Example - Decorating Pizza with Ingredients
- Acceptance Criteria
    - We must have the main interface that all decorators will implement. This interface will be called ```IngredientAdd```, and it will have the ```AddIngredient() string``` method.
    - We must have a core ```PizzaDecorator``` type (the decorator) that we will add ingredients to.
    - We must have an ingredient “onion” implementing the same ```IngredientAdd``` interface that will add the string ```onion``` to the returned pizza.
    - We must have a ingredient “meat” implementing the ```IngredientAdd``` interface that will add the string ```meat``` to the returned pizza.
    - When calling ```AddIngredient``` method on the top object, it must return a fully decorated ```pizza``` with the text ```Pizza with the following ingredients: meat, onion.```


## Real-life example - server middleware
- Decorator can be used combined with de adaptor pattern to create a server with log and authentication middleware (see Go Patterns book page 138)