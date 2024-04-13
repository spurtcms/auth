# Auth Package

This Auth package stands as a versatile, standalone module, tailored for seamless integration into any Golang project, promising flexibility and convenience. It isn't confined solely to any particular platform; it offers full freedom to incorporate it into any Golang platform. We have seamlessly integrated it with the spurtCMS admin system, ensuring a smooth integration process.

## Features

- Smooth admin login 
- Generates and returns an authentication token, typically used for user authentication and authorization processes.
- Validates the authenticity and integrity of a given authentication token.
- Check if a user login attempt is valid by verifying credentials against stored user information.
- Parses and decrypts JWT tokens (Json Web Tokens) from the HTTP request's 'authorization' header
- Verifies the validity of tokens using the provided jwtSecret


# Installation

This package makes it possible to use a JSON Web Token (JWT) to securely authenticate a valid user requesting access to your spurtCMS.
JSON Web Tokens are an open, industry standard RFC 7519 method for representing claims securely between two parties.

# Check UserName Password

The checklogin function takes a username and password,checks if they are valid from the database.

```bash
func (auth *Auth) Checklogin(Username string, Password string) (string, int, error) {

	username := Username

	password := Password

	user, err := CheckLogin(username, password, auth.DBString)

	if err != nil {

		log.Println(err)

	}
```

# Generates a token

This creates a JWT with some claims(e.g.,user information,expiration time) and signs it with a secret key

```bash
func (auth *Auth) CreateToken() (string, error) {

	atClaims := jwt.MapClaims{}

	atClaims["user_id"] = auth.UserId

	atClaims["expiry_time"] = time.Now().Add(time.Duration(auth.ExpiryTime) * time.Hour).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	return token.SignedString([]byte(auth.SecretKey))
}
```

# Verifies a token

This code verifies the token generated and is used to parse and verify the token

```bash
func (auth *Auth) VerifyToken(token string, secret string) (userid int, err error) {

	Claims := jwt.MapClaims{}

	tkn, err := jwt.ParseWithClaims(token, Claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	
```
# Getting help
If you encounter a problem with the package,please refer [Please refer [(https://www.spurtcms.com/documentation/cms-admin)] or you can create a new Issue in this repo[https://github.com/spurtcms/auth/issues]. 
