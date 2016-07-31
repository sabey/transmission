package transmission

// Copyright 2015, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

type Tracker struct {
	ID       int    `json:"id,omitempty"`
	Announce string `json:"announce,omitempty"`
	Scrape   string `json:"scrape,omitempty"`
	Tier     int    `json:"tier,omitempty"`
}
type TrackerStat struct {
	ID                    int    `json:"id,omitempty"`
	Announce              string `json:"announce,omitempty"`
	AnnounceState         int    `json:"announceState,omitempty"`
	DownloadCount         int    `json:"downloadCount,omitempty"`
	HasAnnounced          bool   `json:"hasAnnounced,omitempty"`
	HasScraped            bool   `json:"hasScraped,omitempty"`
	Host                  string `json:"host,omitempty"`
	IsBackup              bool   `json:"isBackup,omitempty"`
	LastAnnouncePeerCount int    `json:"lastAnnouncePeerCount,omitempty"`
	LastAnnounceResult    string `json:"lastAnnounceResult,omitempty"`
	LastAnnounceStartTime int    `json:"lastAnnounceStartTime,omitempty"`
	LastAnnounceSucceeded bool   `json:"lastAnnounceSucceeded,omitempty"`
	LastAnnounceTime      int    `json:"lastAnnounceTime,omitempty"`
	LastAnnounceTimedOut  bool   `json:"lastAnnounceTimedOut,omitempty"`
	LastScrapeResult      string `json:"lastScrapeResult,omitempty"`
	LastScrapeStartTime   int    `json:"lastScrapeStartTime,omitempty"`
	LastScrapeSucceeded   bool   `json:"lastScrapeSucceeded,omitempty"`
	LastScrapeTime        int    `json:"lastScrapeTime,omitempty"`
	LastScrapeTimedOut    int    `json:"lastScrapeTimedOut,omitempty"`
	LeecherCount          int    `json:"leecherCount,omitempty"`
	NextAnnounceTime      int    `json:"nextAnnounceTime,omitempty"`
	NextScrapeTime        int    `json:"nextScrapeTime,omitempty"`
	Scrape                string `json:"scrape,omitempty"`
	ScrapeState           int    `json:"scrapeState,omitempty"`
	SeederCount           int    `json:"seederCount,omitempty"`
	Tier                  int    `json:"tier,omitempty"`
}
