# ggja
[![Build Status](https://travis-ci.org/fractalqb/ggja.svg)](https://travis-ci.org/fractalqb/ggja)
[![codecov](https://codecov.io/gh/fractalqb/ggja/branch/master/graph/badge.svg)](https://codecov.io/gh/fractalqb/ggja)
[![Go Report Card](https://goreportcard.com/badge/github.com/fractalqb/ggja)](https://goreportcard.com/report/github.com/fractalqb/ggja)
[![GoDoc](https://godoc.org/github.com/fractalqb/ggja?status.svg)](https://godoc.org/github.com/fractalqb/ggja)

`import "git.fractalqb.de/fractalqb/ggja"`

---
# Intro

Go Generic JSON Api is a small helper library to make working with Go's generic
JSON types more comfortable.

Instead of writing this coed fragment

```go
if tmp, ok := genJson["key"]; !ok {
	log.Panic("no member key in JSON object")
}
i := tmp.(int)
```

with ggja it can be somewhat shorter:

```go
i := ggjaObj.MInt("key")
```

OK, one first has to wrap the generic JSON `genJson map[string]interface{}`
into a ggja object wrapper:

```go
ggjaObj := ggja.Obj{
	Bare: genJson,
	OnError: func(err error) { log.Panic(err) },
}
```

But this is only done once and—also useful—you can pass the error handling
strategy around with the `ggjaObj`.

# JSON Object Member by `string` or `fmt.Stringer`

**TODO:** Access methods ending with 's', e.g. `.Bools` accept a `fmt.Stringer`
as member name. This might be useful with consts that have a `Stringer`.

# Access Methods for Basic Types

Access methods for basic types come in two flavours:

1. A conditional one, that returns a default value `nvl` if the requested value
   is not in the JSON object or array. Those methods are simply called after
   type they return, e.g. `ggjaObj.Str("name", "-")` returns a string.

2. A mandatory flavour that calls your provided `OnError` function if the
   requested value is not preset. If you provide no error function it will
   `panic` instead.

# Access Methods for Array and Objects

**TODO**