# ggja

[**Repository moved to codeberg.org**](https://codeberg.org/fractalqb/ggja)

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

For more details read on or consult
[godoc](https://godoc.org/github.com/fractalqb/ggja).

# JSON Object Member by `string` or `fmt.Stringer`

**TODO:** Access methods ending with 's', e.g. `.Bools` accept a `fmt.Stringer`
as member name. This might be useful with consts that have a `Stringer`.

# JSON Array Element by Index

Simply by index (that's it for now):

```go
strVal := ggjaArr.Str(3, strVal)
```

Updates strVal in case ggjaArr has at least four elements. 

# Access Methods for Basic Types

Access methods for basic types come in two flavours:

1. A _conditional_ one, that returns a default value `nvl` if the requested value
   is not in the JSON object or array. Those methods are simply called after the
   type they return, e.g. `ggjaObj.Str("name", "-")` returns a string.

2. A _mandatory_ flavour that calls your provided `OnError` function if the
   requested value is not preset. If you provide no error function it will
   `panic` instead. Mandatory access methods start with 'M', e.g. `MInt(…)`

# Access Methods for Array and Objects

Besides the access methods for basic types there are also access methods for
objects and arrays. Both come in the two flavours _conditional_ and _mandatory_.
The _conditional_ access for arrays and objects does not have a default value
but returns `nil` in case the is nothing that can be accessed. 

However array element and object member access have a third _generating_
flavour that creates an empty array resp. object on access if needed. Generating
access methods start with 'C' for _create_, e.g. `CObj(…)`.

```go
ggjaArr := Arr{OnError: fail}
objMbr3 := ggjaArr.CObj(3)
objMbr3.Put("name", "John Doe")
jStr, _ := json.Marshal(ggjaArr.Bare)
fmt.Println(string(jStr))
jStr, _ = json.Marshal(objMbr3.Bare)
fmt.Println(string(jStr))
```
writes output

```
[null,null,null,{"name":"John Doe"}]
{"name":"John Doe"}
```
