package transmission

// Copyright 2015, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"encoding/json"
	"fmt"
	"log"
)

/*
Method name

session-stats
*/
type Session_Stats_Response struct {
	ActiveTorrentCount int                  `json:"activeTorrentCount,omitempty"`
	DownloadSpeed      int                  `json:"downloadSpeed,omitempty"`
	PausedTorrentCount int                  `json:"pausedTorrentCount,omitempty"`
	TorrentCount       int                  `json:"torrentCount,omitempty"`
	UploadSpeed        int                  `json:"uploadSpeed,omitempty"`
	CumulativeStats    *Session_Stats_Units `json:"cumulative-stats,omitempty"`
	CurrentStats       *Session_Stats_Units `json:"current-stats,omitempty"`
}
type Session_Stats_Units struct {
	UploadedBytes   int `json:"uploadedBytes,omitempty"`
	DownloadedBytes int `json:"downloadedBytes,omitempty"`
	FilesAdded      int `json:"filesAdded,omitempty"`
	SessionCount    int `json:"sessionCount,omitempty"`
	SecondsActive   int `json:"secondsActive,omitempty"`
}

func (self *Client) GetStats() (
	*Session_Stats_Response,
	error,
) {
	request := &Request{}
	request.Method = "session-stats"
	var result *Response
	var err error
	for i := 0; i < 5; i++ {
		if result, err = self.Request(request); err == nil && result != nil {
			stats := &Session_Stats_Response{}
			if err2 := json.Unmarshal(result.Args, &stats); err2 != nil {
				log.Printf("Transmission.GetStats(): failed to unmarshal %s | raw: %s\n", err2, result.Args)
			} else {
				return stats, nil
			}
		}
	}
	return nil, fmt.Errorf("Transmission.GetStats(): failed to get stats: %s", err)
}
