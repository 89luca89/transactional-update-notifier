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
