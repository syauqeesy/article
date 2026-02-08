package foundation

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"ahmadsyauqi.dev/article/common"
	"ahmadsyauqi.dev/article/configuration"
	"ahmadsyauqi.dev/article/handler"
	"ahmadsyauqi.dev/article/repository"
	"ahmadsyauqi.dev/article/service"
)

type httpFoundation struct {
	configuration *configuration.Configuration
	mux           *http.ServeMux
	server        *http.Server
	database      *databaseFoundation
	repository    *repository.Repository
	service       *service.Service
	handler       *handler.Handler
}

func (f *httpFoundation) Setup() error {
	f.mux = http.NewServeMux()

	f.database = &databaseFoundation{
		configuration: f.configuration,
	}

	err := f.database.Setup()
	if err != nil {
		return err
	}

	err = f.database.Boot()
	if err != nil {
		return err
	}

	f.repository = repository.New(f.database.database)

	f.service = service.New(f.configuration, f.repository)

	f.handler = handler.New(f.mux, f.configuration, f.service)

	var t time.Time
	result := f.database.database.Raw("SELECT NOW() AS THIS_MOMENT").Scan(&t)
	if result.Error != nil {
		log.Fatal("database error : ", result.Error)
	}

	f.mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		common.HttpErrorHandler(w, common.CreateException(http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed)), nil)
	})

	f.server = &http.Server{
		Addr:    f.configuration.HTTP.Port,
		Handler: f.mux,
	}

	return nil
}

func (f *httpFoundation) Boot() error {
	fmt.Println("http server started")

	err := f.server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

func (f *httpFoundation) Shutdown() error {
	err := f.database.Shutdown()
	if err != nil {
		return err
	}

	fmt.Println("shutting down http server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = f.server.Shutdown(ctx)
	if err != nil {
		return err
	}

	fmt.Println("http server exited")

	return nil
}
