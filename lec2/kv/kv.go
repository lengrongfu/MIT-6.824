package main

import (
	"fmt"
	"net"
	"net/rpc"
	"sync"
)

var (
	OK     string = "ok"
	NotKey string = "not find key"
)

type KV struct {
	data sync.Map
}

type GetArgs struct {
	Key string
}

type GetReply struct {
	Reply string
	Err   string
}

type PutArgs struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type PutReply struct {
	Err string `json:"err"`
}

func server() {
	kv := new(KV)
	rpcServer := rpc.NewServer()
	rpcServer.Register(kv)
	listen, e := net.Listen("tcp", ":1234")
	if e != nil {
		panic(e)
	}
	go func() {
		for {
			c, e := listen.Accept()
			if e != nil {
				break
			}
			go rpcServer.ServeConn(c)
		}
	}()
}

func (kv *KV) Get(args *GetArgs, reply *GetReply) error {
	if v, ok := kv.data.Load(args.Key); !ok {
		reply.Err = fmt.Sprintf("%s : %s", NotKey, args.Key)
	} else {
		reply.Reply = v.(string)
		reply.Err = OK
	}
	return nil
}

func (kv *KV) Put(args *PutArgs, reply *PutReply) error {
	kv.data.Store(args.Key, args.Value)
	reply.Err = OK
	return nil
}

func connect() *rpc.Client {
	client, e := rpc.Dial("tcp", ":1234")
	if e != nil {
		panic(e)
	}
	return client
}

func get(key string) (value string) {
	client := connect()
	defer client.Close()
	args := GetArgs{
		Key: key,
	}
	reply := GetReply{}
	call := client.Call("KV.Get", &args, &reply)
	if call != nil {
		fmt.Println(call.Error())
	}
	if reply.Err == OK {
		value = reply.Reply
	} else {
		value = ""
		fmt.Printf("call KV.Get error %s\n", reply.Err)
	}
	return
}

func put(key, value string) {
	client := connect()
	defer client.Close()
	args := PutArgs{
		Key:   key,
		Value: value,
	}
	reply := PutReply{}
	client.Call("KV.Put", &args, &reply)
	if reply.Err != OK {
		fmt.Printf("call KV.Put error %s\n", reply.Err)
	}
}

func init() {
	server()
}

func main() {
	get("lrf")
	put("lrf", "lrf")
	value := get("lrf")
	fmt.Println("value ", value)
}
