# POST /book/rent/{bookId}

*Generated from the portolan catalog · commit `498f577` · at 2026-09-09T01:20:07Z. Do not edit by hand.*

- **Id:** `flow.api-http-post-book-rent-bookid`
- **Owner:** [microservice-template-ddd](../microservice-template-ddd/README.md)
- **Trigger:** `http` · POST /book/rent/{bookId}
- **Root confidence:** high
- **Source:** `internal/api/http/routes_book.go`

Source-derived http execution path for `POST /book/rent/{bookId}`. Source-backed cross-protocol continuations are included.

## Participants

| Participant | Kind | Context |
| --- | --- | --- |
| `microservice-template-ddd.api` | service | [microservice-template-ddd](../microservice-template-ddd/README.md) |
| `client` | actor | — |
| `microservice-template-ddd.book` | service | [microservice-template-ddd](../microservice-template-ddd/README.md) |
| `microservice-template-ddd.user` | service | [microservice-template-ddd](../microservice-template-ddd/README.md) |
| `microservice-template-ddd.billing` | service | [microservice-template-ddd](../microservice-template-ddd/README.md) |
| `book-redis` | store | [microservice-template-ddd](../microservice-template-ddd/README.md) |

## Sequence

```mermaid
sequenceDiagram
    autonumber
    participant p0 as microservice-template-ddd.api
    actor p1 as client
    participant p2 as microservice-template-ddd.book
    participant p3 as microservice-template-ddd.user
    participant p4 as microservice-template-ddd.billing
    participant p5 as book-redis
    p1->>p0: POST /book/rent/{bookId}
    p0->>p2: Rent
    p2->>p2: Rent
    p2->>p3: Get
    p3->>p3: Get
    p3-->>p2: GetResponse
    p2->>p4: Get
    p4->>p4: Get
    p4-->>p2: GetResponse
    p2->>p5: READ `{id}`
    alt err != nil
        p2->>p5: WRITE `{in.Title}`
        p2-->>p0: RentResponse
        Note over p0: flow ends here
    else otherwise
    end
    p2->>p5: WRITE `{in.Title}`
    p2-->>p0: RentResponse
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
    p0-->>p1: 200 · RentResponse
```

## Steps

<a id="step-s1"></a>
1. **client** → **microservice-template-ddd.api** — POST /book/rent/{bookId}
   status: declared · `internal/api/http/routes_book.go:23`
<a id="step-s2"></a>
2. **microservice-template-ddd.api** → **microservice-template-ddd.book** — Rent
   `book_rpc.BookRPC/Rent` · status: declared · `internal/api/http/routes_book.go:65`
<a id="step-continuation-book-grpc-book-rpc-bookrpc-rent-s2-s1"></a>
3. **microservice-template-ddd.book** ↺ **microservice-template-ddd.book** — Rent
   status: declared · `internal/book/infrastructure/rpc/rent.go:8`
<a id="step-continuation-book-grpc-book-rpc-bookrpc-rent-s2-s2"></a>
4. **microservice-template-ddd.book** → **microservice-template-ddd.user** — Get
   `user_rpc.UserRPC/Get` · status: declared · `internal/book/application/rent.go:53`
<a id="step-continuation-book-grpc-book-rpc-bookrpc-rent-s2-continuation-user-grpc-user-rpc-userrpc-get-s2-s1"></a>
5. **microservice-template-ddd.user** ↺ **microservice-template-ddd.user** — Get
   status: declared · `internal/user/infrastructure/rpc/grpc.go:46`
<a id="step-continuation-book-grpc-book-rpc-bookrpc-rent-s2-response-s2"></a>
6. **microservice-template-ddd.user** → **microservice-template-ddd.book** — GetResponse
   status: declared · Synthesized from the proven synchronous unary gRPC return.
<a id="step-continuation-book-grpc-book-rpc-bookrpc-rent-s2-s3"></a>
7. **microservice-template-ddd.book** → **microservice-template-ddd.billing** — Get
   `billing_rpc.BillingRPC/Get` · status: declared · `internal/book/application/rent.go:59`
<a id="step-continuation-book-grpc-book-rpc-bookrpc-rent-s2-continuation-billing-grpc-billing-rpc-billingrpc-get-s3-s1"></a>
8. **microservice-template-ddd.billing** ↺ **microservice-template-ddd.billing** — Get
   status: declared · `internal/billing/infrastructure/rpc/grpc.go:46`
<a id="step-continuation-book-grpc-book-rpc-bookrpc-rent-s2-response-s3"></a>
9. **microservice-template-ddd.billing** → **microservice-template-ddd.book** — GetResponse
   status: declared · Synthesized from the proven synchronous unary gRPC return.
<a id="step-continuation-book-grpc-book-rpc-bookrpc-rent-s2-s4"></a>
10. **microservice-template-ddd.book** → **book-redis** — READ `{id}`
   status: declared · `internal/book/application/rent.go:36`

> **One of**
>
> *err != nil — *ends the flow**
>
> <a id="step-continuation-book-grpc-book-rpc-bookrpc-rent-s2-s5"></a>
> 11. **microservice-template-ddd.book** → **book-redis** — WRITE `{in.Title}`
>    status: declared · `internal/book/application/rent.go:39`
> <a id="step-response-s2"></a>
> 12. **microservice-template-ddd.book** → **microservice-template-ddd.api** — RentResponse
>    status: declared · Synthesized from the proven synchronous unary gRPC return.
>
> *otherwise*

<a id="step-continuation-book-grpc-book-rpc-bookrpc-rent-s2-s7"></a>
13. **microservice-template-ddd.book** → **book-redis** — WRITE `{in.Title}`
   status: declared · `internal/book/application/rent.go:76`
<a id="step-response-s2-exit-2"></a>
14. **microservice-template-ddd.book** → **microservice-template-ddd.api** — RentResponse
   status: declared · Synthesized from the proven synchronous unary gRPC return.

> **One of**
>
> *err != nil — *ends the flow**
>
> <a id="step-s3"></a>
> 15. **microservice-template-ddd.api** → **client** — 200 · Error
>    status: declared · `internal/api/http/routes_book.go:68` · HTTP · 200 · `application/json` · json · error · fields: error: string · warning: Error response has no explicit non-2xx status; net/http will send 200.
>
> *otherwise*


> **One of**
>
> *err != nil*
>
> <a id="step-s5"></a>
> 16. **microservice-template-ddd.api** → **client** — 200 · Error
>    status: declared · `internal/api/http/routes_book.go:76` · HTTP · 200 · `application/json` · json · error · fields: error: string · warning: Error response has no explicit non-2xx status; net/http will send 200. Execution continues after writing the error body and may append another response.
>
> *otherwise*

<a id="step-s7"></a>
17. **microservice-template-ddd.api** → **client** — 200 · RentResponse
   status: declared · `internal/api/http/routes_book.go:79` · HTTP · 200 · `application/json` · protojson · success · body derived from `book_rpc.BookRPC/Rent`
