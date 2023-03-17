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
  - [x] Ent Schema Folder with
  - [x] User + Tenant Model
  - [x] Setup Client Connection on Start
  - [ ] ( ) Sqlite (CGo) or (x) Postgres Support
- [ ] Optional Authentication Middleware
  - [ ] JWT + Refresh Tokens (Maybe DB Tokens)
- [x] Docker Support (goreleaser)
- [x] Struct Validation
- [ ] Open API Docs