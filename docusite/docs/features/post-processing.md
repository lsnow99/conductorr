# Post-Processing

Conductorr post-processes releases by scanning the destination folder/file path given to it by the downloader and performs the following:

- Recursively searches for `.rar`, `.tar.gz`, `.zip`, and `.7z` files and extracts them into a temporary folder.
- Recursively locates the largest `.mkv` or `.mp4` file in the folders.
- If no file can be found, the download is marked as an error, and if any other releases are pending in the queue for the requested media, the next release in the queue will be attempted.
- Conductorr selects the largest file among the `.mkv` and `.mp4` files, and moves or copies it to the media's designated path within its designated library folder.
- Conductorr renames the file to conform to the settings specified in the Post-Processing configuration tab.

<!-- For custom preprocessing, Conductorr's [integration]() system can be used. -->
