# Cookie Example

This example demonstrates how to set and get cookies using the Gin framework.

## Steps to Run the Example

1. **Build and Run the Server:**

```bash
go run main.go
```

1. **Login to Set the Cookie:**
   Open your browser and visit the login page:

```sh
http://localhost:8080/login
```

1. **Access the Home Page within 30 Seconds:**
   After logging in, visit the home page within 30 seconds to see the cookie in action:

```sh
http://localhost:8080/home
```

1. **Access the Home Page after 30 Seconds:**
   If you try to visit the home page after 30 seconds, you will see a forbidden error due to the expired cookie:

```sh
http://localhost:8080/home
```

## Code Explanation

- **main.go:**
  - The `main.go` file contains the server setup and route definitions.
  - The `/login` route sets a cookie with a label "ok" and a max age of 30 seconds.
  - The `/home` route is protected by the `CookieTool` middleware, which checks for the presence of the cookie.

```go
package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

func CookieTool() gin.HandlerFunc {
  return func(c *gin.Context) {
    // Get cookie
    if cookie, err := c.Cookie("label"); err == nil {
      if cookie == "ok" {
        c.Next()
        return
      }
    }

    // Cookie verification failed
    c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden with no cookie"})
    c.Abort()
  }
}

func main() {
  route := gin.Default()

  route.GET("/login", func(c *gin.Context) {
    // Set cookie {"label": "ok" }, maxAge 30 seconds.
    c.SetCookie("label", "ok", 30, "/", "localhost", false, true)
    c.String(200, "Login success!")
  })

  route.GET("/home", CookieTool(), func(c *gin.Context) {
    c.JSON(200, gin.H{"data": "Your home page"})
  })

  route.Run(":8080")
}
```

## Conclusion

This example shows how to use cookies for simple session management in a Gin web application. By following the steps above, you can see how cookies are set and validated in a real-world scenario.
