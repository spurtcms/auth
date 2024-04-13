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

``` bash
go get github.com/spurtcms/auth
```


- Setup the Auth config
``` bash
config := Config{
		UserId:     1,
		ExpiryTime: 2, // It should be in hours not minutes or seconds
		SecretKey:  "test@123",
		DB:         &gorm.DB{},
	}

	auth := AuthSetup(config)
```


# Usage Example
``` bash
func main (){

		config := Config{
			UserId:     1,
			ExpiryTime: 2, // It should be in hours not minutes or seconds
			SecretKey:  "test@123",
			DB:         &gorm.DB{},
		}
		
		auth := AuthSetup(config)
		
		//create token - generates new  JWtoken
		token, _ := auth.CreateToken()
		
		//checklogin - verifies user credentials
		token, userid,err :=auth.Checklogin("Username", "Password")
		
		//verifytoken - parses and verifies the token generated
		userid, err :=	auth.VerifyToken("token","secretkey","currenttime")
	}
```

# Getting help
If you encounter a problem with the package,please refer [Please refer (https://www.spurtcms.com/documentation/cms-admin) or you can create a new Issue in this repo (https://github.com/spurtcms/auth/issues). 
