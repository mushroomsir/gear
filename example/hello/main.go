package main

import (
	"github.com/fluent/fluent-logger-golang/fluent"
	"github.com/teambition/gear"
	"github.com/teambition/gear/logging"
)

// go run example/hello/main.go
func main() {
	app := gear.New()

	fc, err := fluent.New(fluent.Config{FluentPort: 24224, FluentHost: "127.0.0.1", MarshalAsJSON: true})
	if err != nil {
		panic(err)
	}
	logger := logging.Default()
	logger.SetOutput(fc.EncodeAndPostData)
	// Add logging middleware
	app.UseHandler(logger)

	// Add router middleware
	router := gear.NewRouter()

	// try: http://127.0.0.1:3000/hello
	router.Get("/hello", func(ctx *gear.Context) error {
		return ctx.HTML(200, "<h1>Hello, Gear!</h1>")
	})

	// try: http://127.0.0.1:3000/test?query=hello
	router.Otherwise(func(ctx *gear.Context) error {
		return ctx.JSON(200, map[string]interface{}{
			"Host":    ctx.Host,
			"Method":  ctx.Method,
			"Path":    ctx.Path,
			"URL":     ctx.Req.URL.String(),
			"Headers": ctx.Req.Header,
		})
	})
	app.UseHandler(router)
	app.Error(app.Listen(":3000"))
}
