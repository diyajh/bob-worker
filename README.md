# bob-worker

The **Worker agent being improved** by [Bob](https://github.com/). The Worker writes a Redis-compatible
server in Go; it is graded byte-for-byte against Valkey.

Four improvable layers (Bob edits exactly one per change):

| Layer | File | Purpose |
|---|---|---|
| Skills | `skills.md` | rules the Worker follows |
| Tools | `tools/` | tool definitions the Worker can call |
| Routing | `routing.yaml` | which model handles which task type |
| Context | `context.yaml` | which files/docs to load per task type |

Other directories:

- `server/` — the Go server the Worker writes (skeleton: answers `PING` only).
- `tasks/` — command groups the Worker must implement (strings, keys, lists, hashes).
