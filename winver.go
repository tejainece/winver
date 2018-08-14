package winver

import (
	"fmt"
	"syscall"
)

// Version contains data about a Windows version
type Version struct {
	Major uint8
	Minor uint8
	Build uint16
}

func (v Version) String() string {
	return fmt.Sprintf("Windows version %d.%d (Build %d)\n", v.Major, v.Minor, v.Build)
}

// IsVista returns if the version is Vista
func (v Version) IsVista() bool {
	return v.Major == 6 && v.Minor == 0
}

// IsAtleastWindows7 returns if the version is  atleast Windows 7
func (v Version) IsAtleastWindows7() bool {
	return v.Major >= 6 && v.Minor >= 1
}

// IsWindows7 returns if the version is Windows 7
func (v Version) IsWindows7() bool {
	return v.Major == 6 && v.Minor == 1
}

// IsAtleastWindows8 returns if the version is atleast Windows 8
func (v Version) IsAtleastWindows8() bool {
	return v.Major >= 6 && v.Minor >= 2
}

// IsWindows8 returns if the version is Windows 8
func (v Version) IsWindows8() bool {
	return v.Major == 6 && (v.Minor == 2 || v.Minor == 3)
}

// IsWindows10 returns if the version is Windows 10
func (v Version) IsWindows10() bool {
	return v.Major == 10 && v.Minor == 0
}

func GetVersion() Version {
	v, _ := syscall.GetVersion()
	ver := Version{
		Major: byte(v),
		Minor: uint8(v >> 8),
		Build: uint16(v >> 16),
	}
	return ver
}
