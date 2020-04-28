package main

import "github.com/lsnow2017/conductorr/internal/schema"

// LoadDownloaderConfigurations from the database
func LoadDownloaderConfigurations() []schema.DownloaderConfiguration {
	configs := []schema.DownloaderConfiguration{}
	if err := db.Model(&configs).Select(); err != nil {
		panic(err)
	}
	return configs
}

// SaveDownloaderConfigurations to the database
func SaveDownloaderConfigurations(configs []schema.DownloaderConfiguration) {
	if _, err := db.Model((*schema.DownloaderConfiguration)(nil)).Exec(`
		DELETE FROM downloader_configuration WHERE 1=1;
	`); err != nil {
		panic(err)
	}
	if _, err := db.Model(&configs).OnConflict("(downloader_configuration_id) DO UPDATE").Set("name = EXCLUDED.name").Set("download_dir = EXCLUDED.download_dir").Insert(); err != nil {
		panic(err)
	}
}

func GetDownloaderByName(name string) schema.DownloaderConfiguration {
	config := schema.DownloaderConfiguration{}
	config.Name = name
	if err := db.Model(&config).Where("name = ?", name).Limit(1).Select(); err != nil {
		panic(err)
	}
	return config
}