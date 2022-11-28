// main package
package main

import "github.com/godbus/dbus/v5"

// NotifyDaemonClient will search througt all files in /run/user, to find all
// running transactionalupdatenotification socket files, then send a message
// to each one of them in order to trigger the notification for all.
func NotifyDaemonClient(success string) {
	conn, err := dbus.SystemBus()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	obj := conn.Object(
		Iface,
		dbus.ObjectPath(FullPath),
	)
	call := obj.Call(
		Iface+".Notify",
		0,
		success,
	)

	if call.Err != nil {
		panic(call.Err)
	}
}
