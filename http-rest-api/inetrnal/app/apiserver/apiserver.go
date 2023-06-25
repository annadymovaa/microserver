package apiserver

import (
	//"io"
	"log"
	"net/http"
	"text/template"

	//"path/filepath"
	"github.com/annadymovaa/microserver/inetrnal/app/model"
	"github.com/annadymovaa/microserver/inetrnal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

var (
	tmpl1 = template.Must(template.ParseFiles("/Volumes/ssd adata/microserver/http-rest-api/templates/forms.html"))
	tmpl  = template.Must(template.ParseFiles("/Volumes/ssd adata/microserver/http-rest-api/templates/main.html"))
)

func handler1(w http.ResponseWriter, r *http.Request) {
	data1 := model.User{
		Name:    r.FormValue("name"),
		Surname: r.FormValue("surname"),
		Email:   r.FormValue("email"),
		Pass:    r.FormValue("password"),
	}

	tmpl1.Execute(w, data1)
	//

}

// tmpl2 = template.Must(template.ParseFiles("/Volumes/ssd adata/microserver/http-rest-api/templates/transaction_forms.html"))

// var (
// 		price = r.FormValue("price")
// 		from_acc = r.FormValue("email_from")
// )
// if s, err := strconv.ParseFloat(price, 64); err != nil {
// 	if f, err := store.Repository.FindByEmail(from_acc); err != nil{
// 		data2 := model.Transaction{

// 			From_acc: f,
// 			To_acc: store.Repository.FindByEmail(r.FormValue("email_to")),
// 			Price: s,
// 		}
// 	}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	if err := s.configureStore(); err != nil {
		return err
	}

	s.logger.Info("starting api server")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil

}

func (s *APIServer) configureRouter() {
	http.HandleFunc("/register", handler1)
	log.Fatal(http.ListenAndServe(":8080", nil))
	http.Handle("/", http.FileServer(http.Dir("main.html")))

}

func (s *APIServer) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st

	return nil
}

// func (s *APIServer) handleHello() http.HandlerFunc {

// 	return func(w http.ResponseWriter, r *http.Request) {
// 		 tmpl.Execute(w, )
// 	}
// }
