# httpdump
tiny program to dump raw http request

# usage

```shell
$go get github.com/imzjy/httpdump
$./bin/httpdump
```

# todo

* response based on regex, support type of string and file, like

```text
/getfile/, file, c:\test.txt
/ip/, string, 192.168.1.1
/hostname/, string, myhost
```