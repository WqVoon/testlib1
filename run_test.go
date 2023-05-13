package testlib1

import "testing"

func TestGetAll(t *testing.T) {
	Register("name", "hygao")
	GetAll()
}
