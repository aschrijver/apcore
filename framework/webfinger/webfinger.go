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

package webfinger

import (
	"fmt"
)

type Link struct {
	Rel      string `json:"rel,omitempty"`
	Type     string `json:"type,omitempty"`
	Href     string `json:"href,omitempty"`
	Template string `json:"template,omitempty"`
}

type Webfinger struct {
	Subject string   `json:"subject,omitempty"`
	Aliases []string `json:"aliases,omitempty"`
	Links   []Link   `json:"links,omitempty"`
}

func ToWebfinger(scheme, host, username, idPath string) (w Webfinger, err error) {
	w = Webfinger{
		Subject: fmt.Sprintf("acct:%s@%s", username, host),
		Aliases: []string{
			fmt.Sprintf("%s://%s%s", scheme, host, idPath),
		},
		Links: []Link{
			{
				Rel:  "self",
				Type: "application/activity+json",
				Href: fmt.Sprintf("%s://%s%s", scheme, host, idPath),
			},
		},
	}
	return
}
