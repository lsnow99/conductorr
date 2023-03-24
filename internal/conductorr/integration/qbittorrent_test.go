package integration

import "testing"

const username = "admin"
const password = "adminadmin"
const baseUrl = "http://localhost:6881"

func TestConnect(t *testing.T) {
	q, err := NewQBittorrent(username, password, baseUrl)
	if err != nil {
		t.Fatal(err)
	}
	if err := q.TestConnection(); err != nil {
		t.Fatal(err)
	}
}

func TestAddToqrrent(t *testing.T) {
	q, err := NewQBittorrent(username, password, baseUrl)
	if err != nil {
		t.Fatal(err)
	}
	str, err := q.AddRelease(Release{
		DownloadURL: "https://prowlarr.snowcloud.one/5/download?apikey=4915d4f0062f43ce8b38dcd6ff05445b&link=MzAwWUV4TWtHcSszKzFib2lHTlUrODZDZHJYZjJMZVJ0UUpoeHhtcjI2UCtHRE9VRzd1Tlh5MGpqRTFDMXJSR1d6cU94eCtzdVQweXN3bC9VQXlsMzBkU0F3S2xQamVLSlVGcDhUV3lYbHFzenNPYXljeWc3Tk1vMHhEeHJnR1lqM0loMDNNVHJOa29oeEo0VUsyYnVRMGkzQ1dydURIT09VM3FLTXNrYk9hN1RIQnFuNnEzUDNMVEt3R0xRUm5ZSDZYNUhMV1h2QktXWjZweGtLaGl0MmM2M20zWVNBUjRZb1RPbEdZUEp3RT0&file=American+Dad+S16E05+Jeff+and+the+Dank+Ass+Weed+Factory+1080p+AMZN+WEB-DL+DD%2B5+1+H+264-CtrlHD",
	})
  if err != nil {
    t.Fatal(err)
  }
	t.Logf("%s: %v", str, err)
}

func TestListTorrent(t *testing.T) {
	q, err := NewQBittorrent(username, password, baseUrl)
  if err != nil {
    t.Fatal(err)
  }
	str, err := q.AddRelease(Release{
		DownloadURL: "https://prowlarr.snowcloud.one/5/download?apikey=4915d4f0062f43ce8b38dcd6ff05445b&link=MzAwWUV4TWtHcSszKzFib2lHTlUrODZDZHJYZjJMZVJ0UUpoeHhtcjI2UCtHRE9VRzd1Tlh5MGpqRTFDMXJSR1d6cU94eCtzdVQweXN3bC9VQXlsMzBkU0F3S2xQamVLSlVGcDhUV3lYbHFzenNPYXljeWc3Tk1vMHhEeHJnR1lqM0loMDNNVHJOa29oeEo0VUsyYnVRMGkzQ1dydURIT09VM3FLTXNrYk9hN1RIQnFuNnEzUDNMVEt3R0xRUm5ZSDZYNUhMV1h2QktXWjZweGtLaGl0MmM2M20zWVNBUjRZb1RPbEdZUEp3RT0&file=American+Dad+S16E05+Jeff+and+the+Dank+Ass+Weed+Factory+1080p+AMZN+WEB-DL+DD%2B5+1+H+264-CtrlHD",
	})
  if err != nil {
    t.Fatal(err)
  }
  downloads, err := q.PollDownloads([]string{str})
  if err != nil {
    t.Fatal(err)
  }
  t.Logf("downloads: %v", downloads)
}

