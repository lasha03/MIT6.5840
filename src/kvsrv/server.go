package kvsrv

import (
	"log"
	"sync"
	"fmt"
)

const Debug = false

func DPrintf(format string, a ...interface{}) (n int, err error) {
	if Debug {
		log.Printf(format, a...)
	}
	return
}

type MapInput struct{
	value string
	id int
}

func (m MapInput) String() string {
	return fmt.Sprintf("MapInput{id: %d, value: %s}", m.id, m.value)
}

type KVServer struct {
	mu sync.Mutex
	data map[string]string
	replies map[int64]MapInput
	// Your definitions here.
}


func (kv *KVServer) Get(args *GetArgs, reply *GetReply) {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	// fmt.Println("get", args)

	// if oldReply, ok := kv.replies[args.ClientId]; ok && oldReply.id == args.RequestId {
		// reply.Value = oldReply.value
		// fmt.Println("old")
		// fmt.Println(kv.data)
	// 	return
	// }

	requestReply := kv.data[args.Key]
	// kv.replies[args.ClientId] = MapInput{
	// 	value:	requestReply,
	// 	id:		args.RequestId,
	// }

	reply.Value = requestReply
	// fmt.Println(kv.data)
	// fmt.Println(reply.Value)
}

func (kv *KVServer) Put(args *PutAppendArgs, reply *PutAppendReply) {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	// fmt.Println("put", args)

	if oldReply, ok := kv.replies[args.ClientId]; ok && oldReply.id == args.RequestId {
		reply.Value = ""
		// fmt.Println("old")
		// fmt.Println(kv.data)
		// fmt.Println(reply.Value)
		return
	}

	requestReply := ""
	kv.replies[args.ClientId] = MapInput{
		value:	requestReply,
		id:		args.RequestId,
	}

	// reply.Value = ""
	kv.data[args.Key] = args.Value
	// fmt.Println(kv.data)
	// fmt.Println(reply.Value)
}

func (kv *KVServer) Append(args *PutAppendArgs, reply *PutAppendReply) {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	// fmt.Println("append", args)

	if oldReply, ok := kv.replies[args.ClientId]; ok && oldReply.id == args.RequestId {
		reply.Value = oldReply.value
		// fmt.Println("old")
		// fmt.Println(kv.data)
		// fmt.Println(reply.Value)
		return
	}
	requestReply := kv.data[args.Key]
	kv.replies[args.ClientId] = MapInput{
		value:	requestReply,
		id:		args.RequestId,
	}

	reply.Value = requestReply
	kv.data[args.Key] += args.Value
	// fmt.Println(kv.data)
	// fmt.Println(reply.Value)
}

func StartKVServer() *KVServer {
	kv := &KVServer{
		data:		make(map[string]string),
		replies: 	make(map[int64]MapInput),
	}
	// You may need initialization code here.

	return kv
}
