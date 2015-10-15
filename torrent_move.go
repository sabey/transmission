// Copyright 2015, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.
package transmission

/*
Method name

torrent-set-location
*/
/*
All torrents are used if the "ids" argument is omitted.
*/
type Move_Argument struct {
	IDs      []int  `json:"ids,omitempty"`
	Location string `json:"location,omitempty"`
	Move     bool   `json:"move,omitempty"`
}
