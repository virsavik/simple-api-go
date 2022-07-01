package model

type ResponseEntity struct {
	Header     interface{}
	Body       interface{}
	StatusCode int
}

func NewResponseEntity(header interface{}, body interface{}, statusCode int) ResponseEntity {
	return ResponseEntity{
		Header:     header,
		Body:       body,
		StatusCode: statusCode,
	}
}
