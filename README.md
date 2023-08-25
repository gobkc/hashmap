# go-hash-map
A simple golang hashmap library


### Contributing
You can commit PR to this repository

### How to get it?
````
go get -u github.com/gobkc/hashmap
````

### Quick start
````
package main

import (
	"github.com/gobkc/hashmap"
	"fmt"
)

func main() {
	hs := hashmap.NewHash()
    hs.Set("Alple", "123456")
    fmt.Println(hs.Get("Alple"))
}
````
result:
````
123456
````

### License
© Gobkc, 2023~time.Now

Released under the Apache License