package entity

import "time"

// MedicalRecord represents medical record.
type MedicalRecord struct {
	Symptom   string
	Diagnosis string
	Therapy   string
	Result    string
	User      *User
	CreatedAt time.Time
	UpdatedAt time.Time
}

// User represents user.
type User struct {
	Name  string
	Email string
}
