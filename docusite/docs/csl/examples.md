# Examples

This page contains code examples designed to demonstrate the core functionality of CSL. For an exhaustive list of all available CSL functions you can refer to the [reference](/csl/reference.html#built-in-functions) page.

## Basics
### Hello World


```lisp
("Hello world")
; Hello world
```

### Arithmetic

```lisp
; Multiplication
(* 2 3)
; 6
; You can add more than 2 arguments
(* 2 3 4)
; 24

; Addition
(+ 4 10) 
; 14
; You can add more than 2 arguments
(+ 4 10 16)
; 30

; Division
(/ 24 4)
; 6
; You can add more than two arguments, such as in the following example it will be executed like ((24 / 4) / 3)
(/ 24 4 3)
; 2

; Subtraction
(- 10 5)
; 5
; You can add more than two arguments, such as in the following example it will be executed like ((20 - 10) - 4)
(- 20 10 4)
; 6
```

### String Manipulation

```lisp
; Length
(lenstr "Hello world")
; 11

; Nth character
(nthstr 4 "Hello world")
; o
```
> **NOTE**: Make sure to use the `lenstr` and `nthstr` functions instead of the `len` and `nth` functions when dealing with strings, as the latter two functions expect lists.

## Using Imports


