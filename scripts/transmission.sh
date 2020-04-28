#!/bin/bash

ARG_PATH="$TR_TORRENT_DIR"
ARG_NAME="$TR_TORRENT_NAME"
ARG_HASH="$TR_TORRENT_HASH"

if [ "$ARG_LABEL" = "Movies" ] || [ "$ARG_LABEL" = "Series" ]; then
    curl --header "Content-Type: application/json" \
	--max-time 900 \
        --request POST \
        --data "{\"dl_dir\":\"$ARG_PATH\",\"dl_content_id\":\"$ARG_HASH\",\"content_name\":\"$ARG_NAME\",\"dl_client_id\":\"rTorrent\"}" \
        http://conductorr.plex.svc.cluster.local/_import
else
	if [ "$ARG_LABEL" = "Books" ]; then
		#mv "$ARG_PATH" "/valinor/plex/books/$ARG_NAME"
		echo "book"
	else
    		echo "Not a movie or show"
	fi
fi
