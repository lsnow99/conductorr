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

### Data Types and Literals

CSL has supports the following datatypes:
- Integers
- Strings
- Booleans
- Lists

#### Integers

All CSL integers are 64-bit numbers. They can be negative, positive, or zero. One special feature of CSL is the ability to use suffixes that are a shorthand to multiply by an order of magnitude. For example, `3G` is equivalent to `3000000000`.

The full list of available suffixes is below:
- `G` --> 1000000000
- `Gi` --> 2^30
- `M` --> 1000000
- `Mi` --> 2^20
- `K` --> 1000
- `Ki` --> 2^10
- `B` --> 8

#### Strings

String literals must be enclosed in double quotes (`"`)

#### Booleans

Boolean literals are `true` and `false`, and are case-sensitive

#### Lists

Lists are implicitly defined any time expressions are joined within parentheses such as `(1 4 "hello" 5 2 "world" true)`. List elements are not required to have the same data type. Single-value lists may be treated as regular values or lists. (ie, `(+ (1) (2))` is equivalent to `(+ 1 2)` and `(in 1 (1))` evaluates correctly as well).

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
- `(>< x y v ...)` Returns true iff for all integer arguments `v` that `v` ∈ (`x`, `y`)
- `(><= x y v ...)` Returns true iff for all integer arguments `v` that `v` ∈ (`x`, `y`]
- `(>=< x y v ...)` Returns true iff for all integer arguments `v` that `v` ∈ [`x`, `y`)
- `(>=<= x y v ...)` Returns true iff for all integer arguments `v` that `v` ∈ [`x`, `y`]
- `(<> x y v ...)` Returns true iff for all integer arguments `v` that `v` ∈ [-2<sup>63</sup>, x)∩(`y`, 2<sup>63</sup>-1]
- `(<=>= x y v ...)` Returns true iff for all integer arguments `v` that `v` ∈ [-2<sup>63</sup>, `x`]∩[`y`, 2<sup>63</sup>-1]
- `(<=> x y v ...)` Returns true iff for all integer arguments `v` that `v` ∈ [-2<sup>63</sup>, `x`]∩(`y`, 2<sup>63</sup>-1]
- `(<>= x y v ...)` Returns true iff for all integer arguments `v` that `v` ∈ [-2<sup>63</sup>, `x`)∩[`y`, 2<sup>63</sup>-1]
- `(eq x y ...)` Returns true if all arguments are equal to each other in both type and value. In the case of lists, the elements must be in the same order to be considered equal. (Under the hood we are just using Go's `reflect.DeepEqual`)
- `(define x expr)` Defines a variable `x` initialized with the result of `expr`
- `(in v l)` Returns true if the value of `v` appears in list `l`
- `(in s1 s2)` Returns true iff `s1` is a substring of `s2`
- `(nth i l)` Returns the `i`th value in list `l` or error if out of bounds 
- `(nths i s)` Returns the `i`th character in the string `s`
- `(len l)` Returns the length of list `l`
- `(lens s)` Returns the length of string `s`
- `(append l v ...)` Appends `v` and all subsequent arguments in order to the right end of list `l`. If the list `l` does not yet exist, then it is initialized
- `(appendleft l v ...)` Appends `v` and all subsequent arguments in order to the left end of list `l`. If the list `l` does not yet exist, then it is initialized
- `(pop l ...)` Removes the rightmost element in list `l`, returning the removed element
- `(popleft l ...)` Removes the leftmost element in list `l`, returning the removed element
- `(if p expr1 expr2)` Returns the result of `expr1` if `p` evaluates to `true`, and returns the result of `expr2` otherwise
- `(and p ...)` Returns the result of a conditional AND applied to all operands in the expression. Returns `true` if there is only one operand
- `(or p ...)` Returns the result of a conditional OR applied to all operands in the expression. Returns `true` if there is only one operand
- `(not p)` Returns the result of conditional inverse applied to p
- `(join s v ...)` Joins elements elements `v1`, `v2`, ... into a string using separator `s`. Does not add a separator to the end of the resultant string. Non-string elements such as integers will display using `fmt.Sprintf(%v, v)`
- `(split s d)` Splits a string into a list of strings using delimiter `d`

> NOTE: All of the above functions are evaluated *eagerly*, meaning their arguments are evaluated before the function is called. The exceptions to this are the `if`, `append`, `appendleft`, `pop`, and `popleft` functions, which only evaluate their arguments as they are needed.
