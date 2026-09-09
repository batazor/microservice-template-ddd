# Microservice Template Ddd

*Generated from the portolan catalog · commit `baebd80` · at 2026-09-09T21:28:15+07:00. Do not edit by hand.*


## Contexts

| Context | Services |
| --- | --- |
| [Microservice Template Ddd](microservice-template-ddd/README.md) | [API](microservice-template-ddd/api/README.md), [Billing](microservice-template-ddd/billing/README.md), [Book](microservice-template-ddd/book/README.md), [User](microservice-template-ddd/user/README.md) |

## Schema modules

| Module | Publisher | Packages |
| --- | --- | --- |
| [internal/billing/domain](modules/internal-billing-domain.md) | [microservice-template-ddd.billing](microservice-template-ddd/billing/README.md) | 1 package |
| [internal/billing/infrastructure/rpc](modules/internal-billing-infrastructure-rpc.md) | [microservice-template-ddd.billing](microservice-template-ddd/billing/README.md) | 1 package |
| [internal/book/domain](modules/internal-book-domain.md) | [microservice-template-ddd.book](microservice-template-ddd/book/README.md) | 1 package |
| [internal/book/infrastructure/rpc](modules/internal-book-infrastructure-rpc.md) | [microservice-template-ddd.book](microservice-template-ddd/book/README.md) | 1 package |
| [internal/user/domain](modules/internal-user-domain.md) | [microservice-template-ddd.user](microservice-template-ddd/user/README.md) | 1 package |
| [internal/user/infrastructure/rpc](modules/internal-user-infrastructure-rpc.md) | [microservice-template-ddd.user](microservice-template-ddd/user/README.md) | 1 package |

## Flows

| Flow | Owner | Summary |
| --- | --- | --- |
| [GET /billing/](flows/api-http-get-billing.md) | [microservice-template-ddd](microservice-template-ddd/README.md) | Source-derived http execution path for `GET /billing/`. Source-backed cross-protocol continuations are included. |
| [GET /book/](flows/api-http-get-book.md) | [microservice-template-ddd](microservice-template-ddd/README.md) | Source-derived http execution path for `GET /book/`. Source-backed cross-protocol continuations are included. |
| [GET /user/](flows/api-http-get-user.md) | [microservice-template-ddd](microservice-template-ddd/README.md) | Source-derived http execution path for `GET /user/`. Source-backed cross-protocol continuations are included. |
| [POST /book/rent/{bookId}](flows/api-http-post-book-rent-bookid.md) | [microservice-template-ddd](microservice-template-ddd/README.md) | Source-derived http execution path for `POST /book/rent/{bookId}`. Source-backed cross-protocol continuations are included. |
