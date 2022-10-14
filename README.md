# size

size provides simple byte unit constants. These include decimal units (kilobyte, megabyte) and binary units (kibibyte, mebibyte).

This is useful wherever you may need to provide a number of bytes, helping with readability. For example:

```go
func myHandler(w http.ResponseWriter, r *http.Request) {
    limitedReq := io.LimitReader(r.Body, 3 * size.KB)

	data, err := io.ReadAll(limitedReq)
	if err != nil {
		...
	}
    ...
}
```

### Reference

https://en.wikipedia.org/wiki/Megabyte
