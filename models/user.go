package blog

type Access int

const (
	ADMIN    Access = 0
	EDITOR   Access = 1
	STANDARD Access = 2
)

type User struct {
	Alias       string
	Firstname   string
	Lastname    string
	AccessLevel Access
}
