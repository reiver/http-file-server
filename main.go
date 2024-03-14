package main

import (
	"flag"
	"fmt"
	"mime"
	"net/http"
	"os"
)

var port string
var directory string

func init() {
	flag.StringVar(&port, "http-port", "8080", "TCP-port to use to serve HTTP")
	flag.StringVar(&directory, "directory", ".", "the local directory of the static files to serve over HTTP")

	flag.Parse()
}

func init() {
	mime.AddExtensionType(".js", "application/javascript")
}

func main() {
	log("Hello world! 🙂")

	logf("http-port = %q", port)
	logf("directory = %q", directory)

	var handler http.Handler
	{
		var dir http.Dir = http.Dir(directory)
		var filesystem http.FileSystem = dir
		handler = http.FileServer(filesystem)
		if nil == handler {
			logerror("problem creating file-system HTTP-handler: received a nil HTTP-handler")
			return
		}
	}

	var addr string
	{
		addr = ":" + port
		logf("addr = %q", port)
	}

	{
		err := http.ListenAndServe(addr, handler)
		if nil != err {
			logerrorf("err = (%T) %s", err, err)
			panic(err)
		}
	}
}

func log(a ...interface{}) {
	s := fmt.Sprint(a...)

	logf("%s", s)
}

func logf(format string, a ...interface{}) {
	format = "[LOG] " + format + "\n"

	fmt.Fprintf(os.Stdout, format, a...)
}

func logerror(a ...interface{}) {
	s := fmt.Sprint(a...)

	logerrorf("%s", s)
}

func logerrorf(format string, a ...interface{}) {
	format = "[ERROR] " + format + "\n"

	fmt.Fprintf(os.Stderr, format, a...)
}
