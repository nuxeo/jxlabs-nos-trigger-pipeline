package common

import (
	"os"
	"strconv"

	"github.com/jenkins-x/jx/v2/pkg/cmd/opts"
	"github.com/jenkins-x/jx/v2/pkg/log"
	"github.com/jenkins-x/jx/v2/pkg/util"
	"github.com/spf13/cobra"
)

// BinaryName the binary name to use in help docs
var BinaryName string

// TopLevelCommand the top level command name
var TopLevelCommand string

func init() {
	BinaryName = os.Getenv("BINARY_NAME")
	if BinaryName == "" {
		BinaryName = "tp"
	}
	TopLevelCommand = os.Getenv("TOP_LEVEL_COMMAND")
	if TopLevelCommand == "" {
		TopLevelCommand = "tp"
	}
}

func SetLoggingLevel(cmd *cobra.Command) {
	verbose := false
	flag := cmd.Flag(opts.OptionVerbose)
	if flag != nil {
		var err error
		verbose, err = strconv.ParseBool(flag.Value.String())
		if err != nil {
			log.Logger().Errorf("Unable to check if the verbose flag is set")
		}
	}

	level := os.Getenv("JX_LOG_LEVEL")
	if level != "" {
		if verbose {
			log.Logger().Trace("The JX_LOG_LEVEL environment variable took precedence over the verbose flag")
		}

		err := log.SetLevel(level)
		if err != nil {
			log.Logger().Errorf("Unable to set log level to %s", level)
		}
	} else {
		if verbose {
			err := log.SetLevel("debug")
			if err != nil {
				log.Logger().Errorf("Unable to set log level to debug")
			}
		} else {
			err := log.SetLevel("info")
			if err != nil {
				log.Logger().Errorf("Unable to set log level to info")
			}
		}
	}
}

// SplitCommand helper command to ignore the options object
func SplitCommand(cmd *cobra.Command, options interface{}) *cobra.Command {
	return cmd
}

// GetIOFileHandles lazily creates a file handles object if the input is nil
func GetIOFileHandles(h *util.IOFileHandles) util.IOFileHandles {
	if h == nil {
		h = &util.IOFileHandles{
			Err: os.Stderr,
			In:  os.Stdin,
			Out: os.Stdout,
		}
	}
	return *h
}
