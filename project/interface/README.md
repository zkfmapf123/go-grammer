## Interface

### Type Check use interface
```go
    var v interface{}
	switch v.(type) {
	case int:
		fmt.Println("int")

	default:
		fmt.Println("unknown")
	}
```