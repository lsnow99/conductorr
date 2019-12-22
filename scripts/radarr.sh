#!/bin/bash

torrent_linker_id="$radarr_download_id"
nzb_linker_id="$radarr_release_title"
xarr_id="$radarr_movie_id"
imdb_id="$radarr_movie_imdbid"
content_name="$radarr_movie_title"
grabbed_quality="$radarr_release_quality"
grabbed_size="$radarr_release_size"

if [ "$torrent_linker_id" == "" ]; then
    torrent_linker_id="N/A"
fi

curl --header "Content-Type: application/json" \
  --request POST \
  --data "{\"content_id\":$xarr_id,\"content_name\":\"$content_name\",\"content_category\":\"Series\",\"torrent_linker_id\":\"$torrent_linker_id\",\"nzb_linker_id\":\"$nzb_linker_id\",\"grabbed_quality\":\"$grabbed_quality\",\"grabbed_size\":\"$grabbed_size\",\"imdb_id\":\"$imdb_id\"}" \
  http://conductorr.plex.svc.cluster.local/_link