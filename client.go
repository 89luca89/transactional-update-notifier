// main package
package main

import (
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// NotifyDaemonClient will search througt all files in /run/user, to find all
// running transactionalupdatenotification socket files, then send a message
// to each one of them in order to trigger the notification for all.
func NotifyDaemonClient(success string) {
	transactionalUpdateSockets := []string{}

	err := filepath.Walk("/run/user", func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, "transactionalupdatenotification.socket") {
			transactionalUpdateSockets = append(transactionalUpdateSockets, path)
		}

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, socketFile := range transactionalUpdateSockets {
		connection, err := net.Dial("unix", socketFile)
		if err != nil {
			log.Fatal(err)
		}
		defer connection.Close()

		_, err = connection.Write([]byte(Message + ":" + success))
		if err != nil {
			log.Println("write error:", err)
		}

		time.Sleep(100 * time.Millisecond)
	}
}
