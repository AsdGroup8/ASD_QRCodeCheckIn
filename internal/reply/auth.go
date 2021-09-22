package reply

// UserAuth when user pass auth, this will be replied
type UserAuth struct {
	Token string `json:"token"`
	Name  string `json:"name"`
}
