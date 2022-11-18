// main package
package main

import (
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/godbus/dbus/v5"
)

var sockAddr = "/run/user/" +
	strconv.Itoa(os.Geteuid()) +
	"/transactionalupdatenotification.socket"

func notifySend(input string) {
	success := strings.Split(input, ":")[1]

	// Customize message based on success state
	message := "Updates successfully installed"
	submessage := "System has been upgraded, on " +
		string(time.Now().Format(time.RFC1123)) +
		" please reboot to take effect."
	icon := "appointment-soon"
	if strings.Compare(success, "failure") == 0 {
		message = "Update process failed"
		submessage = "An error was encountered while upgrading on " +
			string(time.Now().Format(time.RFC1123))
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
}

func handleMessage(connection net.Conn) {
	log.Printf("Client connected [%s]", connection.RemoteAddr().Network())

	inputBuffer := make([]byte, 1024)
	data, err := connection.Read(inputBuffer)

	if err != nil {
		panic("Receiving error")
	}

	if strings.Contains(string(inputBuffer[:data]), Message) {
		notifySend(string(inputBuffer[:data]))
	}

	_, err = io.Copy(connection, connection)
	if err != nil {
		panic("Receiving error")
	}

	connection.Close()
}

// NotifyDaemon is the user-facing running daemon that will be sending the graphical
// notifications.
func NotifyDaemon() {
	if err := os.RemoveAll(sockAddr); err != nil {
		log.Fatal(err)
	}

	listener, err := net.Listen("unix", sockAddr)
	if err != nil {
		log.Fatal("listen error:", err)
	}
	defer listener.Close()

	for {
		// Accept new connections, dispatching them to echoServer
		// in a goroutine.
		conn, err := listener.Accept()
		if err != nil {
			log.Println("accept error:", err)
		}

		go handleMessage(conn)
	}
}
