package task

import (
	"context"
)

// Optimize NIC driver via registry

const (
	registryNICSettingsPath          = "HKEY_LOCAL_MACHINE\\SYSTEM\\ControlSet001\\Control\\Class\\{4d36e972-e325-11ce-bfc1-08002be10318}"
	registryNetworkGroupSettingsPath = "HKEY_LOCAL_MACHINE\\SOFTWARE\\Microsoft\\Windows NT\\CurrentVersion\\NetworkList\\Profiles"
)

// taskDisableEEE Registry optimize task. Implements Optimzier interface
type taskDisableEEE struct {
}

func (t *taskDisableEEE) Run(ctx context.Context) error {
	// TODO: レジストリからEEEを無効化する
	return nil
}

func NewTaskDisableEEE() Task {
	return &taskDisableEEE{}
}

type taskMakeNetworkPrivate struct {
}

func (t *taskMakeNetworkPrivate) Run(ctx context.Context) error {
	// TODO: PowerShellからネットワークをプライベートに設定する
	return nil
}

func NewTaskMakeNetworkPrivate() Task {
	return &taskMakeNetworkPrivate{}
}
