package main

import (
	"dict/dict"
	"net"
	"testing"
	"time"
)

/*
type FakeCon interface {
	Write(b []byte) (n int, e error)
	Close() error
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
}
*/

type SpyCon struct {
	Calls       int
	Closed      bool
	ReadCalled  bool
	WriteCalled bool
}

type TestAddr struct{}

func (t *TestAddr) Network() string {
	return "test"
}

func (t *TestAddr) String() string {
	return "test"
}

func (s *SpyCon) Read(b []byte) (int, error) {
	s.ReadCalled = true
	return 1, nil
}

func (s *SpyCon) Write(b []byte) (int, error) {
	s.WriteCalled = true
	return 1, nil
}

func (s *SpyCon) Close() error {
	s.Closed = true
	return nil
}

func (s *SpyCon) LocalAddr() net.Addr {
	a := &TestAddr{}
	return a
}

func (s *SpyCon) RemoteAddr() net.Addr {
	a := &TestAddr{}
	return a
}

func (s *SpyCon) SetDeadline(t time.Time) error {
	return nil
}

func (s *SpyCon) SetReadDeadline(t time.Time) error {
	return nil
}

func (s *SpyCon) SetWriteDeadline(t time.Time) error {
	return nil
}

func TestConWriter(t *testing.T) {
	c := &SpyCon{}
	connWriter(c, "message")

	if c.WriteCalled != true {
		t.Errorf("WriteCalled must be true")
	}
}

func TestSet(t *testing.T) {
	d := dict.Dictionary{}
	s := &SpyCon{}

	t.Run("Value stored successfully", func(t *testing.T) {
		cmds := []string{"set", "k", "0", "0", "1"}
		res := handleSet(s, cmds, d)

		if res != "STORED\r\n" {
			t.Errorf("Value not stored, expected to be stored")
		}
	})

	t.Run("Value not stored", func(t *testing.T) {
		cmds := []string{"set", "k", "0", "0", "1"}
		handleSet(s, cmds, d)
		res := handleSet(s, cmds, d)

		if res != "NOT_STORED\r\n" {
			t.Errorf("Expected value to be not stored!")
		}
	})
}

func TestGet(t *testing.T) {

	t.Run("Test handleGet", func(t *testing.T) {
		d := dict.Dictionary{}
		wantedVal := "v"
		d.Add("k", wantedVal)
		cmds := []string{"get", "k"}
		_, val := handleGet(cmds, d)

		if val != wantedVal {
			t.Errorf("Value must match with %s", wantedVal)
		}
	})
}

func TestDelete(t *testing.T) {
	t.Run("Test handleDelete (should delete)", func(t *testing.T) {
		d := dict.Dictionary{}
		d.Add("k", "v")
		cmds := []string{"delete", "k"}
		wanted := "DELETED\r\n"

		info := handleDelete(cmds, d)

		if info != wanted {
			t.Errorf("Expected %s, but got %s", wanted, info)
		}
	})

	t.Run("Test handleDelete (not found)", func(t *testing.T) {
		d := dict.Dictionary{}
		cmds := []string{"delete", "k"}
		wanted := "NOT_FOUND\r\n"

		info := handleDelete(cmds, d)

		if info != wanted {
			t.Errorf("Expected %s, but got %s", wanted, info)
		}
	})
}

///type FuncType func(string)

/*
func TestConnectionHandler(t *testing.T) {
	i := 0
//	ogConnWriter := connWriter

	t.Run("ttt", func(t *testing.T) {

		//connWriter = callFuncc(func(string) {}, "ddd")

		newfunc := func() func(net.Conn, string) {
			return func(net.Conn, string) {}
		}

		connWriter = newfunc()
		c := &SpyCon{}
		d := dict.Dictionary{}

		handleConnection(c, 1, d)

	})
	connWriter = ogConnWriter
}

func callFuncc(f func(string), s string) {
	f(s)
}
*/
