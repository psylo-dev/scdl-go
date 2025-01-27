package scdl

import "github.com/imthaghost/scdl/pkg/soundcloud"

// Download downloads a song from a given SoundCloud URL.
// This function is exported (uppercase) so `gomobile bind` can use it.
func Download(url string) error {
	// Create a new SoundCloud client
	sc := soundcloud.NewClient("", nil)
	// Download the song
	return sc.Download(url)
}
