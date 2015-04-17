## Simple Command Line to validate Cloud-Config
Based off https://github.com/coreos/coreos-cloudinit

###Dependencies
```
go get gopkg.in/yaml.v1
go get github.com/coreos/yaml
go get github.com/coreos/coreos-cloudinit/config
```

###Run
```
go run main.go -config config.yml | echo $?
```