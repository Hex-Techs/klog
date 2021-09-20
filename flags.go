package klog

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/spf13/pflag"
)

const logFlushFreqFlagName = "log-flush-frequency"

var logFlushFreq = pflag.Duration(logFlushFreqFlagName, 5*time.Second, "Maximum number of seconds between log flushes")

func AddGlobalFlags(fs *pflag.FlagSet, name string) {
	addKlogFlags(fs)
	addLogFlushFlags(fs)

	fs.BoolP("help", "h", false, fmt.Sprintf("help for %s", name))
}

// addKlogFlags adds flags from k8s.io/klog
func addKlogFlags(fs *pflag.FlagSet) {
	local := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	InitFlags(local)
	normalizeFunc := fs.GetNormalizeFunc()
	local.VisitAll(func(fl *flag.Flag) {
		fl.Name = string(normalizeFunc(fs, fl.Name))
		fs.AddGoFlag(fl)
	})
}

func addLogFlushFlags(fs *pflag.FlagSet) {
	fs.AddFlag(pflag.Lookup(logFlushFreqFlagName))
}
