package internalsetup

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

type SetupOrganization struct {
	Name string
	Slug string
}

type SetupUser struct {
	Email            string
	Password         string
	Roles            []string
	OrganizationSlug string
}

type InternalSetupConfig struct {
	Enabled            bool
	NamedOrganizations []SetupOrganization
	NamedUsers         []SetupUser
}

func (s *InternalSetupConfig) Check() error {
	return nil
}

func (s *InternalSetupConfig) SetupOrganizations() error {
	return nil
}

func (s *InternalSetupConfig) SetupUsers() error {
	return nil
}

func NewInternalSetupConfig(ctx context.Context) *InternalSetupConfig {
	internalConfig := g.Cfg().MustGet(ctx, "internalSetup").MapDeep()
	enabled := false
	if v, ok := internalConfig["enabled"].(bool); ok {
		enabled = v
	}
	return &InternalSetupConfig{
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
				Slug: toString(m["slug"]),
			})
		}
	case []interface{}:
		for _, e := range arr {
			if m, ok := e.(map[string]interface{}); ok {
				out = append(out, SetupOrganization{
					Name: toString(m["name"]),
					Slug: toString(m["slug"]),
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
		Email:            toString(m["email"]),
		Password:         toString(m["password"]),
		OrganizationSlug: toString(m["organizationSlug"]),
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
