# unpkg alternative service in GoLang

## Directory Structure

unpkg-alternative/
├── cmd/
│   └── server/
│       └── main.go         # Entrypoint
├── http/
│   ├── router.go           # HTTP routing logic
│   └── handler.go          # Request handler functions
├── cache/
│   └── cache.go            # Cache manager
├── npm/
│   ├── resolver.go         # Resolves package metadata
│   └── fetcher.go          # Downloads & extracts tarballs
├── utils/
│   └── file.go             # Helpers (MIME types, path safety)
└── go.mod
