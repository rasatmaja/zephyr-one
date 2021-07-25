package models

import "testing"

func TestContact(t *testing.T) {
	t.Run("MarshalJSON", func(t *testing.T) {
		contact := &Contact{
			IsPrimary:     true,
			Contact:       "test@test.test",
			ContactTypeID: "1",
		}
		contact.MarshalJSON()
	})

	t.Run("parseContactTypeID", func(t *testing.T) {
		c1 := &Contact{
			IsPrimary:     true,
			Contact:       "test@test.test",
			ContactTypeID: "1",
		}
		c1.ParseContactTypeID()

		c1.ContactTypeID = "2"
		c1.ParseContactTypeID()

		c1.ContactTypeID = "0"
		c1.ParseContactTypeID()
	})

	t.Run("parseContactTypeID", func(t *testing.T) {
		c1 := &Contact{
			IsPrimary: true,
			Contact:   "test@test.test",
		}
		c1.ParseContactType("email")
		c1.ParseContactType("phone")
		c1.ParseContactType("unknown")
	})
}
