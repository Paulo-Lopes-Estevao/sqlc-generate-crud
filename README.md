# sqlc generate crud

is a library that generates CRUD sql operations for a struct using the sqlc structure.

## Case of use

was created in order to decrease the time spent writing queries using the sqlc: A SQL Compiler.
[sqlc](https://sqlc.dev/)


## Installation

```bash
go get github.com/Paulo-Lopes-Estevao/sqlc-generate-crud
```

## Usage

### Required create a file sqlc.yaml

```yaml
version: "2"
sql:
  - schema: "postgresql/schema.sql"
    queries: "postgresql"
    engine: "postgresql"
    gen:
      go:
        package: "authors"
        out: "postgresql"
```

> generated sql files will be saved in the queries path folder in the sqlc.yaml file

function Generate has 2 parameters:

- **data** : struct that will be used to generate the CRUD operations
- **options** : struct that will be used to configure the generation _tag_ and _pathTarget_

```golang
// GenerateConfig is a struct that will be used to configure the generation
type GenerateConfig struct {
    Tag        string
    PathTarget string
}
```

### Example

```golang
package main

import (
    "github.com/Paulo-Lopes-Estevao/sqlc-generate-crud"
)

type User struct {
    ID int64 `json:"id"`
    Name string `json:"name"`
    Email string `json:"email"`
}

func main() {
	err := sqlcgeneratecrud.Generate(User{}, &sqlcgeneratecrud.GenerateConfig{})
	if err != nil {
        panic(err)
    }
}
```



## Contributing

Before opening an issue or pull request, please check the project's contribution documents.

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details about our code of conduct, and the process for submitting pull requests.

## Support Donate

If you find this project useful, you can buy author a glass of juice 🧃

[![ko-fi](https://ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/E1E2L169R)

also a coffee ☕️

<a href="https://www.buymeacoffee.com/pl1745240p" target="_blank"><img src="https://cdn.buymeacoffee.com/buttons/v2/default-yellow.png" alt="Buy Me A Coffee" style="height: 60px !important;width: 217px !important;" ></a>

will be very grateful to you for your support 😊.
