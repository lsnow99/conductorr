#!/bin/bash

ARG_PATH="$1"
ARG_LABEL="$2"
ARG_NAME="$3"
ARG_HASH="$4"

if [ "$ARG_LABEL" = "Movies" ] || [ "$ARG_LABEL" = "Series" ]; then
    curl --header "Content-Type: application/json" \
	--max-time 900 \
        --request POST \
        --data "{\"dl_dir\":\"$ARG_PATH\",\"dl_content_id\":\"$ARG_HASH\",\"content_category\":\"$ARG_LABEL\",\"content_name\":\"$ARG_NAME\",\"dl_client_id\":\"rTorrent\"}" \
        http://conductorr.plex.svc.cluster.local/_import
else
	if [ "$ARG_LABEL" = "Books" ]; then
		#mv "$ARG_PATH" "/valinor/plex/books/$ARG_NAME"
		echo "book"
	else
    		echo "Not a movie or show"
	fi
fi