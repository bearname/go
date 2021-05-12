package main

type subject interface {
	register(observer Observer)
	deregister(observer Observer)
	notifyAll()
}
