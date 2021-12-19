# Backups

Conductorr provides a backup system that allows you to create backups of the database that you can restore from in the future. These backups are useful if you accidently lose your database file or if you want to migrate your installation to a new database backend (ie. to transition from SQLite to PostgreSQL). 

## Creating a Backup

In Conductorr, navigate to the System tab and click the button to download a system backup. Depending on the size of your library and hardware, this may be instant or it may take a few seconds. When the backup is done being created, a zip file should download in your browser.

## Restoring a Backup

To restore a backup, you must have Conductorr running this can be a brand new instance or an existing one. If it is an existing one, all existing data will be wiped. In Conductorr, navigate to the system tab and click the button to restore from a backup. A window will appear with a warning and a button to select your backup file. After selecting your file click the restore button and the restoration will begin. 

## Caveats

### Restoring from a Different Version

Conductorr only allows restoration of backups that were created with the same version as the currently-running version of Conductorr. This is to ensure the guarantees that backup restoration provide. It may be the case that you wish to restore a backup that was made on an older version of Conductorr. To do this, you will need to temporarily downgrade your current instance to restore the backup. After the backup has been restored, you can upgrade Conductorr to the desired version.

## Automatic Backup Creation (advanced) 

Backups can be automated by using the `/api/v1/backup` API route. Details can be found in the (API documentation)[addlink].
