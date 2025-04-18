package commons

import (
	"math"
	"math/rand"
	"snapp/config"
	"strconv"
	"time"
)

func MakeOtp() string {
	cfg := config.GetConfig()
	rand.Seed(time.Now().UTC().UnixNano())
	min := int(math.Pow(float64(10), float64(cfg.Otp.Digit)-1))
	max := int(math.Pow(float64(10), float64(cfg.Otp.Digit)))-1
	otp := rand.Intn(max - min) + min
	return strconv.Itoa(otp)
}