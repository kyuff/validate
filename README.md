# validate

[![Build Status](https://github.com/kyuff/validate/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/kyuff/validate/actions/workflows/go.yml)
[![Report Card](https://goreportcard.com/badge/github.com/kyuff/es)](https://goreportcard.com/report/github.com/kyuff/validate/)
[![Go Reference](https://pkg.go.dev/badge/github.com/kyuff/validate.svg)](https://pkg.go.dev/github.com/kyuff/validate)
[![codecov](https://codecov.io/gh/kyuff/validate/graph/badge.svg?token=EY0LT9XASR)](https://codecov.io/gh/kyuff/validate)

A composable validation library for Go domain types. Validators are plain functions that return errors, compose freely via `All()`, and work with named types out of the box thanks to generics. Every validation run collects all errors rather than stopping at the first failure.

## Installation

```bash
go get github.com/kyuff/validate
```

## Quick Start

Implement the `Validator` interface on your domain type and compose rules with `All()`:

```go
type UserID string
type Email string

type CreateUserCommand struct {
    ID    UserID
    Email Email
    Age   int
}

func (c CreateUserCommand) Validate() error {
    return validate.All(
        validate.Required(c.ID),
        validate.StringNotEmpty(c.Email),
        validate.StringMatches(c.Email, regexp.MustCompile(`^[^@]+@[^@]+$`)),
        validate.NumberMin(c.Age, 18),
    )
}

func handle(cmd CreateUserCommand) error {
    if err := cmd.Validate(); err != nil {
        return err // all rule failures are collected and returned together
    }
    // ...
}
```

To check whether an error came from validation (as opposed to a downstream failure), use `errors.Is`:

```go
var validationErr validate.Error
if errors.Is(err, validationErr) {
    // respond with 400 Bad Request
}
```

## Validator Reference

Every validator comes in two forms:

- **Plain** — `StringNotEmpty(value)` uses a built-in error message.
- **Formatted** — `StringNotEmptyf(value, "name is required")` lets you supply a custom message (with `fmt.Sprintf`-style formatting).

### String

| Function | Fails when |
|---|---|
| `StringNotEmpty(value)` | value is `""` |
| `StringMinLength(value, min)` | `len(value) < min` |
| `StringMaxLength(value, max)` | `len(value) > max` |
| `StringMatches(value, pattern)` | value does not match the regexp |
| `StringOneOf(value, options...)` | value is not in options |

All accept any type whose underlying kind is `string` (e.g. `type ID string`).

### Number

| Function | Fails when |
|---|---|
| `NumberMin(value, min)` | `value < min` |
| `NumberMax(value, max)` | `value > max` |
| `NumberPositive(value)` | `value <= 0` |
| `NumberNegative(value)` | `value >= 0` |
| `NumberBetween(value, min, max)` | `value < min \|\| value > max` |

Supports all integer and float kinds, including named types.

### Bool

| Function | Fails when |
|---|---|
| `BoolTrue(value)` | value is `false` |
| `BoolFalse(value)` | value is `true` |

### Required / Nil

| Function | Fails when |
|---|---|
| `Required(value)` | value equals its zero value (any `comparable`) |
| `NotNil(ptr)` | pointer is `nil`; validates the pointed-to value otherwise |
| `IfNotNil(ptr)` | pointer is non-`nil` and the pointed-to value is invalid |

### Slices

```go
err := validate.SliceContainsf(roles, "admin", "user must have the admin role")
```

`SliceContainsf` returns an error when `target` is not present in `values`. It returns an `error` directly rather than a `ValidatorFunc`, so it can be used standalone or wrapped in `ValidatorFunc`.

### Struct Fields

`FieldsStrict` validates a struct by calling `Validate()` on every field. All fields must implement `Validator`; the function fails immediately if one does not.

```go
type Address struct {
    Street validate.ValidatorFunc
    City   validate.ValidatorFunc
}

func NewAddress(street, city string) Address {
    return Address{
        Street: validate.StringNotEmpty(street),
        City:   validate.StringNotEmpty(city),
    }
}

func (a Address) Validate() error {
    return validate.FieldsStrict(a)()
}
```

Reflection metadata is cached per type, so repeated calls are cheap.

## Composing Validators

`All` runs every validator and joins all errors with `errors.Join`. Nil validators are silently skipped.

```go
err := validate.All(
    validate.Required(order.ID),
    validate.NumberPositive(order.Amount),
    validate.StringNotEmpty(order.Currency),
)
// err contains all failures, not just the first one
```

## Middleware

`Middleware` wraps any `func(context.Context, T) error` handler. Before calling `next`, it checks that the argument implements `Validator`, is not nil, and passes validation. Non-`validate.Error` errors returned from `Validate()` are automatically promoted to `validate.Error`.

```go
type PlaceOrderCommand struct { /* ... */ }

func (c PlaceOrderCommand) Validate() error {
    return validate.All(
        validate.Required(c.OrderID),
        validate.NumberPositive(c.Amount),
    )
}

func placeOrder(ctx context.Context, cmd PlaceOrderCommand) error {
    // business logic — validation already passed
    return nil
}

var handler = validate.Middleware(placeOrder)

// caller
err := handler(ctx, cmd) // returns validate.Error if cmd is invalid
```

## Custom Validators

Any function matching `func() error` can be used as a `ValidatorFunc`:

```go
validate.ValidatorFunc(func() error {
    if !isValidIBAN(account.IBAN) {
        return validate.Errorf("invalid IBAN: %s", account.IBAN)
    }
    return nil
})
```

To make a type self-validating, implement the `Validator` interface:

```go
type Validator interface {
    Validate() error
}
```

Self-validating types work with `NotNil`, `IfNotNil`, `FieldsStrict`, and `Middleware` out of the box.
