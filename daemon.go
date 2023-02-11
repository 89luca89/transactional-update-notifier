// main package
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/godbus/dbus/v5"
)

func notify(input string) {
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

	// Set hints of the notification, in this case we set it to
	// an urgent notification, so we're sure that it will stick
	// in the tray, and the user is notified.
	hints := map[string]dbus.Variant{
		"urgency":  dbus.MakeVariant(byte(1)),
		"category": dbus.MakeVariant("device"),
	}

	call := obj.Call(
		"org.freedesktop.Notifications.Notify",
		0,
		"",
		uint32(0),
		icon,
		message,
		submessage,
		[]string{},
		hints,
		int32(-1),
	)

	if call.Err != nil {
		panic(call.Err)
	}
}

// NotifyDaemon will wait for a message on org.opensuse.tukit.Updated and trigger
// a graphical notification accordingly.
func NotifyDaemon() {
	conn, err := dbus.SystemBus()
	if err != nil {
		panic(err)
	}

	if err = conn.AddMatchSignal(
		dbus.WithMatchSender(Iface),
		dbus.WithMatchObjectPath(dbus.ObjectPath(FullPath)),
		dbus.WithMatchInterface(Iface),
		dbus.WithMatchMember(Member),
	); err != nil {
		panic(err)
	}

	c := make(chan *dbus.Signal, 10)
	conn.Signal(c)

	for v := range c {
		body := fmt.Sprintf("%s", v.Body...)
		notify(body)
	}
}
