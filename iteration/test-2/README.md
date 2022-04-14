# Learn Go With Tests:

### Writing benchmarks in Go is another first-class feature of the language and it is very similar to writing tests.

* You'll see the code is very similar to a test.
* The testing.B gives you access to the cryptically named b.N.
* When the benchmark code is executed, it runs b.N times and measures how long it takes.
* To run the benchmarks do go test -bench=. (or if you're in Windows Powershell go test -bench=".")


#### What `NUMBER -> n` ns/op means is our function takes on average n nanoseconds to run (on my computer). Which is pretty ok! To test this it ran it 10000000 times.