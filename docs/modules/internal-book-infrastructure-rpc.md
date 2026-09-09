# internal/book/infrastructure/rpc

*Generated from the portolan catalog · commit `baebd80` · at 2026-09-09T21:28:15+07:00. Do not edit by hand.*

- **Id:** `local:internal/book/infrastructure/rpc`
- **Publisher:** [microservice-template-ddd.book](../microservice-template-ddd/book/README.md)
- **Source:** `internal/book/infrastructure/rpc`

## Packages

| Package |
| --- |
| `book_rpc` |

## Files

| File |
| --- |
| `book_rpc.proto` |

## Used by

| Service | Access |
| --- | --- |
| [Book](../microservice-template-ddd/book/README.md) | publishes |

## Interfaces

### book_rpc.BookRPC

- **Source:** `internal/book/infrastructure/rpc/book_rpc.proto:9`
- **Module:** [local:internal/book/infrastructure/rpc](internal-book-infrastructure-rpc.md)

| Method | Request | Response |
| --- | --- | --- |
| `Get` | `GetRequest` | `GetResponse` |
| `Rent` | `RentRequest` | `RentResponse` |

<a id="message-book"></a>
<details><summary>Book</summary>

| Field | Type |
| --- | --- |
| `Title` | `string` |
| `Author` | `string` |
| `IsRent` | `bool` |

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
| `Book` | [`Book`](../types.md#type-book) |

</details>

<a id="message-rentrequest"></a>
<details><summary>RentRequest</summary>

| Field | Type |
| --- | --- |
| `Id` | `string` |

</details>

<a id="message-rentresponse"></a>
<details><summary>RentResponse</summary>

| Field | Type |
| --- | --- |
| `Book` | [`Book`](../types.md#type-book) |

</details>
