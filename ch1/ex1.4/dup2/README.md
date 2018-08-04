Input from STDIN.

```
$ echo 'foo
bar
baz
foo' | go run ch1/ex1.4/dup2/main.go
2       foo     /dev/stdin
```

Input from files.

```
gopl-book-exercise (master) $ go run ch1/ex1.4/dup2/main.go sample1.txt sample2.txt
4       foo     sample1.txt,sample2.txt
3       bar     sample1.txt,sample2.txt
2       baz     sample1.txt,sample2.txt
```
