package user

import (
	"fmt"
	"time"
	"encoding/json"
	"testing"
)

const (
	name = "Mayama"
	sex = "male"
	addr = "Ebisu"
)

func TestUser(t *testing.T) {
	dob := time.Now()
	u := &User{Name: name, Birthday: dob, Sex: sex, Address: addr}
	if u.Name != name {
		t.Errorf("Name %v is not same %v", u.Name, name)
	}
	if !u.Birthday.Equal(dob) {
		t.Errorf("Birthday %v is not same %v", u.Birthday, dob)
	}
	if u.Sex != sex {
		t.Errorf("Sex %v is not same %v", u.Sex, sex)
	}
	if u.Address != addr {
		t.Errorf("Address %v is not same %v", u.Address, addr)
	}
	j, _ := json.Marshal(u)
	fmt.Printf("%v", string(j))
}
