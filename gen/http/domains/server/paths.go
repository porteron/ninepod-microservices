// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// HTTP request path constructors for the domains service.
//
// Command:
// $ goa gen ninepod/design

package server

import (
	"fmt"
)

// DomainsDomainsPath returns the URL path to the domains service domains HTTP endpoint.
func DomainsDomainsPath(a int, b string) string {
	return fmt.Sprintf("/domains/%v/%v", a, b)
}
