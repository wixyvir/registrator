package main

import (
	_ "github.com/wixyvir/registrator/consul"
	_ "github.com/wixyvir/registrator/consulkv"
	_ "github.com/wixyvir/registrator/etcd"
	_ "github.com/wixyvir/registrator/skydns2"
	_ "github.com/wixyvir/registrator/zookeeper"
)
