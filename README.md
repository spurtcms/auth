# Auth Package

The auth package is an independent module that can be individually enabled and configured to suit the needs of the website owner.
One may use the spurtCMS auth package, or if the current application already includes it, there's the option to bypass our auth package, retaining the existing auth package setup, and continue using our other packages without hassle.


# Description

This package makes it possible to use a JSON Web Token (JWT) to securely authenticate a valid user requesting access to your spurtCMS.
JSON Web Tokens are an open, industry standard RFC 7519 method for representing claims securely between two parties.

![Screenshot of spurtCMS log in screen](https://www.spurtcms.com/spurtcms-starter-template.jpg)



## Features

- Smooth admin login 
- Generates and returns an authentication token, typically used for user authentication and authorization processes.
- Validates the authenticity and integrity of a given authentication token.
- Check if a user login attempt is valid by verifying credentials against stored user information.
- Parses and decrypts JWT tokens (Json Web Tokens) from the HTTP request's 'authorization' header
- Verifies the validity of tokens using the provided jwtSecret

# Check UserName Password

```bash
func (auth *Auth) Checklogin(Username string, Password string) (string, int, error) {

	username := Username

	password := Password

	user, err := CheckLogin(username, password, auth.DBString)

	if err != nil {

		log.Println(err)

	}
```

# creates a token

```bash
func (auth *Auth) CreateToken() (string, error) {

	atClaims := jwt.MapClaims{}

	atClaims["user_id"] = auth.UserId

	atClaims["expiry_time"] = time.Now().Add(time.Duration(auth.ExpiryTime) * time.Hour).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	return token.SignedString([]byte(auth.SecretKey))
}
```

# verifies a token
```bash
func (auth *Auth) VerifyToken(token string, secret string) (userid int, err error) {

	Claims := jwt.MapClaims{}

	tkn, err := jwt.ParseWithClaims(token, Claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	
```
# Getting help
If you encounter a problem with the package,please refer [Please refer [(https://www.spurtcms.com/documentation/cms-admin)] or you can create a new Issue in this repo[https://github.com/spurtcms/auth/blob/main/auth.go]. 
