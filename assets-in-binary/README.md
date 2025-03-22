# Assets in Binary

This project demonstrates how to embed static files and file trees into a Go executable using the `//go:embed` directive introduced in Go 1.16. The embedded files can be accessed at runtime, allowing for easy distribution of assets within a single binary.

## Project Structure

- **assets/**: Contains static assets such as images and icons.
  - `favicon.ico`: The favicon for the website.
  - `images/`: Directory containing example images.
- **templates/**: Contains HTML templates used by the application.
  - `index.tmpl`: The main template for the homepage.
  - `foo/bar.tmpl`: The template for the "Foo" page.
- **go.mod**: The Go module file, listing the dependencies required for the project.
- **go.sum**: The Go checksum file, ensuring the integrity of the dependencies.
- **main.go**: The main application file, setting up the web server and routes.

## Dependencies

The project uses the following dependencies:

- `github.com/gin-gonic/gin`: A web framework for Go.
- `github.com/bytedance/sonic`: A high-performance JSON library.
- `github.com/gabriel-vasile/mimetype`: A library for detecting MIME types.
- And various other indirect dependencies listed in `go.mod`.

## Running the Application

To run the application, use the following command:

```bash
go run main.go
```

The application will start a web server on `http://localhost:8080`. You can access the following routes:

- `/`: The homepage, rendered using `index.tmpl`.
- `/foo`: The "Foo" page, rendered using `bar.tmpl`.
- `/public/assets/images/example.png`: An example image served from the embedded assets.
- `/favicon.ico`: The favicon served from the embedded assets.

## Embedding Files

The `//go:embed` directive is used to embed the contents of the `assets` and `templates` directories into the Go binary. The embedded files are accessed using the `embed.FS` type.

Example usage in `main.go`:

```go
//go:embed assets/* templates/*
var f embed.FS

func main() {
  router := gin.Default()
  templ := template.Must(template.New("").ParseFS(f, "templates/*.tmpl", "templates/foo/*.tmpl"))
  router.SetHTMLTemplate(templ)

  router.StaticFS("/public", http.FS(f))

  router.GET("/", func(c *gin.Context) {
    c.HTML(http.StatusOK, "index.tmpl", gin.H{
      "title": "Main website",
    })
  })

  router.GET("/foo", func(c *gin.Context) {
    c.HTML(http.StatusOK, "bar.tmpl", gin.H{
      "title": "Foo website",
    })
  })

  router.GET("favicon.ico", func(c *gin.Context) {
    file, _ := f.ReadFile("assets/favicon.ico")
    c.Data(
      http.StatusOK,
      "image/x-icon",
      file,
    )
  })

  router.Run(":8080")
}
```

## License

This project is licensed under the MIT License. See the [LICENSE](../LICENSE) file for details.

## Contributing

Contributions are welcome! Please see the [CONTRIBUTING](../CONTRIBUTING.md) file for guidelines.

## References

- [Go 1.16 Release Notes](https://tip.golang.org/doc/go1.16#embed)
- [embed package documentation](https://tip.golang.org/pkg/embed/)
