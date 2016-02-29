package main

type Config struct {
	Email    string `json:"login_email"`
	Password string `json:"login_password"`
	Token    string `json:"login_token"`
	Domain   string `json:"domain"`
}

func (this *Config) Check() bool {
	return this.Domain != "" && ((this.Email != "" && this.Password != "") || this.Token != "")
}
