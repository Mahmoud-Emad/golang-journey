# Learn Go With Tests:

You will notice that we're using %d as our format strings rather than %q. That's because we want it to print an integer rather than a string.
Also note that we are no longer using the main package, instead we've defined a package named integers, as the name suggests this will group functions for working with integers such as Add.

##### in this part we'll use the examples:
* If you publish your code with examples to a public URL, you can share the documentation of your code at pkg.go.dev. 
* This web interface allows you to search for documentation of standard library packages and third-party packages.
* To run and see it just write `godoc -http=:6060`, and if you have not install godoc just use `sudo apt install golang-golang-x-tools` to install it # For Linux
* [Now go through to see it running](http://localhost:6060/pkg/)