package main

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
)

type Server struct {
    store *EmployeeStore
}

func NewServer() *Server {
    return &Server{
        store: NewEmployeeStore(),
    }
}

func (s *Server) listEmployees(w http.ResponseWriter, r *http.Request) {
    query := r.URL.Query()
    pageStr := query.Get("page")
    sizeStr := query.Get("size")

    page, _ := strconv.Atoi(pageStr)
    size, _ := strconv.Atoi(sizeStr)

    if page < 1 {
        page = 1
    }
    if size < 1 {
        size = 10
    }

    s.store.mu.Lock()
    defer s.store.mu.Unlock()

    employees := make([]Employee, 0, len(s.store.employees))
    for _, employee := range s.store.employees {
        employees = append(employees, employee)
    }

    start := (page - 1) * size
    end := start + size

    if start > len(employees) {
        start = len(employees)
    }
    if end > len(employees) {
        end = len(employees)
    }

    response := employees[start:end]
    json.NewEncoder(w).Encode(response)
}

func (s *Server) createEmployee(w http.ResponseWriter, r *http.Request) {
    var emp Employee
    if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    createdEmployee := s.store.CreateEmployee(emp.Name, emp.Position, emp.Salary)
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(createdEmployee)
}

func (s *Server) getEmployeeByID(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, _ := strconv.Atoi(vars["id"])
    emp, err := s.store.GetEmployeeByID(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(emp)
}

func (s *Server) updateEmployee(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, _ := strconv.Atoi(vars["id"])

    var emp Employee
    if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    updatedEmployee, err := s.store.UpdateEmployee(id, emp.Name, emp.Position, emp.Salary)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(updatedEmployee)
}

func (s *Server) deleteEmployee(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, _ := strconv.Atoi(vars["id"])

    if err := s.store.DeleteEmployee(id); err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}
