package task

import (
	"context"
	"log"

	"golang.org/x/sys/windows/registry"
)

// Optimize NIC driver via registry

const (
	registryNICSettingsPath          = "SYSTEM\\ControlSet001\\Control\\Class\\{4d36e972-e325-11ce-bfc1-08002be10318}"
	registryNetworkGroupSettingsPath = "SOFTWARE\\Microsoft\\Windows NT\\CurrentVersion\\NetworkList\\Profiles"
)

// taskDisableEEE Registry optimize task. Implements Task interface
type taskDisableEEE struct {
}

func (t *taskDisableEEE) Run(ctx context.Context) error {

	return nil
}

// NewTaskDisableEEE get taskDisableEEE
func NewTaskDisableEEE() Task {
	return &taskDisableEEE{}
}

// taskMakeNetworkPrivate Make network private task. Implements Task interface
type taskMakeNetworkPrivate struct {
}

func (t *taskMakeNetworkPrivate) Run(ctx context.Context) error {
	// TODO: レジストリからネットワークをプライベートに設定する
	return nil
}

// NewTaskMakeNetworkPrivate Get taskMakeNetworkPrivate
func NewTaskMakeNetworkPrivate() Task {
	return &taskMakeNetworkPrivate{}
}

// taskDisableJumboPacket Disable jumbo packet task. Implements Task interface
type taskDisableJumboPacket struct {
}

func (t *taskDisableJumboPacket) Run(ctx context.Context) error {
	const defaultVal = "1514"

	k, err := registry.OpenKey(registry.LOCAL_MACHINE, registryNICSettingsPath+"\\0004", registry.ALL_ACCESS)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()

	s, _, err := k.GetStringValue("*JumboPacket")
	if err != nil {
		log.Fatal(err)
	}
	if s != defaultVal {
		log.Printf("*JumboPacket value is %q\n", s)
		log.Println("Enforcing default...")
		if err := k.SetStringValue("*JumboPacket", defaultVal); err != nil {
			return err
		}
	}

	return nil
}

// NewTaskDisableJumboPacket Get taskDisableJumboPacket
func NewTaskDisableJumboPacket() Task {
	return &taskDisableJumboPacket{}
}
