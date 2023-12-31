package intermediate

import (
	"io"
	"log"
	"os"
)

// handling temporary resources created by programs e.g files and network connections
// * a simple version of the linux `cat` command
func Defer() {
	if len(os.Args) < 2 {
		log.Fatal("no file specified")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	data := make([]byte, 2048)
	for {
		count, err := f.Read(data)
		os.Stdout.Write(data[:count])
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}
}
