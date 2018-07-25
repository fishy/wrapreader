package wrapreader_test

import (
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/fishy/wrapreader"
)

type DummyReader struct {
	ReadCalled bool
}

func (reader *DummyReader) Read(p []byte) (n int, err error) {
	reader.ReadCalled = true
	return
}

type DummyReadCloser struct {
	ReadCalled  bool
	CloseCalled bool
}

func (closer *DummyReadCloser) Read(p []byte) (n int, err error) {
	closer.ReadCalled = true
	return
}

func (closer *DummyReadCloser) Close() (err error) {
	closer.CloseCalled = true
	return
}

func TestRead(t *testing.T) {
	reader := new(DummyReader)
	closer := new(DummyReadCloser)
	wrap := wrapreader.Wrap(reader, closer)
	wrap.Read(nil)
	if !reader.ReadCalled {
		t.Error("WrapReader.Read should call the underlying reader's Read function")
	}
	if closer.ReadCalled {
		t.Error(
			"WrapReader.Read should not call the underlying closer's Read function",
		)
	}
}

func TestClose(t *testing.T) {
	reader := new(DummyReader)
	closer := new(DummyReadCloser)
	wrap := wrapreader.Wrap(reader, closer)
	wrap.Close()
	if !closer.CloseCalled {
		t.Error(
			"WrapReader.Close should call the underlying closer's Close function",
		)
	}
}

func TestCloseBoth(t *testing.T) {
	reader := new(DummyReadCloser)
	closer := new(DummyReadCloser)
	wrap := wrapreader.Wrap(reader, closer)
	wrap.Close()
	if !closer.CloseCalled {
		t.Error(
			"WrapReader.Close should call the underlying closer's Close function",
		)
	}
	if !reader.CloseCalled {
		t.Error(
			"WrapReader.Close should call the underlying reader's Close function",
		)
	}
}

func ExampleWrap() {
	filename := filepath.Join(os.TempDir(), "_test.gz")
	content := "Hello, world!"
	defer os.Remove(filename)

	f, err := os.Create(filename)
	if err != nil {
		// TODO: handle error properly
		panic(err)
	}
	writer := gzip.NewWriter(f)
	func() {
		defer func() {
			writer.Close()
			f.Close()
		}()
		_, err = writer.Write([]byte(content))
		if err != nil {
			// TODO: handle error properly
			panic(err)
		}
	}()

	f, err = os.Open(filename)
	if err != nil {
		// TODO: handle error properly
		panic(err)
	}
	reader, err := gzip.NewReader(f)
	if err != nil {
		// TODO: handle error properly
		panic(err)
	}
	readCloser := wrapreader.Wrap(reader, f)
	// readCloser.Close() will close both f and reader
	defer readCloser.Close()
	// Read from readCloser is actually read from reader
	read, err := ioutil.ReadAll(readCloser)
	if err != nil {
		// TODO: handle error properly
		panic(err)
	}
	fmt.Println(string(read))
	// Output:
	// Hello, world!
}
