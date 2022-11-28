// main package
package main

import (
	"log"
	"time"

	"github.com/godbus/dbus/v5"
	"github.com/godbus/dbus/v5/introspect"
)

type notify string

func (f notify) Notify(input string) (string, *dbus.Error) {

	log.Printf("Update finished: %s", input)
	// Customize message based on success state
	message := "Updates successfully installed"
	submessage := "System has been upgraded, on " +
		time.Now().Format(time.RFC1123) +
		" please reboot to take effect."
	icon := "appointment-soon"

	if input == "failure" {
		message = "Update process failed"
		submessage = "An error was encountered while upgrading on " +
			time.Now().Format(time.RFC1123)
		icon = "appointment-missed"
	}

	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	obj := conn.Object(
		"org.freedesktop.Notifications",
		"/org/freedesktop/Notifications",
	)
	call := obj.Call(
		"org.freedesktop.Notifications.Notify",
		0,
		"",
		uint32(0),
		icon,
		message,
		submessage,
		[]string{},
		map[string]dbus.Variant{},
		int32(5000),
	)

	if call.Err != nil {
		panic(call.Err)
	}

	return string(f), nil
}

// NotifyDaemon is the user-facing running daemon that will be sending the graphical
// notifications.
func NotifyDaemon() {
	conn, err := dbus.SystemBus()

	// couldnt connect to session bus
	if err != nil {
		panic(err)
	}

	reply, err := conn.RequestName(Iface, dbus.NameFlagDoNotQueue)
	if err != nil {
		panic(err)
	}

	if reply != dbus.RequestNameReplyPrimaryOwner {
		panic("Name already taken")
	}

	m := notify("Ok!")

	err = conn.Export(m, dbus.ObjectPath(FullPath), Iface)
	if err != nil {
		panic(err)
	}

	n := &introspect.Node{
		Interfaces: []introspect.Interface{
			{
				Name:    Iface,
				Methods: introspect.Methods(m),
				Signals: []introspect.Signal{},
			},
		},
	}

	root := &introspect.Node{
		Children: []introspect.Node{
			{
				Name: "org/test/tu",
			},
		},
	}

	err = conn.Export(introspect.NewIntrospectable(n),
		dbus.ObjectPath(FullPath),
		"org.freedesktop.DBus.Introspectable")
	if err != nil {
		panic(err)
	}

	err = conn.Export(introspect.NewIntrospectable(root),
		"/",
		"org.freedesktop.DBus.Introspectable")
	if err != nil {
		panic(err)
	}

	log.Printf("Bridge is Running.")

	select {}
}
