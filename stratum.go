package Stratum

import (
	"encoding/json"
	"errors"
)

// Stratum has three types of messages: notification, request, and response.

// Notification for methods that do not require a response.
type Notification struct {
	Method string        `json:"method"`
	Params []interface{} `json:"params"`
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

// request is for methods that require a response.
type Request struct {
	MessageID MessageID     `json:"id"`
	Method    string        `json:"method"`
	Params    []interface{} `json:"params"`
}

func NewRequest(id MessageID, m Method, params []interface{}) Request {
	n, _ := EncodeMethod(m)
	return Request{
		MessageID: id,
		Method:    n,
		Params:    params,
	}
}

func (n *Request) GetMethod() Method {
	m, _ := DecodeMethod(n.Method)
	return m
}

// Response is what is sent back in response to requests.
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

func NewErrorResponse(id MessageID, e Error) Response {
	return Response{
		MessageID: id,
		Result:    nil,
		Error:     &e,
	}
}

func (r *Request) MarshallJSON() ([]byte, error) {
	if !ValidMessageID(r.MessageID) {
		return []byte{}, errors.New("invalid id")
	}

	if r.Method == "" {
		return []byte{}, errors.New("invalid method")
	}

	return json.Marshal(r)
}

func (r *Request) UnmarshallJSON(j []byte) error {
	err := json.Unmarshal(j, r)
	if err != nil {
		return err
	}

	if !ValidMessageID(r.MessageID) {
		return errors.New("invalid id")
	}

	if r.GetMethod() == Unset {
		return errors.New("invalid method")
	}

	return nil
}

func (r *Response) MarshallJSON() ([]byte, error) {
	if !ValidMessageID(r.MessageID) {
		return []byte{}, errors.New("invalid id")
	}

	return json.Marshal(r)
}

func (r *Response) UnmarshallJSON(j []byte) error {
	err := json.Unmarshal(j, r)
	if err != nil {
		return err
	}

	if !ValidMessageID(r.MessageID) {
		return errors.New("invalid id")
	}

	return nil
}

func (r *Notification) MarshallJSON() ([]byte, error) {
	if r.Method == "" {
		return []byte{}, errors.New("invalid method")
	}

	return json.Marshal(r)
}

func (r *Notification) UnmarshallJSON(j []byte) error {
	err := json.Unmarshal(j, r)
	if err != nil {
		return err
	}

	if r.GetMethod() == Unset {
		return errors.New("invalid method")
	}

	return nil
}
