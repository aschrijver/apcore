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

package app

import (
	"net/http"
	"net/url"

	"github.com/go-fed/activity/streams/vocab"
	"github.com/go-fed/apcore/paths"
	"github.com/go-fed/apcore/util"
)

// Framework provides request-time hooks for use in handlers.
type Framework interface {
	Context(r *http.Request) util.Context

	UserIRI(userUUID paths.UUID) *url.URL

	// Validate attempts to obtain and validate the OAuth token or first
	// party credential in the request. This can be called in your handlers
	// at request-handing time.
	//
	// If an error is returned, both the token and authentication values
	// should be ignored.
	//
	// TODO: Scopes
	Validate(w http.ResponseWriter, r *http.Request) (userID paths.UUID, authenticated bool, err error)

	// Send will send an Activity or Object on behalf of the user.
	//
	// Note that a new ID is not needed on the activity and/or objects that
	// are being sent; they will be generated as needed.
	Send(c util.Context, userID paths.UUID, toSend vocab.Type) error

	Session(r *http.Request) (Session, error)

	// TODO: Determine if we need this.
	GetByIRI(c util.Context, id *url.URL) (vocab.Type, error)
}

type Session interface {
	UserID() (string, error)
}
