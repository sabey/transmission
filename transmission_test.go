package transmission

// Copyright 2015, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"encoding/json"
	"fmt"
	"log"
	"sabey.co/unittest"
	"testing"
)

/*
	for these unittests to pass you have to have the transmission api running
	you're expected to have atleast one existing torrent
	these unittests overall don't offer much in integrity
	I'm open to accepting a proper set of unittests with mock torrents
*/
func TestTransmission(t *testing.T) {
	fmt.Println("TestTransmission()")
	client := Construct(
		nil, // http client
		"",  // address
		"",  // username
		"",  // password
	)
	unittest.NotNil(t, client)

	// GetTorrents
	torrents, err := client.GetTorrents()
	unittest.IsNil(t, err)
	unittest.Equals(t, len(torrents) > 0, true)
	// print examples of a real populated torrents
	bs, err := json.MarshalIndent(torrents, "", "\t")
	if err != nil {
		log.Printf("Transmission Error: %s\n", err)
	}
	unittest.IsNil(t, err)
	fmt.Printf("Torrents:\n%s\n", bs)
	for _, t := range torrents {
		// if you don't know the magnet hash of your torrents
		// you should be looping and determining which of these torrents you want to work with
		// the library requires magnet hashes and not the ID in transmission because that was too complex
		// a magnet hash will never change and the transmission ID is unreliable to the end user
		//
		// my usecase is to check t.Creator to see which services I tagged the torrent to
		found := false
		done := false
		// t.IsDone() or:
		for _, f := range t.Files {
			found = true
			// checking the Bytescompleted on every file seems to be the only accurate way to see if a torrent is done
			// other values were incorrect
			if f.Bytescompleted == f.Length {
				done = true
			} else {
				done = false
				break
			}
		}
		if found {
			if done {
				log.Printf("%s %s | all files done\n", t.HashString, t.Name)
			} else {
				log.Printf("%s %s | not done\n", t.HashString, t.Name)
			}
		}
	}

	/*
		// don't run these automatically
		// we don't want our users first experience with this library to be anger
		client.RemoveTorrents(
			"remove_magnet",
			"remove this",
			"remove doesn't delete files",
		)
		client.RemoveUnknownTorrents(
			"dontremove_magnet",
			"remove everything else",
		)
		client.StopTorrents(
			"stop_magnet",
			"stop this",
		)
		client.StopUnknownTorrents(
			"dontstop_magnet",
			"stop everything else",
		)
		client.StartTorrents(
			"start_magnet",
			"start this",
		)
		client.StartUnknownTorrents(
			"dontstart_magnet",
			"start everything else",
		)
		client.StartNowTorrents(
			"startnow_magnet",
			"start this now!",
		)
		client.StartNowUnknownTorrents(
			"dontstartnow_magnet",
			"start everything else now!",
		)
		client.VerifyTorrents(
			"verify_magnet",
			"verify this",
			"Hello, is it me you're looking for?",
		)
		client.VerifyUnknownTorrents(
			"dontverify_magnet",
			"don't verify this",
			"I trust this",
		)
		client.ReannounceTorrents(
			"reannounce_magnet",
			"reannounce this",
			"shoutout",
		)
		client.ReannounceUnknownTorrents(
			"dontreannounce_magnet",
			"don't announce this",
			"don't mention me",
		)
	*/

	// session
	session, err := client.GetSession()
	unittest.IsNil(t, err)
	unittest.NotNil(t, session)
	bs, _ = json.MarshalIndent(session, "", "\t")
	fmt.Printf("GetSession:\n%s\n", bs)

	// stats
	stats, err := client.GetStats()
	unittest.IsNil(t, err)
	unittest.NotNil(t, session)
	bs, _ = json.MarshalIndent(stats, "", "\t")
	fmt.Printf("GetStats:\n%s\n", bs)

	// add torrent
	/*
		torrent, err := client.AddTorrent(
			"", // `magnet:?xt=urn:btih:673b3cc3f5a4d6b1fecdf64eeaa8733887de2433&dn=localhost-1434676188.tar.gz.pgp&tr=http%3A%2F%2Flocalhost%3A6969%2Fannounce&tr=http%3A%2F%2Flocalhost%3A6970%2Fannounce`,
			"", // "/path/to/downloads/",
		)
		unittest.IsNil(t, err)
		unittest.NotNil(t, torrent)
		bs, _ = json.MarshalIndent(torrent, "", "\t")
		fmt.Printf("AddTorrent:\n%s\n", bs)
	*/
}
