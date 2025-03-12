package dto

type Example struct {
	ID   int64  `json:"id"`
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
	Meta string `json:"meta,omitempty"`
}

type GetExampleRequest struct {
	ID int64 `fastmicro:"id,path"`
}

type GetExampleResponse struct {
	Payload *Example `json:"payload,omitempty"`
}

type ReadExampleRequest struct {
}

type ReadExampleResponse struct {
	Payload []*Example `json:"payload,omitempty"`
}

type CreateExampleRequest struct {
	Payload *Example `json:"payload,omitempty"`
}

type CreateExampleResponse struct {
	Payload *Example `json:"payload,omitempty"`
}

type UpdateExampleRequest struct {
	ID      int64    `fastmicro:"id,path"`
	Payload *Example `json:"payload,omitempty"`
}

type UpdateExampleResponse struct {
}

type DeleteExampleRequest struct {
	ID int64 `fastmicro:"id,path"`
}

type DeleteExampleResponse struct {
}
