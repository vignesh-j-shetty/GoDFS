//go:build !windows
package platform

import "golang.org/x/sys/unix"

func GetFreeSpace(path string) (uint64, error) {
    var stat unix.Statfs_t
    err := unix.Statfs(path, &stat)
    if err != nil {
        return 0, err
    }
    return stat.Bavail * uint64(stat.Bsize), nil
}