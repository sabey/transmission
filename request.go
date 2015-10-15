// Copyright 2015, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.
package transmission

/*
Requests support three keys:

(1) A required "method" string telling the name of the method to invoke
(2) An optional "arguments" object of key/value pairs
(3) An optional "tag" number used by clients to track responses.
If provided by a request, the response MUST include the same tag.
*/
type Request struct {
	Method string      `json:"method,omitempty"`
	Args   interface{} `json:"arguments,omitempty"`
}
