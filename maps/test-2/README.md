# Learn Go With Tests

### Our Add is looking good. Except, we didn't consider what happens when the value we are trying to add already exists!
* Map will not throw an error if the value already exists. Instead, they will go ahead and overwrite the value with the newly provided value. This can be convenient in practice, but makes our function name less than accurate. Add should not modify existing values. It should only add new words to our dictionary.