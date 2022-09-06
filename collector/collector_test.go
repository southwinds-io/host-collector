/*
  host telemetry collector - © 2018-Present - SouthWinds Tech Ltd - www.southwinds.io
  Licensed under the Apache License, Version 2.0 at http://www.apache.org/licenses/LICENSE-2.0
  Contributors to this project, hereby assign copyright in this code to the project,
  to be licensed under the same terms as the rest of the code.
*/

package collector

import (
	"context"
	"testing"
)

func TestCollectorRun(t *testing.T) {
	collector := NewCollector([]string{"./telem.yaml"}, "0.0.0", nil)
	err := collector.Run(context.Background())
	if err != nil {
		t.Fatal(err.Error())
	}
	status := <-collector.Status()
	if !status.Running {
		t.Fatal("service should be running")
	}
	if status.Err != nil {
		t.Fatal(status.Err.Error())
	}
	collector.Stop()
	status = <-collector.Status()
	if status.Running {
		t.Fatal("service should be stopped")
	}
}
