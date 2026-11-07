# server

Go Redis-compatible server. The skeleton replies to `PING` only; the Worker implements the rest.

## RESP expectations

- Requests arrive as RESP arrays of bulk strings: `*N\r\n$len\r\narg\r\n...`.
- Replies are compared byte for byte with Valkey:
  - simple string `+OK\r\n`, integer `:1\r\n`, bulk `$3\r\nfoo\r\n`, null bulk `$-1\r\n`,
    array `*2\r\n...`, empty array `*0\r\n`, error `-ERR ...\r\n`.
- Error text must match Valkey exactly, e.g.
  `-WRONGTYPE Operation against a key holding the wrong kind of value\r\n`.
- Start with `./server --port <port>`.
