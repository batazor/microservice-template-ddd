# Shared types

*Generated from the portolan catalog · commit `baebd80` · at 2026-09-09T21:28:15+07:00. Do not edit by hand.*

Types named by more than one aggregate, event or message. A field that
refers to one of these is knowably the same shape everywhere it appears.

<a id="type-billing"></a>

## Billing

| Field | Type |
| --- | --- |
| `Balance` | `float` |

<a id="type-book"></a>

## Book

| Field | Type |
| --- | --- |
| `Title` | `string` |
| `Author` | `string` |
| `IsRent` | `bool` |

<a id="type-user"></a>

## User

| Field | Type |
| --- | --- |
| `Login` | `string` |
| `Password` | `string` |
| `Email` | `string` |
| `IsActive` | `bool` |
