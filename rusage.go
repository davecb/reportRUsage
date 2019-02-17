package reportRUsage

import (
	"fmt"
	"log"
	"os"
	"syscall"
	"time"
)


// reportRusage reports cpu-seconds, memory and IOPS used, from /proc/stats
func reportRUsage(name string, start time.Time) {
	var r syscall.Rusage

	err := syscall.Getrusage(syscall.RUSAGE_SELF, &r)
	if err != nil {
		log.Fatal(err)
		log.Printf("%s %s %d no resource usage available\n",
			start.Format("2006-01-02 15:04:05.000"), name, os.Getpid())
		return
	}
	fmt.Fprint(os.Stderr, "#date      time         name        pid  utime stime maxrss inblock outblock\n")
	fmt.Fprintf(os.Stderr, "%s %s %d %f %f %d %d %d\n", start.Format("2006-01-02 15:04:05.000"),
		name, os.Getpid(), seconds(r.Utime), seconds(r.Stime), r.Maxrss*1024, r.Inblock, r.Oublock)
}

// seconds converts a syscall.Timeval to seconds
func seconds(t syscall.Timeval) float64 {
	return float64(time.Duration(t.Sec)*time.Second+time.Duration(t.Usec)*time.Microsecond) / float64(time.Second)
}

