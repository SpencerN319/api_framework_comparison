package signals

import (
	"log"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
	"time"
)

// Interrupted waits until a signal interrupts us and then returns the name of the signal.
//
// However, if interrupted by SIGINT (i.e. interactive ctrl-c typed),
// then the user is told "interrupt again to dump call stack. If not done in 2 seconds panic(
// "Interrupted") is called
func Interrupted() string {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	got := <-sig
	signal.Stop(sig)
	if syscall.SIGTERM == got {
		log.Println("Interrupted by SIGTERM")
		return "SIGTERM"
	}

	log.Println("Interrupt again to prevent dumping of all call stacks")
	time.Sleep(2 * time.Second)
	debug.SetTraceback("all")
	panic("Interrupted")
}
