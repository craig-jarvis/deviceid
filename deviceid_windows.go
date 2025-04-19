//go:build windows

package deviceid

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/yusufpapurcu/wmi"
	"golang.org/x/sys/windows/registry"
)

var computerSystemData *WmiComputerSystemProduct

type WmiComputerSystemProduct struct {
	IdentifyingNumber string
	Name              string
	Vendor            string
	Version           string
	Caption           string
	Description       string
	SKUNumber         string
	UUID              string
}

func (b *Builder) AddOsVersion() error {
	osVer, err := getWmiOperatingSystem()
	if err != nil {
		b.components[_COMPONENT_OSVERSION_] = _OSVER_UNKNOWN_
	}

	b.components[_COMPONENT_OSVERSION_] = osVer

	if b.Debug {
		fmt.Println(osVer)
	}

	return nil
}

func (b *Builder) AddSerialNumber() error {
	if computerSystemData == nil {
		getWmiComputerSystemProduct()
	}

	b.components[_COMPONENT_SERIALNUMBER_] = computerSystemData.IdentifyingNumber
	if b.Debug {
		fmt.Println(computerSystemData.IdentifyingNumber)
	}
	return nil
}

func (b *Builder) AddSystemUuid() error {
	if computerSystemData == nil {
		getWmiComputerSystemProduct()
	}

	r := strings.ToLower(computerSystemData.UUID)

	b.components[_COMPONENT_SYSTEMUUID_] = r
	if b.Debug {
		fmt.Println(r)
	}
	return nil
}

func (b *Builder) AddWindowsDeviceId() error {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\SQMClient`, registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()

	s, _, err := k.GetStringValue(_COMPONENT_MACHINEID_)
	if err != nil {
		log.Fatal(err)
	}

	r := strings.TrimLeft(s, "{")
	r = strings.TrimRight(r, "}")
	r = strings.ToLower(r)

	b.components["WindowsDeviceId"] = r
	if b.Debug {
		fmt.Println(r)
	}
	return nil
}

func (b *Builder) AddWindowsMachineGuid() error {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Cryptography`, registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()

	s, _, err := k.GetStringValue(_COMPONENT_MACHINEGUID_)
	if err != nil {
		log.Fatal(err)
	}

	r := strings.TrimLeft(s, "{")
	r = strings.TrimRight(r, "}")

	b.components["WindowsMachineGuid"] = r
	if b.Debug {
		fmt.Println(r)
	}
	return nil
}

func getWmiComputerSystemProduct() error {
	dst := []WmiComputerSystemProduct{}

	namespace := `root\cimv2`
	qcsp := "SELECT * FROM Win32_ComputerSystemProduct"
	ctx := context.Background()
	err := wmiQueryWithContext(ctx, qcsp, &dst, namespace)
	if err != nil {
		return err
	}

	if len(dst) != 1 {
		return fmt.Errorf("error getting Win32_ComputerSystemProduct")
	}

	computerSystemData = &dst[0]

	return nil
}

func getWmiOperatingSystem() (string, error) {
	var dst []struct {
		Caption string
		Version string
	}

	namespace := `root\cimv2`
	qOs := "SELECT Version, Caption FROM Win32_OperatingSystem"
	ctx := context.Background()
	err := wmiQueryWithContext(ctx, qOs, &dst, namespace)
	if err != nil {
		return "", err
	}

	if len(dst) != 1 {
		return "", fmt.Errorf("error getting Win32_OperatingSystem")
	}

	data := &dst[0]
	result := fmt.Sprintf("%s [%s]", data.Caption, data.Version)

	return result, nil
}

func wmiQueryWithContext(ctx context.Context, query string, dst interface{}, namespace string) error {
	if _, ok := ctx.Deadline(); !ok {
		ctxTimeout, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		ctx = ctxTimeout
	}

	errChan := make(chan error, 1)
	go func() {
		errChan <- wmi.QueryNamespace(query, dst, namespace)
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-errChan:
		return err
	}
}
