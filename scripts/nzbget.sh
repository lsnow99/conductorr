#!/bin/sh 

###########################################
### NZBGET POST-PROCESSING SCRIPT       ###

# A post processing script to run on 
# completion of all media grabbed from 
# Sonarr/Radarr
#
# The script calls /import

################################
### OPTIONS                  ###

# fbserver URL (ignore this setting for now, edit in script instead)
#Server=http://conductorr.plex.svc.cluster.local/_import

### NZBGET POST-PROCESSING SCRIPT       ###
###########################################

if [[ $NZBPP_TOTALSTATUS != "SUCCESS" ]]; then
	exit 93
fi

curl --header "Content-Type: application/json" \
  --max-time 900 \
  --request POST \
  --data "{\"dl_dir\":\"$NZBPP_DIRECTORY\",\"dl_content_id\":\"$NZBPP_NZBNAME\",\"content_category\":\"$NZBPP_CATEGORY\",\"content_name\":\"$NZBPP_NZBFILENAME\",\"dl_client_id\":\"NZBGet\"}" \
  http://conductorr.plex.svc.cluster.local/_import

exit 93