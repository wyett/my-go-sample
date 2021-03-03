// @author     : wyettLei
// @date       : Created in 2021/3/2 11:58
// @description: TODO

package flagsample

import (
	"commonpkg/utils"
	"flag"
)

const (
	tConfig  = "config"
	tVersion = "version"
	tUsage   = "usage"

	// base
)

var (
	configure = flag.String(tConfig, "", "config file name")
	version   = flag.String(tVersion, "", "version")
	usage     = flag.Bool(tUsage, false, "print usage info")

	// base

)

func parseCommand() {
	flag.Parse()
	defer utils.HandleExit()

	// check args
	if *configure == "" || *version == "" {
		utils.Crash("args has no values", -1)
	}

	// check args and print

}
