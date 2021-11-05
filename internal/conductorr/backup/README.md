# Conductorr Backup

Conductorr has builtin functionality to create backups that can be used to transfer Conductorr to a new instance.

## Guarantees

Conductorr's backup system is guaranteed to successfully restore from backups that were made on the same version of Conductorr, but backup files made on different versions of Conductorr are not guaranteed to successfully or fully restore.

Therefore, if you wish to restore from a backup created on an older version of Conductorr, it is recommended to do the following:

- Uninstall the newer version of Conductorr
- Install the version of Conductorr that matches the backup file
- Upload the backup file through the interface
- Upgrade Conductorr to the newer version

Conductorr's database migration system will handle the data manipulation needed to bring the database up to spec.
