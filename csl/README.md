# CSL (Conductorr-Specific Language)

CSL is a stripped-down Lisp implementation intended to allow users of Conductorr to extend its capabilities and finely tune release processing.

The parsing capabilities are largely borrowed from [Jan Andersson](https://github.com/janne)'s [go-lisp](https://github.com/janne/go-lisp) project.

The main differences between go-lisp and CSL are that CSL is *not* Turing-Complete, such that loops or recursion are impossible to implement. This allows us to keep the language simple and avoid the halting problem when running user-defined scripts. This means user scripts in CSL cannot contain loops or function definitions. The other main difference is that CSL has a number of predefined functions that are specifically written to aid in the processing of Newznab releases. This helps avoid the loss of functionality due to the restricted iteration and function-defining capabilities of CSL.

## Syntax

If you are familiar with Lisp syntax, CSL follows it directly. Here is an example of a script in CSL:

```lisp
;; Define a variable
(define x 17)

;; Return the result of x * 37
(* x 37)
```
The value `629` is returned.

## Built-In Functions

- `(+ x y ...)` Adds all integer arguments together
- `(- x y ...)` Subtracts subsequent integer arguments from the first one
- `(* x y ...)` Multiplies all integer arguments together
- `(/ x y ...)` Divides subsequent integer arguments from the first one
- `(> x y ...)` Returns true iff `x` is strictly greater than all subsequent integer arguments
- `(>> x y ...)` Returns true iff each integer argument is strictly greater than the ones before it
- `(< x y ...)`  Returns true iff `x` is strictly less than all subsequent integer arguments
- `(<< x y ...)` Returns true iff each integer argument is strictly less than the ones before it
- `(>= x y ...)`  Returns true iff `x` is greater than or equal to all subsequent integer arguments
- `(>>= x y ...)` Returns true iff each integer argument is strictly greater than or equal to the ones before it
- `(<= x y ...)`  Returns true iff `x` is less than or equal to all subsequent integer arguments
- `(<<= x y ...)` Returns true iff each integer argument is strictly less than or equal to the ones before it
- `(eq x y ...)` Returns true if all arguments are equal to each other in both type and value. In the case of lists, the elements must be in the same order to be considered equal. (Under the hood we are just using Go's `reflect.DeepEqual`)
- `(define x expr)` Defines a variable `x` initialized with the result of `expr`
- `(in v l)` Returns true if the value of `v` appears in list `l`
- `(in s1 s2)` Returns true iff `s1` is a substring of `s2`
- `(nth i l)` Returns the `i`th value in list `l` or error if out of bounds 