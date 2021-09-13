package reg

import (
	"context"

	"github.com/FlowingSPDG/windows-ndi-optimizer/optimizer"
)

// Optimize NIC driver via registry

const (
	registryNICSettingsPath = "HKEY_LOCAL_MACHINE\\SYSTEM\\ControlSet001\\Control\\Class\\{4d36e972-e325-11ce-bfc1-08002be10318}"
	// registryNetworkGroupSettingsPath = ""
)

// taskDisableEEE Registry optimize task. Implements Optimzier interface
type taskDisableEEE struct {
}

func (t *taskDisableEEE) Run(ctx context.Context) error {
	// TODO: レジストリからEEEを無効化する
	return nil
}

func NewTaskDisableEEE() optimizer.Task {
	return &taskDisableEEE{}
}

type makeNetworkPrivate struct {
}

func (t *makeNetworkPrivate) Run(ctx context.Context) error {
	// TODO: PowerShellからネットワークをプライベートに設定する
	return nil
}

func NewMakeNetworkPrivate() optimizer.Task {
	return &makeNetworkPrivate{}
}
