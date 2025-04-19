# deviceid

A small library to generate a 'device ID' to uniquely identify a computer.

Heavily inspired by the C# library written by [Matthew King](https://github.com/MatthewKing) - [DeviceId](https://github.com/MatthewKing/DeviceId)

## Motivation

I couldn't find a device id library for golang that created a unique ID with the level of control and options I was looking for.

I'd used DeviceId by Matthew King in some C# projects, so I'd been spoiled in that language, and wanted something similar for an internal work project. So I created this.

## Installation

Just use go get.

```
go get github.com/craig-jarvis/deviceid
```

## Usage

Create a builder, add at least one component, then generate an id

```
var err error
bldr := deviceid.NewBuilder()
err = bldr.AddSerialNumber()
if err != nil {
	// handle err
}

d, err = bldr.GetDeviceId()
if err != nil {
	// handle err
}
```

alternatively, you can use the opinionated `GetDeviceIdWithDefaults` method.

This method includes the `AddMachineName`, `AddSerialNumber`, and `AddSystemUuid` elements. On windows, it also includes `AddWindowsDeviceId` and `AddWindowsMachineGuid`

```
bldr := deviceid.NewBuilder()
d, err = bldr.GetDeviceIdWithDefaults()
if err != nil {
	// handle err
}
```

### What can you include in a device identifier

- `AddMachineName` adds the host name to the device identifier.
- `AddMacAddress` adds all the MAC addresses to the device identifier. Method supports excluding the default docker bridge interface named "docker0".
- `AddOsVersion` adds the current OS version to the device identifier.
- `AddSerialNumber` adds the motherboard serial number to the device identifier.
- `AddSystemUuid` adds the system UUID to the device identifier.
- `AddWindowsDeviceId` adds the Windows Device ID (also known as Machine ID or Advertising ID) to the device identifier. This value is the one displayed as "Device ID" in the Windows Device Specifications UI.
- `AddWindowsMachineGuid` adds the machine GUID from `HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Cryptography` to the device identifier.

## License and copyright

Copyright (c) 2025 Craig Jarvis
Distributed under the [MIT License](http://opensource.org/licenses/MIT). Refer to LICENSE for more information.
