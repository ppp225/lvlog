package main

import log "github.com/ppp225/lvlog"

func main() {
	log.Info("info msg")              // == std.Println(" [INFO]", "info msg")
	log.Infof("msg=%q", "info msg")   // == std.Printf("%s msg=%q", " [INFO]", "info msg")
	log.Debug("debug msg")            // won't be displayed
	log.SetLevel(log.ALL)             // default level is INFO and up
	log.Debug("debug msg")            // will be displayed now
	log.Fatalf("msg=%q", "fatal msg") // will be displayed and exit with status code 1
}
