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

session-get
session-set
*/
/*
session-set:
any attribute except:

"config-dir",
"rpc-version",
"rpc-version-minimum",
"version",
*/
type Session_Arguments struct {
	AltSpeedDown              int            `json:"alt-speed-down,omitempty"`
	AltSpeedEnabled           bool           `json:"alt-speed-enabled,omitempty"`
	AltSpeedTimeBegin         int            `json:"alt-speed-time-begin,omitempty"`
	AltSpeedTimeEnabled       bool           `json:"alt-speed-time-enabled,omitempty"`
	AltSpeedTimeEnd           int            `json:"alt-speed-time-end,omitempty"`
	AltSpeedTimeDay           int            `json:"alt-speed-time-day,omitempty"`
	AltSpeedUp                int            `json:"alt-speed-up,omitempty"`
	BlocklistUrl              string         `json:"blocklist-url,omitempty"`
	BlocklistEnabled          bool           `json:"blocklist-enabled,omitempty"`
	BlocklistSize             int            `json:"blocklist-size,omitempty"`
	CacheSizeMB               int            `json:"cache-size-mb,omitempty"`
	ConfigDir                 string         `json:"config-dir,omitempty"`
	DownloadDir               string         `json:"download-dir,omitempty"`
	DownloadQueueSize         int            `json:"download-queue-size,omitempty"`
	DownloadQueueEnabled      bool           `json:"download-queue-enabled,omitempty"`
	DHTEnabled                bool           `json:"dht-enabled,omitempty"`
	Encryption                string         `json:"encryption,omitempty"`
	IdleSeedingLimit          int            `json:"idle-seeding-limit,omitempty"`
	IdleSeedingLimitEnabled   bool           `json:"idle-seeding-limit-enabled,omitempty"`
	IncompleteDir             string         `json:"incomplete-dir,omitempty"`
	IncompleteDirEnabled      bool           `json:"incomplete-dir-enabled,omitempty"`
	LPDEnabled                bool           `json:"lpd-enabled,omitempty"`
	PeerLimitGlobal           int            `json:"peer-limit-global,omitempty"`
	PeerLimitPerTorrent       int            `json:"peer-limit-per-torrent,omitempty"`
	PEXEnabled                bool           `json:"pex-enabled,omitempty"`
	PeerPort                  int            `json:"peer-port,omitempty"`
	PeerPortRandomOnStart     bool           `json:"peer-port-random-on-start,omitempty"`
	PortForwardingEnabled     bool           `json:"port-forwarding-enabled,omitempty"`
	QueueStalledEnabled       bool           `json:"queue-stalled-enabled,omitempty"`
	QueueStalledMinutes       int            `json:"queue-stalled-minutes,omitempty"`
	RenamePartialFiles        bool           `json:"rename-partial-files,omitempty"`
	RPCVersion                int            `json:"rpc-version,omitempty"`
	RPCVersionMinimum         int            `json:"rpc-version-minimum,omitempty"`
	ScriptTorrentDoneFilename string         `json:"script-torrent-done-filename,omitempty"`
	ScriptTorrentDoneEnabled  bool           `json:"script-torrent-done-enabled,omitempty"`
	SeedRatioLimit            float64        `json:"seedRatioLimit,omitempty"`
	SeedRatioLimited          bool           `json:"seedRatioLimited,omitempty"`
	SeedQueueSize             int            `json:"seed-queue-size,omitempty"`
	SeedQueueEnabled          bool           `json:"seed-queue-enabled,omitempty"`
	SpeedLimitDown            int            `json:"speed-limit-down,omitempty"`
	SpeedLimitDownEnabled     bool           `json:"speed-limit-down-enabled,omitempty"`
	SpeedLimitUp              int            `json:"speed-limit-up,omitempty"`
	SpeedLimitUpEnabled       bool           `json:"speed-limit-up-enabled,omitempty"`
	StartAddedTorrents        bool           `json:"start-added-torrents,omitempty"`
	TrashOriginalTorrentFiles bool           `json:"trash-original-torrent-files,omitempty"`
	Units                     *Session_Units `json:"units,omitempty"`
	UTPEnabled                bool           `json:"utp-enabled,omitempty"`
	Version                   string         `json:"version,omitempty"`
}
type Session_Units struct {
	SpeedUnits  []string `json:"speed-units,omitempty"`
	SpeedBytes  int      `json:"speed-bytes,omitempty"`
	SizeUnits   []string `json:"size-units,omitempty"`
	SizeBytes   int      `json:"size-bytes,omitempty"`
	MemoryUnits []string `json:"memory-units,omitempty"`
	MemoryBytes int      `json:"memory-bytes,omitempty"`
}

func (self *Client) GetSession() (
	*Session_Arguments,
	error,
) {
	request := &Request{}
	request.Method = "session-get"
	var result *Response
	var err error
	for i := 0; i < 5; i++ {
		if result, err = self.Request(request); err == nil && result != nil {
			session := &Session_Arguments{}
			if err2 := json.Unmarshal(result.Args, &session); err2 != nil {
				log.Printf("Transmission.GetSession(): failed to unmarshal %s | raw: %s\n", err2, result.Args)
			} else {
				return session, nil
			}
		}
	}
	return nil, fmt.Errorf("Transmission.GetSession(): failed to get session: %s", err)
}
