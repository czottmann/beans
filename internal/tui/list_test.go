package tui

import (
	"testing"

	"hmans.dev/beans/internal/bean"
)

func TestSortBeans(t *testing.T) {
	// Define the expected order from DefaultStatuses and DefaultTypes
	statusNames := []string{"backlog", "todo", "in-progress", "completed", "scrapped"}
	typeNames := []string{"milestone", "epic", "bug", "feature", "task"}

	t.Run("sorts by status order first", func(t *testing.T) {
		beans := []*bean.Bean{
			{ID: "1", Status: "completed", Type: "task", Title: "A"},
			{ID: "2", Status: "backlog", Type: "task", Title: "B"},
			{ID: "3", Status: "in-progress", Type: "task", Title: "C"},
			{ID: "4", Status: "todo", Type: "task", Title: "D"},
			{ID: "5", Status: "scrapped", Type: "task", Title: "E"},
		}

		sortBeans(beans, statusNames, typeNames)

		expected := []string{"backlog", "todo", "in-progress", "completed", "scrapped"}
		for i, want := range expected {
			if beans[i].Status != want {
				t.Errorf("index %d: got status %q, want %q", i, beans[i].Status, want)
			}
		}
	})

	t.Run("sorts by type order within same status", func(t *testing.T) {
		beans := []*bean.Bean{
			{ID: "1", Status: "todo", Type: "task", Title: "A"},
			{ID: "2", Status: "todo", Type: "milestone", Title: "B"},
			{ID: "3", Status: "todo", Type: "bug", Title: "C"},
			{ID: "4", Status: "todo", Type: "epic", Title: "D"},
			{ID: "5", Status: "todo", Type: "feature", Title: "E"},
		}

		sortBeans(beans, statusNames, typeNames)

		expected := []string{"milestone", "epic", "bug", "feature", "task"}
		for i, want := range expected {
			if beans[i].Type != want {
				t.Errorf("index %d: got type %q, want %q", i, beans[i].Type, want)
			}
		}
	})

	t.Run("sorts by title within same status and type", func(t *testing.T) {
		beans := []*bean.Bean{
			{ID: "1", Status: "todo", Type: "task", Title: "Zebra"},
			{ID: "2", Status: "todo", Type: "task", Title: "Apple"},
			{ID: "3", Status: "todo", Type: "task", Title: "Mango"},
		}

		sortBeans(beans, statusNames, typeNames)

		expected := []string{"Apple", "Mango", "Zebra"}
		for i, want := range expected {
			if beans[i].Title != want {
				t.Errorf("index %d: got title %q, want %q", i, beans[i].Title, want)
			}
		}
	})

	t.Run("title sort is case-insensitive", func(t *testing.T) {
		beans := []*bean.Bean{
			{ID: "1", Status: "todo", Type: "task", Title: "zebra"},
			{ID: "2", Status: "todo", Type: "task", Title: "Apple"},
			{ID: "3", Status: "todo", Type: "task", Title: "MANGO"},
		}

		sortBeans(beans, statusNames, typeNames)

		expected := []string{"Apple", "MANGO", "zebra"}
		for i, want := range expected {
			if beans[i].Title != want {
				t.Errorf("index %d: got title %q, want %q", i, beans[i].Title, want)
			}
		}
	})

	t.Run("combined sort order: status > type > title", func(t *testing.T) {
		beans := []*bean.Bean{
			{ID: "1", Status: "completed", Type: "bug", Title: "Z"},
			{ID: "2", Status: "todo", Type: "task", Title: "A"},
			{ID: "3", Status: "todo", Type: "bug", Title: "B"},
			{ID: "4", Status: "todo", Type: "bug", Title: "A"},
			{ID: "5", Status: "backlog", Type: "epic", Title: "X"},
		}

		sortBeans(beans, statusNames, typeNames)

		// Expected order:
		// 1. backlog/epic/X
		// 2. todo/bug/A
		// 3. todo/bug/B
		// 4. todo/task/A
		// 5. completed/bug/Z
		expectedIDs := []string{"5", "4", "3", "2", "1"}
		for i, want := range expectedIDs {
			if beans[i].ID != want {
				t.Errorf("index %d: got ID %q, want %q (status=%s, type=%s, title=%s)",
					i, beans[i].ID, want, beans[i].Status, beans[i].Type, beans[i].Title)
			}
		}
	})

	t.Run("unrecognized status sorts last", func(t *testing.T) {
		beans := []*bean.Bean{
			{ID: "1", Status: "unknown", Type: "task", Title: "A"},
			{ID: "2", Status: "todo", Type: "task", Title: "B"},
			{ID: "3", Status: "backlog", Type: "task", Title: "C"},
		}

		sortBeans(beans, statusNames, typeNames)

		// unknown status should be last
		if beans[2].Status != "unknown" {
			t.Errorf("unrecognized status should be last, got %q at position 2", beans[2].Status)
		}
	})

	t.Run("unrecognized type sorts last within status", func(t *testing.T) {
		beans := []*bean.Bean{
			{ID: "1", Status: "todo", Type: "unknown", Title: "A"},
			{ID: "2", Status: "todo", Type: "task", Title: "B"},
			{ID: "3", Status: "todo", Type: "bug", Title: "C"},
		}

		sortBeans(beans, statusNames, typeNames)

		// unknown type should be last within todo status
		if beans[2].Type != "unknown" {
			t.Errorf("unrecognized type should be last, got %q at position 2", beans[2].Type)
		}
	})

	t.Run("empty slice does not panic", func(t *testing.T) {
		beans := []*bean.Bean{}
		sortBeans(beans, statusNames, typeNames)
		// No assertion needed, just checking it doesn't panic
	})

	t.Run("single bean does not panic", func(t *testing.T) {
		beans := []*bean.Bean{
			{ID: "1", Status: "todo", Type: "task", Title: "A"},
		}
		sortBeans(beans, statusNames, typeNames)
		if beans[0].ID != "1" {
			t.Error("single bean should remain unchanged")
		}
	})
}
