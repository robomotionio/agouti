package appium

import "github.com/robomotionio/agouti"

func NewTestDevice(session mobileSession) *Device {
	return &Device{
		Page:    &agouti.Page{},
		session: session,
	}
}
