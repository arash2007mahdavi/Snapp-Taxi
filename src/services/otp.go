package services

import (
	"fmt"
	"snapp/config"
	"snapp/database"
	"snapp/loggers"
	"time"

	"github.com/redis/go-redis/v9"
)

type otpService struct {
	logger loggers.Logger
	cfg *config.Config
	redis *redis.Client
}

type otpValue struct {
	Value string
	Used bool
}

func NewotpService(cfg *config.Config) *otpService {
	logger := loggers.NewLogger(cfg)
	redis := database.GetRedis()
	return &otpService{
		logger: logger,
		redis: redis,
		cfg: cfg,
	}
}

func (s *otpService) SetOtp(mobileNumber string, otp string) error {
	value := &otpValue{
		Value: otp,
		Used: false,
	}

	res, err := database.Get[otpValue](s.redis, mobileNumber)
	if err == nil && !res.Used {
		return fmt.Errorf("otp exists")
	} else if err == nil && res.Used {
		return fmt.Errorf("otp used")
	}
	err = database.Set(s.redis, mobileNumber, value, s.cfg.Otp.Expire*time.Second)
	if err != nil {
		return err
	}
	return nil
}

func (s *otpService) ValidateOtp(mobileNumber string, otp string) error {
	res, err := database.Get[otpValue](s.redis, mobileNumber)
	if err != nil {
		return err
	} else if err == nil && res.Used {
		return fmt.Errorf("otp used")
	} else if err == nil && !res.Used && res.Value != otp {
		return fmt.Errorf("invalid otp")
	} else if err == nil && !res.Used && res.Value == otp {
		res.Used = true
		err = database.Set(s.redis, mobileNumber, otp, s.cfg.Otp.Expire*time.Second)
		if err != nil {
			return err
		}
	}
	return nil
}