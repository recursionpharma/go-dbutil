# go-dbutil
Golang utilities for interacting with databases

## Install

    go get github.com/recursionpharma/go-dbutil

## Usage

See also `examples_test.go`.

### Given a DB URL, connect to the DB

This is useful since `sql.Open()` requires you to pass the driver in addition to the URL,
even though the driver is embedded in the URL.

    import (
        "os"
        "fmt"
        "github.com/recursionpharma/go-dbutil"
    )

    func main() {
        db, err := dbutil.Connect("postgres://USER:PASSWORD@HOST:PORT/DBNAME")
        if err != nil {
            fmt.Println(err)
            os.Exit(1)
        }
        fmt.Println(db != nil)
        // Output: true
    }

### Extract the driver from a DB URL

    import (
        "os"
        "fmt"
        "github.com/recursionpharma/go-dbutil"
    )

    func main() {
        driver, err := GetDriver("postgres://USER:PASSWORD@HOST:PORT/DBNAME")
        if err != nil {
            fmt.Println(err)
            os.Exit(1)
        }
        fmt.Println(driver)
        // Output: postgres
    }

## Directory Structure

```
code/go-dbutil/
|-- dbutil.go
|   Code
|-- dbutil_test.go
|   Tests
|-- examples_test.go
|   Examples
|-- .gitignore
|   Files git will ignore
|-- LICENSE
|   MIT License
|-- README.md
|   This file
`-- .travis.yml
    Travis config
```
The above file tree was generated with `tree -a -L 1 --charset ascii`.
