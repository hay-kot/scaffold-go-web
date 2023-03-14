# Scaffold for Go Web HTTP Applications

- Configured with [ardanlabs/conf](https://github.com/ardanlabs/conf)
- Logging via [rs/zerolog](https://github.com/rs/zerolog)
- HTTP Routing with [go-chi/chi](https://github.com/go-chi/chi)
- Error Handling with [hay-kot/safeserve](https://github.com/hay-kot/safeserve)
- Error Tracing with [hay-kot/safeserve](https://github.com/hay-kot/safeserve)
- Server Utilities [hay-kot/safeserve](https://github.com/hay-kot/safeserve)

## Todo's

- [ ] Optional Database Support
  - [ ] Ent with Atlas Migrations
  - [ ] User + Tenant Model
  - [ ] Repository Pattern w/ DTOs
- [ ] Optional Authentication Middleware
  - [ ] JWT + Refresh Tokens (Maybe DB Tokens)
- [ ] Docker Support
  - [ ] Publish to Github Packages
  - [ ] Multi-arch Support (should be possible with goreleaser)
- [ ] Struct Validation
- [ ] Open API Docs