package sample

import (
	"fmt"
	"github.com/akamensky/argparse"
	"github.com/multipackage_1/package_2"
	"os"

	. "github.com/logrusorgru/aurora"
)

func SampleTwo() {
	fmt.Println(package_2.OneIntPlusAnother(3, 4))
	fmt.Println("Hello,", Magenta("Aurora"))
	fmt.Println(Bold(Cyan("Cya!")))
}

func Sample() {
	parser := argparse.NewParser("print", "Prints provided string to stdout")
	s := parser.String("s", "string", &argparse.Options{Required: true, Help: "String to print"})
	// Parse input
	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
	}
	// Finally print the collected string
	fmt.Println(*s)
}
