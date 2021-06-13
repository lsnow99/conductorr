# CSL (Conductorr-Specific Language)

CSL is a stripped-down Lisp implementation intended to allow users of Conductorr to extend its capabilities and finely tune release processing.

The parsing capabilities are largely borrowed from [Jan Andersson](https://github.com/janne)'s [go-lisp](https://github.com/janne/go-lisp) project.

The main differences between go-lisp and CSL are that CSL is *not* Turing-Complete, such that loops or recursion are impossible to implement. This allows us to keep the language simple and avoid the halting problem when running user-defined scripts. This means user scripts in CSL cannot contain loops or function definitions. The other main difference is that CSL has a number of predefined functions that are specifically written to aid in the processing of Newznab releases. This helps avoid the loss of functionality due to the restricted iteration and function-defining capabilities of CSL.

## Syntax

If you are familiar with Lisp syntax, CSL follows it directly. Here is an example of a script in CSL:



## Built-In Functions