package main

import "fmt"

type sms struct {
	otp
}

func (s *sms) genRandomOTP(len int) string {
	randomOTP := "1234"
	fmt.Printf("SMS: generating random one time password %s \n", randomOTP)
	return randomOTP
}

func (s *sms) saveOTPCache(otp string) {
	fmt.Printf("SMS: saving one time password: %s to cache! \n", otp)
}

func (s *sms) getMessage(otp string) string {
	return "SMS OTP for login is: " + otp
}

func (s *sms) sendNotification(message string) error {
	fmt.Printf("SMS: sending SMS: %s \n", message)
	return nil
}

func (s *sms) publishMetric() {
	fmt.Printf("SMS: publishing metrics! \n")
}
