package constant

type ParsedOption struct {
	Priority     int
	ParseStrings []string
}

var (
	Genres = []string{
		"Action",
		"Adult",
		"Adventure",
		"Animation",
		"Biography",
		"Comedy",
		"Crime",
		"Documentary",
		"Drama",
		"Family",
		"Fantasy",
		"Film-Noir",
		"Game-Show",
		"History",
		"Horror",
		"Musical",
		"Music",
		"Mystery",
		"News",
		"Reality-TV",
		"Romance",
		"Sci-Fi",
		"Short",
		"Sport",
		"Talk-Show",
		"Thriller",
		"War",
		"Western",
	}
	RipTypes = map[string]ParsedOption{
		"CAM": {
			Priority:     10,
			ParseStrings: []string{"CAM", "CAM-Rip", "HDCAM"},
		},
		"TELESYNC": {
			Priority:     20,
			ParseStrings: []string{"TELESYNC", "TS", "HDTS", "PDVD", "PreDVDRip"},
		},
		"SCR": {
			Priority:     30,
			ParseStrings: []string{"SCR", "SCREENER", "DVDSCR", "DVDSCREENER", "BDSCR"},
		},
		"HDTV": {
			Priority:     40,
			ParseStrings: []string{"HDTV", "PDTV", "TVRip", "HDTVRip", "DTVRip", "DVBRip", "DTHRip", "SATRip", "DSRip", "DSR", "HC", "HD-Rip"},
		},
		"DVDRip": {
			Priority:     50,
			ParseStrings: []string{"DVDR", "DVD-Full", "Full-Rip", "DVD-5", "DVD-9", "DVDRip", "DVDMux", "DVD5", "DVD9"},
		},
		"HDRip": {
			Priority:     60,
			ParseStrings: []string{"HDRip", "WEB-DLRip"},
		},
		"WEBCap": {
			Priority:     70,
			ParseStrings: []string{"WEBCAP", "WEB-CAP", "WEB CAP"},
		},
		"WEBRip": {
			Priority:     80,
			ParseStrings: []string{"WEBRip", "WEB Rip", "WEB-RIP"},
		},
		"WEBDL": {
			Priority:     90,
			ParseStrings: []string{"WEBDL", "WEB DL", "WEB-DL", "WEB"},
		},
		"BDRip": {
			Priority:     100,
			ParseStrings: []string{"BDRip", "BRip", "BRRip", "BDMV", "BDR", "BD25", "BD50", "BD5", "BD9", "Blu-Ray", "BluRay", "BLURAY"},
		},
	}
	ResolutionTypes = map[string]ParsedOption{
		"352p": {
			Priority:     10,
			ParseStrings: []string{"352p"},
		},
		"360p": {
			Priority:     20,
			ParseStrings: []string{"360p"},
		},
		"480i": {
			Priority:     30,
			ParseStrings: []string{"480i"},
		},
		"480p": {
			Priority:     40,
			ParseStrings: []string{"480p"},
		},
		"576i": {
			Priority:     50,
			ParseStrings: []string{"576i"},
		},
		"576p": {
			Priority:     60,
			ParseStrings: []string{"576p"},
		},
		"720p": {
			Priority:     70,
			ParseStrings: []string{"720"},
		},
		"1080p": {
			Priority:     80,
			ParseStrings: []string{"1080"},
		},
		"2160p": {
			Priority:     90,
			ParseStrings: []string{"2160"},
		},
	}
	EncodingTypes = map[string]ParsedOption{
		"Xvid": {
			Priority:     10,
			ParseStrings: []string{"Xvid", "DivX"},
		},
		"x264": {
			Priority:     20,
			ParseStrings: []string{"x264", "H.264", "AVC"},
		},
		"x265": {
			Priority:     30,
			ParseStrings: []string{"x265", "HEVC", "H.265"},
		},
		"VP9": {
			Priority:     40,
			ParseStrings: []string{"VP9"},
		},
	}
	DownloaderTypes = map[string]string{
		"transmission": "torrent",
		"nzbget": "nzb",
	}
)

const (
	// Whenever the downloader has the download queued, but is not currently downloading
	StatusWaiting = "waiting"
	// Whenever the download has been set to paused in the downloader
	StatusPaused = "paused"
	// Whenever the downloader is actively downloading the download
	StatusDownloading = "downloading"
	// Whenever the downloader is unpacking/moving/processing the download
	StatusProcessing = "processing"
	// Whenever the downloader believes its job to be completely done for the download
	StatusComplete = "complete"
	// The downloader failed to download the release
	StatusError = "error"
	// Conductorr is processing the download
	StatusCProcessing = "cprocessing"
	// Conductorr failed to process the download
	StatusCError = "cerror"
	// Conductorr is done with the download
	StatusDone = "done"
)