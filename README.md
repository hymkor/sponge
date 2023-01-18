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

Install
-------

Download the binary package from [Releases](https://github.com/hymkor/sponge/releases) and extract the executable.

### for scoop-installer

```
scoop install https://raw.githubusercontent.com/hymkor/sponge/master/sponge.json
```

or

```
scoop bucket add hymkor https://github.com/hymkor/scoop-bucket
scoop install sponge
```

history
-------

* v0.1.0 - repackaging for the scoop-installer and Change License: BSD-3 to MIT
* v0.0.2 (2017.12.14) - 2017.12.14 fix error if target file does not exist.
* v0.0.1 (2016.03.22) - The first version

License: MIT LICENSE
