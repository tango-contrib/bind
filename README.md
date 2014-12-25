Bind [![Build Status](https://drone.io/github.com/tango-contrib/bind/status.png)](https://drone.io/github.com/tango-contrib/bind/latest) [![](http://gocover.io/_badge/github.com/tango-contrib/bind)](http://gocover.io/github.com/tango-contrib/bind)
======

** HEAVILLY DEVELOPMENT **
Bind is a query param auto mapping middleware for [Tango](https://github.com/lunny/tango). 

## Installation

    go get github.com/tango-contrib/bind

## Simple Example

```Go
package main

import (
    "github.com/lunny/tango"
    "github.com/tango-contrib/bind"
)

ype BindExample struct {
    Id   int64
    Name string
}

func (a *BindExample) Get() string {
    return fmt.Sprintf("%d-%s", a.Id, a.Name)
}

func main() {
    o := tango.Classic()
    o.Use(bind.Default())
    o.Get("/", new(BindExample))
}
```

## Getting Help

- [API Reference](https://gowalker.org/github.com/tango-contrib/bind)
