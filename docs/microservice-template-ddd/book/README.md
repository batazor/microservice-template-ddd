# Book

*Generated from the portolan catalog · commit `498f577` · at 2026-09-09T01:20:07Z. Do not edit by hand.*

- **Id:** `microservice-template-ddd.book`
- **Group:** [Microservice Template Ddd](../README.md)
- **Repo:** [`microservice-template-ddd`](https://microservice-template-ddd)
- **Path:** `./`
- **Kind:** service
- **Technologies:** Go, Docker

## microservice-template by ddd

DDD example use micriservices.

#### Getting Started

```
# Get help
make help

# Run services
make run

# Stop services
make down
```

###### Prerequisites

+ docker
+ docker-compose
+ protoc 3.7.1+

#### Service

| Name        | Port  | Description Endpoint          |
|-------------|-------|-------------------------------|
| traefik     | 80    | HTTP                          |
| traefik     | 443   | HTTPS                         |
| traefik     | 8060  | Dashboard                     |
| api         | 7070  | HTTP API                      |
| user        | 50051 | gRPC Server                   |
| billing     | 50051 | gRPC Server                   |


#### Architecture

```
.
├── /cmd/                       # Run service endpoint
│   ├── /user/                  # User service
│   ├── /book/                  # Book service
│   └── /billing/               # Billing service
├── /pkg/                       # The public source code of the application
├── /internal/                  # The private source code of the application
│   └── /bookService/           # Book service source code
│       ├── /useCases/          # Write business logic [./application]
│       ├── /domian/            # Entity struct that represent mapping to data model
│       └── /infrastructure/    # Solves backend technical topics
│           ├── /store/         # Store delivery [../repository]
│           ├── /rpc/           # RPC delivery
│           └── /mq/            # MQ delivery
├── /ops/                       # All infrastructure configuration for IoC
├── .gitignore                  # A gitignore file specifies untracked files
└── README.md                   # README
```

#### HTTP API

+ Import [Postman link](./docs/microservice-template-ddd.postman_collection.json) for
  test HTTP API

#### Request example

![example](./docs/request.png)

###### Opentracing example request

![example](./docs/tracer1.png)
![example](./docs/tracer2.png)

## Aggregates

| Aggregate | Root | Commands | Queries | Events |
| --- | --- | --- | --- | --- |
| [Book](aggregates/book.md) | `Book` | 0 commands | 0 queries | 0 events |

## Provides

### book_rpc.BookRPC

- **Source:** `internal/book/infrastructure/rpc/book_rpc.proto:9`
- **Module:** [local:internal/book/infrastructure/rpc](../../modules/internal-book-infrastructure-rpc.md)

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
| `Book` | [`Book`](../../types.md#type-book) |

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
| `Book` | [`Book`](../../types.md#type-book) |

</details>

## Consumes

| Call | Peer | Status | Source |
| --- | --- | --- | --- |
| `billing_rpc.BillingRPC/Get` | [microservice-template-ddd.billing](../billing/README.md) | declared | `internal/billing/infrastructure/rpc/billing_rpc_grpc.pb.go` |
| `user_rpc.UserRPC/Get` | [microservice-template-ddd.user](../user/README.md) | declared | `internal/user/infrastructure/rpc/user_rpc_grpc.pb.go` |

## Schema modules

| Module | Access | Packages |
| --- | --- | --- |
| [internal/book/domain](../../modules/internal-book-domain.md) | publishes | book |
| [internal/book/infrastructure/rpc](../../modules/internal-book-infrastructure-rpc.md) | publishes | book_rpc |

## Stores

| Store | Kind | Access | Schema |
| --- | --- | --- | --- |
| [book Redis](stores/redis.md) | redis | owns | 2 key patterns |

## Commands

| Run | Does | Body | Source |
| --- | --- | --- | --- |
| `make default` | — | — | `Makefile:8` |
| `make Makefile` | — | — | `Makefile:9` |
| `make help` | Display this help screen | `@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)` | `Makefile:25` |
