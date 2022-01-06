# stringtyper

stringtype tries to determimine the Go type that will best satisfy
representing a series of one or more example strings supplied by the
user.

Where "best" is the most specific Go type that can represent what is
presented, using the range of the type where appropriate. 
So for int and float types, the range can constrain the type that is
selected, with int and float type ranges defined in
[https://pkg.go.dev/builtin](https://pkg.go.dev/builtin). 
Because of this, the example strings presented need to represent both
the type and the range that is to be used. 

For examples, see the tests in
[https://github.com/gnewton/stringtyper/blob/main/pkg/stringtyper/stringtyper_test.go](https://github.com/gnewton/stringtyper/blob/main/pkg/stringtyper/stringtyper_test.go).

# Usage







