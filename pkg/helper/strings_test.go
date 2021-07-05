package helper

import (
	"fmt"
	"testing"
	"testing/iotest"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRandomString(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		str := NewStrings()
		rndm, err := str.GenerateRandomString(32)
		if err != nil {
			t.Error(err)
			t.Fail()
		}
		t.Log(rndm)
	})

	t.Run("error", func(t *testing.T) {
		str := &Strings{
			reader:  iotest.ErrReader(fmt.Errorf("error")),
			letters: "asasdasfafd",
		}
		result, err := str.GenerateRandomString(10)
		assert.Error(t, err)
		assert.Equal(t, "", result)
	})
}

func BenchmarkGenrateRandomString(b *testing.B) {
	str := NewStrings()
	for i := 0; i < b.N; i++ {
		str.GenerateRandomString(32)
	}
}

type ioMock struct {
	mock.Mock
}

func (x *ioMock) Read(p []byte) (n int, err error) {
	args := x.Called(p)
	return args.Int(0), args.Error(1)
}
