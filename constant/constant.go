package constant

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
	RipTypes = map[string][]string{
		"BDRip":    {"BDRip", "BRip", "BRRip", "BDMV", "BDR", "BD25", "BD50", "BD5", "BD9", "Blu-Ray", "BluRay", "BLURAY"},
		"WEBRip":   {"WEBRip", "WEB Rip", "WEB-RIP"},
		"WEBDL":    {"WEBDL", "WEB DL", "WEB-DL", "WEB"},
		"WEBCap":   {"WEBCAP", "WEB-CAP", "WEB CAP"},
		"HDTV":     {"HDTV", "PDTV", "TVRip", "HDTVRip", "DTVRip", "DVBRip", "DTHRip", "SATRip", "DSRip", "DSR", "HC", "HD-Rip"},
		"DVDRip":   {"DVDR", "DVD-Full", "Full-Rip", "DVD-5", "DVD-9", "DVDRip", "DVDMux", "DVD5", "DVD9"},
		"HDRip":    {"HDRip", "WEB-DLRip"},
		"SCR":      {"SCR", "SCREENER", "DVDSCR", "DVDSCREENER", "BDSCR"},
		"CAM":      {"CAM", "CAM-Rip", "HDCAM"},
		"TELESYNC": {"TELESYNC", "TS", "HDTS", "PDVD", "PreDVDRip"},
	}
	ResolutionTypes = map[string][]string{
		"352p":  {"352p"},
		"360p":  {"360p"},
		"480i":  {"480i"},
		"480P":  {"480P"},
		"576i":  {"576i"},
		"576p":  {"576p"},
		"720P":  {"720P"},
		"1080P": {"1080P"},
		"2160P": {"2160P"},
	}
	EncodingTypes = map[string][]string{
		"Xvid": {"Xvid", "DivX"},
		"x264": {"x264", "H.264", "AVC"},
		"x265": {"x265", "HEVC", "H.265"},
		"VP9":  {"VP9"},
	}
)
