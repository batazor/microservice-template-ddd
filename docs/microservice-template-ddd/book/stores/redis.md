# book Redis

*Generated from the portolan catalog · commit `498f577` · at 2026-09-09T01:20:07Z. Do not edit by hand.*

- **Id:** `microservice-template-ddd.book.redis`
- **Kind:** redis
- **Owner:** [microservice-template-ddd.book](../README.md)
- **Source:** `internal/db/redis/redis.go:27`

## Redis key patterns

| Pattern | Operations | Value | TTL | Aggregate | Source |
| --- | --- | --- | --- | --- | --- |
| `{id}` | read | `domain.Book` | — | [microservice-template-ddd.book.book](../aggregates/book.md) | `internal/book/infrastructure/store/redis/redis.go:28` |
| `{in.Title}` | write | `domain.Book` | `none` | [microservice-template-ddd.book.book](../aggregates/book.md) | `internal/book/infrastructure/store/redis/redis.go:59` |

## Redis accesses

| Pattern | Operation | Method | Value | TTL | Source |
| --- | --- | --- | --- | --- | --- |
| `{id}` | read | `Store.Get` | `domain.Book` | — | `internal/book/infrastructure/store/redis/redis.go:28` |
| `{in.Title}` | write | `Store.Add` | `domain.Book` | `none` | `internal/book/infrastructure/store/redis/redis.go:59` |
| `{in.Title}` | write | `Store.Update` | `domain.Book` | `none` | `internal/book/infrastructure/store/redis/redis.go:75` |
