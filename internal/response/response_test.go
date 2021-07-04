package response

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestFactory(t *testing.T) {
	Factory()
}

func TestSuccess(t *testing.T) {
	res := Factory()
	res.Success("Test Success")
	res.SetData("Test Data")
	json, _ := json.MarshalIndent(res, "", "\t")
	fmt.Printf("%s \n", json)
}

func TestInternalServerError(t *testing.T) {
	res := Factory()
	res.InternalServerError("Test 500 error")
	json, _ := json.MarshalIndent(res, "", "\t")
	fmt.Printf("%s \n", json)
}

func TestBadRequest(t *testing.T) {
	res := Factory()
	res.BadRequest("Test 400 error")
	res.SetValidationErr(ValidationErr{
		Field:   "email",
		Message: "email cannot be empty",
	})
	json, _ := json.MarshalIndent(res, "", "\t")
	fmt.Printf("%s \n", json)
}

func TestNotFound(t *testing.T) {
	res := Factory()
	res.NotFound("Test 404 error")
}

func TestCreated(t *testing.T) {
	res := Factory()
	res.Created("Test 201 created")
}

func TestError(t *testing.T) {
	res := Factory()
	res.Created("Testing")
	_ = res.Error()
}

func BenchmarkResponse(b *testing.B) {
	b.Run("success", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			res := Factory()
			res.Success("Success")
			res.SetData("Set Data")
		}
	})

	b.Run("internal-server-error", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			res := Factory()
			res.InternalServerError("Error")
		}
	})

	b.Run("bad-request", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			res := Factory()
			res.BadRequest("Test 400 error")
			res.SetValidationErr(ValidationErr{
				Field:   "email",
				Message: "email cannot be empty",
			})
		}
	})
}
