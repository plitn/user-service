package model

type User struct {
	Id       int64  `db:"id" json:"id" goku:"skipinsert"`
	Name     string `db:"name" json:"name"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
	ImageUrl string `db:"image_url" json:"image_url"`
	Gender   int    `db:"gender" json:"gender"`
	Age      int    `db:"age" json:"age"`
	Weight   int    `db:"weight" json:"weight"`
}

type UserLogin struct {
	Name     string `db:"name" json:"name"`
	Password string `db:"password" json:"password"`
}

type Stylist struct {
	Id                    int64  `db:"id" json:"id" goku:"skipinsert"`
	Name                  string `db:"name" json:"name"`
	Email                 string `db:"email" json:"email"`
	Password              string `db:"password" json:"password"`
	ImageUrl              string `db:"image_url" json:"image_url"`
	Gender                int    `db:"gender" json:"gender"`
	Age                   int    `db:"age" json:"age"`
	ExperienceTime        int    `db:"experience_time" json:"experience_time"`
	Skills                []int  `db:"skills" json:"skills"`
	Portfolio             string `db:"portfolio" json:"portfolio"`
	ExperienceDescription string `db:"experience_description" json:"experience_description"`
}

type UserStylist struct {
	Id        int64 `db:"id" json:"id" goku:"skipinsert"`
	UserId    int64 `db:"user_id" json:"user_id"`
	StylistId int64 `db:"stylist_id" json:"stylist_id"`
}
