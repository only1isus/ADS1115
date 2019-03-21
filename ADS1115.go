package ADS1115

import (
	"encoding/binary"
	"fmt"
	"i2c/consts"
	"strings"
	"time"

	"golang.org/x/exp/io/i2c"
)

// ADS1115 ...
type ADS1115 struct {
	Device   *i2c.Device
	config   uint16
	dataRate int
	gain     uint16
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

func getDataRate(rate int) (int, uint16, error) {
	dr := map[int]uint16{
		8:   ADS1115DataRate8,
		16:  ADS1115DataRate16,
		32:  ADS1115DataRate32,
		64:  ADS1115DataRate64,
		128: ADS1115DataRate128,
		250: ADS1115DataRate250,
		475: ADS1115DataRate475,
		860: ADS1115DataRate860,
	}
	datarate, ok := dr[rate]
	if !ok {
		return 0, 0, fmt.Errorf("no value found for the key entered. Use %v", dr)
	}
	// ads.dataRate = datarate
	return rate, datarate, nil
}

// getChannel takes an integer from 0-3 which represents the channel to use for reading the analog signal
func getChannel(channel int) (uint16, error) {
	if channel < 0 || channel > 3 {
		return 0, fmt.Errorf("the channel should be between 0 and 3")
	}
	channels := map[int]uint16{
		0: ADS1115MultiplexerConfigurationAIN0,
		1: ADS1115MultiplexerConfigurationAIN1,
		2: ADS1115MultiplexerConfigurationAIN2,
		3: ADS1115MultiplexerConfigurationAIN3,
	}
	c := channels[channel]
	return c, nil
}

// Gain sets the voltage gain of the device.
func (ads *ADS1115) Gain(gain uint16) {
	ads.gain = uint16(gain)
}

// DataRate sets the datarate of the channel
func (ads *ADS1115) DataRate(rate int) error {
	ads.dataRate = rate
	return nil

}

// Read reads the analog value at pin (channel). Note, Read uses the default settings found in consts.
func (ads *ADS1115) Read(channel int) (float32, error) {
	holder := make([]byte, 2)
	var out []byte

	if ads.dataRate == 0 {
		ads.dataRate = 128
	}
	_, dataRate, err := getDataRate(ads.dataRate)
	if err != nil {
		return 0, err
	}

	channelToRead, err := getChannel(channel)
	if err != nil {
		return 0, nil
	}
	if ads.gain == 0 {
		ads.gain = uint16(ADS1115ProgramableGainAmplifier6144)
	}

	ads.config = uint16(ADS1115Config) | ads.gain | channelToRead | dataRate
	ads.config = ads.config + 100 // add 100 to the config value in order to get the correct channel reading
	binary.BigEndian.PutUint16(holder, (ads.config>>8)&0xFF)
	out = append(out, holder[1])
	binary.BigEndian.PutUint16(holder, (ads.config & 0xFF))
	out = append(out, holder[1])

	err = ads.Device.Write([]byte{0x00})
	if err != nil {
		return float32(-1), err
	}

	err = ads.Device.WriteReg(ADS1115RegisterPointerConfig, out)
	if err != nil {
		return float32(-1), err
	}

	time.Sleep(time.Duration(1000000/ads.dataRate+100) * time.Microsecond)

	result := make([]byte, 2)
	err = ads.Device.ReadReg(consts.ADS1115RegisterConversionConfig, result)
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
	return data, nil
}

// Close should be called when the device will no longer be used
func (ads *ADS1115) Close() error {
	return ads.Device.Close()
}
