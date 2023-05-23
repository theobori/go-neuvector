# ✨ go-neuvector

NeuVector Controller API wrapper with abstraction for useful features.

## 📖 How to build and run ?

1. Install the dependencies
    - `go`
    - `make`, `docker-compose` (for tests)

## 🤝 Contribute

If you want to help the project, you can follow the guidelines in [CONTRIBUTING.md](./CONTRIBUTING.md).

## ℹ️ Usage example

```go
import (
    "github.com/theobori/go-neuvector/client"
    "github.com/theobori/go-neuvector/neuvector/scan"
)

func main() {
    client, err := client.NewDefaultClient()

    if err != nil {
        return
    }

    registries, err := scan.GetRegistries(client)
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
Bypass expired token --> single session refresh | ✅
Service configuration (patch) | ✅
Admission (create/delete) + fed | ✅
Policy (patch) + fed | ✅
Federation promote as master cluster | ✅
