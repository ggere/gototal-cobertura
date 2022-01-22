package totalcover

import (
	"bufio"
	"encoding/xml"
	"io"
	"io/ioutil"
	"strconv"
)

type coberturaXML struct {
	LineRate string `xml:"line-rate,attr"`
}

func TotalCover(input io.Reader) (total float64, err error) {
	reader := bufio.NewReader(input)
	byteValue, err := ioutil.ReadAll(reader)
	if err != nil {
		return
	}

	var c coberturaXML
	err = xml.Unmarshal(byteValue, &c)
	if err != nil {
		return
	}

	percent, err := strconv.ParseFloat(c.LineRate, 64)
	if err != nil {
		return
	}
	return percent * 100, nil
}
