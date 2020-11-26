// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type AuthResponse struct {
	Token string `json:"token"`
	User  *User  `json:"user"`
}

type LeaveHistory struct {
	ID        int         `json:"id" gorm:"primary_key;index;"`
	User      *User       `json:"user" gorm:"not null"`
	Date      time.Time   `json:"date" gorm:"type:timestamptz;not null"`
	Reason    *string     `json:"reason"`
	Type      LeaveType   `json:"type"`
	Status    LeaveStatus `json:"status" gorm:"default:'applied'"`
	Approver  *User       `json:"approver"`
	CreatedAt time.Time   `json:"createdAt" gorm:"type:timestamptz;not null;->;default:now()"`
	UpdatedAt time.Time   `json:"updatedAt" gorm:"type:timestamptz;not null;default:now()"`
}

type User struct {
	ID              int             `json:"id" gorm:"primary_key;auto_increment"`
	Email           string          `json:"email" gorm:"type:varchar(100);unique_index;not null"`
	Password        string          `json:"password" gorm:"not null"`
	Name            string          `json:"name" gorm:"size:32;not null"`
	Bio             *string         `json:"bio"` // 자기소개
	Department      string          `json:"department" gorm:"not null"` //부서
	Position        string          `json:"position" gorm:"not null"` //직급
	WorkSpace       string          `json:"wrokSpace" gorm:"not null"` //근무지
	Contact         string          `json:"contact"  gorm:"not null"` //연락처
	Role            UserRole        `json:"role" gorm:"not null;default:'normal'"` // 일반 혹은 매니저
	Status          UserStatus      `json:"status" gorm:"default:'inOffice'"`
	ProfileImage     string          `json:"profileImage" gorm:"not null;default:'https://next-airbnb.s3.ap-northeast-2.amazonaws.com/profile_image_green2.svg'"`
	Birthday        time.Time       `json:"birthday" gorm:"type:timestamptz;not null"`
	EnteredDate     time.Time       `json:"enteredDate" gorm:"type:timestamptz;not null"`
	ResignationDate *time.Time      `json:"resignationDate" gorm:"type:timestamptz;"`
	RemainLeaves    int             `json:"remainLeaves" gorm:"not null"` //남은 연차 일
	LeaveHistories  []*LeaveHistory `json:"leaveHistories" gorm:"not null"`
	CreatedAt       time.Time       `json:"createdAt" gorm:"type:timestamptz;not null;->;default:now()"`
	UpdatedAt       time.Time       `json:"updatedAt" gorm:"type:timestamptz;not null;default:now()"`
}

type LeaveStatus string

const (
	LeaveStatusApplied  LeaveStatus = "applied"
	LeaveStatusAccepted LeaveStatus = "accepted"
	LeaveStatusRejected LeaveStatus = "rejected"
)

var AllLeaveStatus = []LeaveStatus{
	LeaveStatusApplied,
	LeaveStatusAccepted,
	LeaveStatusRejected,
}

func (e LeaveStatus) IsValid() bool {
	switch e {
	case LeaveStatusApplied, LeaveStatusAccepted, LeaveStatusRejected:
		return true
	}
	return false
}

func (e LeaveStatus) String() string {
	return string(e)
}

func (e *LeaveStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = LeaveStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid LeaveStatus", str)
	}
	return nil
}

func (e LeaveStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type LeaveType string

const (
	LeaveTypeDay       LeaveType = "day"
	LeaveTypeMorning   LeaveType = "morning"
	LeaveTypeAfternoon LeaveType = "afternoon"
)

var AllLeaveType = []LeaveType{
	LeaveTypeDay,
	LeaveTypeMorning,
	LeaveTypeAfternoon,
}

func (e LeaveType) IsValid() bool {
	switch e {
	case LeaveTypeDay, LeaveTypeMorning, LeaveTypeAfternoon:
		return true
	}
	return false
}

func (e LeaveType) String() string {
	return string(e)
}

func (e *LeaveType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = LeaveType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid LeaveType", str)
	}
	return nil
}

func (e LeaveType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type UserRole string

const (
	UserRoleMaster  UserRole = "master"
	UserRoleManager UserRole = "manager"
	UserRoleNormal  UserRole = "normal"
)

var AllUserRole = []UserRole{
	UserRoleMaster,
	UserRoleManager,
	UserRoleNormal,
}

func (e UserRole) IsValid() bool {
	switch e {
	case UserRoleMaster, UserRoleManager, UserRoleNormal:
		return true
	}
	return false
}

func (e UserRole) String() string {
	return string(e)
}

func (e *UserRole) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserRole(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserRole", str)
	}
	return nil
}

func (e UserRole) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type UserStatus string

const (
	UserStatusInOffice UserStatus = "inOffice"
	UserStatusResign   UserStatus = "resign"
)

var AllUserStatus = []UserStatus{
	UserStatusInOffice,
	UserStatusResign,
}

func (e UserStatus) IsValid() bool {
	switch e {
	case UserStatusInOffice, UserStatusResign:
		return true
	}
	return false
}

func (e UserStatus) String() string {
	return string(e)
}

func (e *UserStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserStatus", str)
	}
	return nil
}

func (e UserStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
