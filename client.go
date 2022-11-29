// main package
package main

import (
	"github.com/godbus/dbus/v5"
)

// NotifyDaemonClient will search througt all files in /run/user, to find all
// running transactionalupdatenotification socket files, then send a message
// to each one of them in order to trigger the notification for all.
func NotifyDaemonClient(success string) {
	conn, err := dbus.SystemBus()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	err = conn.Emit(dbus.ObjectPath(FullPath), Iface+".Notify", success)
	if err != nil {
		panic(err)
	}
}
