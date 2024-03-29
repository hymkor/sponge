package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

var (
	flagBackupPostfix = flag.String("b", "", "Postfix for backup of original files")
	flagVerbose       = flag.Bool("v", false, "verbose")
)

type OutputT struct {
	tmpFile *LazyFile
	orgName string
}

func mains(in io.Reader, args []string) error {
	logger := func(...any) {}
	if *flagVerbose {
		logger = func(args ...any) {
			fmt.Fprintln(os.Stderr, args...)
		}
	}
	outputList := make([]*OutputT, 0, len(args))
	for _, fname := range args {
		perm := os.FileMode(0600)
		stat, err := os.Stat(fname)
		if err == nil {
			perm = stat.Mode()
		}
		outputList = append(outputList, &OutputT{
			tmpFile: &LazyFile{
				Name: fmt.Sprintf("%s-sponge%d", fname, os.Getpid()),
				Flag: os.O_WRONLY | os.O_CREATE | os.O_EXCL,
				Perm: perm,
			},
			orgName: fname,
		})
	}

	for {
		var buffer [4096]byte
		n, err := in.Read(buffer[:])
		if err != nil && err != io.EOF {
			for _, p := range outputList {
				p.tmpFile.Close()
			}
			return err
		}
		if n > 0 {
			for _, p := range outputList {
				p.tmpFile.Write(buffer[:n])
			}
		}
		if err == io.EOF {
			break
		}
	}

	for _, p := range outputList {
		p.tmpFile.Close()
		postfix := *flagBackupPostfix
		var backupName string
		if postfix == "" {
			postfix = time.Now().Format("~20060102_150405~")
			defer func() {
				logger("rm", backupName)
				os.Remove(backupName)
			}()
		}
		backupName = p.orgName + postfix
		logger("rename", p.orgName, backupName)
		err := os.Rename(p.orgName, backupName)
		if err != nil && !os.IsNotExist(err) {
			fmt.Fprintln(os.Stderr, err.Error())
			continue
		}
		logger("rename", p.tmpFile.Name, p.orgName)
		err = os.Rename(p.tmpFile.Name, p.orgName)
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
