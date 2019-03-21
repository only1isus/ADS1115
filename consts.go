package ADS1115

// The following values were taken from the official product datasheet.
// http://www.ti.com/lit/ds/symlink/ads1115.pdf

const (
	// ADS1115ComparatorQueueAssertAfterOne Assert after one conversion
	ADS1115ComparatorQueueAssertAfterOne uint16 = 0x0000
	// ADS1115ComparatorQueueAssertAfterTwo Assert after two conversions
	ADS1115ComparatorQueueAssertAfterTwo uint16 = 0x0001
	// ADS1115ComparatorQueueAssertAfterFour Assert after four conversions
	ADS1115ComparatorQueueAssertAfterFour uint16 = 0x0002
	// ADS1115ComparatorQueueDisable Disable comparator and set ALERT/RDY pin to high-impedance (default)
	ADS1115ComparatorQueueDisable uint16 = 0x0003
	// ADS1115LatchingComparatorLatching The ALERT/RDY pin does not latch when asserted (default)
	ADS1115LatchingComparatorLatching uint16 = 0x0000
	// ADS1115LatchingComparatorNonLatching The asserted ALERT/RDY pin remains latched until
	// conversion data are read by the master or an appropriate SMBus alert response
	// is sent by the master
	ADS1115LatchingComparatorNonLatching uint16 = 0x0004
	// ADS1115ComparatorPolarityActiveLow This bit controls the polarity of the ALERT/RDY pin (default)
	ADS1115ComparatorPolarityActiveLow uint16 = 0x0008
	// ADS1115ComparatorPolarityActiveHigh This bit controls the polarity of the ALERT/RDY pin
	ADS1115ComparatorPolarityActiveHigh uint16 = 0x0001
	// ADS1115ComparatorModeTraditional this bit configures the comparator operating mode. (default)
	ADS1115ComparatorModeTraditional uint16 = 0x0010
	// ADS1115ComparatorModeWindow this bit configures the comparator operating mode.
	ADS1115ComparatorModeWindow uint16 = 0x0001

	// ADS1115DeviceOperationModeContinous Continuous-conversion mode
	ADS1115DeviceOperationModeContinous uint16 = 0x0000
	// ADS1115DeviceOperationModeSingleShot  Single-shot mode or power-down state
	ADS1115DeviceOperationModeSingleShot uint16 = 0x0100

	// ADS1115OperationalStatus determines the operational status of the device. OS can only be written
	// when in power-down state and has no effect when a conversion is ongoing
	ADS1115OperationalStatus uint16 = 0x8000

	// ADS1115RegisterPointerConfig ...
	ADS1115RegisterPointerConfig byte = 0x01
	// ADS1115RegisterConversionConfig Conversion register contains the result of the last conversion in binary two's complement format.
	ADS1115RegisterConversionConfig byte = 0x00

	// ADS1115DataRate8 control the data rate setting. 8 Sample Per Seconds
	ADS1115DataRate8 uint16 = 0x0000
	// ADS1115DataRate16 control the data rate setting. 16 Sample Per Seconds
	ADS1115DataRate16 uint16 = 0x0001
	// ADS1115DataRate32 control the data rate setting. 32 Sample Per Seconds
	ADS1115DataRate32 uint16 = 0x0002
	// ADS1115DataRate64 control the data rate setting. 64 Sample Per Seconds
	ADS1115DataRate64 uint16 = 0x0003
	// ADS1115DataRate128  control the data rate setting. 128 Sample Per Seconds
	ADS1115DataRate128 uint16 = 0x0004
	// ADS1115DataRate250 control the data rate setting. 250  Sample Per Seconds
	ADS1115DataRate250 uint16 = 0x0005
	// ADS1115DataRate475 control the data rate setting. 475 Sample Per Seconds
	ADS1115DataRate475 uint16 = 0x0006
	// ADS1115DataRate860 control the data rate setting. 860 Sample Per Seconds
	ADS1115DataRate860 uint16 = 0x0007

	// ADS1115ProgramableGainAmplifier6144 These bits set the FSR of the programmable gain amplifier. For voltages in the range ±6.144
	ADS1115ProgramableGainAmplifier6144 uint16 = 0x0000
	// ADS1115ProgramableGainAmplifier4096 set the FSR of the programmable gain amplifier. For voltages in the range ±4.096
	ADS1115ProgramableGainAmplifier4096 uint16 = 0x0200
	// ADS1115ProgramableGainAmplifier2048 set the FSR of the programmable gain amplifier. For voltages in the range ±2.048
	ADS1115ProgramableGainAmplifier2048 uint16 = 0x0002
	// ADS1115ProgramableGainAmplifier1024 set the FSR of the programmable gain amplifier. For voltages in the range ±1.024
	ADS1115ProgramableGainAmplifier1024 uint16 = 0x0003
	// ADS1115ProgramableGainAmplifier0512 set the FSR of the programmable gain amplifier. For voltages in the range ±0.512
	ADS1115ProgramableGainAmplifier0512 uint16 = 0x0004
	// ADS1115ProgramableGainAmplifier0256_0 set the FSR of the programmable gain amplifier. For voltages in the range ±0.256
	ADS1115ProgramableGainAmplifier0256_0 uint16 = 0x0005
	// ADS1115ProgramableGainAmplifier0256_1 set the FSR of the programmable gain amplifier. For voltages in the range ±0.256
	ADS1115ProgramableGainAmplifier0256_1 uint16 = 0x0006
	// ADS1115ProgramableGainAmplifier0256_2 set the FSR of the programmable gain amplifier. For voltages in the range ±0.256
	ADS1115ProgramableGainAmplifier0256_2 uint16 = 0x0007

	// ADS1115MultiplexerConfigurationAIN0 AINP = AIN0 and AINN = GND
	ADS1115MultiplexerConfigurationAIN0 uint16 = 0x4000
	// ADS1115MultiplexerConfigurationAIN1 AINP = AIN1 and AINN = GND
	ADS1115MultiplexerConfigurationAIN1 uint16 = 0x5000
	// ADS1115MultiplexerConfigurationAIN2 AIN2 = AIN2 and AINN = GND
	ADS1115MultiplexerConfigurationAIN2 uint16 = 0x6000
	// ADS1115MultiplexerConfigurationAIN3 AIN3 = AIN3 and AINN = GND
	ADS1115MultiplexerConfigurationAIN3 uint16 = 0x7000

	// ADS1115Config default configuration
	ADS1115Config = ADS1115ComparatorQueueDisable |
		ADS1115LatchingComparatorNonLatching |
		ADS1115ComparatorPolarityActiveLow |
		ADS1115ComparatorModeTraditional |
		ADS1115DeviceOperationModeSingleShot |
		ADS1115OperationalStatus
)
