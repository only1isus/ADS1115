package ADS1115

import (
	"encoding/binary"
	"fmt"
	"strings"
	"time"

	"golang.org/x/exp/io/i2c"
)

// ADS1115 ...
type ADS1115 struct {
	Device   *i2c.Device
	dataRate uint16
	gain     uint16
	channel  uint16
}

// NewConnection creates a connection to the I2C device on the given address.
func NewConnection(name string, address int) (*i2c.Device, error) {
	name = strings.ToLower(name)
	device, err := i2c.Open(&i2c.Devfs{Dev: fmt.Sprintf("/dev/%v", name)}, address)
	if err != nil {
		return nil, err
	}
	return device, nil
}

// NewADS1115Device creates a new instance of the ADS1115 struct.
func NewADS1115Device(device *i2c.Device) *ADS1115 {
	adc := new(ADS1115)
	adc.Device = device
	return adc
}

// Gain sets the voltage gain of the device.
func (ads *ADS1115) Gain(gain AmplifierGain) {
	ads.gain = uint16(gain)
}

// DataRate sets the datarate of the channel
func (ads *ADS1115) DataRate(rate DataRate) {
	ads.dataRate = uint16(rate)
}

// Channel takes an integer from 0-3 which represents the channel to use for reading the analog signal
func (ads *ADS1115) Channel(channel int) {
	channels := map[int]MultiplexerInput{
		0: ADS1115MultiplexerConfigurationAIN0,
		1: ADS1115MultiplexerConfigurationAIN1,
		2: ADS1115MultiplexerConfigurationAIN2,
		3: ADS1115MultiplexerConfigurationAIN3,
	}
	c := channels[channel]
	ads.channel = uint16(c)
}

// Read reads the analog value at pin AIN0. Note, Read uses the default settings found in consts.
func (ads *ADS1115) Read() (float32, error) {
	holder := make([]byte, 2)
	var out []byte

	config := uint16(ADS1115Config) | uint16(ads.gain) | uint16(ads.channel)
	// fmt.Println("final config bytes ", (ADS1115Config>>8)&0xFF)
	// fmt.Println("final config bytes ", (ADS1115Config & 0xFF))
	binary.BigEndian.PutUint16(holder, (config>>8)&0xFF)
	out = append(out, holder[1])
	binary.BigEndian.PutUint16(holder, (config & 0xFF))
	out = append(out, holder[1])
	err := ads.Device.WriteReg(ADS1115RegisterPointerConfig, out)
	fmt.Println(out, config)
	if err != nil {
		return float32(-1), err
	}

	delay := 8
	time.Sleep(time.Duration(delay) * time.Millisecond)

	var result = make([]byte, 2)
	err = ads.Device.ReadReg(ADS1115RegisterConversionConfig, result)
	if err != nil {
		return float32(-1), err
	}
	val := (uint32(result[0]) << 8) | uint32(result[1])
	var data float32
	if val > 0x7FFF {
		data = float32((val-0xFFFF)*6144/1000) / 32768.0
	} else {
		data = float32(val*6144/1000) / 32768.0
	}
	result = make([]byte, 2)
	return data, nil
}

// Close should be called when the device will no longer be used
func (ads *ADS1115) Close() error {
	return ads.Device.Close()
}
