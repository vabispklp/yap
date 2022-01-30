package rest

import (
	"compress/gzip"
	"github.com/go-chi/chi/v5"
	"github.com/vabispklp/yap/api/rest/handlers"
	"github.com/vabispklp/yap/internal/app/service/shortener"
	"io"
	"net/http"
	"strings"
)

func initRouter(shortener *shortener.Shortener) (*chi.Mux, error) {
	router := chi.NewRouter()
	router.Use(gzipHandle)

	h, err := handlers.NewHandler(shortener)
	if err != nil {
		return nil, err
	}

	router.Get("/{id}", h.GetHandleGetURL())
	router.Post("/", h.GetHandlerAddURL())
	router.Post("/api/shorten", h.GetHandlerAddShorten())

	return router, nil
}

type gzipWriter struct {
	http.ResponseWriter
	Writer io.Writer
}

func (w gzipWriter) Write(b []byte) (int, error) {
	// w.Writer будет отвечать за gzip-сжатие, поэтому пишем в него
	return w.Writer.Write(b)
}

func gzipHandle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// проверяем, что клиент поддерживает gzip-сжатие
		if r.Header.Get("Content-Encoding") == "" {
			next.ServeHTTP(w, r)
			return
		}

		if r.Header.Get("Content-Encoding") != "gzip" {
			io.WriteString(w, "unsupported compression")
			return
		}

		//var buf bytes.Buffer
		//
		//gz1 := gzip.NewWriter(&buf)
		//
		//defer gz1.Close()
		//
		//b, err := io.ReadAll(r.Body)
		//if err != nil {
		//	io.WriteString(w, err.Error())
		//	return
		//}
		//_, err = gz1.Write(b)
		//if err != nil {
		//	io.WriteString(w, err.Error())
		//	return
		//}
		//
		//reader1, err := gzip.NewReader(io.NopCloser(bytes.NewReader(buf.Bytes())))
		//if err != nil {
		//	io.WriteString(w, err.Error())
		//	return
		//}

		reader, err := gzip.NewReader(r.Body)
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}

		r.Body = reader
		defer reader.Close()

		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
			return
		}

		gz, err := gzip.NewWriterLevel(w, gzip.BestSpeed)
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}
		defer gz.Close()

		w.Header().Set("Content-Encoding", "gzip")

		// передаём обработчику страницы переменную типа gzipWriter для вывода данных
		next.ServeHTTP(gzipWriter{ResponseWriter: w, Writer: gz}, r)
	})
}
