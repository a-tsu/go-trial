my personal setting of go

```
$ export PATH=$PATH:/usr/local/opt/go/libexec/bin
$ export GOPATH=$HOME/Documents/lab/go
$ mkdir -p ${GOPATH}/src/github.com/a-tsu/go-trial
$ go install github.com/a-tsu/go-trial
$ ${GOPATH}/bin/go-trial
```

or git clone & try below
```
$ go run hello.go
```
