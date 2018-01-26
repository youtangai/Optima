package model

type Auth struct {
	Identity Identity `json:"identity"`
	Scope    Scope    `json:"identity"`
}

type Identity struct {
	Methods  []string `json:"methods"`
	Password Password `json:"password"`
}

type Password struct {
	User User `json:"user"`
}

type User struct {
	Domain   Domain `json:"domain"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Domain struct {
	Name string `json:"name"`
}

type Scope struct {
	Project Project `json:"project"`
}

type Project struct {
	Domain Domain `json:"domain"`
	Name   string `json:"name"`
}
