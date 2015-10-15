// Copyright 2015, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.
package transmission

import (
	"fmt"
)

/*
Method name

torrent-start
torrent-start-now
torrent-stop
torrent-verify
torrent-reannounce
*/
/*
All torrents are used if the "ids" argument is omitted.
*/
type Action_Argument struct {
	IDs []int `json:"ids,omitempty"`
}

func (self *Client) StartTorrents(
	magnets ...string,
) (
	[]string,
	error,
) {
	ids, failed, err := self.IDs(magnets...)
	if err != nil {
		return nil, err
	}
	if len(ids) == 0 {
		return failed, nil
	}
	request := &Request{}
	request.Method = "torrent-start"
	access := Remove_Argument{}
	access.IDs = ids
	request.Args = access
	var result *Response
	for i := 0; i < 5; i++ {
		if result, err = self.Request(request); err == nil && result != nil {
			return failed, nil
		}
	}
	return failed, fmt.Errorf("Transmission.StartTorrents(): failed to start torrents: %s", err)
}
func (self *Client) StartUnknownTorrents(
	ignore_magnets ...string,
) (
	[]string,
	error,
) {
	ids, removed, err := self.IDsNot(ignore_magnets...)
	if err != nil {
		return nil, err
	}
	if len(ids) == 0 {
		return removed, nil
	}
	request := &Request{}
	request.Method = "torrent-start"
	access := Remove_Argument{}
	access.IDs = ids
	access.DeleteLocalData = false
	request.Args = access
	var result *Response
	for i := 0; i < 5; i++ {
		if result, err = self.Request(request); err == nil && result != nil {
			return removed, nil
		}
	}
	return removed, fmt.Errorf("Transmission.StartUnknownTorrents(): failed to start unknown torrents: %s", err)
}
func (self *Client) StartNowTorrents(
	magnets ...string,
) (
	[]string,
	error,
) {
	ids, failed, err := self.IDs(magnets...)
	if err != nil {
		return nil, err
	}
	if len(ids) == 0 {
		return failed, nil
	}
	request := &Request{}
	request.Method = "torrent-start-now"
	access := Remove_Argument{}
	access.IDs = ids
	request.Args = access
	var result *Response
	for i := 0; i < 5; i++ {
		if result, err = self.Request(request); err == nil && result != nil {
			return failed, nil
		}
	}
	return failed, fmt.Errorf("Transmission.StartNowTorrents(): failed to start-now torrents: %s", err)
}
func (self *Client) StartNowUnknownTorrents(
	ignore_magnets ...string,
) (
	[]string,
	error,
) {
	ids, removed, err := self.IDsNot(ignore_magnets...)
	if err != nil {
		return nil, err
	}
	if len(ids) == 0 {
		return removed, nil
	}
	request := &Request{}
	request.Method = "torrent-start-now"
	access := Remove_Argument{}
	access.IDs = ids
	access.DeleteLocalData = false
	request.Args = access
	var result *Response
	for i := 0; i < 5; i++ {
		if result, err = self.Request(request); err == nil && result != nil {
			return removed, nil
		}
	}
	return removed, fmt.Errorf("Transmission.StartNowUnknownTorrents(): failed to start-now unknown torrents: %s", err)
}
func (self *Client) StopTorrents(
	magnets ...string,
) (
	[]string,
	error,
) {
	ids, failed, err := self.IDs(magnets...)
	if err != nil {
		return nil, err
	}
	if len(ids) == 0 {
		return failed, nil
	}
	request := &Request{}
	request.Method = "torrent-stop"
	access := Remove_Argument{}
	access.IDs = ids
	request.Args = access
	var result *Response
	for i := 0; i < 5; i++ {
		if result, err = self.Request(request); err == nil && result != nil {
			return failed, nil
		}
	}
	return failed, fmt.Errorf("Transmission.StopTorrents(): failed to stop torrents: %s", err)
}
func (self *Client) StopUnknownTorrents(
	ignore_magnets ...string,
) (
	[]string,
	error,
) {
	ids, removed, err := self.IDsNot(ignore_magnets...)
	if err != nil {
		return nil, err
	}
	if len(ids) == 0 {
		return removed, nil
	}
	request := &Request{}
	request.Method = "torrent-stop"
	access := Remove_Argument{}
	access.IDs = ids
	access.DeleteLocalData = false
	request.Args = access
	var result *Response
	for i := 0; i < 5; i++ {
		if result, err = self.Request(request); err == nil && result != nil {
			return removed, nil
		}
	}
	return removed, fmt.Errorf("Transmission.StopUnknownTorrents(): failed to stop unknown torrents: %s", err)
}
func (self *Client) VerifyTorrents(
	magnets ...string,
) (
	[]string,
	error,
) {
	ids, failed, err := self.IDs(magnets...)
	if err != nil {
		return nil, err
	}
	if len(ids) == 0 {
		return failed, nil
	}
	request := &Request{}
	request.Method = "torrent-verify"
	access := Remove_Argument{}
	access.IDs = ids
	request.Args = access
	var result *Response
	for i := 0; i < 5; i++ {
		if result, err = self.Request(request); err == nil && result != nil {
			return failed, nil
		}
	}
	return failed, fmt.Errorf("Transmission.VerifyTorrents(): failed to verify torrents: %s", err)
}
func (self *Client) VerifyUnknownTorrents(
	ignore_magnets ...string,
) (
	[]string,
	error,
) {
	ids, removed, err := self.IDsNot(ignore_magnets...)
	if err != nil {
		return nil, err
	}
	if len(ids) == 0 {
		return removed, nil
	}
	request := &Request{}
	request.Method = "torrent-verify"
	access := Remove_Argument{}
	access.IDs = ids
	access.DeleteLocalData = false
	request.Args = access
	var result *Response
	for i := 0; i < 5; i++ {
		if result, err = self.Request(request); err == nil && result != nil {
			return removed, nil
		}
	}
	return removed, fmt.Errorf("Transmission.VerifyUnknownTorrents(): failed to verify unknown torrents: %s", err)
}
func (self *Client) ReannounceTorrents(
	magnets ...string,
) (
	[]string,
	error,
) {
	ids, failed, err := self.IDs(magnets...)
	if err != nil {
		return nil, err
	}
	if len(ids) == 0 {
		return failed, nil
	}
	request := &Request{}
	request.Method = "torrent-reannounce"
	access := Remove_Argument{}
	access.IDs = ids
	request.Args = access
	var result *Response
	for i := 0; i < 5; i++ {
		if result, err = self.Request(request); err == nil && result != nil {
			return failed, nil
		}
	}
	return failed, fmt.Errorf("Transmission.ReannounceTorrents(): failed to reannounce torrents: %s", err)
}
func (self *Client) ReannounceUnknownTorrents(
	ignore_magnets ...string,
) (
	[]string,
	error,
) {
	ids, removed, err := self.IDsNot(ignore_magnets...)
	if err != nil {
		return nil, err
	}
	if len(ids) == 0 {
		return removed, nil
	}
	request := &Request{}
	request.Method = "torrent-reannounce"
	access := Remove_Argument{}
	access.IDs = ids
	access.DeleteLocalData = false
	request.Args = access
	var result *Response
	for i := 0; i < 5; i++ {
		if result, err = self.Request(request); err == nil && result != nil {
			return removed, nil
		}
	}
	return removed, fmt.Errorf("Transmission.ReannounceUnknownTorrents(): failed to reannounce unknown torrents: %s", err)
}
