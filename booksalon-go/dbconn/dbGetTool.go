package dbconn

// GetUserByPwd login user return User
func GetUserByPwd(act *UserAccount) (retUser User, err error) {
	var tt UserAccount

	if err = db.First(&tt, "account = ? AND password = ?", act.Account, act.Password).Error; err != nil {
		return
	}
	//fmt.Println(db.Where("account = ? AND password = ?", act.Account, act.Password).First(&tt).Value)

	if err = db.Model(&tt).Related(&retUser, "account_id").Error; err != nil {
		return
	}
	return
}

// GetUsers get all user
func GetUsers() (users []User, err error) {
	if err = db.Find(&users).Error; err != nil {
		return
	}
	// fmt.Println(users[0].Name)
	// var retUsers []userCanShow
	// for _, user := range users {
	// 	println(user.userCanShow.Name, user.userCanShow.Teams)
	// 	retUsers = append(retUsers, user.userCanShow)
	// 	fmt.Println(retUsers)
	// }

	return users, nil
}

// GetUserTeams get user's all team
func GetUserTeams(userid string) (retTeams []Team, err error) {
	var user User

	db.First(&user, "id = ?", userid)
	db.Model(&user).Related(&retTeams, "Teams").Find(&retTeams)
	return
}

// GetUserObjByID get user obj by userid
func GetUserObjByID(userid string) (user User, err error) {
	if err = db.First(&user, "id = ?", userid).Error; err != nil {
		return
	}
	return
}

// GetUserTeamObj get user team obj by userid and teamid
func GetUserTeamObj(userid, teamid string) (retUser User, retTeam Team, err error) {
	if retUser, err = GetUserObjByID(userid); err != nil {
		return
	}

	if err = db.First(&retTeam, "id = ? AND leader_id = ?", teamid, retUser.ID).Error; err != nil {
		return
	}
	return
}

// GetTeamObjByID get user team obj by userid and teamid
func GetTeamObjByID(teamid string) (retTeam Team, err error) {
	if err = db.Where("id = ?", teamid).First(&retTeam).Error; err != nil {
		return
	}
	return
}

// GetTeamMember get team all member
func GetTeamMember(userid, teamid string) (members []User, err error) {
	if _, team, err := GetUserTeamObj(userid, teamid); err != nil {
		return nil, err
	} else if err := db.Model(&team).Related(&members, "Users").Error; err != nil {
		return nil, err
	}

	return members, err
}
