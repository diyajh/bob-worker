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
## LPUSH element order
- LPUSH pushes its values onto the head ONE AT A TIME in argument order, so `LPUSH key a b c` yields the list [c, b, a]: the LAST argument is the first element. Do not prepend the argument list as a block. `LRANGE key 0 -1` after `LPUSH key a b c` must return c, b, a. (RPUSH appends in argument order and is already correct.)
