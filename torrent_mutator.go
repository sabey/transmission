// Copyright 2015, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.
package transmission

/*
Method name

torrent-set
*/
/*
Just as an empty "ids" value is shorthand for "all ids", using an empty array
for "files-wanted", "files-unwanted", "priority-high", "priority-low", or
"priority-normal" is shorthand for saying "all files".
*/
type Mutator_Argument struct {
	IDs                 []int    `json:"ids,omitempty"`
	BandwidthPriority   int      `json:"bandwidthPriority,omitempty"`
	DownloadLimit       int      `json:"downloadLimit,omitempty"`
	DownloadLimited     bool     `json:"downloadLimited,omitempty"`
	FilesWanted         []*File  `json:"files-wanted,omitempty"`
	FilesUnwanted       []*File  `json:"files-unwanted,omitempty"`
	HonorsSessionLimits bool     `json:"honorsSessionLimits,omitempty"`
	Location            string   `json:"location,omitempty"`
	PeerLimit           int      `json:"peer-limit,omitempty"`
	PriorityHigh        []*File  `json:"priority-high,omitempty"`
	PriorityLow         []*File  `json:"priority-low,omitempty"`
	PriorityNormal      []*File  `json:"priority-normal,omitempty"`
	QueuePosition       int      `json:"queuePosition,omitempty"`
	SeedIdleLimit       int      `json:"seedIdleLimit,omitempty"`
	SeedIdleMode        int      `json:"seedIdleMode,omitempty"`
	SeedRatioLimit      float64  `json:"seedRatioLimit,omitempty"`
	SeedRatioMode       int      `json:"seedRatioMode,omitempty"`
	TrackerAdd          []string `json:"trackerAdd,omitempty"`
	TrackerRemove       []int    `json:"trackerRemove,omitempty"`
	TrackerReplace      []string `json:"trackerReplace,omitempty"`
	UploadLimit         int      `json:"uploadLimit,omitempty"`
	UploadLimited       bool     `json:"uploadLimited,omitempty"`
}
