package bean

import (
	"sort"
	"strings"
)

// SortByStatusAndType sorts beans by status order, then by type order, then by title.
// This is the default sorting used by both CLI and TUI.
// Unrecognized statuses and types are sorted last within their category.
func SortByStatusAndType(beans []*Bean, statusNames []string, typeNames []string) {
	statusOrder := make(map[string]int)
	for i, s := range statusNames {
		statusOrder[s] = i
	}
	typeOrder := make(map[string]int)
	for i, t := range typeNames {
		typeOrder[t] = i
	}

	// Helper to get order with unrecognized values sorted last
	getStatusOrder := func(status string) int {
		if order, ok := statusOrder[status]; ok {
			return order
		}
		return len(statusNames) // Unrecognized statuses come last
	}
	getTypeOrder := func(typ string) int {
		if order, ok := typeOrder[typ]; ok {
			return order
		}
		return len(typeNames) // Unrecognized types come last
	}

	sort.Slice(beans, func(i, j int) bool {
		// Primary: status order
		oi, oj := getStatusOrder(beans[i].Status), getStatusOrder(beans[j].Status)
		if oi != oj {
			return oi < oj
		}
		// Secondary: type order
		ti, tj := getTypeOrder(beans[i].Type), getTypeOrder(beans[j].Type)
		if ti != tj {
			return ti < tj
		}
		// Tertiary: title (case-insensitive) for stable, user-friendly ordering
		return strings.ToLower(beans[i].Title) < strings.ToLower(beans[j].Title)
	})
}
