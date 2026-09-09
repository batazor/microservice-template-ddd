# GET /book/

*Generated from the portolan catalog · commit `498f577` · at 2026-09-09T01:20:07Z. Do not edit by hand.*

- **Id:** `flow.api-http-get-book`
- **Owner:** [microservice-template-ddd](../microservice-template-ddd/README.md)
- **Trigger:** `http` · GET /book/
- **Root confidence:** high
- **Source:** `internal/api/http/routes_book.go`

Source-derived http execution path for `GET /book/`. Source-backed cross-protocol continuations are included.

## Participants

| Participant | Kind | Context |
| --- | --- | --- |
| `microservice-template-ddd.api` | service | [microservice-template-ddd](../microservice-template-ddd/README.md) |
| `client` | actor | — |
| `microservice-template-ddd.book` | service | [microservice-template-ddd](../microservice-template-ddd/README.md) |
| `book-redis` | store | [microservice-template-ddd](../microservice-template-ddd/README.md) |

## Sequence

```mermaid
sequenceDiagram
    autonumber
    participant p0 as microservice-template-ddd.api
    actor p1 as client
    participant p2 as microservice-template-ddd.book
    participant p3 as book-redis
    p1->>p0: GET /book/
    p0->>p2: Get
    p2->>p2: Get
    p2->>p3: READ `{id}`
    alt err != nil
        p2->>p3: WRITE `{in.Title}`
        p2-->>p0: GetResponse
        Note over p0: flow ends here
    else otherwise
    end
    p2-->>p0: GetResponse
    alt err != nil
        rect rgba(183, 100, 107, 0.12)
            p0-->>p1: 200 · Error
        end
        Note over p1: flow ends here
    else otherwise
    end
    alt err != nil
        rect rgba(183, 100, 107, 0.12)
            p0-->>p1: 200 · Error
        end
    else otherwise
    end
    p0-->>p1: 200 · GetResponse
```

## Steps

<a id="step-s1"></a>
1. **client** → **microservice-template-ddd.api** — GET /book/
   status: declared · `internal/api/http/routes_book.go:18`
<a id="step-s2"></a>
2. **microservice-template-ddd.api** → **microservice-template-ddd.book** — Get
   `book_rpc.BookRPC/Get` · status: declared · `internal/api/http/routes_book.go:36`
<a id="step-continuation-book-grpc-book-rpc-bookrpc-get-s2-s1"></a>
3. **microservice-template-ddd.book** ↺ **microservice-template-ddd.book** — Get
   status: declared · `internal/book/infrastructure/rpc/book.go:8`
<a id="step-continuation-book-grpc-book-rpc-bookrpc-get-s2-s2"></a>
4. **microservice-template-ddd.book** → **book-redis** — READ `{id}`
   status: declared · `internal/book/application/rent.go:36`

> **One of**
>
> *err != nil — *ends the flow**
>
> <a id="step-continuation-book-grpc-book-rpc-bookrpc-get-s2-s3"></a>
> 5. **microservice-template-ddd.book** → **book-redis** — WRITE `{in.Title}`
>    status: declared · `internal/book/application/rent.go:39`
> <a id="step-response-s2"></a>
> 6. **microservice-template-ddd.book** → **microservice-template-ddd.api** — GetResponse
>    status: declared · Synthesized from the proven synchronous unary gRPC return.
>
> *otherwise*

<a id="step-response-s2-exit-2"></a>
7. **microservice-template-ddd.book** → **microservice-template-ddd.api** — GetResponse
   status: declared · Synthesized from the proven synchronous unary gRPC return.

> **One of**
>
> *err != nil — *ends the flow**
>
> <a id="step-s3"></a>
> 8. **microservice-template-ddd.api** → **client** — 200 · Error
>    status: declared · `internal/api/http/routes_book.go:39` · HTTP · 200 · `application/json` · json · error · fields: error: string · warning: Error response has no explicit non-2xx status; net/http will send 200.
>
> *otherwise*


> **One of**
>
> *err != nil*
>
> <a id="step-s5"></a>
> 9. **microservice-template-ddd.api** → **client** — 200 · Error
>    status: declared · `internal/api/http/routes_book.go:47` · HTTP · 200 · `application/json` · json · error · fields: error: string · warning: Error response has no explicit non-2xx status; net/http will send 200. Execution continues after writing the error body and may append another response.
>
> *otherwise*

<a id="step-s7"></a>
10. **microservice-template-ddd.api** → **client** — 200 · GetResponse
   status: declared · `internal/api/http/routes_book.go:50` · HTTP · 200 · `application/json` · protojson · success · body derived from `book_rpc.BookRPC/Get`
