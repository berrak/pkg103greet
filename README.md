## pkg103greet

### Simple golang repository structure for a single library
Shows howto debianize the pkg103greet library without any vendor dependencies.
The packaging process can be easy or it may require more effort based on the repository layout.
The required commands is under *docs* and is based on Debian stretch (stable) and use dh-make-golang and other standard Debian tools.

```bash
├── docs
│   └── debanize_pkg103greet.md
├── greet.go
├── greet_test.go
├── LICENSE
└── README.md
```

### API example
```bash

package main

import  "fmt"
import  grt "github.com/berrak/pkg103greet"

// A spanish greeting
func main(){
    fmt.Println(grt.GreetMe("sp"))
}
```
### Usage
The created package only serves as a Debian golang library packaging example. 
This library function will be used in 104 application example.

### Background
The purpose of this series is described in the [https://github.com/berrak/enterprise101].

### License
This project is under the Apache License. See the LICENSE file for the full license text.
