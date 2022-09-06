/*
  host telemetry collector - © 2018-Present - SouthWinds Tech Ltd - www.southwinds.io
  Licensed under the Apache License, Version 2.0 at http://www.apache.org/licenses/LICENSE-2.0
  Contributors to this project, hereby assign copyright in this code to the project,
  to be licensed under the same terms as the rest of the code.
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"southwinds.dev/artisan/core"
)

type RootCmd struct {
	Cmd *cobra.Command
}

// NewRootCmd creates the root command
func NewRootCmd() *RootCmd {
	// https://textkool.com/en/ascii-art-generator?hl=default&vl=default&font=Red%20Phoenix&text=host%20collector
	c := &RootCmd{
		Cmd: &cobra.Command{
			Use:   "telemetry",
			Short: "Telemetry",
			Long: fmt.Sprintf(`
+--------------------------------------------------------------------------------------------------+
| .__                      __                      .__   .__                   __                  | 
| |  |__    ____   _______/  |_       ____   ____  |  |  |  |    ____   ____ _/  |_  ____ _______  | 
| |  |  \  /  _ \ /  ___/\   __\    _/ ___\ /  _ \ |  |  |  |  _/ __ \_/ ___\\   __\/  _ \\_  __ \ | 
| |   Y  \(  <_> )\___ \  |  |      \  \___(  <_> )|  |__|  |__\  ___/\  \___ |  | (  <_> )|  | \/ | 
| |___|  / \____//____  > |__|       \___  >\____/ |____/|____/ \___  >\___  >|__|  \____/ |__|    | 
|      \/             \/                 \/                         \/     \/                      |  
+--------------------------------------------------------------------------------------------------+

version: %s`, core.Version),
			Version: core.Version,
		},
	}
	c.Cmd.SetVersionTemplate("Host version: {{.Version}}\n")
	return c
}
