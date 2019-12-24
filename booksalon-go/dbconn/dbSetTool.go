package dbconn

// CreateUser create a new user
func CreateUser(user *User) (retUser User, err error) {
	tx := db.Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()

	if err = tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()
	return *user, nil
}

// CreateTeam create a new team
func CreateTeam(userID string, team *Team) (retTeam Team, err error) {
	tx := db.Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()

	if err = tx.Create(team).Error; err != nil {
		tx.Rollback()
		return
	}

	var user User
	if tx.First(&user, "id = ?", userID).RecordNotFound() {
		tx.Rollback()
		err = tx.Error
		return
	}
	if err = tx.Model(&user).Association("Teams").Append(team).Error; err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()
	return *team, nil
}

// UpdateTeam update a team by teamid
func UpdateTeam(userID, teamID string, team *Team) (retTeam Team, err error) {
	var user User
	if user, err = GetUserObjByID(userID); err != nil {
		return
	}

	tx := db.Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()

	// TODO:there is a error. teamid useless, didn't change old team.
	if err = tx.Model(user).Association("Teams").Replace(team).Error; err != nil {

		tx.Rollback()
		return
	}

	tx.Commit()
	return *team, nil
}

// AddTeamMember add a new team member
func AddTeamMember(userID, teamID, addUserID string) (retUser User, err error) {
	var team Team
	if _, team, err = GetUserTeamObj(userID, teamID); err != nil {
		return
	}

	if retUser, err = GetUserObjByID(addUserID); err != nil {
		return
	}

	tx := db.Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()

	// TODO: ERROR!!!! table inster userid is worn
	if err = tx.Model(&team).Association("Users").Append(&retUser).Error; err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()
	return
}

// DelTeamMember del a team member
func DelTeamMember(userID, teamID, delUserID string) (err error) {
	var team Team
	if _, team, err = GetUserTeamObj(userID, teamID); err != nil {
		return
	}

	var delUser User // TODO: using userID search and del
	tx := db.Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()
	if err = tx.Model(&team).Association("Users").Delete(&delUser).Error; err != nil {

		tx.Rollback()
		return
	}

	tx.Commit()
	return
}

// DelTeam del a team
func DelTeam(userID, delTeamID string) (err error) {
	var team Team

	if _, team, err = GetUserTeamObj(userID, delTeamID); err != nil {
		return
	}

	tx := db.Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()

	if err = tx.Unscoped().Delete(&team).Error; err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()
	return
}

// DelUser del a user
func DelUser(userID string) (retUser User, err error) {

	tx := db.Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()

	if db.Where("id = ?", userID).Unscoped().Delete(&retUser).RecordNotFound() {
		tx.Rollback()
		return
	}

	tx.Commit()
	return
}
