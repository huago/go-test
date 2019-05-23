package test

import "testing"

func TestNoBufferChan(t *testing.T) {
	NoBufferChan()
}

func TestPipe(t *testing.T) {
	Pipe()
}

func TestBufferedChan(t *testing.T) {
	BufferedChan()
}

func TestOneDirection(t *testing.T) {
	OneDirection()
}

func TestSelect(t *testing.T) {
	Select()
}

func TestDoOrder(t *testing.T) {
	DoOrder()
}

func TestCloseChan(t *testing.T) {
	CloseChan()
}