sponge (clone)
==============

soak up standard input and write to a file.
([original version](https://joeyh.name/code/moreutils/),
 [my Rust version](https://github.com/hymkor/sponge-rs))

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

### Use "go install"

```
go install github.com/hymkor/sponge@latest
```

### Use "scoop-installer"

```
scoop install https://raw.githubusercontent.com/hymkor/sponge/master/sponge.json
```

or

```
scoop bucket add hymkor https://github.com/hymkor/scoop-bucket
scoop install sponge
```

Usage
-----

`sponge {options} {FILENAME(s)...}`

+ `-b string`
    + Postfix for backup of original files
+ `-v
    + verbose`

---

[Release Note(English)](./release_note_en.md)
[Release Note(Japanese)](./release_note_ja.md)
