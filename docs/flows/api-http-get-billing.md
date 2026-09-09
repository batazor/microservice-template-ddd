# GET /billing/

*Generated from the portolan catalog · commit `baebd80` · at 2026-09-09T21:28:15+07:00. Do not edit by hand.*

- **Id:** `flow.api-http-get-billing`
- **Owner:** [microservice-template-ddd](../microservice-template-ddd/README.md)
- **Trigger:** `http` · GET /billing/
- **Root confidence:** high
- **Source:** `internal/api/http/routes_billing.go`

Source-derived http execution path for `GET /billing/`. Source-backed cross-protocol continuations are included.

## Participants

| Participant | Kind | Context |
| --- | --- | --- |
| `microservice-template-ddd.api` | service | [microservice-template-ddd](../microservice-template-ddd/README.md) |
| `client` | actor | — |
| `microservice-template-ddd.billing` | service | [microservice-template-ddd](../microservice-template-ddd/README.md) |

## Sequence

```mermaid
sequenceDiagram
    autonumber
    participant p0 as microservice-template-ddd.api
    actor p1 as client
    participant p2 as microservice-template-ddd.billing
    p1->>p0: GET /billing/
    p0->>p2: Get
    p2->>p2: Get
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
1. **client** → **microservice-template-ddd.api** — GET /billing/
   status: declared · `internal/api/http/routes_billing.go:18`
<a id="step-s2"></a>
2. **microservice-template-ddd.api** → **microservice-template-ddd.billing** — Get
   `billing_rpc.BillingRPC/Get` · status: declared · `internal/api/http/routes_billing.go:33`
<a id="step-continuation-billing-grpc-billing-rpc-billingrpc-get-s2-s1"></a>
3. **microservice-template-ddd.billing** ↺ **microservice-template-ddd.billing** — Get
   status: declared · `internal/billing/infrastructure/rpc/grpc.go:46`
<a id="step-response-s2"></a>
4. **microservice-template-ddd.billing** → **microservice-template-ddd.api** — GetResponse
   status: declared · Synthesized from the proven synchronous unary gRPC return.

> **One of**
>
> *err != nil — *ends the flow**
>
> <a id="step-s3"></a>
> 5. **microservice-template-ddd.api** → **client** — 200 · Error
>    status: declared · `internal/api/http/routes_billing.go:36` · HTTP · 200 · `application/json` · json · error · fields: error: string · warning: Error response has no explicit non-2xx status; net/http will send 200.
>
> *otherwise*


> **One of**
>
> *err != nil*
>
> <a id="step-s5"></a>
> 6. **microservice-template-ddd.api** → **client** — 200 · Error
>    status: declared · `internal/api/http/routes_billing.go:44` · HTTP · 200 · `application/json` · json · error · fields: error: string · warning: Error response has no explicit non-2xx status; net/http will send 200. Execution continues after writing the error body and may append another response.
>
> *otherwise*

<a id="step-s7"></a>
7. **microservice-template-ddd.api** → **client** — 200 · GetResponse
   status: declared · `internal/api/http/routes_billing.go:47` · HTTP · 200 · `application/json` · protojson · success · body derived from `billing_rpc.BillingRPC/Get`
