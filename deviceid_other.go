//go:build !(windows || linux)

package deviceid

import (
	"fmt"
	"runtime"
)

func (b *Builder) AddOsVersion() error {
	return fmt.Errorf("AddOsVersion not supported on %s\n", runtime.GOOS)
}

func (b *Builder) AddSerialNumber() error {
	return fmt.Errorf("AddSerialNumber not supported on %s\n", runtime.GOOS)
}

func (b *Builder) AddSystemUuid() error {
	return fmt.Errorf("AddSystemUuid not supported on %s\n", runtime.GOOS)
}

func (b *Builder) AddWindowsDeviceId() error {
	return fmt.Errorf("AddWindowsDeviceId not supported on %s\n", runtime.GOOS)
}

func (b *Builder) AddWindowsMachineGuid() error {
	return fmt.Errorf("AddWindowsMachineGuid not supported on %s\n", runtime.GOOS)
}
