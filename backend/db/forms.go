package db

type (
	// ContentCreateForm Insert new row of things to do together.
	// Can also be used to query items.
	ContentCreateForm struct {
		Title   string `form:"Title" json:"Title" binding:"required"`
		Comment string `form:"Comment" json:"Comment"`
		Status  bool   `form:"ItemStatus" json:"ItemStatus" sql:"default:false"`
	}

	// ContentDeleteForm Delete existing items using their ID.
	ContentDeleteForm struct {
		ItemIDs []int64 `form:"ItemIDs" json:"ItemIDs" binding:"required"`
	}

	// ContentTitleUpdateForm ...
	ContentTitleUpdateForm struct {
		ItemID  int64  `form:"ItemID" json:"ItemID" binding:"required"`
		Title   string `form:"Title" json:"Title" binding:"required"`
		Comment string `form:"Comment" json:"Comment"`
	}

	//ContentStatusUpdateForm ...
	ContentStatusUpdateForm struct {
		ItemID int64 `form:"ItemID" json:"ItemID" binding:"required"`
		Status bool  `form:"ItemStatus" json:"ItemStatus" sql:"default:false"`
	}
)

type (
	// UserSignupForm ...
	UserSignupForm struct {
		Username    string `form:"Username" json:"Username" binding:"required"`
		Password    string `form:"Password" json:"Password" binding:"required,min=6"`
		ConfirmPass string `form:"ConfirmPass" json:"ConfirmPass" binding:"eqfield=Password"`
		Email       string `form:"Email" json:"Email" binding:"omitempty,email"`
		Gender      string `form:"Gender" json:"Gender" binding:"required"`
	}
	// UserQueryForm Check if an user exists.
	UserQueryForm struct {
		Username string `form:"Username" json:"Username" binding:"required"`
	}
	// UserDeleteForm Delete an existing user.
	UserDeleteForm struct {
		Password string `form:"Password" json:"Password" binding:"required,min=6"`
	}
	// UserUpdateForm Change password or email.
	UserUpdateForm struct {
		Password string `form:"Password" json:"Password" binding:"required,min=6"`
		NewPass  string `form:"NewPass" json:"NewPass" binding:"omitempty,min=6,nefield=Password"`
		Email    string `form:"Email" json:"Email" binding:"omitempty,email"`
		Gender   string `form:"Gender" json:"Gender" binding:"omitempty"`
	}
	// UserLoginForm ...
	UserLoginForm struct {
		Username string `form:"Username" json:"Username" binding:"required"`
		Password string `form:"Password" json:"Password" binding:"required,min=6"`
	}
)
