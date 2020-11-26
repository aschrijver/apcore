// apcore is a server framework for implementing an ActivityPub application.
// Copyright (C) 2019 Cory Slep
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package services

type NodeInfoStats struct {
	TotalUsers     int
	ActiveHalfYear int
	ActiveMonth    int
	ActiveWeek     int
	NLocalPosts    int
	NLocalComments int
}

type ServerProfile struct {
	OpenRegistrations bool
	ServerBaseURL     string
	ServerName        string
	OrgName           string
	OrgContact        string
	OrgAccount        string
}

type NodeInfo struct{}

func (n *NodeInfo) GetAnonymizedStats() (t NodeInfoStats, err error) {
	// TODO
	return
}

func (n *NodeInfo) GetServerProfile() (p ServerProfile, err error) {
	// TODO
	return
}
