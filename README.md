# ✨ go-neuvector

NeuVector Controller API wrapper with full abstraction.

## 📖 How to build and run ?

1. Install the dependencies
    - `go`
    - `make` (for tests)

## ℹ️ Usage example

```go
import (
    "fmt"

    "github.com/theobori/go-neuvector/client"
    "github.com/theobori/go-neuvector/neuvector/scan"
)

func main() {
    client, err := client.NewDefaultClient()

    if err != nil {
        return
    }

    registries, err := scan.GetRegistries(client)

    if err != nil {
        return
    }

    fmt.Println(*registries)
}
```

## 🧪 Tests

We performed tests with a containerized instance of NeuVector via the Docker platform.
To run NeuVector and the Go tests, use:

```bash
make test
```

## ✅ Todo

Name           | State
-------------  | :-------------:
Auto refresh token | ❌
Service configuration (patch) | ✅
Admission (create/delete) + fed | ❌
Policy (patch) + fed | ❌
