package can

import "syscall"

// allow compilation
func NewSockaddr(proto uint16, Ifindex int) syscall.Sockaddr {
	return nil
}

