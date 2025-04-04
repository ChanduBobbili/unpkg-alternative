# unpkg alternative service in GoLang

## Directory Structure

```
unpkg-alternative/
├── cmd/
│   └── server/
│       └── main.go         # Entrypoint with logger setup
├── http/
│   ├── router.go           # HTTP routing logic with logging middleware
│   └── handler/
│       └── handler.go      # CDN handler
├── cache/
│   └── cache.go            # Cache manager
├── npm/
│   ├── resolver.go         # Resolves package metadata
│   └── fetcher.go          # Downloads & extracts tarballs
├── utils/
│   └── file.go             # Helpers (MIME types, path safety)
├── logs/
│   └── zap.go              # Zap logger config
└── go.mod
```