package models

import "testing"

func TestAccountInfo(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ai := new(AccountInfo)
		_, err := ai.MarshalJSON()
		if err != nil {
			t.Errorf("Erro occur, got: %s", err)
			t.Fail()
		}
	})
}
