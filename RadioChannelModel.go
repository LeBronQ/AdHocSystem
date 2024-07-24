package AdHocSystem

import (
	"github.com/LeBronQ/RadioChannelModel"
)

const (
	PacketSize = 10000
	Redundancy = 1
)

func ChannelCalculation(LinkID int64, Distance float64, LargeScaleModel string, SmallScaleModel string, Frequency float64, BitRate float64, Mod string, BW float64, M float64, PowerInDbm float64) float64 {
	PathLoss, Fading, BER := 0.0, 0.0, 0.0
	switch LargeScaleModel {
	case "FreeSpacePathLossModel":
		PLParam := RadioChannelModel.FreeSpaceParam{
			Distance:     Distance,
			Frequency:    Frequency,
			TXPowerInDbm: PowerInDbm,
		}
		PathLoss = RadioChannelModel.FreeSpacePathLoss(PLParam)
		break
	case "LogDistancePathLossModel":
	case "":
	}
	switch SmallScaleModel {
	case "NakagamiFadingModel":
		FParam := RadioChannelModel.NakagamiParam{
			TXPowerInDbm: PathLoss,
			Scenario:     "open_filed",
			Elevation:    0,
		}
		Fading = RadioChannelModel.NakagamiFadingModel(FParam)
	}
	SNR := RadioChannelModel.CalculateSNR(BW, Fading, 0)
	switch Mod {
	case "BPSK":
		BParam := RadioChannelModel.BPSKParam{
			Bandwidth: BW,
			SNR:       SNR,
			BitRate:   BitRate,
		}
		BER = RadioChannelModel.CalculateBPSKBER(BParam)
	case "QAM":
		QParam := RadioChannelModel.QAMParam{
			Bandwidth: BW,
			SNR:       SNR,
			BitRate:   BitRate,
			M:         M,
		}
		BER = RadioChannelModel.CalculateQAMBER(QParam)
	}
	TParam := RadioChannelModel.TransportParam{
		BER:             BER,
		PacketSizeInBit: PacketSize,
		Redundancy:      Redundancy,
	}
	PLR := RadioChannelModel.CalculatePLR(TParam)
	return PLR
}
