//go:build windows

package platform

import (
    "syscall"
    "unsafe"
)

func GetFreeSpace(path string) (uint64, error) {
    var freeBytesAvailable, totalNumberOfBytes, totalNumberOfFreeBytes uint64
    pathPtr, _ := syscall.UTF16PtrFromString(path)
    r, _, err := syscall.NewLazyDLL("kernel32.dll").
        NewProc("GetDiskFreeSpaceExW").
        Call(
            uintptr(unsafe.Pointer(pathPtr)),
            uintptr(unsafe.Pointer(&freeBytesAvailable)),
            uintptr(unsafe.Pointer(&totalNumberOfBytes)),
            uintptr(unsafe.Pointer(&totalNumberOfFreeBytes)),
        )
    if r == 0 {
        return 0, err
    }
    return freeBytesAvailable, nil
}
