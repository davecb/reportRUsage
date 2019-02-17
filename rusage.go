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

// The rusage struct on linux has the following "live" elements
//
//       ru_utime
//              This is the total amount of time spent executing in user mode,
//              expressed in a timeval structure (seconds plus microseconds).
//
//       ru_stime
//              This is the total amount of time spent executing in kernel
//              mode, expressed in a timeval structure (seconds plus microsec‐
//              onds).
//
//       ru_maxrss (since Linux 2.6.32)
//              This is the maximum resident set size used (in kilobytes).
//              For RUSAGE_CHILDREN, this is the resident set size of the
//              largest child, not the maximum resident set size of the
//              process tree.
//
//       ru_minflt
//              The number of page faults serviced without any I/O activity;
//              here I/O activity is avoided by “reclaiming” a page frame from
//              the list of pages awaiting reallocation.
//
//       ru_majflt
//              The number of page faults serviced that required I/O activity.
//
//       ru_inblock (since Linux 2.6.22)
//              The number of times the filesystem had to perform input.
//
//       ru_oublock (since Linux 2.6.22)
//              The number of times the filesystem had to perform output.
//
//       ru_nvcsw (since Linux 2.6)
//              The number of times a context switch resulted due to a process
//              voluntarily giving up the processor before its time slice was
//              completed (usually to await availability of a resource).
//
//       ru_nivcsw (since Linux 2.6)
//              The number of times a context switch resulted due to a higher
//              priority process becoming runnable or because the current
//              process exceeded its time slice.
//
// The following are unmaintained on Linux
//       ru_ixrss (unmaintained)
//       ru_idrss (unmaintained)
//       ru_isrss (unmaintained)
//       ru_nswap (unmaintained)
//       ru_msgsnd (unmaintained)
//       ru_msgrcv (unmaintained)
//       ru_nsignals (unmaintained)
