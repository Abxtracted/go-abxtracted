# go-abxtracted

## Usage

```go
var abx = abxtracted.New()
exp, err := abx.Get("project-id", "customer-id", "experiment-key")
if err != nil {
  // deal with it
}
// later on..
if err := abx.Complete("project-id", exp); err != nil {
  // deal with it
}
```

## Questions

If you have any question, [drop us a message](https://abxtracted.com/contact).
