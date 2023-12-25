package mr

//
// RPC definitions.
//
// remember to capitalize all names.
//

import "os"
import "strconv"

//
// example to show how to declare the arguments
// and reply for an RPC.
//

type MapJob struct {
	FileName string
	MapJobID int
	ReduceN  int
}

type ReduceJob struct {
	IntermediateFiles []string
	ReduceJobID       int
}

type RequestJobReply struct {
	Done      bool
	MapJob    *MapJob
	ReduceJob *ReduceJob
}

type ReportMapJobArgs struct {
	Filename          string
	IntermediateFiles []string
}

type ReportReduceJobArgs struct {
	Filename          string
	IntermediateFiles []string
}
type ExampleArgs struct {
	X int
}

type ExampleReply struct {
	Y int
}

// Add your RPC definitions here.

// Cook up a unique-ish UNIX-domain socket name
// in /var/tmp, for the coordinator.
// Can't use the current directory since
// Athena AFS doesn't support UNIX-domain sockets.
func coordinatorSock() string {
	s := "/var/tmp/824-mr-"
	s += strconv.Itoa(os.Getuid())
	return s
}
