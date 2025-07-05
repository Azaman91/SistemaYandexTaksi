package main

import (
	"errors"
	"fmt"
	"sort"
)

type Sotrudnic struct {
	name       string
	position   string
	salary     int
	experience int
}
type CompanyInterface interface {
	AddWorkerInfo(name, position string, salary, experience uint) error
	SortWorkers() ([]string, error)
}

func (s *Sotrudnic) AddWorkerInfo(name, position string, salary, experience uint) error {
	if name == "" {
		return errors.New("Error")
	}
	if position == "" {
		return errors.New("Error")
	}
	if salary <= 0 {
		return errors.New("Error")
	}
	if experience < 0 {
		return errors.New("Error")
	}
	s.name = name
	s.position = position
	s.salary = int(salary)
	s.experience = int(experience)
	return nil
}

func (s Sotrudnic) SortWorkers(sots []Sotrudnic) ([]string, error) {
	position := map[string]int{
		"директор":       5,
		"зам. директора": 4,
		"начальник цеха": 3,
		"мастер":         2,
		"рабочий":        1,
	}
	sort.Slice(sots, func(i, j int) bool {
		if sots[i].salary != sots[i].salary {
			return sots[i].salary > sots[j].salary
		}
		return position[sots[i].position] > position[sots[j].position]
	})
	var result []string
	for _, num := range sots {
		result = append(result, fmt.Sprintf("%s-%d-%s", num.name, num.salary, num.experience))
	}
	return result, nil
}
