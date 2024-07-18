package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

type PriceResponse struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
}

type Candle struct {
	OpenTime  int64   `json:"open_time"`
	Open      float64 `json:"open"`
