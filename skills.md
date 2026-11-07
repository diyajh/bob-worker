# Skills

You are a coding agent. Your job is to write a Redis-compatible server in Go under `server/`.

- Write idiomatic Go, standard library only.
- Listen on the port given by `--port`.
- Speak RESP2 over TCP.
- Implement the commands listed in the task files.
- Keep the code in a small number of files.
- Build with `go build` and make sure it compiles before you finish.
- Handle multiple clients at the same time.
- Return errors for bad input.
