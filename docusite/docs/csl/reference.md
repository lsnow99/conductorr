---
title: Reference
---

# CSL

CSL is a stripped-down Lisp implementation intended to allow users of Conductorr to extend its capabilities and finely tune release processing.

It stands for {{ standsFor }} (<a @click="orDoesIt" href="javascript:void(0);">or does it?</a>).

The parsing capabilities are largely borrowed from [Jan Andersson](https://github.com/janne)'s [go-lisp](https://github.com/janne/go-lisp) project.

The main differences between go-lisp and CSL are that CSL is *not* Turing-Complete, such that loops or recursion are impossible to implement. This allows us to keep the language simple and avoid the halting problem when running user-defined scripts. This means user scripts in CSL cannot contain loops or function definitions. The other main difference is that CSL has a number of predefined functions that are specifically written to aid in the processing of Newznab releases. This helps avoid the loss of functionality due to the restricted iteration and function-defining capabilities of CSL.

## Syntax

If you are familiar with Lisp syntax, CSL follows it directly. Here is an example of a script in CSL:

```lisp
; Define a variable
(define x 17)

; Return the result of x * 37
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

String literals must be enclosed in double quotes (`"`). To represent a double quote inside of your string literal, use `\"`. For example, `("The boy said, \"Hello, world\"")` evaluates to `The boy said, "Hello, world"`. 

All escape patterns:
- `\n` - A newline
- `\\` - A backslash
- `\"` - A double quote

#### Booleans

Boolean literals are `true` or `false`, and are case-sensitive.

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
- `(contains s1 s2)` Returns true iff `s1` contains `s2`
- `(nth i l)` Returns the `i`th value in list `l` or error if out of bounds 
- `(nthstr i s)` Returns the `i`th character in the string `s`
- `(len l)` Returns the length of list `l`
- `(lenstr s)` Returns the length of string `s`
- `(substr s i j=lenstr(s))` Returns the substring `s[i:j]`. If `j` is ommitted, it will default to the length of string `s`
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
- `(matches pattern str n=-1)` Returns the result of running Go's [FindAllStringSubmatch](https://pkg.go.dev/regexp#example-Regexp.FindAllStringSubmatch) function on `str` using the `pattern` argument as the regular expression. The `n` argument is optional and if ommitted will use a default value of `-1`.
- `(match pattern str)` Returns the result of running Go's [MatchString](https://pkg.go.dev/regexp#example-Regexp.MatchString) function on `str` using the `pattern` argument as the regular expression
- `(find pattern str)` Returns the result of running Go's [FindString](https://pkg.go.dev/regexp#example-Regexp.FindString) function on `str` using the `pattern` argument as the regular expression
- `(findall pattern str n=-1)` Returns the result of running Go's [FindAll](https://pkg.go.dev/regexp#example-Regexp.FindAll) function on `str` using the `pattern` argument as the regular expression. The `n` argument is optional and if ommitted will use a default value of `-1`.

> NOTE: All of the above functions are evaluated *eagerly*, meaning their arguments are evaluated before the function is called. The exceptions to this are the `if`, `append`, `appendleft`, `pop`, and `popleft` functions, which only evaluate their arguments as they are needed.

## Extension Functions

Conductorr extends the base CSL specification with functions specific to its own functionality for processing releases.

A `release` type is just a list defined like so:

```lisp
(define a
  (
    "Manos.The.Hands.of.Fate.1966.THEATRiCAL.1080p.BluRay.x264-SADPANDA" 
    "TorrentLeech" 
    "torrent" 
    "movie" 
    "BDRip" 
    "1080p" 
    "x264" 
    32 
    794 
    10737418240 
    288
  )
)
```

The properties are in order: Release title, indexer, download type (either torrent or nzb), content type (either movie or series), rip type, resolution, encoding, number of seeders, age in days, size in bytes, and runtime in minutes.

Possible values for rip type are `CAM`, `TELESYNC`, `SCR`, `HDTV`, `DVDRip`, `HDRip`,`WEBCap`, `WEBRip`, `WEBDL`, `BDRip`.

Possible values for resolution are `352p`, `360p`, `480i`, `480p`, `576i`, `576p`, `720p`, `1080p`, `2160p`

Possible values for encoding are `Xvid`, `x264`, `x265`, `VP9`

The following functions are defined on `release` types:

- `r-title` Retrieves release title
- `r-indexer` Retrieves release indexer
- `r-downloadtype` Retrieves release download type
- `r-contenttype` Retrieves release content type
- `r-riptype` Retrieves release rip type
- `r-resolution` Retrieves release resolution
- `r-encoding` Retrieves release encoding
- `r-seeders` Retrieves number of seeders for release
- `r-age` Retrieves age in days for release
- `r-size` Retrieves size of release in bytes
- `r-runtime` Retrieves the runtime of the release in minutes
- `r-riptype-order` Returns a number designating the order of how preferrable a riptype is (`CAM < TELESYNC < SCR < HDTV < DVDRip < HDRip < WEBCap < WEBRip < WEBDL < BDRip`)
- `r-resolution-order` Returns a number designating the order of how preferrable a resolution is (`352p < 360p < 480i < 480p < 576i < 576p < 720p < 1080p < 2160p`) 
- `r-encoding-order` Returns a number designating the order of how preferrable an encodign is (`Xvid < x264 < x265 < VP9`)

To call use one of these functions, do the following:

`(r-title a)` where `a` is a release as defined above

The above snippet returns `"Manos.The.Hands.of.Fate.1966.THEATRiCAL.1080p.BluRay.x264-SADPANDA"` when `a` is defined like in the earlier example

## Importing Scripts as Functions

In a CSL script, the value of the final expression is considered to be the "return value" of the script. Scripts can be imported as functions.

The syntax to import a script function is `(import "https://conductorr.github.io/demoscript.csl" demo-script)`.

Available protocols:

| Source Type | Syntax | Availability | Notes |
| ----------- | ------ | ------------ | ----- |
| filesystem  | `file:/home/user/scripts/test.csl` | cli | Will load a file from disk. It is only available in the CLI, so it is useful for local testing of CSL scripts. The exception to this is that remote modules in a Git repository can reference local files in the repository. |
| git         | `github.com/user/repo:test.csl@v2` | cli, conductorr, playground | Works with GitLab, GitHub, and Gitea (others possibly as well). The version tag is optional, and any commit hash or branch name or tag is valid. Note the colon separating the hostname and web path from the path within the repository to the script. |
| web         | `https://example.com/test.csl` | cli, conductorr, playground | The CSL package manager will perform a simple HTTP GET request at the given URL. No matter if https or http is specified, it will always try https first, and upon failure, it will only try http if insecure connections have been enabled in the environment. |
| profile     | `profile:(sorter\|filter):myprof` | cli, conductorr, playground | Allows you to access a sorter or profile of the named profile. For example, `profile:sorter:myprof` will import the sorter function in the `myprof` profile. |

You can then invoke the script at any time using `(demo-script)`. The script will execute, and the above expression will evaluate as the script's return value. Arguments to a script are accessible via an ordered list called `args` which is always accessible. Script functions are run in a separate scope and only inherit the arguments that are passed to it.

### Note on Importing Profile Scripts
As shown above, you can import sorter and filter scripts from a profiles within Conductorr. By default, an import path of the form `profile:sorter:myprof` will only work when a script is run within the same instance of Conductorr that has the profile called "myprof". If you are running a script outside of Conductorr, or if you want to import a script from a different instance of Conductorr, then you need to append credentials and connection details in the import path as query parameters.
Available query parameters are:
- `host` the host (without port) for the Conductorr instance.
- `port` the port for the Conductorr instance. (Usually 6416 for local instances and 443 if you are connecting over public Https)
- `auth_token` an auth token for the Conductorr instance.
- `username` admin username for the Conductorr instance.
- `password` admin password for the Conductorr instance.

The `host` parameter is required to connect to an external Conductorr instance. Either the `username` and `password` pair or `auth_token` are required. The port is not required, as a default value of 6416 will be used.

Note that you should only use this feature of importing external profile scripts when absolutely necessary. For example, you can use the [CSL CLI](/csl/cli.html) to run a profile filter script on a running Conductorr instance by doing `csl run profile:filter:myprof?host=localhost&port=6416&username=admin&password=verysecret somearg1 somearg2`. You should not use this feature in scripts that will be published anywhere, especially to public git repositories or webservers, as leaking these credentials can compromise your Conductorr instance.

## Important Notes and Caveats
- Your imports must come before any other code in your script.
- Attempting to pass anything other than a string **literal** as the import statement's second argument will result in an parsing failure.
- Depending on your execution environment (i.e. Conductorr), imported scripts may be cached. Conductorr caches scripts imported from the web and periodically refetches them to check for updates. If you know your dependent scripts have updated, you can force Conductorr to reload them, or if using the git-style imports, you can bump the version tag and the next time the script executes it will pull the updated version of the script.
- If a script fails to load, your execution environment may instead load the script from its cache. If your script is tagged with a version, it should only look for scripts tagged with that same version in the cache.
- For any http request made by your execution environment to resolve dependencies, it will first attempt to use https, and depending on the environment configuration, http may be used as a fallback. By default insecure requests are disabled in all execution environments, except for the playground.

<script>
// In no particular order ;)
const STANDS_FOR = new Set([
  "Custom Scripting Language",
  "Content Searching Language",
  "Conductorr-Specific Language",
  "Custom Sorting Language",
  "Characteristic Searching Language",
  "Custom Shitty Lisp",
])

const pickRandom = (current) => {
  const toSubtract = new Set([current])
  const remaining = new Set([...STANDS_FOR].filter(x => !toSubtract.has(x)))

  if (remaining.size < 1) {
    console.log('returning current', current)
    return current
  }
  const pick = Array.from(remaining)[Math.floor(Math.random() * remaining.size)]
  console.log('returning pick', pick)
  return pick
}

export default {
  data() {
    return {
      standsFor: pickRandom("")
    }
  },
  methods: {
    orDoesIt() {
      this.standsFor = pickRandom(this.standsFor)
    }
  }
}
</script>
