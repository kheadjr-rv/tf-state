package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hashicorp/terraform/addrs"
	"github.com/hashicorp/terraform/states"
	"github.com/hashicorp/terraform/states/statefile"
)

var (
	dumpOnly *bool
)

func main() {

	dumpOnly = flag.Bool("dump", false, "show state debug")

	flag.Parse()

	fi, _ := os.Stdin.Stat()

	var sf *statefile.File
	var err error
	if (fi.Mode() & os.ModeCharDevice) == 0 {
		sf, err = statefile.Read(os.Stdin)
		if err != nil {
			panic(err)
		}
	} else {
		flag.Usage()
		fmt.Println()
		fmt.Println("cat terraform.tfstate | tf-state")
		fmt.Println("terraform state pull | tf-state -dump")
		os.Exit(1)
	}

	if *dumpOnly {
		dump(sf.State)
		os.Exit(0)
	}

	diagram(sf.State)

}

func mode(mode addrs.ResourceMode) string {
	switch mode {
	case addrs.DataResourceMode:
		return "data"
	case addrs.ManagedResourceMode:
		return "resource"
	default:
		return "unknown"
	}
}

func diagram(state *states.State) {
	fmt.Println("graph LR")
	for key, module := range state.Modules {

		if module.Addr.Parent().String() == "" && !module.Addr.IsRoot() {
			fmt.Printf("  root-->%s(%s)\n", key, key)
			continue
		}

		if module.Addr.Parent().String() != "" {
			_, call := module.Addr.Call()
			fmt.Printf("  %s-->%s(%s)\n", module.Addr.Parent().String(), key, call)
		}
	}
}

func dump(sf *states.State) {
	// Module Level
	for k, m := range sf.Modules {

		fmt.Printf("\n%s\n", k)

		// Resource Level
		for name, r := range m.Resources {

			fmt.Printf("\t%s\n", mode(r.Addr.Resource.Mode))
			fmt.Printf("\t%s\n", name)
			fmt.Printf("\t%s\n", r.Addr.Resource.Type)
			fmt.Printf("\t%d\n", len(r.Instances))
			fmt.Printf("\t%s\n", r.ProviderConfig.Provider.Type)

			for ik, iv := range r.Instances {

				if ik == nil {
					// Lets assume its TF 11
					for key, value := range iv.Current.AttrsFlat {
						fmt.Println(key, value)
					}
					continue
				}

				fmt.Printf("\t%s%s\n", name, ik)

				for _, dep := range iv.Current.Dependencies {
					fmt.Printf("\t\t Depends On %s.%s\n", dep.Resource.Type, dep.Resource.Name)

				}

			}

			fmt.Println()

		}
	}
}
