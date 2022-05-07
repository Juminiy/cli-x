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
> Change CLI name
If  you want to change the cli-x name to other ,please correct the `go.mod` module name,some other like: my-cli and run 
````bash
    go mod tidy 
````
> Finished tools
1. http request
2. yaml-file/env config 
3. args 1 parameter receive