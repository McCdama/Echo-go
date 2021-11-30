package payload

import (
	"fmt"
	"strings"
	"time"
)

type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name" param:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email" param:"email"`
}

func (u User) GetDetail() string {
	s := fmt.Sprintf("The name is %s And the Email is %s", u.Name, u.Email)
	return s
}
func (u User) PanicDetail() string {
	s := fmt.Sprintf("please provide a valid name: %s or email: %s", u.Name, u.Email)
	return s
}

type CustomTime struct {
	time.Time
}

type TestModel struct {
	Date CustomTime `json:"date"`
}

func (t CustomTime) MarshalJSON() ([]byte, error) {
	date := t.Time.Format("2006-01-02")
	date = fmt.Sprintf(`"%s"`, date)
	return []byte(date), nil
}

func (t *CustomTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")

	date, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	t.Time = date
	return
}
