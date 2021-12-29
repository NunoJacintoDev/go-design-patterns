[Design Patterns](../../README.md) > [Structural Patterns](../README.md)

# Facade
A facade, in architectural terms, is the front wall that hides the rooms and corridors of a building. It protects its inhabitants from cold and rain, and provides them privacy. It orders and divides the dwellings.
The Facade design pattern does the same, but in our code. It shields the code from unwanted access, orders some calls, and hides the complexity scope from the user.

## Objectives
- When you want to decrease the complexity of some parts of our code. You hide that complexity behind the facade by providing a more easy-to-use method.
- When you want to group actions that are cross-related in a single place.
- When you want to build a library so that others can use your products without worrying about how it all works.

## Example - HTTP REST API
As an example, we are going to take the first steps toward writing our own library that accesses OpenWeatherMaps service. In case you are not familiar with OpenWeatherMap service, it is an HTTP service that provides you with live information about weather, as well as historical data on it. The HTTP REST API is very easy to use, and will be a good example on how to create a Facade pattern for hiding the complexity of the network connections behind the REST service.
### Aceptance criteria
- Provide a single type to access the data. All information retrieved from OpenWeatherMap service will pass through it.
- Create a way to get the weather data for some city of some country.
- Create a way to get the weather data for some latitude and longitude position.
- Only second and thrird point must be visible outside of the package; everything else must be hidden (including all connection-related data).