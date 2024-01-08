package main

import (
	"fmt"
	"io"
	"os"
)

type OutputT struct {
	Fd      *os.File
	TmpName string
	Fname   string
}

func mains(in io.Reader, args []string) error {
	outputList := make([]*OutputT, 0, len(args))
	for _, fname := range args {
		tmpName := fname + ".sponge"
		fd, err := os.Create(tmpName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %s\n", tmpName, err.Error())
		} else {
			outputList = append(outputList, &OutputT{
				Fd:      fd,
				TmpName: tmpName,
				Fname:   fname,
			})
		}
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
		err := os.Remove(p.Fname)
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
	if err := mains(os.Stdin, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
