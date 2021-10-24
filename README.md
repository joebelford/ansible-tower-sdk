# ansible-tower-sdk

AWX (Ansible Tower) SDK for the Go programming language.


This SDK has been developed against AWX `14.0.0`.

## Installing

```
go get -u github.com/Kaginari/ansible-tower-sdk
```

## Example

We can simply import `goawx` and call its services, such as the PingService:

```
import (
    "log"
    tower "github.com/Kaginari/ansible-tower-sdk"
)

func main() {
    client := tower.NewAWX("http://awx.your_server_host.com", "your_awx_username", "your_awx_passwd", nil)
    result, err := client.PingService.Ping()
    if err != nil {
        log.Fatalf("Ping awx err: %s", err)
    }

    log.Println("Ping awx: ", result)
}
```

More examples can be found at [here](https://github.com/Kaginari/ansible-tower-sdk/tree/main/examples).

## Roadmap

ansible-tower-sdk is still in development, and its roadmap could be found at [here](https://github.com/Kaginari/ansible-tower-sdk/blob/main/ROADMAP.md).

## Contribute

There are many ways to contribute to awx-go.

* Submit bugs via [Github issues](https://github.com/Kaginari/ansible-tower-sdk/issues);
* Submit a [pull request](https://github.com/Kaginari/ansible-tower-sdk/pulls) for fixes or features;

## awx sdk

* this sdk is forked from awx github.com/mrcrilly/goawx  to add some other api supports
