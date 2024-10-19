package kvsrv

import "fmt"

// Put or Append
type PutAppendArgs struct {
	Key       string
	Value     string
	RequestId int
	ClientId  int64
}

// String returns a string representation of PutAppendArgs.
func (p PutAppendArgs) String() string {
	return fmt.Sprintf("PutAppendArgs{Key: %s, Value: %s, RequestId: %d, ClientId: %d}", p.Key, p.Value, p.RequestId, p.ClientId)
}

type PutAppendReply struct {
	Value string
}

// String returns a string representation of PutAppendReply.
func (p PutAppendReply) String() string {
	return fmt.Sprintf("PutAppendReply{Value: %s}", p.Value)
}

type GetArgs struct {
	Key       string
	RequestId int
	ClientId  int64
}

// String returns a string representation of GetArgs.
func (g GetArgs) String() string {
	return fmt.Sprintf("GetArgs{Key: %s, RequestId: %d, ClientId: %d}", g.Key, g.RequestId, g.ClientId)
}

type GetReply struct {
	Value string
}

// String returns a string representation of GetReply.
func (g GetReply) String() string {
	return fmt.Sprintf("GetReply{Value: %s}", g.Value)
}
