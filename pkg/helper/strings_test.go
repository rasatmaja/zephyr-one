package helper

import "testing"

func TestRandomString(t *testing.T) {
	rndm, err := GenerateRandomString(32)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	t.Log(rndm)
}

func BenchmarkGenrateRandomString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateRandomString(32)
	}
}
