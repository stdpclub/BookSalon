package dbconn

// CreateUser create a new user
func CreateUser(loginInfo *LoginInfo) (retUser User, err error) {
	tx := db.Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	user := User{Name: loginInfo.Name}
	userAccount := UserAccount{
		Account:  loginInfo.Account,
		Password: loginInfo.Password,
	}

	if err = tx.Create(&userAccount).Error; err != nil {
		return // tx.Rollback()
	}

	user.AccountID = userAccount.ID
	user.UserAccount = userAccount
	if err = tx.Create(&user).Error; err != nil {
		return // tx.Rollback()
	}

	retUser = user // 暂时放在这里，由于并不知道是应该返回多少东西。而账户和密码一般来说不应该返回
	retUser.UserAccount = UserAccount{}
	return // tx.Commit()
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
	if err = tx.First(&user, "id = ?", userID).Error; err != nil {
		tx.Rollback()
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
	if user, retTeam, err = GetUserTeamObj(userID, teamID); err != nil {
		return
	}

	tx := db.Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()

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
	if err = db.First(&delUser, "ID = ?", delUserID).Error; err != nil {
		return
	}

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
	if err = db.First(&retUser, userID).Error; err != nil {
		return
	}

	tx := db.Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()

	// TODO: 怎么简化
	if err = db.Where("id = ?", retUser.AccountID).Unscoped().Delete(&UserAccount{}).Error; err != nil {
		tx.Rollback()
		return
	}
	if err = db.Where("id = ?", userID).Unscoped().Delete(&retUser).Error; err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()
	return
}
