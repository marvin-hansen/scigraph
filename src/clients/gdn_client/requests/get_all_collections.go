package requests

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForGetAllCollections(fabric string) *RequestForGetAllCollections {
	return &RequestForGetAllCollections{
		path: fmt.Sprintf("_fabric/%v/_api/collection", fabric),
	}
}

type RequestForGetAllCollections struct {
	path string
}

func (req *RequestForGetAllCollections) Path() string {
	return req.path
}

func (req *RequestForGetAllCollections) Method() string {
	return http.MethodGet
}

func (req *RequestForGetAllCollections) Query() string {
	return ""
}

func (req *RequestForGetAllCollections) HasQueryParameter() bool {
	return true
}
func (req *RequestForGetAllCollections) GetQueryParameter() string {
	return "?excludeSystem=true"
}

func (req *RequestForGetAllCollections) Payload() []byte {
	return nil
}

func (req *RequestForGetAllCollections) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForGetAllCollections() *ResponseForGetAllCollections {
	return new(ResponseForGetAllCollections)
}

func (r ResponseForGetAllCollections) String() string {
	return fmt.Sprintf("Result: %v", r.Result)
}

type ResponseForGetAllCollections struct {
	Result     []ResultGetAllCollections `json:"result"`
	RawMessage []byte
}

func (r *ResponseForGetAllCollections) GetRawMessage() []byte {
	return r.RawMessage
}

func (r *ResponseForGetAllCollections) SetRawMessage(raw []byte) {
	r.RawMessage = raw
}

type ResultGetAllCollections struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	Status           int    `json:"status"`
	Type             int    `json:"type"`
	CollectionModel  string `json:"collectionModel"`
	IsSpot           bool   `json:"isSpot"`
	IsLocal          bool   `json:"isLocal"`
	HasStream        bool   `json:"hasStream"`
	WaitForSync      bool   `json:"waitForSync"`
	IsSystem         bool   `json:"isSystem"`
	GloballyUniqueId string `json:"globallyUniqueId"`
	SearchEnabled    bool   `json:"searchEnabled"`
}

func (r ResultGetAllCollections) String() string {
	return fmt.Sprintf("ID: %v, Name: %v,  Status: %v, Type: %v, CollectionModel: %v, IsSpot: %v, IsLocal: %v, HasStream: %v, WaitForSync: %v, IsSystem: %v, GloballyUniqueId: %v, SearchEnabled: %v",
		r.Id,
		r.Name,
		r.Status,
		r.Type,
		r.CollectionModel,
		r.IsSpot,
		r.IsLocal,
		r.HasStream,
		r.WaitForSync,
		r.IsSystem,
		r.GloballyUniqueId,
		r.SearchEnabled,
	)
}
