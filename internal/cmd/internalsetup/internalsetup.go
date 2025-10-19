package internalsetup

import (
	"context"
	"database/sql"
	"v1consortium/internal/model/do"
	"v1consortium/internal/service"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type SetupOrganization struct {
	Name string
	Id   string
	Type string // "internal" or "client", default to "client"
}

type SetupUser struct {
	Email          string
	Password       string
	Roles          []string
	FirstName      string
	LastName       string
	OrganizationId string
}

type InternalSetupConfig struct {
	InternalCtx        context.Context
	Enabled            bool
	NamedOrganizations []SetupOrganization
	NamedUsers         []SetupUser
}

func (s *InternalSetupConfig) Check() error {
	return nil
}

func (s *InternalSetupConfig) SetupOrganizations() error {
	g.Log().Infof(s.InternalCtx, "Setting up %d organizations", len(s.NamedOrganizations))
	for _, v := range s.NamedOrganizations {
		// Check if organization already exists
		org, err := service.OrganizationService().GetOrganization(s.InternalCtx, v.Id)
		if err != nil && !gerror.HasCode(err, gcode.CodeNotFound) && err != sql.ErrNoRows {
			g.Log().Errorf(s.InternalCtx, "Error checking organization %s: %v", v.Id, err)
			return err
		}

		// Create organization only if it doesn't exist
		if org == nil || org.Id == "" || err == sql.ErrNoRows || gerror.HasCode(err, gcode.CodeNotFound) {
			_, err := service.OrganizationService().CreateOrganization(s.InternalCtx, &do.Organizations{
				Id:   v.Id,
				Name: v.Name,
				Type: v.Type,
			})
			if err != nil {
				// Check if it's a duplicate key error (organization already exists)
				if gerror.HasCode(err, gcode.CodeValidationFailed) ||
					(err.Error() != "" && (err.Error() == "duplicate key value violates unique constraint" ||
						err.Error() == "organizations_pkey" ||
						err.Error() == "UNIQUE constraint failed")) {
					g.Log().Infof(s.InternalCtx, "Organization %s with ID %s already exists, skipping", v.Name, v.Id)
					continue
				}
				g.Log().Errorf(s.InternalCtx, "Failed to create organization %s: %v", v.Name, err)
				return err
			}
			g.Log().Infof(s.InternalCtx, "Created organization %s with ID %s", v.Name, v.Id)
		} else {
			g.Log().Infof(s.InternalCtx, "Organization %s with ID %s already exists, skipping", v.Name, v.Id)
		}
	}
	g.Log().Infof(s.InternalCtx, "Finished setting up %d organizations", len(s.NamedOrganizations))
	return nil
}

func (s *InternalSetupConfig) SetupUsers() error {
	g.Log().Infof(s.InternalCtx, "Setting up %d users", len(s.NamedUsers))

	for _, v := range s.NamedUsers {
		// Check if user already exists
		userProfile, err := service.Auth().GetUserProfileByEmail(s.InternalCtx, v.Email)

		// Handle real errors (not "user not found" errors)
		if err != nil && !gerror.HasCode(err, gcode.CodeNotFound) && err != sql.ErrNoRows {
			g.Log().Errorf(s.InternalCtx, "Error checking user %s: %v", v.Email, err)
			return err
		}

		// Determine if user exists
		userExists := userProfile != nil && userProfile.Id != ""
		g.Log().Debugf(s.InternalCtx, "User %s exists: %v", v.Email, userExists)

		if !userExists {
			role := "user"
			if len(v.Roles) > 0 {
				role = v.Roles[0]
			}

			// Try to register the user
			_, err := service.Auth().RegisterUser(s.InternalCtx, v.Email, v.Password, map[string]interface{}{
				"organization_id": v.OrganizationId,
				"role":            role,
				"first_name":      v.FirstName,
				"last_name":       v.LastName,
				"email_confirm":   true,
			})

			g.Log().Debugf(s.InternalCtx, "User %s registration response: %v", v.Email, err)

			if err != nil {
				// Check if it's a "user already exists" type error or "no rows" during registration
				if err == sql.ErrNoRows ||
					gerror.HasCode(err, gcode.CodeValidationFailed) ||
					(err.Error() != "" && (err.Error() == "user already exists" ||
						err.Error() == "duplicate key value violates unique constraint" ||
						err.Error() == "UNIQUE constraint failed")) {
					g.Log().Infof(s.InternalCtx, "User %s already exists or registration skipped, continuing", v.Email)
					continue
				}
				g.Log().Errorf(s.InternalCtx, "Failed to register user %s: %v", v.Email, err)
				return err
			}
			g.Log().Infof(s.InternalCtx, "Created user %s with roles %v in organization %s", v.Email, v.Roles, v.OrganizationId)
		} else {
			g.Log().Infof(s.InternalCtx, "User %s already exists, skipping creation", v.Email)
		}
	}
	g.Log().Infof(s.InternalCtx, "Finished setting up %d users", len(s.NamedUsers))
	return nil
}

func NewInternalSetupConfig(ctx context.Context) *InternalSetupConfig {
	internalConfig := g.Cfg().MustGet(ctx, "internalSetup").MapDeep()
	enabled := false
	if v, ok := internalConfig["enabled"].(bool); ok {
		enabled = v
	}
	return &InternalSetupConfig{
		InternalCtx:        ctx,
		Enabled:            enabled,
		NamedOrganizations: parseOrganizations(internalConfig["namedOrganizations"]),
		NamedUsers:         parseUsers(internalConfig["namedUsers"]),
	}
}

func parseOrganizations(v interface{}) []SetupOrganization {
	var out []SetupOrganization
	if v == nil {
		return out
	}
	switch arr := v.(type) {
	case []map[string]interface{}:
		for _, m := range arr {
			out = append(out, SetupOrganization{
				Name: toString(m["name"]),
				Id:   toString(m["id"]),
				Type: toString(m["type"]),
			})
		}
	case []interface{}:
		for _, e := range arr {
			if m, ok := e.(map[string]interface{}); ok {
				out = append(out, SetupOrganization{
					Name: toString(m["name"]),
					Id:   toString(m["id"]),
					Type: toString(m["type"]),
				})
			}
		}
	}
	return out
}

func parseUsers(v interface{}) []SetupUser {
	var out []SetupUser
	if v == nil {
		return out
	}
	switch arr := v.(type) {
	case []map[string]interface{}:
		for _, m := range arr {
			out = append(out, buildUserFromMap(m))
		}
	case []interface{}:
		for _, e := range arr {
			if m, ok := e.(map[string]interface{}); ok {
				out = append(out, buildUserFromMap(m))
			}
		}
	}
	return out
}

func buildUserFromMap(m map[string]interface{}) SetupUser {
	u := SetupUser{
		Email:          toString(m["email"]),
		Password:       toString(m["password"]),
		OrganizationId: toString(m["organizationId"]),
		FirstName:      toString(m["firstName"]),
		LastName:       toString(m["lastName"]),
	}
	// parse roles from various possible types
	switch r := m["roles"].(type) {
	case []interface{}:
		for _, ri := range r {
			if s := toString(ri); s != "" {
				u.Roles = append(u.Roles, s)
			}
		}
	case []string:
		u.Roles = append(u.Roles, r...)
	case string:
		if r != "" {
			u.Roles = append(u.Roles, r)
		}

	}
	return u
}

func toString(v interface{}) string {
	if v == nil {
		return ""
	}
	if s, ok := v.(string); ok {
		return s
	}
	return ""
}

func toBool(v interface{}) bool {
	if v == nil {
		return false
	}
	if b, ok := v.(bool); ok {
		return b
	}
	if s, ok := v.(string); ok {
		if s == "true" || s == "1" {
			return true
		}
		return false
	}
	return false
}
