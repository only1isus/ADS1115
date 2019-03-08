package ADS1115

type (
	// DataRate is the data rate of the device
	DataRate uint16
	// AmplifierGain sets the gain of the channel
	AmplifierGain uint16
	// MultiplexerInput is the channel being used
	MultiplexerInput uint16
)

const (
	ADS1115ComparatorQueue     uint16 = 0x0003
	ADS1115LatchingComparator  uint16 = 0x0000
	ADS1115ComparatorPolarity  uint16 = 0x0000
	ADS1115ComparatorMode      uint16 = 0x0001
	ADS1115DeviceOperationMode uint16 = 0x0000
	ADS1115OperationalStatus   uint16 = 0x8000

	ADS1115RegisterPointerConfig    byte          = 0x01
	ADS1115RegisterConversionConfig byte          = 0x00
	ADS1115DataRate                 DataRate      = 0x0007
	ADS1115ProgramableGainAmplifier AmplifierGain = 0x0000

	ADS1115MultiplexerConfigurationAIN0 MultiplexerInput = 0x4000
	ADS1115MultiplexerConfigurationAIN1 MultiplexerInput = 0x5000
	ADS1115MultiplexerConfigurationAIN2 MultiplexerInput = 0x6000
	ADS1115MultiplexerConfigurationAIN3 MultiplexerInput = 0x7000

	ADS1115Config = ADS1115ComparatorQueue |
		ADS1115LatchingComparator |
		ADS1115ComparatorPolarity |
		ADS1115ComparatorMode |
		ADS1115DeviceOperationMode |
		ADS1115OperationalStatus
)
