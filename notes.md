# Variables
Outside a function: `var dog string = 'niko'`
Inside a function: `dog := 'niko'`

# Types
## Array vs Slice
- Array: fixed length  
- Slice: length can change. All types have to be the same:  
`dogs := []string{"niko", "lola"}`
  slice[startPositionIncluded : upToNotIncluded]
### **append**
`vars = append(var1, var2)`

### Iteration
`for i, card:= range cards { ... }`

# Functions
```
func newDog() string { ... }
```

# Custom type declarations
`type deck []string`

## Receiver
When we want to be able to call the receiver with a method.
```
func (d deck) print() { ... }
```
A variable of type `deck` now gets access to the `print` method. The `d` (receiver argument) is similar to `self`.