package transmission

// Copyright 2015, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"fmt"
)

/*
Method name

torrent-remove
*/
/*
All torrents are used if the "ids" argument is omitted.
*/
type Remove_Argument struct {
	IDs             []int `json:"ids,omitempty"`
	DeleteLocalData bool  `json:"delete-local-data,omitempty"`
}

func (self *Client) RemoveTorrents(
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
	request.Method = "torrent-remove"
	access := Remove_Argument{}
	access.IDs = ids
	access.DeleteLocalData = false
	request.Args = access
	var result *Response
	for i := 0; i < 5; i++ {
		if result, err = self.Request(request); err == nil && result != nil {
			return failed, nil
		}
	}
	return failed, fmt.Errorf("Transmission.RemoveTorrents(): failed to remove torrents: %s", err)
}
func (self *Client) RemoveUnknownTorrents(
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
	request.Method = "torrent-remove"
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
	return removed, fmt.Errorf("Transmission.RemoveUnknownTorrents(): failed to remove unknown torrents: %s", err)
}
