// template
// skeleton of algorithm in base class and subclasses override
// steps w/out changing algorithim's structure

package main

import "fmt"

func main() {
	smsOTP := &sms{}
	o := otp{
		iOtp: smsOTP,
	}
	o.genAndSendOTP(4)
	fmt.Println("")
	emailOTP := &email{}
	o = otp{
		iOtp: emailOTP,
	}
	o.genAndSendOTP(4)
}
