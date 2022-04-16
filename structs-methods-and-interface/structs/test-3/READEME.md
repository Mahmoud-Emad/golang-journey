# Learn Go With Tests

* Now that you have some understanding of structs we can introduce "table driven tests".
Table driven tests are useful when you want to build a list of test cases that can be tested in the same manner

</p>
The only new syntax here is creating an "anonymous struct", areaTests. We are declaring a slice of structs by using
</p>

`[]struct` with two fields, the shape and the want. Then we fill the slice with cases.

We then iterate over them just like we do any other slice, using the struct fields to run our tests.

You can see how it would be very easy for a developer to introduce a new shape, implement Area and then add it to 
the test cases. In addition, if a bug is found with Area it is very easy to add a new test case to exercise it before fixing it.
Table driven tests can be a great item in your toolbox, but be sure that you have a need for the extra noise in the tests. They are a great fit when you wish to test various implementations of an interface, or if the data being passed in to a function has lots of different requirements that need testing.
Let's demonstrate all this by adding another shape and testing it; a triangle.