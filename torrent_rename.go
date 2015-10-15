// Copyright 2015, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.
package transmission

/*
Method name

torrent-rename-path
*/
/*
All torrents are used if the "ids" argument is omitted.
*/
/*
if this call succeeds you'll want to update the torrent's "files" and "name" field with torrent-get.
*/
type Rename_Argument struct {
	IDs  []int  `json:"ids,omitempty"`
	Path string `json:"path,omitempty"`
	Name bool   `json:"name,omitempty"`
}
