// Package stratum provides methods to assist with implementing the Stratum v1 protocol.
package stratum

import (
	"errors"

	"github.com/bytedance/sonic"
)

// MessageID is a unique numerical identifier that is different for each notification and request / response.
// TODO: do we even need this as a "type"?
type MessageID uint64

// Base interface behind [Request], [Response], and [Notification].
type Message interface {
	GetMethod() Method
	Marshal() ([]byte, error)
	Unmarshal(b []byte) error
}

// Stratum has three types of messages: notification, request, and response.
// notification: unprompted, server to client
// request: client to server OR server to client
// response: server to client

// Notification is for methods that do not require a [Response].
// Automatically appends a newline when marshalling.
// Implements [Message].
type Notification struct {
	Method string        `json:"method"`
	Params []interface{} `json:"params"`
}

// NewNotification creates a new [Notification] with the given method and params.
func NewNotification(m Method, params []interface{}) *Notification {
	n, _ := EncodeMethod(m)
	return &Notification{
		Method: n,
		Params: params,
	}
}

// Request is for methods that require a [Response].
// Automatically appends a newline when marshalling.
// Implements [Message].
type Request struct {
	MessageID MessageID     `json:"id"`
	Method    string        `json:"method"`
	Params    []interface{} `json:"params"`
}

// internal helper, exposed for advanced use
//
// you probably want the methods ToRequest/ToNotification/ToResponse functions
// NewRequest creates a new [Request] with the given id, method, and params.
func NewRequest(id MessageID, method Method, params []interface{}) *Request {
	n, _ := EncodeMethod(method)
	return &Request{
		/// FIXME/MAYBE: cast to uint64?
		MessageID: id,
		Method:    n,
		Params:    params,
	}
}

// Response is what is sent back in response to a [Request].
// Automatically appends a newline when marshalling.
// Implements [Message].
type Response struct {
	MessageID MessageID   `json:"id"`
	Result    interface{} `json:"result"`
	Error     *Error      `json:"error,omitempty"`
}

// NewResponse creates a new [Response] with the given id and result.
func NewResponse(id MessageID, r interface{}) *Response {
	return &Response{
		MessageID: id,
		Result:    r,
	}
}

type BooleanResult struct {
	Result bool
}

// NewBooleanResponse creates a new [Response] with a boolean result.
func NewBooleanResponse(id MessageID, x bool) *Response {
	return NewResponse(id, x)
}

// NewErrorResponse creates a new [Response] with an error.
func NewErrorResponse(id MessageID, e Error) *Response {
	return &Response{
		MessageID: id,
		Result:    false,
		Error:     &e,
	}
}

// Respond creates a [Response] to the [Request] with the given result.
func (r *Request) Respond(d interface{}) *Response {
	return NewResponse(r.MessageID, d)
}

// RespondError creates an error [Response] to the [Request].
func (r *Request) RespondError(e Error) *Response {
	return NewErrorResponse(r.MessageID, e)
}

// GetMethod returns the [Method] of the [Request].
func (req *Request) GetMethod() Method {
	return DecodeMethod(req.Method)
}

// Marshal returns a JSON-encoded [Request] with a trailing newline.
func (r *Request) Marshal() ([]byte, error) {
	if r.Method == "" {
		return []byte{}, errors.New("invalid method")
	}
	marshalled, err := sonic.Marshal(r)
	if err != nil {
		return []byte{}, err
	}
	return append(marshalled, '\n'), nil
}

// Unmarshal parses JSON into a [Request].
func (r *Request) Unmarshal(j []byte) error {
	err := sonic.Unmarshal(j, r)
	if err != nil {
		return err
	}

	if r.GetMethod() == MethodUnknown {
		return errors.New("invalid method")
	}

	return nil
}

// GetMethod returns the [Method] of the [Response].
func (res *Response) GetMethod() Method {
	return MethodUnknown
}

// Marshal returns the JSON-encoded [Response] with a trailing newline.
func (r *Response) Marshal() ([]byte, error) {
	marshalled, err := sonic.Marshal(r)
	if err != nil {
		return []byte{}, err
	}
	return append(marshalled, '\n'), nil
}

// Unmarshal parses JSON into the [Response].
func (r *Response) Unmarshal(j []byte) error {
	err := sonic.Unmarshal(j, r)
	if err != nil {
		return err
	}

	return nil
}

// GetMethod returns the [Method] of the [Notification].
func (n *Notification) GetMethod() Method {
	return DecodeMethod(n.Method)
}

// Marshal returns the JSON-encoded [Notification] with a trailing newline.
func (r *Notification) Marshal() ([]byte, error) {
	if r.Method == "" {
		return []byte{}, errors.New("invalid method")
	}

	marshalled, err := sonic.Marshal(r)
	if err != nil {
		return nil, err
	}
	return append(marshalled, '\n'), nil
}

// Unmarshal parses JSON into the [Notification].
func (r *Notification) Unmarshal(j []byte) error {
	err := sonic.Unmarshal(j, r)
	if err != nil {
		return err
	}

	if r.GetMethod() == MethodUnknown {
		return errors.New("invalid method")
	}

	return nil
}

// FromResponse parses the [BooleanResult] from a [Response].
func (b *BooleanResult) FromResponse(r *Response) error {
	var ok bool
	b.Result, ok = r.Result.(bool)
	if !ok {
		return errors.New("invalid value")
	}

	return nil
}

// ToResponse creates a [Response] from the [BooleanResult].
func (b *BooleanResult) ToResponse(id MessageID) *Response {
	return NewResponse(id, b.Result)
}
