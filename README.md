# Grua :: Maintenance Page

A simple maintenance page: every path, any method, answers `503 Service
Unavailable` with the same static page (`index.html`, embedded in the
binary), `Retry-After: 60` and `Cache-Control: no-store`. The page
reloads itself every 60 seconds.

    PORT=8080 go run .     # the port defaults to 8080

## Image

    bin/build 0.1.0        # ghcr.io/grua-io/maintenance-in-progress, both arches
