package shodan

type PingResponse struct {
	Ip              string    `json:"ip"`
	IsAlive         bool      `json:"is_alive"`
	MinRtt          float64   `json:"min_rtt"`
	AvgRtt          float64   `json:"avg_rtt"`
	MaxRtt          float64   `json:"max_rtt"`
	Rtts            []float64 `json:"rtts"`
	PacketsSent     int       `json:"packets_sent"`
	PacketsReceived int       `json:"packets_received"`
	PacketLoss      int       `json:"packet_loss"`
	FromLoc         struct {
		City    string `json:"city"`
		Country string `json:"country"`
		Latlon  string `json:"latlon"`
	} `json:"from_loc"`
}

type GeoNetResponse []PingResponse

const (
	Ping    = geoNetUri + "/ping/"
	GeoPing = geoNetUri + "/geoping/"
)
