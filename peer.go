// Copyright 2015, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.
package transmission

type Peer struct {
	Address            string  `json:"address,omitempty"`
	ClientName         string  `json:"clientName,omitempty"`
	ClientIsChoked     bool    `json:"clientIsChoked,omitempty"`
	ClientIsInterested bool    `json:"clientIsInterested,omitempty"`
	FlagStr            string  `json:"flagStr,omitempty"`
	IsDownloadingFrom  bool    `json:"isDownloadingFrom,omitempty"`
	IsEncrypted        bool    `json:"isEncrypted,omitempty"`
	IsIncoming         bool    `json:"isIncoming,omitempty"`
	IsUploadingTo      bool    `json:"isUploadingTo,omitempty"`
	IsUTP              bool    `json:"isUTP,omitempty"`
	PeerIsChoked       bool    `json:"peerIsChoked,omitempty"`
	PeerIsInterested   bool    `json:"peerIsInterested,omitempty"`
	Port               int     `json:"port,omitempty"`
	Progress           float64 `json:"progress,omitempty"`
	RateToClient       int     `json:"rateToClient,omitempty"`
	RateToPeer         int     `json:"rateToPeer,omitempty"`
}
type PeersFrom struct {
	FromCache    int `json:"fromCache,omitempty"`
	FromDHT      int `json:"fromDht,omitempty"`
	FromIncoming int `json:"fromIncoming,omitempty"`
	FromLPD      int `json:"fromLpd,omitempty"`
	FromLTEP     int `json:"fromLtep,omitempty"`
	FromPEX      int `json:"fromPex,omitempty"`
	FromTracker  int `json:"fromTracker,omitempty"`
}
