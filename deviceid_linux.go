//go:build linux

package deviceid

import (
	"fmt"
	"runtime"

	"github.com/zcalusic/sysinfo"
)

func (b *Builder) AddOsVersion() error {
	var si sysinfo.SysInfo
	var osVer string

	si.GetSysInfo()

	if si.OS.Release == "" {
		osVer = si.OS.Vendor + " " + si.OS.Version
	} else {
		osVer = si.OS.Vendor + " " + si.OS.Release
	}

	if osVer == "" {
		b.components[_COMPONENT_OSVERSION_] = _OSVER_UNKNOWN_
	} else {
		b.components[_COMPONENT_OSVERSION_] = osVer
	}

	if b.Debug {
		fmt.Println(b.components[_COMPONENT_OSVERSION_])
	}

	return nil
}

func (b *Builder) AddSerialNumber() error {
	var si sysinfo.SysInfo

	si.GetSysInfo()

	if si.Product.Serial == "System Serial Number" || si.Product.Serial == "" {
		b.components[_COMPONENT_SERIALNUMBER_] = _SERIALNUMBER_UNKNOWN_
	} else {
		b.components[_COMPONENT_SERIALNUMBER_] = si.Product.Serial
	}

	if b.Debug {
		fmt.Println(b.components[_COMPONENT_SERIALNUMBER_])
	}

	return nil
}

func (b *Builder) AddSystemUuid() error {
	var si sysinfo.SysInfo

	si.GetSysInfo()

	if si.Node.MachineID == "" {
		b.components["MachineUuid"] = _MACHINE_UUID_UNKNOWN_
	} else {
		b.components["MachineUuid"] = si.Node.MachineID
	}

	if b.Debug {
		fmt.Println(b.components["MachineUuid"])
	}

	return nil
}

func (b *Builder) AddWindowsDeviceId() error {
	return fmt.Errorf("AddWindowsDeviceId not supported on %s", runtime.GOOS)
}

func (b *Builder) AddWindowsMachineGuid() error {
	return fmt.Errorf("AddWindowsMachineGuid not supported on %s", runtime.GOOS)
}
