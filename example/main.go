package main

import (
	"log"
	"os/exec"

	"github.com/crosbymichael/cgroups"
	specs "github.com/opencontainers/runtime-spec/specs-go"
)

var shares = uint64(100)

func demoCgroups() error {
	control, err := cgroups.V1(cgroups.Unified, cgroups.StaticPath("/testv1"), &specs.Resources{
		CPU: &specs.CPU{
			Shares: &shares,
		},
	})
	if err != nil {
		return err
	}
	cmd := exec.Command("sleep", "2")
	if err := cmd.Start(); err != nil {
		return err
	}
	if err := control.Add(cmd.Process.Pid); err != nil {
		return err
	}
	stats, err := control.Stat(false)
	if err != nil {
		return err
	}
	log.Printf("%#v\n", stats)
	cmd.Wait()
	return control.Delete()
}

func main() {
	if err := demoCgroups(); err != nil {
		log.Fatal(err)
	}
}