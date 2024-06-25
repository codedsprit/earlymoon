package help

import (
	"flag"
	"fmt"
	"os"
	
)

const (
	resetColor   = "\033[0m"
	boldText     = "\033[1m"
	redColor     = "\033[31m"
	greenColor   = "\033[32m"
	yellowColor  = "\033[33m"
	blueColor    = "\033[34m"
	magentaColor = "\033[35m"
	cyanColor    = "\033[36m"
	whiteColor   = "\033[37m"
)

const Version = "0.1.0"

// PrintBanner displays the banner for the earlymoon project.
func PrintBanner() {
	banner := boldText + yellowColor + `
      ┓           
┏┓┏┓┏┓┃┓┏┏┳┓┏┓┏┓┏┓
┗ ┗┻┛ ┗┗┫┛┗┗┗┛┗┛┛┗ v` + Version + `
        ┛` + resetColor + `
` + boldText + cyanColor + " Query dns fast with earlymoon " + resetColor + `

` + boldText + "Usage:" + resetColor + ` ` + greenColor + boldText + "earlymoon [Flags]" + resetColor + `

` + boldText + "Flags:" + resetColor + `
    ` + yellowColor + boldText + "-d" + resetColor + `       Domain to query ` + magentaColor + boldText + `(required)` + resetColor + `
    ` + yellowColor + boldText + "-t" + resetColor + `       Type of Dns record ` + magentaColor + boldText + `(required)` + resetColor + `

` + boldText + "Options:" + resetColor + ` ` + greenColor + resetColor + `
    ` + yellowColor + boldText + "-V" + resetColor + `       Print Version ` + resetColor + `
    ` + yellowColor + boldText + "help" + resetColor + `     Display this help message `
	fmt.Println(banner)
}

// Args holds the command-line arguments
type Args struct {
	Domain      string
	RecordType  string
	ShowVersion bool
}

// ParseArgs parses the command-line arguments and checks for unknown arguments
func ParseArgs() Args {
	var args Args
	flag.StringVar(&args.Domain, "domain", "", "The domain to query (required)")
	flag.StringVar(&args.Domain, "d", "", "The domain to query (required)")
	flag.StringVar(&args.RecordType, "type", "", "The type of DNS record to query (required)")
	flag.StringVar(&args.RecordType, "t", "", "The type of DNS record to query (required)")
	flag.BoolVar(&args.ShowVersion, "V", false, "print version information")

	// Override default usage
	flag.Usage = func() {
		PrintBanner()
	}

	// Parse flags
	flag.Parse()

	return args
}

// HandleArgs processes the parsed arguments and performs the corresponding actions
func HandleArgs(args Args) {
	if args.ShowVersion {
		fmt.Println("earlymoon", Version)
		os.Exit(0)
	}

	if len(os.Args) == 1 || os.Args[1] == "help" || os.Args[1] == "-h" {
		PrintBanner()
		os.Exit(0)
	}

	if args.Domain == "" || args.RecordType == "" {
		PrintBanner()
		os.Exit(1)
	}
}
