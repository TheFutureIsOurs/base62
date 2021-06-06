# base62
Go base62 encoding and decoding

[go doc](https://pkg.go.dev/github.com/TheFutureIsOurs/base62)

## Install

	go get github.com/TheFutureIsOurs/base62

## How to use


```go
// encode
num := 264688217293324297
str := base62.Encode(int64(num))
fmt.Println(str)

// decodde
res, err := base62.Decode(str)
fmt.Println(res, err)

```