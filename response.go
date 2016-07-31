package transmission

// Copyright 2015, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"encoding/json"
)

/*
Reponses support three keys:

(1) A required "result" string whose value MUST be "success" on success,
or an error string on failure.
(2) An optional "arguments" object of key/value pairs
(3) An optional "tag" number as described in 2.1.
*/
type Response struct {
	Args   json.RawMessage `json:"arguments,omitempty"`
	Result string          `json:"result,omitempty"`
}
