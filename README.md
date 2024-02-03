# Setlib (Work in progress)
Setlib allows you to define a collection of objects that behaves like a set. You can perform set theoretic computations on your sets and manage the elements in the sets. The elements of a set can be of any type that satisfies the comparable type as well as sets of sets.

## Installation
```bash
$ go get github.com/cjdupless/setlib
```

## Usage
```go
import sl "github.com/cjdupless/setlib"
```

## Iterations for first version
Only sets with non-pointer types will be considered.
### Iteration 1: Basic functionality of a set (Set)
1.1 Add an element to a set  
1.2 Remove an element of a set  
1.3 Check membership  
1.4 Retrieve the elements within a set  
1.5 Compute a union, intersection and difference between sets  
1.6 Unit tests  

### Iteration 2. Basic functionality for a set of sets (SSet)
2.1 Add an element to an sset  
2.2 Remove an element of an sset  
2.3 Check membership  
2.4 Retrieve the elements within an sset  
2.5 Compute a union, intersection and difference between ssets  
2.6 Unit tests  

### Iteration 3. Additional functionality
3.1 Equality  
3.2 Subset of/Superset of  
3.3 Power set  

