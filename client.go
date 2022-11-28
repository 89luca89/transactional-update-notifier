// main package
package main

import "github.com/godbus/dbus/v5"

// NotifyDaemonClient will search througt all files in /run/user, to find all
// running transactionalupdatenotification socket files, then send a message
// to each one of them in order to trigger the notification for all.
func NotifyDaemonClient(success string) {
	// conn, err := dbus.ConnectSessionBus()
	conn, err := dbus.SystemBus()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	obj := conn.Object(
		"org.test.tu",
		"/org/test/tu",
	)
	call := obj.Call(
		"org.test.tu.Notify",
		0,
		"notify:" + success,
	)

	if call.Err != nil {
		panic(call.Err)
	}
}
