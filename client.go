// main package
package main

import (
	"github.com/godbus/dbus/v5"
)

// NotifyDaemonClient will emit a message on org.opensuse.tukit.Updated
// so each one of the user-facing service will trigger the graphical notification.
func NotifyDaemonClient(success string) {
	conn, err := dbus.SystemBus()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	reply, err := conn.RequestName(Iface,
		dbus.NameFlagDoNotQueue)
	if err != nil {
		panic(err)
	}

	if reply != dbus.RequestNameReplyPrimaryOwner {
		panic("name already taken")
	}

	err = conn.Emit(dbus.ObjectPath(FullPath), Iface+"."+Member, success)
	if err != nil {
		panic(err)
	}
}
