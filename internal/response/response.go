package response

/** RESPONSE
 *  The porpouse of this pacakage is a provide a JSON structure
 *  for http response
 *
 *  EXAMPLE:
 *  :: SUCCESS
 *  {
 *		"code"    : 200,
 *		"status"  : "success",
 *		"message" : "successfully registered",
 *		"data"	  : {
 *			"email"		: "atmaja@rasio.dev",
 *			"username"	: "rasatmaja"
 *		}
 *	}
 *
 *  :: UNAUTHORIZED
 *  {
 *		"code"    : 401,
 *		"status"  : "unauthorized",
 *		"message" : "Username and password doesnt match",
 *	}
 *
 *  :: BAD REQUEST
 *  {
 *		"code"    		: 400,
 *		"status"  		: "bad_request",
 *		"message" 		: "request validation failed",
 *		"validation"  	: [
 *			{
 *				"field" 	: "email",
 *				"message"	: "email format not valid"
 *			}
 *		]
 *	}
 */

// Response define a JSON structure for http response
type Response struct {
	Code       int             `json:"code,omitempty"`
	Status     string          `json:"status,omitempty"`
	Message    string          `json:"message,omitempty"`
	Data       interface{}     `json:"data,omitempty"`
	Validation []ValidationErr `json:"validation,omitempty"`
}

// ValidationErr define field and status when validation fail
type ValidationErr struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
}

// Factory is a function to init Resaponse struct
func Factory() *Response {
	return &Response{}
}

// SetData is a function to set data
func (res *Response) SetData(dt interface{}) *Response {
	res.Data = dt
	return res
}

// SetValidationErr is a function to set validation error
func (res *Response) SetValidationErr(vld ...ValidationErr) *Response {
	res.Validation = append(res.Validation, vld...)
	return res
}

// Success define success response
func (res *Response) Success(msg ...string) *Response {
	res.Code = 200
	res.Status = "success"
	res.Message = res.Status
	if len(msg) > 0 {
		res.Message = msg[0]
	}
	return res
}

// InternalServerError will build error response with 500 http code
func (res *Response) InternalServerError(msg ...string) *Response {
	res.Code = 500
	res.Status = "internal_server_error"
	res.Message = res.Status
	if len(msg) > 0 {
		res.Message = msg[0]
	}
	return res
}

// Unauthorized will build error response with 401 http code
func (res *Response) Unauthorized(msg ...string) *Response {
	res.Code = 401
	res.Status = "unauthorized"
	res.Message = res.Status
	if len(msg) > 0 {
		res.Message = msg[0]
	}
	return res
}

// BadRequest will build error response with 400 http code
func (res *Response) BadRequest(msg ...string) *Response {
	res.Code = 400
	res.Status = "bad_request"
	res.Message = res.Status
	if len(msg) > 0 {
		res.Message = msg[0]
	}
	return res
}

// NotFound will build error response with 404 http code
func (res *Response) NotFound(msg ...string) *Response {
	res.Code = 404
	res.Status = "not_found"
	res.Message = res.Status
	if len(msg) > 0 {
		res.Message = msg[0]
	}
	return res
}

// Created will build error response with 201 http code
func (res *Response) Created(msg ...string) *Response {
	res.Code = 201
	res.Status = "created"
	res.Message = res.Status
	if len(msg) > 0 {
		res.Message = msg[0]
	}
	return res
}

func (res *Response) Error() string {
	return res.Message
}
