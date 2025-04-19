package deviceid

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"os"
	"runtime"
	"slices"
	"strings"
)

type deviceIdBase interface {
	GetDeviceIdWithDefaults() (string, error)
	GetDeviceId() (string, error)
	AddMachineName() error
	AddOsVersion() error
	AddSerialNumber() error
	AddSystemUuid() error
	AddWindowsDeviceId() error
	AddWindowsMachineGuid() error
}

type Builder struct {
	components map[string]string
	Debug      bool
	Trace      bool
}

func NewBuilder() *Builder {
	return &Builder{
		components: make(map[string]string),
		Debug:      false,
	}
}

func (b *Builder) GetDeviceIdWithDefaults() (string, error) {

	b.AddMachineName()
	b.AddSerialNumber()
	b.AddSystemUuid()
	switch runtime.GOOS {
	case "linux":
	case "windows":
		b.AddWindowsDeviceId()
		b.AddWindowsMachineGuid()
	default:
		return "", fmt.Errorf("deviceid does not support %s", runtime.GOOS)
	}

	jStr := orderAndJoinMap(b.components, b.Trace)
	return generateHashString(jStr)
}

func (b *Builder) GetDeviceId() (string, error) {
	if len(b.components) == 0 {
		return "", errors.New("no data provided to generate device id")
	}

	jStr := orderAndJoinMap(b.components, b.Trace)
	return generateHashString(jStr)
}

func (b *Builder) AddMachineName() error {
	hostname, err := os.Hostname()
	if err != nil {
		return err
	}

	b.components[_COMPONENT_COMPUTERNAME_] = hostname
	if b.Debug {
		fmt.Println(hostname)
	}
	return nil
}

func orderAndJoinMap(m map[string]string, trace bool) string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	slices.Sort(keys)

	var values []string
	for _, k := range keys {
		if m[k] != "" {
			values = append(values, m[k])
		}
	}

	r := strings.Join(values, ",")

	if trace {
		fmt.Println(r)
	}

	return r
}

func generateHashString(str string) (string, error) {
	h := sha256.New()
	h.Write([]byte(str))
	bs := h.Sum(nil)

	return fmt.Sprintf("%x", bs), nil
}
