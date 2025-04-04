# High-Level Design: Unpkg-like CDN Service in Go

## Overview

A content delivery service similar to [unpkg.com](https://unpkg.com) that serves JavaScript, CSS, and other static files directly from npm packages. The service fetches package tarballs from the npm registry, extracts them, caches them locally, and serves files on demand.

## 🧭 High-Level Architecture

```
+---------------------+          +-------------------------+
|     HTTP Client     | <----->  |     CDN Go Server       |
+---------------------+          +-------------------------+
                                        |
                                        v
                           +--------------------------+
                           |     Request Router       |
                           +--------------------------+
                                        |
      +---------------------------------+--------------------------------+
      |                                                                  |
      v                                                                  v
+------------------+                                         +--------------------+
|   Cache Manager  |                                         |  Package Resolver  |
| (Disk / Cloud)   |                                         | (NPM Registry API) |
+------------------+                                         +--------------------+
      |                                                                  |
      v                                                                  v
+---------------------+                              +---------------------------+
| Serve Cached File   |                              |  Download & Extract .tgz  |
+---------------------+                              +---------------------------+
                                                              |
                                                              v
                                                  +-------------------------+
                                                  |  Local Cache Directory  |
                                                  +-------------------------+
```

---

## 🔩 Components

### 1. HTTP Server & Request Router

- **Function**: Listens for incoming requests and dispatches them to appropriate handlers.
- **Technology**: Go's `net/http` package
- **Responsibility**:
  - Parse incoming URLs (e.g., `/react@18.2.0/umd/react.production.min.js`)
  - Delegate to cache or resolver

### 2. URL Parser

- **Function**: Extract package name, version, and file path from request URL.
- **Responsibilities**:
  - Support both scoped and unscoped packages
  - Default to `latest` if version is not specified

### 3. Cache Manager

- **Function**: Store and retrieve extracted package files from local or cloud storage.
- **Responsibilities**:
  - Check if requested file exists
  - Serve from cache if available
  - Evict old cache entries (optional future enhancement)

### 4. Package Resolver

- **Function**: Resolve tarball URL from the npm registry.
- **Responsibilities**:
  - Use npm registry API to get metadata
  - Extract tarball URL
  - Handle errors for non-existent packages or versions

### 5. Tarball Downloader and Extractor

- **Function**: Download and extract `.tgz` tarball from npm registry.
- **Responsibilities**:
  - Extract only `package/` directory contents
  - Store files in appropriate cache directory
  - Prevent directory traversal vulnerabilities

### 6. Static File Server

- **Function**: Serve requested files over HTTP.
- **Responsibilities**:
  - Set correct MIME type
  - Add CORS headers
  - Add cache control headers

---

## 📁 File System Layout (After Caching)

```
cdn-go/
├── main.go
├── cache/
│   ├── react@18.2.0/
│   │   └── package/
│   │       └── umd/react.production.min.js
├── go.mod
```

---

## Supported URL Formats

| URL Example                                              | Description                   |
|----------------------------------------------------------|-------------------------------|
| `/react@18.2.0/umd/react.production.min.js`              | Serve React UMD build         |
| `/lodash@4.17.21/lodash.js`                              | Serve lodash build            |
| `/@types/react@18.0.27/index.d.ts`                       | Scoped package                |
| `/react/umd/react.production.min.js`                     | Defaults to latest version    |

---

## Security Considerations

- Sanitize file paths to prevent directory traversal `(../../)`.
- Limit download sizes / rate limiting to prevent abuse.
- Only serve files under `package/` folder inside .tgz.

---

## ⚡ Possible Enhancements (Post-MVP)

| Feature                    | Description                                       |
|----------------------------|---------------------------------------------------|
| CDN layer (Cloudflare/S3)  | Offload bandwidth and improve latency            |
| Directory listings         | Like `https://unpkg.com/react@18.2.0/`           |
| Gzip/Brotli compression    | Serve compressed files for JS/CSS                |
| ETag & Last-Modified       | Enable browser cache validation                  |
| Scoped package handling    | Support for scoped packages like `@scope/pkg`    |
| Usage analytics            | Logs, metrics, rate limits                       |
| Deploy with Docker         | Containerize and deploy with nginx or Caddy      |

---

## Future Enhancements

- Directory listings
- Gzip/Brotli compression
- ETag and Last-Modified support
- Cloudflare or CDN integration
- Metrics and usage analytics
- Docker deployment
