Input from STDIN.

```
$ echo 'foo
quote> bar
quote> baz
quote> foo' | go run ch1/ex1.4/dup2/main.go
2       foo
```

Input from files.

```
$ go run ch1/ex1.4/dup2/main.go sample1.txt sample2.txt
4       foo
3       bar
2       baz
```
