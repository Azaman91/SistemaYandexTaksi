
package main

import (
	"testing"
	"slices"
)

func TestWorkerSort(t *testing.T) {

	type Worker struct {
		Name            string
		Position        string
		Salary          uint
		ExperienceYears uint
	}

	tests := []struct {
		workers  []Worker
		expected []string
	}{
		{
			workers: []Worker{
				{Name: "Михаил", Position: "директор", Salary: 200, ExperienceYears: 5},
				{Name: "Игорь", Position: "зам. директора", Salary: 180, ExperienceYears: 3},
				{Name: "Николай", Position: "начальник цеха", Salary: 120, ExperienceYears: 2},
				{Name: "Андрей", Position: "мастер", Salary: 90, ExperienceYears: 10},
				{Name: "Виктор", Position: "рабочий", Salary: 80, ExperienceYears: 3},
			},
			expected: []string{
				"Михаил — 12000 — директор",
				"Андрей — 10800 — мастер",
				"Игорь — 6480 — зам. директора",
				"Николай — 2880 — начальник цеха",
				"Виктор — 2880 — рабочий",
			},
		},
		{
			workers: []Worker{
				{Name: "Андрей", Position: "директор", Salary: 200, ExperienceYears: 1},
				{Name: "Максим",... File is too long to be displayed fully}
			}
		}
	}
}