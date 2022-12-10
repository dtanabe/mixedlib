mixedlib
========

`mixedlib` is a library that doesn't do much of anything other than demonstrate
how a library written in mulitple languages might look.

Examples:
---------

```py
from mixedlib import add
print(add(1, 1))
```

```go
import github.com/dtanabe/mixedlib

func main() {
	print(mixedlib.add(1, 1))
}
```

About
-----

The standard packaging files (`setup.py`, `go.mod`, `package.json`, `pom.xml`,
etc.) are all located at the root of the file tree so that tools idiosyncractic
to each language _also_ work at the root, and so that most IDEs work sensibly
with all languages when importing the root. It also ensures that top-level
directories, such as `test` are accessible from all projects without too many
issues; many tools take issue with trying to reach outside of a root directory.

```shell
> git clone github.com/dtanabe/mixedlib
> cd mixedlib

# run Python tests
>

# run Go tests
> go test

# run JavaScript tests
> npm test

# run Java tests
> mvn test
```


Python
------

Running t
