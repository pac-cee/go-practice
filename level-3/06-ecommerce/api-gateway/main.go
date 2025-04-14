package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

type Gateway struct {
	productServiceURL string
	orderServiceURL   string
	userServiceURL    string
}

func NewGateway() *Gateway {
	return &Gateway{
		productServiceURL: os.Getenv("PRODUCT_SERVICE_URL"),
		orderServiceURL:   os.Getenv("ORDER_SERVICE_URL"),
		userServiceURL:    os.Getenv("USER_SERVICE_URL"),
	}
}

func (g *Gateway) createProxy(targetURL string) *httputil.ReverseProxy {
	url, err := url.Parse(targetURL)
	if err != nil {
		log.Fatal(err)
	}
	return httputil.NewSingleHostReverseProxy(url)
