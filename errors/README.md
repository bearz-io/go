# 💥 errors 💥

Error struct and functions.

## Overview 📖

The errors module provides a standard error that supports code, cause,
and stacktrace for errors and other common error types. The module is
inspired by the [pkg/errors](https://github.com/pkg/errors) module
and is intended to support bearz.io go projects and make it easy to debug errors.



## Usage 🚀

```go
package main

import "github.com/bearz-io/go-errors"

var (
    // custom error with a default message of 'custom error message' and code of 'CustomError'
    ErrCustom = errors.NewStdError("custom error message", "CustomError")
)

func main() {
	err := errors.New("test")
	fmt.Printf("%+v", err)

    stdErr := errors.NewStdError("test", "TestError")

    ErrCustom.Is(stdErr) // false. The error codes are different
}
```

## License

This project is licensed under the MIT License - see
the [LICENSE](./LICENSE.md) file for details.

The stackframe leverages code from https://github.com/pkg/errors
which is licensed under the 
[BSD-2-Clause License](https://github.com/pkg/errors/blob/master/LICENSE)
