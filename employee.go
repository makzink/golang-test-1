package main

import (
    "errors"
    "sync"
)

type Employee struct {
    ID       int     `json:"id"`
    Name     string  `json:"name"`
    Position string  `json:"position"`
    Salary   float64 `json:"salary"`
}

type EmployeeStore struct {
    mu        sync.Mutex
    employees map[int]Employee
    nextID    int
}

func NewEmployeeStore() *EmployeeStore {
    return &EmployeeStore{
        employees: make(map[int]Employee),
        nextID:    1,
    }
}

func (s *EmployeeStore) CreateEmployee(name, position string, salary float64) Employee {
    s.mu.Lock()
    defer s.mu.Unlock()
    employee := Employee{
        ID:       s.nextID,
        Name:     name,
        Position: position,
        Salary:   salary,
    }
    s.employees[s.nextID] = employee
    s.nextID++
    return employee
}

func (s *EmployeeStore) GetEmployeeByID(id int) (Employee, error) {
    s.mu.Lock()
    defer s.mu.Unlock()
    employee, exists := s.employees[id]
    if !exists {
        return Employee{}, errors.New("employee not found")
    }
    return employee, nil
}

func (s *EmployeeStore) UpdateEmployee(id int, name, position string, salary float64) (Employee, error) {
    s.mu.Lock()
    defer s.mu.Unlock()
    employee, exists := s.employees[id]
    if !exists {
        return Employee{}, errors.New("employee not found")
    }
    employee.Name = name
    employee.Position = position
    employee.Salary = salary
    s.employees[id] = employee
    return employee, nil
}

func (s *EmployeeStore) DeleteEmployee(id int) error {
    s.mu.Lock()
    defer s.mu.Unlock()
    if _, exists := s.employees[id]; !exists {
        return errors.New("employee not found")
    }
    delete(s.employees, id)
    return nil
}
