package models

type UserKey string

func (k UserKey) String() string {
	return string(k)
}
