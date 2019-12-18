# Conductorr
Kubernetes Plex Download Orchestration

## What is it?
Conductorr sits as a deployment in your Kubernetes cluster along with Sonarr, Radarr, Filebot, NZBGet/rTorrent, and Plex. Conductorr is the glue between processing downloads from Radarr/Sonarr and running them through Filebot, and finally triggering a Plex Scan on the item. Condcutorr comes with a beautiful web interface to view logs and other information about your downloads so you can find issues in your own pipeline.

## Why is it Needed?
Running a Plex Server and downloaders in Kubernetes can have its issues, especially with Filebot. For example, after Filebot processes a download, Radarr/Sonarr might disagree with Filebot on the output directory, and if Sonarr/Radarr are not notified of the new path, they may mark the download as a fail and keep trying more downloads indefinitely. Conductorr notifies Sonarr/Radarr of the change to prevent them from continuing to download. Another issue solved by Conductorr is the issue of Plex not noticing new content on certain filesystems (NFS, Ceph, etc). One solution is to use Filebot's built-in Plex scanning feature, but this will scan your entire library, and not just the single file you want it to scan.

## Installation Guide
Please refer to the installation guide in the Wiki. Conductorr requires a significant amount of setup, but the documentation aims to be as straighforward as possible.
