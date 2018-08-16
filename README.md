# ReconJSON-Go

ReconJSON-Go is an implementation of the ReconJSON standard written for Golang application.

It tries to follow as much as possible the standard to provide feedback, updates and modifications to it.

Tests are provided to unsure its correctness.

## How to use

Install with : 
`go get github.com/edznux/ReconJSON-Go`

Use as a lib in your application :
```
...
import (
    rjson "github.com/edznux/ReconJSON-Go"
)

func main(){
    host := rjson.NewHost()
    
    host.Fqdn = "example.com"
    host.Company = "Exampe inc."
    
    rjson.Write(host, "MyAwesomeRECON.json")

}
```

## Caution

The API is instable at the moment : the standard is not yet well defined.
It doesn't handle any lock mechanism for the moment. Not sure if this is it's job.