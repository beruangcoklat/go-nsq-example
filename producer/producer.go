package main

import (
	"net/http"

	"github.com/beruangcoklat/go-nsq/config"
	"github.com/bitly/go-nsq"
	"github.com/gorilla/mux"
)

func producer() *nsq.Producer {
	cfg := nsq.NewConfig()
	producer, err := nsq.NewProducer(config.NSQLOOKUPD_URL, cfg)
	if err != nil {
		panic(err)
	}
	return producer
}

func publish(w http.ResponseWriter, r *http.Request) {
	msg := r.URL.Query().Get("msg")
	producer := producer()
	if err := producer.Publish(config.NSQ_TOPIC, []byte(msg)); err != nil {
		panic(err)
	}
	producer.Stop()
}

func router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/publish", publish)
	return r
}

func main() {
	if err := http.ListenAndServe(":"+config.PORT, router()); err != nil {
		panic(err)
	}
}
