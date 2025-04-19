package deviceid

import (
	"testing"
)

func TestNewBuilder(t *testing.T) {
	bldr := NewBuilder()

	if bldr.components == nil {
		t.Errorf("builder components is not of type map[string]string")
	}
}

func TestGetDeviceIdWithJustComputerName(t *testing.T) {
	bldr := NewBuilder()
	bldr.components[_COMPONENT_COMPUTERNAME_] = "host1"

	d, err := bldr.GetDeviceId()
	if err != nil {
		t.Errorf("Failed to get device id, %v", err)
	}

	want := "c0365b5a3867cc382f6854fdc4f6f10c7857275c8b1e525beb8c399f80949be5"

	if d != want {
		t.Errorf("incorrect hash generated. got: [%s], want: [%s]", d, want)
	}
}

func TestAddMacAddress(t *testing.T) {
	bldr := NewBuilder()
	bldr.AddMacAddress(false)
	bldr.components[_COMPONENT_MACADDRESS_] = "12:34:56:AB:CD:EF"

	d, err := bldr.GetDeviceId()
	if err != nil {
		t.Errorf("Failed to get device id, %v", err)
	}

	want := "0e8aa98797e189d504b13506b27aace3c22f181d4b320ffc9ff9beef4affddb6"

	if d != want {
		t.Errorf("incorrect hash generated. got: [%s], want: [%s]", d, want)
	}
}

func TestGetDeviceIdWithJustOsVersion(t *testing.T) {
	bldr := NewBuilder()
	bldr.components[_COMPONENT_OSVERSION_] = "Microsoft Windows [Version 10.0.26100.3775]"

	d, err := bldr.GetDeviceId()
	if err != nil {
		t.Errorf("Failed to get device id, %v", err)
	}

	want := "c9ad7587bfa9c4c6593454cfd4a5d07b807d8215b4c33aa640253223edab460f"

	if d != want {
		t.Errorf("incorrect hash generated. got: [%s], want: [%s]", d, want)
	}
}

func TestGetDeviceIdWithJustSerialNumber(t *testing.T) {
	bldr := NewBuilder()
	bldr.components[_COMPONENT_SERIALNUMBER_] = "serial12345"

	d, err := bldr.GetDeviceId()
	if err != nil {
		t.Errorf("Failed to get device id, %v", err)
	}

	want := "45eb0e52947e574e4ea6147a7cdaceeb5bedb83305937908dde75b75f9347139"

	if d != want {
		t.Errorf("incorrect hash generated. got: [%s], want: [%s]", d, want)
	}
}

func TestGetDeviceIdWithJustSystemUuid(t *testing.T) {
	bldr := NewBuilder()
	bldr.components[_COMPONENT_SYSTEMUUID_] = "1b18534d-a15c-46b6-b2a2-61132c6e8480"

	d, err := bldr.GetDeviceId()
	if err != nil {
		t.Errorf("Failed to get device id, %v", err)
	}

	want := "826e3c8036655bd1463b06ee8ec855e3386b4ee673e6d43f0abf5a4a82c463c4"

	if d != want {
		t.Errorf("incorrect hash generated. got: [%s], want: [%s]", d, want)
	}
}

func TestGetDeviceIdWithJustWindowsDeviceId(t *testing.T) {
	bldr := NewBuilder()
	bldr.components[_COMPONENT_MACHINEID_] = "f923d167-90b7-4a1a-b113-604eba1f1ec6"

	d, err := bldr.GetDeviceId()
	if err != nil {
		t.Errorf("Failed to get device id, %v", err)
	}

	want := "46ee224fc34004e2a7969282637860fdca1701ef853d3496d54460af31d72420"

	if d != want {
		t.Errorf("incorrect hash generated. got: [%s], want: [%s]", d, want)
	}
}

func TestGetDeviceIdWithJustWindowsMachineGuid(t *testing.T) {
	bldr := NewBuilder()
	bldr.components[_COMPONENT_MACHINEGUID_] = "173d78dc-9cb0-433d-bb3c-e292fa322274"

	d, err := bldr.GetDeviceId()
	if err != nil {
		t.Errorf("Failed to get device id, %v", err)
	}

	want := "46541eab774d4c6b50e705ba21175797514f054f7e376325bc777f1807a58b96"

	if d != want {
		t.Errorf("incorrect hash generated. got: [%s], want: [%s]", d, want)
	}
}

func TestGetDeviceIdWithAllComponents(t *testing.T) {
	bldr := NewBuilder()
	bldr.components[_COMPONENT_COMPUTERNAME_] = "host1"
	bldr.components[_COMPONENT_OSVERSION_] = "Microsoft Windows [Version 10.0.26100.3775]"
	bldr.components[_COMPONENT_SERIALNUMBER_] = "serial12345"
	bldr.components[_COMPONENT_SYSTEMUUID_] = "1b18534d-a15c-46b6-b2a2-61132c6e8480"
	bldr.components[_COMPONENT_MACHINEID_] = "f923d167-90b7-4a1a-b113-604eba1f1ec6"
	bldr.components[_COMPONENT_MACHINEGUID_] = "173d78dc-9cb0-433d-bb3c-e292fa322274"

	d, err := bldr.GetDeviceId()
	if err != nil {
		t.Errorf("Failed to get device id, %v", err)
	}

	want := "3840d9b3ccf061d69c654e9ce867d2a750d1832e2dc9afc1a16c7f73c83b5c02"

	if d != want {
		t.Errorf("incorrect hash generated. got: [%s], want: [%s]", d, want)
	}
}

func TestDeviceIdShouldNotUseEmptyComponent(t *testing.T) {
	bldr := NewBuilder()
	bldr.components[_COMPONENT_COMPUTERNAME_] = "host1"
	bldr.components[_COMPONENT_SERIALNUMBER_] = ""

	d, err := bldr.GetDeviceId()
	if err != nil {
		t.Errorf("Failed to get device id, %v", err)
	}

	want := "c0365b5a3867cc382f6854fdc4f6f10c7857275c8b1e525beb8c399f80949be5"

	if d != want {
		t.Errorf("incorrect hash generated. got: [%s], want: [%s]", d, want)
	}
}

func TestOrderAndJoinMap(t *testing.T) {
	bldr := NewBuilder()
	bldr.components[_COMPONENT_COMPUTERNAME_] = "host1"
	bldr.components[_COMPONENT_SERIALNUMBER_] = ""

	d := orderAndJoinMap(bldr.components, false)

	lstChr := d[len(d)-1:]

	if lstChr == "," {
		t.Error("Joined map of components should not end in a comma [,]")
	}
}
