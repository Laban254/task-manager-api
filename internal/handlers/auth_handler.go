// /internal/handlers/auth_handler.go
package handlers

import (
    "net/http"
    "time"
    "context"
    "golang.org/x/oauth2"
    "golang.org/x/oauth2/google"
    "encoding/json"
	"strings"

    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
    "task_management_api/pkg/models"
    "task_management_api/pkg/database"
)
 

var googleOauthConfig = &oauth2.Config{
    ClientID:     cfg.GOOGLE_CLIENT_ID,
    ClientSecret: cfg.GOOGLE_CLIENT_SECRET,
    RedirectURL:  cfg.GOOGLE_REDIRECT_URL,
    Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
    Endpoint:     google.Endpoint,
}

func RegisterUser(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    trimmedPassword := strings.TrimSpace(user.Password)
    if trimmedPassword == "" || len(trimmedPassword) < 6 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Password must be at least 6 characters long"})
        return
    }

    if err := validate.Struct(user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    hashedPassword, err := hashPassword(trimmedPassword)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
        return
    }
    user.Password = hashedPassword

    if err := database.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// LoginUser handles user login
func LoginUser(c *gin.Context) {
    var loginData models.User
    if err := c.ShouldBindJSON(&loginData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var user models.User
    if err := database.DB.Where("username = ?", loginData.Username).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    if !checkPasswordHash(loginData.Password, user.Password) {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    token := generateToken(user.Username, user.Role)
    c.JSON(http.StatusOK, gin.H{"token": token})
}
// GoogleLogin redirects to Google OAuth
func GoogleLogin(c *gin.Context) {
    url := googleOauthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
    c.Redirect(http.StatusTemporaryRedirect, url)
}

// GoogleCallback handles the callback from Google OAuth
func GoogleCallback(c *gin.Context) {
    code := c.Query("code")
    if code == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "No code in the request"})
        return
    }

    token, err := googleOauthConfig.Exchange(context.Background(), code)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange token"})
        return
    }

    resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user info"})
        return
    }
    defer resp.Body.Close()

    var googleUser struct {
        Email string `json:"email"`
        Name  string `json:"name"`
        ID    string `json:"id"`
    }
    if err := json.NewDecoder(resp.Body).Decode(&googleUser); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode user info"})
        return
    }

    var user models.User
    if err := database.DB.Where("username = ?", googleUser.Email).First(&user).Error; err != nil {
        user = models.User{
            Username: googleUser.Email,
            Password: "", 
            Role:     "user", 
        }

        if err := database.DB.Create(&user).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save user to the database"})
            return
        }
    }

    tokenString := generateToken(user.Username, user.Role)

    c.JSON(http.StatusOK, gin.H{"token": tokenString, "user": googleUser})
}


func generateToken(username string, role string) string {
    claims := jwt.MapClaims{
        "username": username,
        "role":     role,
        "exp":      time.Now().Add(time.Hour * 72).Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, _ := token.SignedString([]byte(cfg.JWTSecret))
    return tokenString
}


// hashPassword hashes the user's password
func hashPassword(password string) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hashedPassword), nil
}

// checkPasswordHash compares the password with the hashed password
func checkPasswordHash(password, hashedPassword string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
    return err == nil
}

// AdminRegisterUser handles admin registration of new users
func AdminRegisterUser(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    trimmedPassword := strings.TrimSpace(user.Password)
    if trimmedPassword == "" || len(trimmedPassword) < 6 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Password must be at least 6 characters long"})
        return
    }

    if user.Role == "" {
        user.Role = "admin" 
    }

    if err := validate.Struct(user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    hashedPassword, err := hashPassword(trimmedPassword)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
        return
    }
    user.Password = hashedPassword

    var existingUser models.User
    if err := database.DB.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
        c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
        return
    }

    if err := database.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}