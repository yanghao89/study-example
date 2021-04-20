package example

import (
	"sync"
)

type ChannelExample struct {
	Channels chan int
}

var (
	mu sync.Mutex
)

func NewExample() *ChannelExample  {
	return &ChannelExample{}
}
var (
	wg sync.WaitGroup
)

func (c *ChannelExample) MakeChan(n int)[]int  {
	c.Channels = make(chan int,n)
	for i:=n;i > 0;i--{
		go func(channel chan int,i int) {
			channel <- i
		}(c.Channels,i)
	}
	var (
		ints []int
	)
	for v := range c.Channels{
			ints = append(ints,v)
			if len(ints) == n{
				break
			}
	}
	close(c.Channels)
	return ints
}

