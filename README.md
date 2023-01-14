sponge(clone)
=============

soak up standard input and write to a file.
([original version](https://joeyh.name/code/moreutils/))

```
$ cat foo.txt
a
b
c
d
$ cat -n foo.txt | sponge foo.txt
$ cat foo.txt
     1 a
     2 b
     3 c
     4 d
```
(Japanese: 標準入力を全て読み取ってから、その内容を引数のファイルに出力します)

history
-------
* 2017.12.14 fix error if target file does not exist.

License: New BSD-License
