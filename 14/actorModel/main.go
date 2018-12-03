package main

import (
	"fmt"
	"github.com/AsynkronIT/goconsole"
	"github.com/AsynkronIT/protoactor-go/actor"
)

type Hello struct{ Who string }
type HelloActor struct{}

func (state *HelloActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case Hello:
		fmt.Printf("Hello %v\n", msg.Who)
	}
}

func main() {
	props := actor.FromProducer(func() actor.Actor { return &HelloActor{} })
	pid := actor.Spawn(props)
	pid.Tell(Hello{Who: "Atrophy"})
	console.ReadLine()
}
