// Copyright 2015, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.
package transmission

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

/*
Method name

torrent-add
*/
/*
Either "filename" OR "metainfo" MUST be included.
*/
/*
Response arguments:

On success, a "torrent-added" object contains
id, name, and hashString.

On failure due to a duplicate torrent existing,
a "torrent-duplicate" object in the same form.
*/
type Add_Argument struct {
	Cookies           string  `json:"cookies,omitempty"`
	DownloadDir       string  `json:"download-dir,omitempty"`
	Filename          string  `json:"filename,omitempty"`
	MetaInfo          string  `json:"metainfo,omitempty"`
	Paused            bool    `json:"paused,omitempty"`
	PeerLimit         int     `json:"peer-limit,omitempty"`
	BandwidthPriority int     `json:"bandwidthPriority,omitempty"`
	FilesWanted       []*File `json:"files-wanted,omitempty"`
	FilesUnwanted     []*File `json:"files-unwanted,omitempty"`
	PriorityHigh      []*File `json:"priority-high,omitempty"`
	PriorityLow       []*File `json:"priority-low,omitempty"`
	PriorityNormal    []*File `json:"priority-normal,omitempty"`
}
type Add_Response struct {
	Torrent *Torrent `json:"torrent-added,omitempty"`
}

func (self *Client) AddTorrent(
	location string,
	path string,
) (
	*Torrent,
	error,
) {
	if location == "" {
		return nil, fmt.Errorf("Transmission.AddTorrent(): location empty")
	}
	if path == "" {
		return nil, fmt.Errorf("Transmission.AddTorrent(): path empty")
	}
	if !strings.HasPrefix(path, "/") {
		return nil, fmt.Errorf("Transmission.AddTorrent(): path doesn't start with /")
	}
	request := &Request{}
	request.Method = "torrent-add"
	access := Add_Argument{}
	access.Filename = location
	access.DownloadDir = path
	request.Args = access
	var result *Response
	var err error
	for i := 0; i < 5; i++ {
		if result, err = self.Request(request); err == nil && result != nil {
			if result.Result == "success" {
				torrent := Add_Response{}
				if err2 := json.Unmarshal(result.Args, &torrent); err2 != nil {
					log.Printf("Transmission.AddTorrent(): failed to unmarshal %s | raw: %v\n", err2, result.Args)
				} else {
					return torrent.Torrent, nil
				}
			} else {
				return nil, fmt.Errorf("Transmission.AddTorrent(): failed to add: %s", result.Result)
			}
		}
	}
	return nil, fmt.Errorf("Transmission.AddTorrent(): failed to get torrents: %s", err)
}
