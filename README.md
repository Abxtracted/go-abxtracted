<p align="center">
  <img alt="Abxtracted Logo" src="https://avatars2.githubusercontent.com/u/26093365?v=3&s=200" height="140" />
  <h3 align="center">go-abxtracted</h3>
  <p align="center">Get experiments complexity abstracted and make your A/B tests trustable.</p>
  <p align="center">
    <a href="https://github.com/Abxtracted/go-abxtracted/releases/latest"><img alt="Release" src="https://img.shields.io/github/release/Abxtracted/go-abxtracted.svg?style=flat-square"></a>
    <a href="/LICENSE.md"><img alt="Software License" src="https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square"></a>
    <a href="https://travis-ci.org/Abxtracted/go-abxtracted"><img alt="Travis" src="https://img.shields.io/travis/Abxtracted/go-abxtracted.svg?style=flat-square"></a>
    <a href="https://codecov.io/gh/Abxtracted/go-abxtracted"><img alt="Codecov branch" src="https://img.shields.io/codecov/c/github/Abxtracted/go-abxtracted/master.svg?style=flat-square"></a>
    <a href="https://goreportcard.com/report/github.com/Abxtracted/go-abxtracted"><img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/Abxtracted/go-abxtracted?style=flat-square"></a>
    <a href="http://godoc.org/github.com/Abxtracted/go-abxtracted"><img alt="Go Doc" src="https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square"></a>
  </p>
</p>


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
