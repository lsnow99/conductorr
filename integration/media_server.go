package integration

type MediaServer interface {
	/*
	ImportMedia imports media from a filepath into the media server
	*/
	ImportMedia(string) error
}