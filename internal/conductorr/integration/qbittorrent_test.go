package integration

import "testing"

func TestConnect(t *testing.T) {
	q, err := NewQBittorrent("admin", "adminadmin", "http://localhost:8080")
	if err != nil {
		t.Fatal(err)
	}
	if err := q.TestConnection(); err != nil {
		t.Fatal(err)
	}
}

func TestAddToqrrent(t *testing.T) {
	q, err := NewQBittorrent("admin", "adminadmin", "http://localhost:8080")
	if err != nil {
		t.Fatal(err)
	}
	str, err := q.AddRelease(Release{
		DownloadURL: "https://jackett.mediaserver.pro:443/dl/iptorrents/?jackett_apikey=8ao88icd8nl1i0amfv62ekb992g4kokm&path=Q2ZESjhQdFhNdFlvcU4xQXUwVXJFcjB6cHRhcFVtejY1emFWWFJHVGIwMkxXV284dGFEdl9TWEotWVBwUmNLazhqSDhOUmZIbnNsQVFfbF9ITmx0bGdfODVMNW9fd3hPQktMenh0a1dVRUJPNFppOEZITmlpQ3NwR2RaX0Z3NTVhMzkzRHZKYTN1bGdjclFWZlVabERDeEFCLWdtVWhadjNDV2xzclQyM2VETGh4bFFmQjh1eEZoNG4tdTFET0VvYmxJZkFya1VmYXhXR1IxUEFmU0xiMVBSRUo1cGlRdVVWeEFkRUNNVU5fWGtMYi1OS1NPc3VpdkVtZDlvTV9qRVlVMm9CX2xLeFp0N1EzT3NkLWlDLUl3ZlMzMTNYZHU5LWtlNmducjJQQndwNE1rTw&file=Family+or+Fiance+S02E09+Lakesha+and+JaQuan+Babies+and+Baggage+720p+HDTV+x264-CRiMSON",
	})
	t.Logf("%s: %v", str, err)
}
