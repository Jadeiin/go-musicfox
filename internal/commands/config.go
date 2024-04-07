package commands

import (
	"fmt"
	"path"

	"github.com/anhoder/foxful-cli/util"
	"github.com/gookit/gcli/v2"
	"github.com/muesli/termenv"

	"github.com/go-musicfox/go-musicfox/internal/types"
	"github.com/go-musicfox/go-musicfox/utils"
)

func NewConfigCommand() *gcli.Command {
	cmd := &gcli.Command{
		Name:   "config",
		UseFor: "Print configuration file to be loaded",
		Func: func(_ *gcli.Command, _ []string) error {
			var configPath = util.SetFgStyle(path.Join(utils.GetLocalDataDir(), types.AppIniFile), termenv.ANSICyan)
			fmt.Printf("Loaded Configuration File:\n\t%s\n", configPath)
			return nil
		},
	}
	return cmd
}
