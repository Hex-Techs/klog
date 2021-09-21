package klog

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/spf13/pflag"
)

func AddGlobalFlags(fs *pflag.FlagSet, name string) {
	addKlogFlags(fs)

	fs.BoolP("help", "h", false, fmt.Sprintf("help for %s", name))
}

// addKlogFlags adds flags from klog
func addKlogFlags(fs *pflag.FlagSet) {
	local := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	InitFlags(local)
	normalizeFunc := fs.GetNormalizeFunc()
	local.VisitAll(func(fl *flag.Flag) {
		fl.Name = string(normalizeFunc(fs, fl.Name))
		fs.AddGoFlag(fl)
	})
}

func FlushSchedule(fn func(flush func(), frequency time.Duration)) {
	go fn(Flush, logging.flushFrequency)
}
