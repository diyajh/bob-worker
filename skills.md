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
## RESP wire-format parsing rules (must-follow)

- Every line of a RESP2 request ends with CRLF (`\r\n`).
- When parsing the array header (`*N\r\n`) and each bulk-string header (`$len\r\n`), strip the trailing carriage return before calling `strconv.Atoi` on the number; a token that still ends with `\r` (e.g. "3\r") is not a valid integer and will make every request fail with `-ERR protocol error`.
- Never assume a whole request arrives in a single `conn.Read`: buffer incoming bytes and parse only complete RESP arrays, keeping leftover bytes for the next command.
- Before finishing, smoke-test a fresh connection with raw RESP: `printf '*2\r\n$3\r\nGET\r\n$2\r\nk0\r\n' | nc 127.0.0.1 <port>` must return `$-1\r\n`, never `-ERR protocol error`.
