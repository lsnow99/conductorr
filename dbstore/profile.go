package dbstore

func CreateProfile(name string) error {
	_, err := db.Exec(`INSERT INTO profile (name) VALUES (?)`, name)
	return err
}

func GetProfiles() ([]*Profile, error) {
	profiles := make([]*Profile, 0)

	rows, err := db.Query(`SELECT id, name, filter, sorter FROM profile;`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		profile := Profile{}
		err := rows.Scan(&profile.ID, &profile.Name, &profile.Filter, &profile.Sorter)
		if err != nil {
			return nil, err
		}
		profiles = append(profiles, &profile)
	}

	return profiles, rows.Err()
}

func GetProfileByID(id int) (Profile, error) {
	row := db.QueryRow(`SELECT id, name, filter, sorter FROM profile WHERE id = ?`, id)

	profile := Profile{}
	err := row.Scan(&profile.ID, &profile.Name, &profile.Filter, &profile.Sorter)

	return profile, err
}

func UpdateProfile(id int, name, filter, sorter string) error {
	_, err := db.Exec(`UPDATE profile SET name = ?, filter = ?, sorter = ? WHERE id = ?`, name, filter, sorter, id)
	return err
}

func DeleteProfile(id int) error {
	_, err := db.Exec(`DELETE FROM profile WHERE id = ?`, id)
	return err
}