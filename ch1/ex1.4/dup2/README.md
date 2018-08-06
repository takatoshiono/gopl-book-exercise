Input from STDIN.

```
$ echo 'foo
bar
baz
foo' | go run main.go
2       foo     /dev/stdin
```

Input from files.

```
$ go run main.go sample1.txt sample2.txt
3       foo     sample1.txt,sample2.txt
3       bar     sample1.txt,sample2.txt
2       baz     sample1.txt
```
