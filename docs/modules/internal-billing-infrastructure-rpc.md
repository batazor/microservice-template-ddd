# internal/billing/infrastructure/rpc

*Generated from the portolan catalog · commit `498f577` · at 2026-09-09T01:20:07Z. Do not edit by hand.*

- **Id:** `local:internal/billing/infrastructure/rpc`
- **Publisher:** [microservice-template-ddd.billing](../microservice-template-ddd/billing/README.md)
- **Source:** `internal/billing/infrastructure/rpc`

## Packages

| Package |
| --- |
| `billing_rpc` |

## Files

| File |
| --- |
| `billing_rpc.proto` |

## Used by

| Service | Access |
| --- | --- |
| [Billing](../microservice-template-ddd/billing/README.md) | publishes |

## Interfaces

### billing_rpc.BillingRPC

- **Source:** `internal/billing/infrastructure/rpc/billing_rpc.proto:13`
- **Module:** [local:internal/billing/infrastructure/rpc](internal-billing-infrastructure-rpc.md)

| Method | Request | Response |
| --- | --- | --- |
| `Get` | `GetRequest` | `GetResponse` |

<a id="message-billing"></a>
<details><summary>Billing</summary>

| Field | Type |
| --- | --- |
| `Balance` | `float` |

</details>

<a id="message-getrequest"></a>
<details><summary>GetRequest</summary>

| Field | Type |
| --- | --- |
| `Id` | `string` |

</details>

<a id="message-getresponse"></a>
<details><summary>GetResponse</summary>

| Field | Type |
| --- | --- |
| `Billing` | `Payload` |

</details>

<a id="message-payload"></a>
<details><summary>Payload</summary>

| Field | Type |
| --- | --- |
| `Billing` | [`Billing`](../types.md#type-billing) |

</details>
