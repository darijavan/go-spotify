package main

import (
	"encoding/json"
	"fmt"

	"github.com/tsirysndr/go-spotify"
)

func main() {
	// Replace with your ACCESS TOKEN
	token := "BQCn4whGh-ZpqYC5CjeLQMczQGtI0ZK_56RypCO7gNTIC5sGIY9ImxtkHlYhs3NJRvN9ql9qrcZNgDmrD8f3i2bYcXQ5mitYvSHW0Gyyiem8v_ntqDpPZAkXshAueNNLSSw8bqfvxwezdSRAMqpOUHL_Xi2tB1dM2Soxh3VmpDPWdwnvWboiIBq6d7tng_SeL9tvWtTG8P1QvVIx2LtgNcESpcmhw1Gmafu3V1ZQVOI-Lsbkbo3YjPWJig_Qo8R88QpcMZ351AJrOlf1sMo5hSLGgrjYXA"
	client := spotify.NewClient(token)
	result, _ := client.Track.AudioFeatures("7ouMYWpwJ422jRcDASZB7P,4VqPOruhp5EdPBeR92t6lQ,2takcwOaAZWiXQijPHIx7B")
	features, _ := json.Marshal(result)
	fmt.Println(string(features))
}
