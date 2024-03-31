v0.2.0
======
Mar 31, 2024

+ Fix error occured on Windows when `cat -n < FILE | sponge FILE` was executed
  because the shell did not close FILE.
  (Just rename once and at least complete the replacement)
+ Implement the new option: `-b SUFFIX` to rename and keep the original file with suffix.
+ Now terminates with an error when a file with the same name as a temporary file exists
+ The permission of new files are set same as original files now.
+ Use `(original-name)-sponge(process-id)` as the format for temporary filenames
+ Zero byte files are not created even when errors or interrupts stop
+ `-h`: print args[0], version, GOOS, GOARCH to the standard error output

v0.1.1
=======
Jan 8, 2024

+ Prevented copying of the block immediately before EOF from being leaked
+ Add tests

v0.1.0
=======
Jan 15, 2023

* Repackaging for the scoop-installer and Change License: BSD-3 to MIT

v0.0.2 (20171214)
=======
Dec 14, 2017

+ Fix error if target-file does not exist

v0.0.1 (20160322)
=======
Mar 22, 2016

+ The first release
