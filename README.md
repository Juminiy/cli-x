> Usage
from the source code 
````bash 
    go mod tidy 
    go install 
````

````bash
    # rolling update by
    go install 
    # go install again and again 
````

If  you want to change the cli-x name to other ,please correct the `go.mod` module name,some other like: my-cli and run 
````bash
    go mod tidy 
````