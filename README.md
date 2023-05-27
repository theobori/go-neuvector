# ✨ go-neuvector

NeuVector Controller API wrapper with abstraction for useful features.

The abstraction

## 📖 How to build and run ?

1. Install the dependencies
    - `go`
    - `make` (for tests)
    - `docker-compose` (for tests)

## 🤝 Contribute

If you want to help the project, you can follow the guidelines in [CONTRIBUTING.md](./CONTRIBUTING.md).

There are some types written by neuvector [here](https://github.com/neuvector/neuvector/blob/main/controller/api), which we've redone some of them for greater clarity and ease of reading. But we're not opposed to using them for future implementations.

## 🧪 Tests

We performed tests with a containerized instance of NeuVector via the Docker platform.
To run NeuVector and the Go tests, use:

```bash
make neuvector
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
