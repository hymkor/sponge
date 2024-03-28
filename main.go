package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	flagBackupPostfix = flag.String("b", "~", "Postfix for backup of original files")
	flagTmpPostfix    = flag.String("t", ".sponge", "Postfix for temporary files")
)

type OutputT struct {
	Fd      *os.File
	TmpName string
	Fname   string
}

func mains(in io.Reader, args []string) error {
	outputList := make([]*OutputT, 0, len(args))
	for _, fname := range args {
		perm := os.FileMode(0600)
		stat, err := os.Stat(fname)
		if err == nil {
			perm = stat.Mode()
		}
		tmpName := fname + *flagTmpPostfix
		fd, err := os.OpenFile(tmpName, os.O_WRONLY|os.O_CREATE|os.O_EXCL, perm)
		if err != nil {
			return err
		}
		outputList = append(outputList, &OutputT{
			Fd:      fd,
			TmpName: tmpName,
			Fname:   fname,
		})
	}

	for {
		var buffer [4096]byte
		n, err := in.Read(buffer[:])
		if err != nil && err != io.EOF {
			for _, p := range outputList {
				p.Fd.Close()
			}
			return err
		}
		if n > 0 {
			for _, p := range outputList {
				p.Fd.Write(buffer[:n])
			}
		}
		if err == io.EOF {
			break
		}
	}

	for _, p := range outputList {
		p.Fd.Close()
		backupName := p.Fname + *flagBackupPostfix
		err := os.Rename(p.Fname, backupName)
		if err != nil && !os.IsNotExist(err) {
			fmt.Fprintln(os.Stderr, err.Error())
			continue
		}
		err = os.Rename(p.TmpName, p.Fname)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}
	}
	return nil
}

func main() {
	flag.Parse()
	if err := mains(os.Stdin, flag.Args()); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
