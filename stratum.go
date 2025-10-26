package stratum

import (
	"encoding/json"
	"errors"
)

// MessageID is a unique numerical identifier that is different for each notification
// and request / response.
type MessageID uint64

// Stratum has three types of messages: notification, request, and response.
// notification: unprompted, server to client
// request: client to server
// response: server to client

// Notification is for methods that do not require a response.
// Automatically includes a newline when marshalling.
type Notification struct {
	MessageID MessageID     `json:"id"`
	Method    string        `json:"method"`
	Params    []interface{} `json:"params"`
}

func NewNotification(m Method, params []interface{}) Notification {
	n, _ := EncodeMethod(m)
	return Notification{
		Method: n,
		Params: params,
	}
}

func (n *Notification) GetMethod() Method {
	m, _ := DecodeMethod(n.Method)
	return m
}

// Request is for methods that require a [Response].
// Automatically includes a newline when marshalling.
type Request struct {
	MessageID MessageID     `json:"id"`
	Method    string        `json:"method"`
	Params    []interface{} `json:"params"`
}

func NewRequest(id MessageID, m Method, params []interface{}) Request {
	n, _ := EncodeMethod(m)
	return Request{
		/// FIXME/MAYBE: cast to uint64?
		MessageID: id,
		Method:    n,
		Params:    params,
	}
}

func (n *Request) GetMethod() Method {
	m, _ := DecodeMethod(n.Method)
	return m
}

// Response is what is sent back in response to [Request]s.
// Automatically includes a newline when marshalling.
type Response struct {
	MessageID MessageID   `json:"id"`
	Result    interface{} `json:"result"`
	Error     *Error      `json:"error"`
}

func NewResponse(id MessageID, r interface{}) Response {
	return Response{
		MessageID: id,
		Result:    r,
		Error:     nil,
	}
}

type BooleanResult struct {
	Result bool
}

func (b *BooleanResult) Read(r *Response) error {
	var ok bool
	b.Result, ok = r.Result.(bool)
	if !ok {
		return errors.New("invalid value")
	}

	return nil
}

func NewBooleanResponse(id MessageID, x bool) Response {
	return NewResponse(id, x)
}
func NewErrorResponse(id MessageID, e Error) Response {
	return Response{
		MessageID: id,
		Result:    nil,
		Error:     &e,
	}
}

func (r *Request) Respond(d interface{}) Response {
	return NewResponse(r.MessageID, d)
}
func (r *Request) Marshal() ([]byte, error) {
	if r.Method == "" {
		return []byte{}, errors.New("invalid method")
	}
	marshalled, err := json.Marshal(r)
	if err != nil {
		return []byte{}, err
	}
	return append(marshalled, '\n'), nil
}
func (r *Request) Unmarshal(j []byte) error {
	err := json.Unmarshal(j, r)
	if err != nil {
		return err
	}

	if r.GetMethod() == Unset {
		return errors.New("invalid method")
	}

	return nil
}

func (r *Response) Marshal() ([]byte, error) {
	marshalled, err := json.Marshal(r)
	if err != nil {
		return []byte{}, err
	}
	return append(marshalled, '\n'), nil
}
func (r *Response) Unmarshal(j []byte) error {
	err := json.Unmarshal(j, r)
	if err != nil {
		return err
	}

	return nil
}

func (r *Notification) Marshal() ([]byte, error) {
	if r.Method == "" {
		return []byte{}, errors.New("invalid method")
	}

	marshalled, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	return append(marshalled, '\n'), nil
}
func (r *Notification) Unmarshal(j []byte) error {
	err := json.Unmarshal(j, r)
	if err != nil {
		return err
	}

	if r.GetMethod() == Unset {
		return errors.New("invalid method")
	}

	return nil
}
