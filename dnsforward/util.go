package dnsforward

import (
	"sort"
	"strings"

	"github.com/AdguardTeam/golibs/utils"
)

func stringArrayDup(a []string) []string {
	a2 := make([]string, len(a))
	copy(a2, a)
	return a2
}

// Find value in a sorted array
func findSorted(ar []string, val string) int {
	i := sort.SearchStrings(ar, val)
	if i == len(ar) || ar[i] != val {
		return -1
	}
	return i
}

func isWildcard(host string) bool {
	return len(host) >= 2 &&
		host[0] == '*' && host[1] == '.'
}

// Return TRUE if host name matches a wildcard pattern
func matchDomainWildcard(host, wildcard string) bool {
	return isWildcard(wildcard) &&
		strings.HasSuffix(host, wildcard[1:])
}

// Return TRUE if client's SNI value matches DNS names from certificate
func matchDNSName(dnsNames []string, sni string) bool {
	if utils.IsValidHostname(sni) != nil {
		return false
	}
	if findSorted(dnsNames, sni) != -1 {
		return true
	}

	for _, dn := range dnsNames {
		if matchDomainWildcard(sni, dn) {
			return true
		}
	}
	return false
}
