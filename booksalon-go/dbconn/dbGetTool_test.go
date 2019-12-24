package dbconn

import (
	"reflect"
	"testing"
)

func TestGetTeamMember(t *testing.T) {
	type args struct {
		userid string
		teamid string
	}
	tests := []struct {
		name        string
		args        args
		wantMembers []User
		wantErr     bool
	}{
		{
			name: "jelech",
			args: args{"1", "1"},
			wantMembers: []User{
				{
					Name:  "jelech",
					Teams: []Team{},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMembers, err := GetTeamMember(tt.args.userid, tt.args.teamid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTeamMember() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotMembers, tt.wantMembers) {
				t.Errorf("GetTeamMember() = %v, want %v", gotMembers, tt.wantMembers)
			}
		})
	}
}
