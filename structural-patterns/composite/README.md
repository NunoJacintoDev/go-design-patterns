[Design Patterns](../../README.md) > [Structural Patterns](../README.md)

# Composite

The Composite design pattern favors composition (commonly defined as a has a relationship) over inheritance (an is a relationship). The composition over inheritance approach has been a source of discussions among engineers since the nineties. We will learn how to create object structures by using a has a approach. All in all, Go doesn't have inheritance because it doesn't need it!

## Direct and Embedding Composition
With Go we can use two types of composition:
- Direct Composition
- Embedding Composition (as in [animal.go](./animal.go) )


## Which Approach?
- Structures [composite-swimmer-a.go](./composite-swimmer-a.go) 
- Interfaces [composite-swimmer-b.go](./composite-swimmer-b.go) 


## Acceptance criteria
- We must have an Athlete struct with a Train method
- We must have a Swimmer with a Swim method
- We must have an Animal struct with an Eat method
- We must have a Fish struct with a Swim method that is shared with the Swimmer, and not have inheritance or hierarchy issues


## Other use case - Binary Tree compositions
```
type Tree struct {
    LeafValue int
    Right     *Tree
    Left      *Tree
}
```
This is some kind of recursive compositing, and, because of the nature of recursivity, we must use pointers so that the compiler knows how much memory it must reserve for this struct. Our Tree struct stored a LeafValue object for each instance and a new Tree in its Right and Left fields.

With this structure, we could create an object like this:
```
root := Tree{
    LeafValue: 0,
    Right:&Tree{
      LeafValue: 5,
      Right: &1Tree{ 6, nil, nil },
      Left: nil,
    },
    Left:&Tree{ 4, nil, nil },
}
```


# ⚠️ Composite pattern vs inheritance
When using the Composite design pattern in Go, you must be very careful not to confuse it with inheritance. For example, when you embed a Parent struct within a Son struct, like in the following example:

```
type Parent struct {
     SomeField int
}

type Son struct {
     Parent
}
```
You cannot consider that the Son struct is also the Parent struct. What this means is that you cannot pass an instance of the Son struct to a function that is expecting a Parent struct like the following:
```
func GetParentField(p *Parent) int{
     fmt.Println(p.SomeField)
}
```
When you try to pass a Son instance to the GetParentField method, you will get the
following error message:

``` cannot use son (type Son) as type Parent in argument to GetParentField```

This, in fact, makes a lot of sense. What's the solution for this? Well, you can simply composite the Son struct with the parent without embedding so that you can access the Parent instance later:
```
type Son struct {
     P Parent
}
```
So now you could use the P field to pass it to the GetParentField method:
```
son := Son{}
GetParentField(son.P)
```