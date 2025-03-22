# OIDC GitHub OAuth Example

This project demonstrates how to implement GitHub OAuth2 authentication using the Gin web framework in Go. It includes routes for logging in with GitHub, handling the OAuth2 callback, and displaying user information.

## Table of Contents

- [OIDC GitHub OAuth Example](#oidc-github-oauth-example)
  - [Table of Contents](#table-of-contents)
  - [Overview](#overview)
  - [Setup](#setup)
    - [Prerequisites](#prerequisites)
    - [Installation](#installation)
  - [Usage](#usage)
    - [Routes](#routes)
    - [Example](#example)
  - [Code Explanation](#code-explanation)
    - [Main Components](#main-components)
    - [Main Routes](#main-routes)
    - [OAuth2 Configuration](#oauth2-configuration)
    - [State Management](#state-management)
    - [User Information](#user-information)
  - [License](#license)

## Overview

This project provides a simple example of how to integrate GitHub OAuth2 authentication into a Go web application using the Gin framework. It includes the following features:

- Redirecting users to GitHub for authentication
- Handling the OAuth2 callback
- Retrieving and displaying user information from GitHub
- Protecting routes that require authentication

## Setup

### Prerequisites

- Go 1.23.1 or later
- GitHub OAuth2 application credentials (Client ID and Client Secret)

### Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/yourusername/oidc.git
   cd oidc
   ```

2. Install dependencies:

   ```sh
   go mod tidy
   ```

3. Set up environment variables:

```sh
export GITHUB_CLIENT_ID=your_github_client_id
export GITHUB_CLIENT_SECRET=your_github_client_secret
```

4. Run the application:

```sh
go run main.go
```

1. Open your browser and navigate to `http://localhost:8080`.

## Usage

### Routes

- `/`: Home page with a link to log in with GitHub.
- `/login`: Redirects to GitHub for authentication.
- `/callback`: Handles the GitHub OAuth2 callback and displays user information.
- `/protected`: A protected route that requires authentication.

### Example

1. Visit the home page at `http://localhost:8080`.
2. Click the "Log in with GitHub" link to be redirected to GitHub for authentication.
3. After logging in, you will be redirected back to the application and your GitHub user information will be displayed.
4. Access the protected route at `http://localhost:8080/protected`.

## Code Explanation

### Main Components

- `main.go`: The main application file containing the routes and OAuth2 configuration.
- `templates/`: Directory containing HTML templates for the home and success pages.

### Main Routes

- **Home Route (`/`)**: Displays the home page with a link to log in with GitHub.

```go
r.GET("/", func(c *gin.Context) {
  c.HTML(http.StatusOK, "index.html", gin.H{
    "title": "GitHub OAuth Example",
  })
})
```

- **Login Route (`/login`)**: Redirects to GitHub for authentication.

```go
r.GET("/login", func(c *gin.Context) {
  state, err := generateRandomState()
  if err != nil {
    c.String(http.StatusInternalServerError, "Unable to generate state value")
    return
  }
  stateCache.Set(state, true, cache.DefaultExpiration)
  authURL := oauth2Config.AuthCodeURL(state)
  c.Redirect(http.StatusFound, authURL)
})
```

- **Callback Route (`/callback`)**: Handles the GitHub OAuth2 callback and displays user information.

```go
r.GET("/callback", func(c *gin.Context) {
  state := c.Query("state")
  if _, exists := stateCache.Get(state); !exists {
    c.String(http.StatusBadRequest, "Invalid state value")
    return
  }
  stateCache.Delete(state)
  code := c.Query("code")
  if code == "" {
    c.String(http.StatusBadRequest, "Authorization code not provided")
    return
  }
  token, err := oauth2Config.Exchange(context.Background(), code)
  if err != nil {
    c.String(http.StatusInternalServerError, "Unable to exchange access token: "+err.Error())
    return
  }
  client := oauth2Config.Client(context.Background(), token)
  resp, err := client.Get(userInfoURL)
  if err != nil {
    c.String(http.StatusInternalServerError, "Unable to retrieve user information: "+err.Error())
    return
  }
  defer resp.Body.Close()
  var user GitHubUser
  if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
    c.String(http.StatusInternalServerError, "Unable to parse user information: "+err.Error())
    return
  }
  c.HTML(http.StatusOK, "success.html", gin.H{
    "title":      "Authentication Successful",
    "username":   user.Login,
    "name":       user.Name,
    "email":      user.Email,
    "avatar_url": user.AvatarURL,
  })
})
```

- **Protected Route (`/protected`)**: A protected route that requires authentication.

```go
r.GET("/protected", func(c *gin.Context) {
  c.String(http.StatusOK, "This is a protected resource!")
})
```

### OAuth2 Configuration

The OAuth2 configuration is set up using the `oauth2.Config` struct, which includes the client ID, client secret, redirect URL, and GitHub OAuth2 endpoint.

```go
var oauth2Config = oauth2.Config{
  ClientID:     clientID,
  ClientSecret: clientSecret,
  RedirectURL:  redirectURL,
  Endpoint:     github.Endpoint,
  Scopes:       []string{"user:email", "read:user"},
}
```

### State Management

A random state value is generated and stored in a cache to prevent CSRF attacks during the OAuth2 authentication process.

```go
func generateRandomState() (string, error) {
  b := make([]byte, 32)
  if _, err := rand.Read(b); err != nil {
    return "", err
  }
  return base64.URLEncoding.EncodeToString(b), nil
}
```

### User Information

The user information is retrieved from GitHub using the access token obtained during the OAuth2 authentication process.

```go
type GitHubUser struct {
  ID        int    `json:"id"`
  Login     string `json:"login"`
  Name      string `json:"name"`
  Email     string `json:"email"`
  AvatarURL string `json:"avatar_url"`
}
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
