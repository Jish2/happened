
Built using [Huma REST API Framework](https://huma.rocks/)
We use the automatically created OpenAPI to generate client Typescript SDK.

# Local Development
- `make watch` to start a local dev server.
- `make test` to run full test suite.
- `make image` to build the Docker image
- `make db` to spin up a local database instance with docker compose
- `make migrate` to run database migrations against the schema.sql script using [Atlas](https://atlasgo.io/docs)
