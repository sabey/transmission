// Copyright 2015, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.
package transmission

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type Torrent struct {
	ID                      int            `json:"id,omitempty"`
	ActivityDate            int            `json:"activityDate,omitempty"`
	AddedDate               int            `json:"addedDate,omitempty"`
	BandwidthPriority       int            `json:"bandwidthPriority,omitempty"`
	Comment                 string         `json:"comment,omitempty"`
	CorruptEver             int            `json:"corruptEver,omitempty"`
	Creator                 string         `json:"creator,omitempty"`
	DateCreated             int            `json:"dateCreated,omitempty"`
	DesiredAvailable        int            `json:"desiredAvailable,omitempty"`
	DoneDate                int            `json:"doneDate,omitempty"`
	DownloadDir             string         `json:"downloadDir,omitempty"`
	DownloadedEver          int            `json:"downloadedEver,omitempty"`
	DownloadLimit           int            `json:"downloadLimit,omitempty"`
	DownloadLimited         bool           `json:"downloadLimited,omitempty"`
	Error                   int            `json:"error,omitempty"`
	ErrorString             string         `json:"errorString,omitempty"`
	Eta                     int            `json:"eta,omitempty"`
	EtaIdle                 int            `json:"etaIdle,omitempty"`
	Files                   []*File        `json:"files,omitempty"`
	FileStats               []*FileStat    `json:"fileStats,omitempty"`
	HashString              string         `json:"hashString,omitempty"`
	HaveUnchecked           int            `json:"haveUnchecked,omitempty"`
	HaveValid               int            `json:"haveValid,omitempty"`
	HonorsSessionLimits     bool           `json:"honorsSessionLimits,omitempty"`
	IsFinished              bool           `json:"isFinished,omitempty"`
	IsPrivate               bool           `json:"isPrivate,omitempty"`
	IsStalled               bool           `json:"isStalled,omitempty"`
	LeftUntilDone           int            `json:"leftUntilDone,omitempty"`
	MagnetLink              string         `json:"magnetLink,omitempty"`
	ManualAnnounceTime      int            `json:"manualAnnounceTime,omitempty"`
	MaxConnectedPeers       int            `json:"maxConnectedPeers,omitempty"`
	MetadataPercentComplete float64        `json:"metadataPercentComplete,omitempty"`
	Name                    string         `json:"name,omitempty"`
	PeerLimit               int            `json:"peer-limit,omitempty"`
	Peers                   []*Peer        `json:"peers,omitempty"`
	PeersConnected          int            `json:"peersConnected,omitempty"`
	PeersFrom               *PeersFrom     `json:"peersFrom,omitempty"`
	PeersGettingFromUs      int            `json:"peersGettingFromUs,omitempty"`
	PeersSendingToUs        int            `json:"peersSendingToUs,omitempty"`
	PercentDone             float64        `json:"percentDone,omitempty"`
	Pieces                  string         `json:"pieces,omitempty"`
	PieceCount              int            `json:"pieceCount,omitempty"`
	PieceSize               int            `json:"pieceSize,omitempty"`
	Priorities              []int          `json:"priorities,omitempty"`
	QueuePosition           int            `json:"queuePosition,omitempty"`
	RateDownload            int            `json:"rateDownload,omitempty"`
	RateUpload              int            `json:"rateUpload,omitempty"`
	RecheckProgress         float64        `json:"recheckProgress,omitempty"`
	SecondsDownloading      int            `json:"secondsDownloading,omitempty"`
	SecondsSeeding          int            `json:"secondsSeeding,omitempty"`
	SeedIdleLimit           int            `json:"seedIdleLimit,omitempty"`
	SeedIdleMode            int            `json:"seedIdleMode,omitempty"`
	SeedRatioLimit          float64        `json:"seedRatioLimit,omitempty"`
	SeedRatioMode           int            `json:"seedRatioMode,omitempty"`
	SizeWhenDone            int            `json:"sizeWhenDone,omitempty"`
	StartDate               int            `json:"startDate,omitempty"`
	Status                  int            `json:"status,omitempty"`
	Trackers                []*Tracker     `json:"trackers,omitempty"`
	TrackerStats            []*TrackerStat `json:"trackerStats,omitempty"`
	TotalSize               int            `json:"totalSize,omitempty"`
	TorrentFile             string         `json:"torrentFile,omitempty"`
	UploadedEver            int            `json:"uploadedEver,omitempty"`
	UploadLimit             int            `json:"uploadLimit,omitempty"`
	UploadLimited           bool           `json:"uploadLimited,omitempty"`
	UploadRatio             float64        `json:"uploadRatio,omitempty"`
	Wanted                  []int          `json:"wanted,omitempty"`
	Webseeds                []string       `json:"webseeds,omitempty"`
	WebseedsSendingToUs     int            `json:"webseedsSendingToUs,omitempty"`
}
type Torrent_Response struct {
	Torrents []*Torrent `json:"torrents,omitempty"`
}

func (self *Client) GetTorrents() (
	[]*Torrent,
	error,
) {
	request := &Request{}
	request.Method = "torrent-get"
	access := Accessor_Argument{}
	access.Fields = accessor_fields
	request.Args = access
	var result *Response
	var err error
	for i := 0; i < 5; i++ {
		if result, err = self.Request(request); err == nil && result != nil {
			torrents := Torrent_Response{}
			if err2 := json.Unmarshal(result.Args, &torrents); err2 != nil {
				log.Printf("Transmission.GetTorrents(): failed to unmarshal %s | raw: %s\n", err2, result.Args)
			} else {
				return torrents.Torrents, nil
			}
		}
	}
	return nil, fmt.Errorf("Transmission.GetTorrents(): failed to get torrents: %s", err)
}
func (self *Client) Exists(
	magnet string,
) bool {
	if magnet == "" {
		return false
	}
	torrents, _ := self.GetTorrents()
	for _, t := range torrents {
		if strings.ToLower(t.HashString) == strings.ToLower(magnet) {
			return true
		}
	}
	return false
}
func (self *Client) IDs(
	magnets ...string,
) (
	[]int,
	[]string,
	error,
) {
	magnets = UniqueStrings(magnets)
	if len(magnets) == 0 {
		return nil, nil, fmt.Errorf("Transmission.IDs(): magnets empty")
	}
	ids := []int{}
	failed := []string{}
	torrents, err := self.GetTorrents()
	if err != nil {
		// failed to get torrents
		return nil, nil, err
	}
	for _, m := range magnets {
		f := false
		for _, t := range torrents {
			if strings.ToLower(t.HashString) == strings.ToLower(m) {
				f = true
				ids = append(ids, t.ID)
			}
		}
		if !f {
			failed = append(failed, m)
		}
	}
	return ids, failed, nil
}
func (self *Client) IDsNot(
	ignore_magnets ...string,
) (
	[]int,
	[]string,
	error,
) {
	ignore_magnets = UniqueStrings(ignore_magnets)
	if len(ignore_magnets) == 0 {
		return nil, nil, fmt.Errorf("Transmission.IDsNot(): ignore_magnets empty")
	}
	ids := []int{}
	removed := []string{}
	torrents, err := self.GetTorrents()
	if err != nil {
		// failed to get torrents
		return nil, nil, err
	}
	for _, t := range torrents {
		f := false
		for _, m := range ignore_magnets {
			if strings.ToLower(t.HashString) == strings.ToLower(m) {
				f = true
			}
		}
		if !f {
			removed = append(removed, t.HashString)
			ids = append(ids, t.ID)
		}
	}
	return ids, removed, nil
}
func (self *Torrent) IsDone() bool {
	if len(self.Files) == 0 {
		return false
	}
	// checking the Bytescompleted on every file seems to be the only accurate way to see if a torrent is done
	// other values were incorrect
	for _, f := range self.Files {
		if f.Bytescompleted != f.Length {
			return false
		}
	}
	return true
}
