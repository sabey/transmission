// Copyright 2015, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.
package transmission

// spec: https://trac.transmissionbt.com/browser/trunk/extras/rpc-spec.txt
// implemented transmission version 2.80

func UniqueStrings(
	ss []string,
) []string {
	uniq := make([]string, 0, len(ss))
	exists := make(map[string]bool)
	for _, s := range ss {
		if s != "" {
			if _, ok := exists[s]; !ok {
				exists[s] = true
				uniq = append(uniq, s)
			}
		}
	}
	return uniq
}
