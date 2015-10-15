// Copyright 2015, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.
package transmission

type File struct {
	Bytescompleted int    `json:"bytesCompleted,omitempty"`
	Length         int    `json:"length,omitempty"`
	Name           string `json:"name,omitempty"`
}
type FileStat struct {
	BytesCompleted int  `json:"bytesCompleted,omitempty"`
	Wanted         bool `json:"wanted,omitempty"`
	Priority       int  `json:"priority,omitempty"`
}
