package main

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
)

var app = fiber.New()

func TestHelloWorld(t *testing.T) {
	app.Get("/", func(ctx fiber.Ctx) error {
		return ctx.SendString("Hello, World!")
	})

	req := httptest.NewRequest("GET", "/", nil)
	res, err := app.Test(req)

	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)

	bytes, err := io.ReadAll(res.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello, World!", string(bytes))
}

func TestCtx(t *testing.T) {
	app.Get("/hello", func(ctx fiber.Ctx) error {
		name := ctx.Query("name", "Guest") //fungsi untuk mengambil nilai parameter query "name" dari permintaan HTTP. Jika parameter tersebut tidak ada, maka akan menggunakan nilai default "Guest".
		return ctx.SendString("Hello " + name + "!")
	})

	req := httptest.NewRequest("GET", "/hello?name=dam", nil)
	res, err := app.Test(req)

	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)

	bytes, err := io.ReadAll(res.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello dam!", string(bytes))

	req = httptest.NewRequest("GET", "/hello", nil)
	res, err = app.Test(req)

	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)

	bytes, err = io.ReadAll(res.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello Guest!", string(bytes))
}

func TestHttpRequest(t *testing.T) {
	app.Get("/request", func(ctx fiber.Ctx) error {
		//fungsi untuk mengambil nilai header "firstname" dari permintaan HTTP menggunakan ctx.Get("firstname")
		// dan nilai cookie "lastname" menggunakan ctx.Cookies("lastname").
		// Kemudian, nilai-nilai tersebut digabungkan dalam string yang dikirim sebagai respons.
		first := ctx.Get("firstname")
		last := ctx.Cookies("lastname")
		return ctx.SendString("Hello " + first + " " + last + "!")
	})

	req := httptest.NewRequest("GET", "/request", nil)
	req.Header.Set("firstname", "dam")
	req.AddCookie(&http.Cookie{Name: "lastname", Value: "dim"})
	res, err := app.Test(req)

	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)

	bytes, err := io.ReadAll(res.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello dam dim!", string(bytes))
}

func TestRouteParameter(t *testing.T) {
	app.Get("/users/:userId/orders/:orderId", func(ctx fiber.Ctx) error {
		//fungsi untuk mengambil nilai parameter rute "userId" dan "orderId" dari URL menggunakan ctx.Params("userId") dan ctx.Params("orderId").
		//Nilai-nilai tersebut kemudian digabungkan dalam string yang dikirim sebagai respons.
		userId := ctx.Params("userId")
		orderId := ctx.Params("orderId")
		return ctx.SendString("User ID: " + userId + ", Order ID: " + orderId)
	})

	req := httptest.NewRequest("GET", "/users/123/orders/456", nil)
	res, err := app.Test(req)

	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)

	bytes, err := io.ReadAll(res.Body)
	assert.Nil(t, err)
	assert.Equal(t, "User ID: 123, Order ID: 456", string(bytes))
}

func TestFormRequest(t *testing.T) {
	app.Post("/form", func(ctx fiber.Ctx) error {
		name := ctx.FormValue("name")
		return ctx.SendString("Hello " + name + "!")
	})

	body := strings.NewReader("name=dam")
	req := httptest.NewRequest("POST", "/form", body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, err := app.Test(req)

	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)

	bytes, err := io.ReadAll(res.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello dam!", string(bytes))
}

//go:embed source/image.jpeg
var contohFile []byte

func TestFormUpload(t *testing.T) {
	app.Post("/upload", func(ctx fiber.Ctx) error {
		//fungsi untuk menangani unggahan file melalui permintaan HTTP POST. Fungsi ini menggunakan ctx.FormFile("file") untuk mengambil file yang diunggah dengan nama "file" dari formulir.
		// Jika terjadi kesalahan saat mengambil file, fungsi akan mengembalikan error tersebut.
		// Setelah berhasil mengambil file, fungsi menggunakan ctx.SaveFile(file, "./target/"+file.Filename)
		// untuk menyimpan file tersebut ke direktori "./target/" dengan nama yang sama seperti nama file yang diunggah.
		// Jika terjadi kesalahan saat menyimpan file, fungsi akan mengembalikan error tersebut.
		// Jika semuanya berhasil, fungsi akan mengirimkan respons "Upload Success" kepada klien.
		file, err := ctx.FormFile("file")
		if err != nil {
			return err
		}

		err = ctx.SaveFile(file, "./target/"+file.Filename)
		if err != nil {
			return err
		}

		return ctx.SendString("Upload Success")
	})

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	file, err := writer.CreateFormFile("file", "image.jpeg")
	assert.Nil(t, err)
	file.Write(contohFile)
	writer.Close()

	request := httptest.NewRequest("POST", "/upload", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Upload Success", string(bytes))
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func TestRequestBody(t *testing.T) {
	app.Post("/login", func(ctx fiber.Ctx) error {
		body := ctx.Body() //fungsi untuk mengambil isi body permintaan HTTP menggunakan ctx.Body(). Isi body tersebut kemudian di-unmarshal dari format JSON ke dalam struct LoginRequest menggunakan json.Unmarshal.

		request := new(LoginRequest)
		err := json.Unmarshal(body, request)
		if err != nil {
			return err
		}

		return ctx.SendString("Hello " + request.Username)
	})

	body := strings.NewReader(`{"username":"dam", "password":"rahasia"}`)
	request := httptest.NewRequest("POST", "/login", body)
	request.Header.Set("Content-Type", "application/json")
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello dam", string(bytes))
}

type RegisterRequest struct {
	Username string `json:"username" xml:"username" form:"username"`
	Password string `json:"password" xml:"password" form:"password"`
	Name     string `json:"name" xml:"name" form:"name"`
}

func TestBodyParser(t *testing.T) {
	app.Post("/register", func(ctx fiber.Ctx) error {
		request := new(RegisterRequest)
		err := ctx.Bind().Body(request)
		if err != nil {
			return err
		}
		//fungsi untuk mengikat (bind) data dari body permintaan HTTP ke dalam struct RegisterRequest menggunakan ctx.Bind(request). Fungsi ini secara otomatis akan mendeteksi format data (JSON, XML, atau form) berdasarkan header Content-Type dari permintaan dan mengisi struct dengan data yang sesuai.

		return ctx.SendString("Register Success " + request.Username)
	})
}

func TestBodyParserJSON(t *testing.T) {
	TestBodyParser(t)

	body := strings.NewReader(`{"username":"Dam", "password":"rahasia", "name":"Dam Dim"}`)
	request := httptest.NewRequest("POST", "/register", body)
	request.Header.Set("Content-Type", "application/json")
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Register Success Dam", string(bytes))
}

func TestBodyParserForm(t *testing.T) {
	TestBodyParser(t)

	body := strings.NewReader(`username=Dam&password=rahasia&name=Dam+Dim`)
	request := httptest.NewRequest("POST", "/register", body)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Register Success Dam", string(bytes))
}

func TestBodyParserXML(t *testing.T) {
	TestBodyParser(t)

	body := strings.NewReader(
		`<RegisterRequest>
			<username>Dam</username>
			<password>rahasia</password>
			<name>Dam Dim</name>
		</RegisterRequest>
		`)
	request := httptest.NewRequest("POST", "/register", body)
	request.Header.Set("Content-Type", "application/xml")
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Register Success Dam", string(bytes))
}

func TestResponseJSON(t *testing.T) {
	app.Get("/user", func(ctx fiber.Ctx) error {
		//Fungsi ini menerima sebuah objek (dalam hal ini, fiber.Map) yang berisi data yang ingin dikirim sebagai respons.
		// Data tersebut akan secara otomatis diubah menjadi format JSON sebelum dikirim ke klien.
		return ctx.JSON(fiber.Map{
			"username": "Dam",
			"name":     "Dam Dim",
		})
	})

	request := httptest.NewRequest("GET", "/user", nil)
	request.Header.Set("Accept", "application/json")
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, `{"name":"Dam Dim","username":"Dam"}`, string(bytes))
}

func TestRoutingGroup(t *testing.T) {
	helloWorld := func(ctx fiber.Ctx) error {
		return ctx.SendString("Hello World")
	}

	// group mirip prefix untuk rute, yang memungkinkan kita untuk mengelompokkan beberapa rute di bawah satu path dasar.
	api := app.Group("/api")
	api.Get("/hello", helloWorld) // /api/hello
	api.Get("/world", helloWorld) // /api/world

	web := app.Group("/web")
	web.Get("/hello", helloWorld) // /web/hello
	web.Get("/world", helloWorld) // /web/world

	request := httptest.NewRequest("GET", "/api/hello", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello World", string(bytes))
}
