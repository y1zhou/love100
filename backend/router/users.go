package router

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/y1zhou/love100/backend/db"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser INSERT INTO users table. Does nothing if username already exists.
func CreateUser(c *gin.Context) {
	var json db.UserSignupForm
	if err := c.ShouldBind(&json); err != nil {
		// Form validation errors
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Form doesn't bind.",
			"err": err.Error(),
		})
		return
	}
	hash, _ := hashPassword(json.Password)
	user := db.Users{
		Username: json.Username,
		Password: hash,
		Email:    json.Email,
		Gender:   json.Gender,
	}
	if err := db.DB.Create(&user).Error; err != nil {
		// User already exists
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "User already exists.",
			"err": err,
		})
	} else {
		// Create new user
		c.JSON(http.StatusCreated, gin.H{
			"msg": user.Username,
			"err": "",
		})
	}
}

// DeleteUser actually a soft delete as there's the "deleted_at" field.
func DeleteUser(c *gin.Context) {
	var json db.UserDeleteForm
	if err := c.ShouldBind(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Form doens't bind.",
			"err": err.Error(),
		})
		return
	}
	session := sessions.Default(c)
	userID := session.Get("userID")
	var user db.Users
	if err := db.DB.Where(userID).
		First(&user).Error; gorm.IsRecordNotFoundError(err) {
		// User not found
		c.JSON(http.StatusOK, gin.H{
			"msg": "User not found in database.",
			"err": err,
		})
		return
	}
	if checkPasswordHash(json.Password, user.Password) {
		session.Clear()
		session.Save()
		// Soft delete user
		db.DB.Where(userID).Delete(&user)
		c.JSON(http.StatusOK, gin.H{
			"msg": user.Username,
			"err": "",
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": fmt.Sprintf("Check password hash failed for user %s", user.Username),
			"err": user.Username,
		})
	}
}

// UpdateUser Change password and/or email.
func UpdateUser(c *gin.Context) {
	var json db.UserUpdateForm
	if errs := c.ShouldBind(&json); errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Form doesn't bind.",
			"err": errs,
		})
		return
	}
	userID := sessions.Default(c).Get("userID")
	var user db.Users
	if err := db.DB.Unscoped().
		Where("id = ?", userID).
		First(&user).Error; err != nil {
		// User not found
		c.JSON(http.StatusOK, gin.H{
			"msg": "User not found in database.",
			"err": err,
		})
		return
	}
	if checkPasswordHash(json.Password, user.Password) {
		var newHash string
		if json.NewPass != "" {
			newHash, _ = hashPassword(json.NewPass)
		} else {
			newHash = user.Password
		}

		db.DB.Model(&user).
			Updates(map[string]interface{}{
				"password": newHash,
				"email":    json.Email,
			}).
			Where("id = ?", userID)
		c.JSON(http.StatusOK, gin.H{
			"msg": user.Username,
			"err": "",
		})
	} else {
		// Old password doesn't match
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": fmt.Sprintf("Check password hash failed for user %s", user.Username),
			"err": user.Username,
		})
	}
}

// FetchAllUsers lists all users in the database.
func FetchAllUsers(c *gin.Context) {
	type userQuery struct {
		gorm.Model
		Username string
		Email    string
	}
	var users []userQuery
	db.DB.Table("users").
		Select("id, username, email, gender, is_active, created_at, updated_at, deleted_at").
		Order("deleted_at asc").
		Scan(&users)
	c.JSON(http.StatusOK, gin.H{
		"msg":  len(users),
		"err":  "",
		"data": users,
	})

}

// LoginUser Start login session
func LoginUser(c *gin.Context) {
	var json db.UserLoginForm
	session := sessions.Default(c)
	if errs := c.ShouldBind(&json); errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Form doesn't bind.",
			"err": errs.Error(),
		})
		return
	}
	var user db.Users
	if err := db.DB.Where("username = ? AND is_active = TRUE", json.Username).
		First(&user).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "User not found in database.",
			"err": err,
		})
		return
	}
	if checkPasswordHash(json.Password, user.Password) {
		session.Clear()
		session.Set("userID", user.ID)
		session.Set("username", user.Username)
		session.Options(sessions.Options{
			MaxAge: 3600 * 12,
			Path:   "/",
		})
		err := session.Save()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "Failed to generate session token",
				"err": err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg": user.Username,
				"err": "",
			})
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": fmt.Sprintf("Check password hash failed for user %s", user.Username),
			"err": user.Username,
		})
	}
}

// LogoutUser ...
func LogoutUser(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("userID")
	if userID == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Session ID not found",
			"err": "Invalid session token",
		})
	} else {
		session.Set("userID", userID)
		session.Options(sessions.Options{
			MaxAge: -1,
			Path:   "/",
		})
		err := session.Save()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "Failed to delete session token",
				"err": err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg": "Successfully logged out",
				"err": "",
			})
		}
	}
}

// CurrentUser Get the current logged in user.
func CurrentUser(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("userID")
	if userID == nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "User not logged in.",
			"err": "User ID not in session.",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg": session.Get("Username"),
			"err": "",
		})
	}
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
