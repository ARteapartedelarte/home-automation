// Code generated by jrpc. DO NOT EDIT.

package deviceregistrydef

import (
	rpc "github.com/jakewright/home-automation/libraries/go/rpc"
)

// Do performs the request
func (m *GetDeviceRequest) Do() (*GetDeviceResponse, error) {
	req := &rpc.Request{
		Method: "GET",
		URL:    "service.device-registry/device",
		Body:   m,
	}

	rsp := &GetDeviceResponse{}
	_, err := rpc.Do(req, rsp)
	return rsp, err
}

// Do performs the request
func (m *ListDevicesRequest) Do() (*ListDevicesResponse, error) {
	req := &rpc.Request{
		Method: "GET",
		URL:    "service.device-registry/devices",
		Body:   m,
	}

	rsp := &ListDevicesResponse{}
	_, err := rpc.Do(req, rsp)
	return rsp, err
}

// Do performs the request
func (m *GetRoomRequest) Do() (*GetRoomResponse, error) {
	req := &rpc.Request{
		Method: "GET",
		URL:    "service.device-registry/room",
		Body:   m,
	}

	rsp := &GetRoomResponse{}
	_, err := rpc.Do(req, rsp)
	return rsp, err
}

// Do performs the request
func (m *ListRoomsRequest) Do() (*ListRoomsResponse, error) {
	req := &rpc.Request{
		Method: "GET",
		URL:    "service.device-registry/rooms",
		Body:   m,
	}

	rsp := &ListRoomsResponse{}
	_, err := rpc.Do(req, rsp)
	return rsp, err
}
