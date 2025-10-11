# go-crypto [![Paypal donate](https://www.paypalobjects.com/en_US/i/btn/btn_donate_LG.gif)](https://www.paypal.com/donate/?business=HZF49NM9D35SJ&no_recurring=0&currency_code=CAD)

Command line tool for encryption function.

### Table Of Content
<!-- TOC -->

- [Purpose](#purpose)
- [Usage](#usage)
- [Package](#package)
- [Command Line](#command-line)
- [Example](#example)
- [Repository](#repository)
- [Contributors](#contributors)
- [Change Log](#change-log)
- [License](#license)

<!-- /TOC -->
<!--more-->
> This is not crypto currency program.

### Purpose

Provide crypto function for command line scripting, eg. [mygit](https://github.com/J-Siu/mygit). Currently only box seal anonymous is implemented, which is used for calculating github repository secret.

Open an issue here or in [go-crypto](https://github.com/J-Siu/go-crypto) if you need additional crypto functions.

### Usage

### Package

```sh
go get github.com/J-Siu/go-crypto
```

```go
import "github.com/J-Siu/go-crypto/crypto"
```

### Command Line

```sh
$ go-crypto box sealanonymous -h
x/crypto box seal anonymous. Output is base64 encoded

Usage:
  go-crypto box sealanonymous [flags]

Aliases:
  sealanonymous, s

Flags:
  -h, --help         help for sealanonymous
  -k, --key string   (required) base64 encoded public key
  -m, --msg string   (required) plain text message
```

### Example

```sh
$ go run main.go box sealanonymous -k 'z492di80U5FuJfY8VH2M26Cnzg4UfRRxlqTXMHSWfyY=' -m "This is a test"
CYlrGgMkPCnkucOeFlFZ68IfyW78SJyHF6o5CkWN7HtvreTVbw6umsoLxX7+5buvDzMAlLzQDJaiqjvcOjY=
```

### Repository

- [go-crypto](https://github.com/J-Siu/go-crypto)

### Contributors

- [John Sing Dao Siu](https://github.com/J-Siu)

### Change Log

- 1.0.0
  - Implement box seal anonymous
- v1.0.1
  - Fix `goreleaser`
- v1.0.2
  - Update to Go 1.20 and dependency
- v1.0.3
  - Fix Github workflows
- v1.0.4
  - Include `crypto` mod
- v1.0.5
  - Update go-helper/v2
- v1.1.0
  - Update go-helper/v2
  - BoxSealAnonymous()
    - Remove key length(32) check. No longer apply.
- v1.1.1
  - Add version const
  - Add cli shorthand
  - Fix debug output
  - Update go-helper/v2

### License

The MIT License (MIT)

Copyright © 2025 John, Sing Dao, Siu <john.sd.siu@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
