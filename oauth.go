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

package apcore

import (
	"fmt"
	"net/http"
	"time"

	"gopkg.in/oauth2.v3"
	"gopkg.in/oauth2.v3/errors"
	"gopkg.in/oauth2.v3/manage"
	oaserver "gopkg.in/oauth2.v3/server"
)

type oAuth2Server struct {
	d *database
	k *sessions
	m *manage.Manager
	s *oaserver.Server
}

func newOAuth2Server(c *config, d *database, k *sessions) (s *oAuth2Server, err error) {
	m := manage.NewDefaultManager()
	// Configure Access token and Refresh token refresh.
	if c.OAuthConfig.AccessTokenExpiry <= 0 {
		err = fmt.Errorf("oauth2 access token expiration duration is <= 0")
		return
	} else if c.OAuthConfig.RefreshTokenExpiry <= 0 {
		err = fmt.Errorf("oauthr2 refresh token expiration duration is <= 0")
		return
	}
	m.SetAuthorizeCodeTokenCfg(&manage.Config{
		AccessTokenExp:    time.Second * time.Duration(c.OAuthConfig.AccessTokenExpiry),
		RefreshTokenExp:   time.Second * time.Duration(c.OAuthConfig.RefreshTokenExpiry),
		IsGenerateRefresh: true,
	})
	m.SetRefreshTokenCfg(&manage.RefreshingConfig{
		// Generate new refresh token
		IsGenerateRefresh: true,
		// Remove previous access token
		IsRemoveAccess: true,
		// Remove previous refreshing token
		IsRemoveRefreshing: true,
	})
	m.MapTokenStorage( /*TODO*/ nil)
	m.MapClientStorage( /*TODO*/ nil)
	// OAuth2 server
	srv := oaserver.NewServer(&oaserver.Config{
		TokenType: "Bearer",
		// Must follow the spec.
		AllowGetAccessRequest: false,
		// Support only the non-implicit flow.
		AllowedResponseTypes: []oauth2.ResponseType{oauth2.Code},
		// Allow:
		// - Authorization Code (for third parties)
		// - Refreshing Tokens
		// - Resource owner secrets
		// - Client secrets
		AllowedGrantTypes: []oauth2.GrantType{
			oauth2.AuthorizationCode,
			oauth2.Refreshing,
			oauth2.PasswordCredentials,
			oauth2.ClientCredentials,
		},
	}, m)
	// Parse tokens in POST body.
	srv.SetClientInfoHandler(oaserver.ClientFormHandler)
	// Determines the user to use when granting an authorization token.
	srv.SetUserAuthorizationHandler(func(w http.ResponseWriter, r *http.Request) (userID string, err error) {
		var s *session
		if s, err = k.Get(r); err != nil {
			return
		}
		if userID, err = s.UserID(); err != nil {
			return
		}
		return
	})
	// When granting an authorization token, overrides the scopes of incoming requests.
	srv.SetAuthorizeScopeHandler(func(w http.ResponseWriter, r *http.Request) (scope string, err error) {
		// TODO
		return
	})
	// Called when requesting a token through the password credential grant flow.
	srv.SetPasswordAuthorizationHandler(func(username, password string) (userID string, err error) {
		// TODO
		return
	})
	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		re = &errors.Response{
			Error:       errors.ErrServerError,
			ErrorCode:   http.StatusInternalServerError,
			Description: "Internal Error",
			StatusCode:  http.StatusInternalServerError,
		}
		return
	})
	srv.SetResponseErrorHandler(func(re *errors.Response) {
		ErrorLogger.Errorf("oauth2 response error: %s", re.Error.Error())
	})
	s = &oAuth2Server{
		d: d,
		k: k,
		m: m,
		s: srv,
	}
	return
}

// TODO: Call/expose this handler
func (o *oAuth2Server) HandleAuthorizationRequest(w http.ResponseWriter, r *http.Request) {
	if err := o.s.HandleAuthorizeRequest(w, r); err != nil {
		// oauth2 library would already have written headers by now.
		ErrorLogger.Errorf("oauth2 HandleAuthorizeRequest error: %s", err)
	}
}

// TODO: Call/expose this handler
func (o *oAuth2Server) HandleAccessTokenRequest(w http.ResponseWriter, r *http.Request) {
	if err := o.s.HandleTokenRequest(w, r); err != nil {
		// oauth2 library would already have written headers by now.
		ErrorLogger.Errorf("oauth2 HandleTokenRequest error: %s", err)
	}
}

// TODO: Call/expose this handler
func (o *oAuth2Server) ValidateOAuth2AccessToken(w http.ResponseWriter, r *http.Request) (token oauth2.TokenInfo, err error) {
	token, err = o.s.ValidationBearerToken(r)
	return
}
