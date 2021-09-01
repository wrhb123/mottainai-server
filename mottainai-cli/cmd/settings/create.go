/*

Copyright (C) 2017-2021  Ettore Di Giacinto <mudler@gentoo.org>
                         Daniele Rondina <geaaru@sabayonlinux.org>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.

*/

package settingcmd

import (
	"fmt"
	"log"
	"os"

	utils "github.com/MottainaiCI/mottainai-server/mottainai-cli/cmd/utils"
	tools "github.com/MottainaiCI/mottainai-server/mottainai-cli/common"
	setting "github.com/MottainaiCI/mottainai-server/pkg/settings"

	cobra "github.com/spf13/cobra"
)

func newSettingCreateCommand(config *setting.Config) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "create <key> <value>",
		Short: "Create a new setting",
		Args:  cobra.RangeArgs(2, 2),

		// TODO: PreRun check of minimal args if --json is not present
		Run: func(cmd *cobra.Command, args []string) {
			fetcher, err := utils.CreateClient(config)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}

			dat := make(map[string]interface{})

			if len(args) != 2 {
				log.Fatalln("You need to define akey and a value to create")
			}
			dat["key"] = args[0]
			dat["value"] = args[1]

			res, err := fetcher.SettingCreate(dat)
			tools.CheckError(err)
			tools.PrintResponse(res)
		},
	}

	return cmd
}