echo a>foo.txt
echo b>>foo.txt
echo c>>foo.txt
echo d>>foo.txt
cat -n foo.txt | sponge foo.txt
type foo.txt
