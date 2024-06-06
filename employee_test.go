package main

import (
    "testing"
)

func TestEmployeeStore(t *testing.T) {
    store := NewEmployeeStore()

    // Test CreateEmployee
    emp := store.CreateEmployee("John Doe", "Developer", 50000)
    if emp.ID != 1 || emp.Name != "John Doe" || emp.Position != "Developer" || emp.Salary != 50000 {
        t.Errorf("CreateEmployee failed: got %+v", emp)
    }

    // Test GetEmployeeByID
    emp, err := store.GetEmployeeByID(1)
    if err != nil || emp.ID != 1 {
        t.Errorf("GetEmployeeByID failed: %v", err)
    }

    // Test UpdateEmployee
    emp, err = store.UpdateEmployee(1, "Jane Doe", "Manager", 60000)
    if err != nil || emp.Name != "Jane Doe" || emp.Position != "Manager" || emp.Salary != 60000 {
        t.Errorf("UpdateEmployee failed: got %+v", emp)
    }

    // Test DeleteEmployee
    err = store.DeleteEmployee(1)
    if err != nil {
        t.Errorf("DeleteEmployee failed: %v", err)
    }
    _, err = store.GetEmployeeByID(1)
    if err == nil {
        t.Errorf("GetEmployeeByID should have failed after deletion")
    }
}
