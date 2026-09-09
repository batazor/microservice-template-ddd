# internal/user/infrastructure/rpc

*Generated from the portolan catalog · commit `baebd80` · at 2026-09-09T21:28:15+07:00. Do not edit by hand.*

- **Id:** `local:internal/user/infrastructure/rpc`
- **Publisher:** [microservice-template-ddd.user](../microservice-template-ddd/user/README.md)
- **Source:** `internal/user/infrastructure/rpc`

## Packages

| Package |
| --- |
| `user_rpc` |

## Files

| File |
| --- |
| `user_rpc.proto` |

## Used by

| Service | Access |
| --- | --- |
| [User](../microservice-template-ddd/user/README.md) | publishes |

## Interfaces

### user_rpc.UserRPC

- **Source:** `internal/user/infrastructure/rpc/user_rpc.proto:9`
- **Module:** [local:internal/user/infrastructure/rpc](internal-user-infrastructure-rpc.md)

| Method | Request | Response |
| --- | --- | --- |
| `Get` | `GetRequest` | `GetResponse` |

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
| `User` | [`User`](../types.md#type-user) |

</details>

<a id="message-user"></a>
<details><summary>User</summary>

| Field | Type |
| --- | --- |
| `Login` | `string` |
| `Password` | `string` |
| `Email` | `string` |
| `IsActive` | `bool` |

</details>
