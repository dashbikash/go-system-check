package main

import (
	"fmt"

	"golang.org/x/sys/unix"
)

func hasIPCLock() (bool, error) {
	var caps unix.CapUserHeader
	caps.Version = unix.LINUX_CAPABILITY_VERSION_3
	caps.Pid = 0 // 0 = self

	var data [2]unix.CapUserData
	err := unix.Capget(&caps, &data[0])
	if err != nil {
		return false, err
	}

	// CAP_IPC_LOCK is capability 14
	const CAP_IPC_LOCK = 14
	effective := (data[0].Effective & (1 << CAP_IPC_LOCK)) != 0
	return effective, nil
}

func main() {
	ok, err := hasIPCLock()
	if err != nil {
		fmt.Println("Error checking CAP_IPC_LOCK:", err)
		return
	}
	if ok {
		fmt.Println("CAP_IPC_LOCK is present.")
	} else {
		fmt.Println("CAP_IPC_LOCK is NOT present.")
	}
}
