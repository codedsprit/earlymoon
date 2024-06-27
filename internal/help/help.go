package help

import (
	"fmt"
	
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
` + boldText + cyanColor + " Query dns fast with earlymoon " + resetColor + `

` + boldText + "Usage:" + resetColor + ` ` + greenColor + boldText + "earlymoon [Flags]" + resetColor + `

` + boldText + "Flags:" + resetColor + `
    ` + yellowColor + boldText + "-d" + resetColor + `       Domain to query ` + magentaColor + boldText + `(required)` + resetColor + `
    ` + yellowColor + boldText + "-t" + resetColor + `       Type of Dns record ` + magentaColor + boldText + `(required)` + resetColor + `

` + boldText + "Options:" + resetColor + ` ` + greenColor + resetColor + `
    ` + yellowColor + boldText + "-v" + resetColor + `       Print Version ` + resetColor + `
    ` + yellowColor + boldText + "help" + resetColor + `     Display this help message `
	fmt.Println(banner)
}

