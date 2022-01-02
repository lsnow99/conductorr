# CLI

The CSL CLI is a command line utility for developing and testing CSL scripts outside of Conductorr or a playground environment. The CLI can load and execute modules from all sources.

## Usage

The CLI has two available commands: `get` and `run`.
For both commands, a source script must be provided as the first argument. This can be any valid [import path](/csl/examples.html#importing-modules), not just local file paths. For example, you can run `csl run github.com/lsnow99/myscripts:main.csl` and it will fetch and run the script.

> NOTE: When importing or executing profile modules that are hosted on an instance of Conductorr, you must provide credentials to allow the CLI to access those modules. For more details, see {TODO ADD LINK}

### Get

`csl get {your_script}`

The `get` command fetches all dependencies required to run the script and saves them to the default global CSL cache directory. When a CSL script that imports any of the fetched modules is executed in the future, and if caching is enabled (default behavior of `run`), then those scripts will be served directly from the cache.

Flags:
- `-insecure` - Allow insecure (http) requests. Https is always attempted first by default, even if Http import paths are specified.

### Run

`csl run {your_script} {your_first_argument} {your_second_argument} ...`

The `run` command fetches all dependencies for a given script and invokes it with the given arguments. By default, dependencies will be served from the cache, and will fall back to fetching in the case of a cache miss.

Flags:
- `-insecure` - Allow insecure (http) requests. Https is always attempted first by default, even if Http import paths are specified.
- `-nocache` - Disable caching of dependencies for this invokation.
