# GO-UNITEST [![Go Report Card](https://goreportcard.com/badge/github.com/verlandz/go-unitest)](https://goreportcard.com/report/github.com/verlandz/go-unitest) [![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/verlandz/go-pkg)](https://golang.org/doc/go1.14) [![License](https://img.shields.io/github/license/verlandz/go-unitest)](https://github.com/verlandz/go-unitest/blob/main/LICENSE)

## Example

| Sample | Desc |
| ------------- | ------------- |
| simple | Simple unitest that most of us already know and get used to it. |
| dependency-injection | Robust unitest that have connection to other layer from result test. It's not merely mock the result, but it will pick the result from dest's test. <br><br> **ex:** if you have changes in `usecase`, you probably not only fix the test in `usecase`, but also in `delivery` since `delivery` test pick the mock result from `usecase`. |
| mock-vendor | To do some mock, we need interface. Sometimes there's some vendor who only provide struct with the functions. To mock that we need to step back, generate our own interface and store it. |

## Rule of thumb
| About | Desc |
| ------------- | ------------- |
| Using test name convention | Our testName using [this convention](https://www.agilealliance.org/glossary/gwt/). Sometimes the devs tend to forget to do it properly, therefore having [this builder](https://github.com/verlandz/go-pkg/blob/9b2d523f26a034e0c843c6b91bce43eaa22e1032/tester/name.go) will reduce the problem. |
| Centralized data test | Sometimes we need to reuse data over and over again which lead dry in unitest. To avoid that, we need to centralized data somewhere. In here, we put under dir [test-data](https://github.com/verlandz/go-unitest/tree/example/dependency-injection/test-data) |
| Don't test `private function` | If there's part of code can't covered in `private function` from `public function` test then perhaps that part doesn't needed / can't be reached. |
| Avoid `mock.Anything` | This is powerful yet will make devs not know the exact param. Sometimes there's case that we need this such as `random`. |
