package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Context struct {
	TitleHost string
	ImageName string
	Image     string
	Host      string
}

// https://www.digitalocean.com/community/tutorials/how-to-define-and-call-functions-in-go-es
func main() {
	fmt.Printf("Programa ejecutandose")
	const doc = `
	<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@4.0.0/dist/css/bootstrap.min.css"
        integrity="sha384-Gn5384xqQ1aoWXA+058RXPxPg6fy4IWvTNh0E263XmFcJlSAwiGgFAW/dAiS6JXm" crossorigin="anonymous">
    <link rel="stylesheet" href="css/index.css">
    <title>{{.TitleHost}}</title>
</head>

<body>
    <div class="container rounded bg-light py-5">
        <center>
            <h1>{{.ImageName}}</h1>
            <div id="carouselMultiItemExample" class="carousel slide carousel-dark text-center"
                data-mdb-ride="carousel">

                <!-- Inner -->
                <center>
                                <div class="col-lg-4">     
                                        <img src= "data:image/png;base64 {{.Image}} "
                                            class="card-img-top" alt="Waterfall" />
                                          
                                </div>

                                <h3>{{.Host}}</h2>
                </center>

        </center>
        <!-- Inner -->
    </div>
    </div>

    <script src="https://code.jquery.com/jquery-3.2.1.slim.min.js"
        integrity="sha384-KJ3o2DKtIkvYIK3UENzmM7KCkRr/rE9/Qpg6aAZGJwFDMVNA/GpGFF93hXpG5KkN"
        crossorigin="anonymous"></script>
    <script src="https://cdn.jsdelivr.net/npm/popper.js@1.12.9/dist/umd/popper.min.js"
        integrity="sha384-ApNbgh9B+Y1QKtv3Rn7W3mgPxhU9K/ScQsAP7hUibX39j7fakFPskvXusvfa0b4Q"
        crossorigin="anonymous"></script>
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@4.0.0/dist/js/bootstrap.min.js"
        integrity="sha384-JZR6Spejh4U02d8jOt6vLEHfe/JQGiRRSQQxSfFWpi1MquVdAyjUar5+76PVCmYl"
        crossorigin="anonymous"></script>
</body>

</html>
	`
	eleccion := flag.String("archivos", "archivos", "Directorio de archivos")
	puerto := flag.String("puerto", "3600", "Servidor HTTP")

	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {

		var imagen string
		var imagenContenido string

		imagen, imagenContenido = leerNombreArchivos(*eleccion)

		w.Header().Add("Content Type", "assets/html")
		templates, _ := template.New("doc").Parse(doc)
		context := Context{ // Ejemplo futuro para llamar variables
			TitleHost: imprimirHost(),
			ImageName: imagen,
			Image:     imagenContenido,
			Host:      imprimirHost(),
		}
		templates.Lookup("doc").Execute(w, context)
	})
	http.ListenAndServe(":"+*puerto, nil)

}

//https://golang.cafe/blog/golang-random-number-generator.html
//https://freshman.tech/snippets/go/image-to-base64/

func leerNombreArchivos(eleccion string) (imagen string, imagenContenido string) {

	archivos, err := ioutil.ReadDir(eleccion)

	imagenes := []string{}
	if err != nil {
		log.Fatal(err)
	}
	for _, archivo := range archivos {

		if archivo.Name()[len(archivo.Name())-3:len(archivo.Name())] == "jpg" ||
			archivo.Name()[len(archivo.Name())-4:len(archivo.Name())] == "jpeg" ||
			archivo.Name()[len(archivo.Name())-3:len(archivo.Name())] == "png" {
			imagenes = append(imagenes, archivo.Name())
		}

	}
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := len(imagenes) - 1
	aux := rand.Intn(max-min+1) + min

	imagen = imagenes[aux]

	bytes, err := ioutil.ReadFile(eleccion + "/" + imagenes[aux])
	if err != nil {
		log.Fatal(err)
	}

	var base64Encoding string

	// Determine the content type of the image file
	mimeType := http.DetectContentType(bytes)

	// Prepend the appropriate URI scheme header depending
	// on the MIME type
	switch mimeType {

	case "image/jpg":
		base64Encoding += "data:image/jpg;base64,"
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}

	// Append the base64 encoded output
	base64Encoding += toBase64(bytes)

	// Print the full base64 representation of the image
	//fmt.Println(base64Encoding)

	imagenContenido = base64Encoding

	return imagen, imagenContenido
}

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

// https://freshman.tech/snippets/go/image-to-base64/
func imprimirHost() string {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//fmt.Printf("Hostname: %s", hostname)
	return hostname
}

func htmlPage() {
	fs := http.FileServer(http.Dir("./"))
	http.Handle("/", fs)

	log.Print("Abierto el puerto 4 lukitas")

	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
