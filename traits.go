package analytics

import "time"

type Traits map[string]interface{}

func NewTraits() Traits {
	return make(Traits, 10)
}

func (t Traits) SetAddress(address string) Traits {
	return t.Set("address", address)
}

func (t Traits) SetAge(age int) Traits {
	return t.Set("age", age)
}

func (t Traits) SetAvatar(url string) Traits {
	return t.Set("avatar", url)
}

func (t Traits) SetBirthday(date time.Time) Traits {
	return t.Set("birthday", date)
}

func (t Traits) SetCreatedAt(date time.Time) Traits {
	return t.Set("createdAt", date.Unix())
}

func (t Traits) SetDescription(desc string) Traits {
	return t.Set("description", desc)
}

func (t Traits) SetEmail(email string) Traits {
	return t.Set("email", email)
}

func (t Traits) SetFirstName(firstName string) Traits {
	return t.Set("firstName", firstName)
}

func (t Traits) SetGender(gender string) Traits {
	return t.Set("gender", gender)
}

func (t Traits) SetLastName(lastName string) Traits {
	return t.Set("lastName", lastName)
}

func (t Traits) SetName(name string) Traits {
	return t.Set("name", name)
}

func (t Traits) SetPhone(phone string) Traits {
	return t.Set("phone", phone)
}

func (t Traits) SetTitle(title string) Traits {
	return t.Set("title", title)
}

func (t Traits) SetUsername(username string) Traits {
	return t.Set("username", username)
}

func (t Traits) SetWebsite(url string) Traits {
	return t.Set("website", url)
}

func (t Traits) Set(field string, value interface{}) Traits {
	t[field] = value
	return t
}
