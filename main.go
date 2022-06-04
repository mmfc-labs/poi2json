package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

type Color string

const (
	Red   Color = "red"
	Blue  Color = "blue"
	Green Color = "green"
	gray  Color = "gray"
)

func (c *Color) String() string {
	return string(*c)
}

func NewColor(c string) (Color, error) {
	switch c {
	case "red":
		return Red, nil
	case "blue":
		return Blue, nil
	case "green":
		return Green, nil
	case "gray":
		return gray, nil
	case "红色":
		return Red, nil
	default:
		return "", fmt.Errorf("invalid color: %s", c)
	}
}

type Pois struct {
	Points []Point `json:"pois"`
}

type Point struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Color Color   `json:"color"`
	Lat   float64 `json:"lat"`
	Lon   float64 `json:"lon"`
}

func init() {
	rootCmd.AddCommand(toJsonCmd)
	toJsonCmd.Flags().StringP("input", "i", "", "input file")

}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{}

var toJsonCmd = &cobra.Command{
	Use:   "tojson",
	Short: "Convert a file to JSON",
	Long:  `Convert a file to JSON`,
	Run: func(cmd *cobra.Command, args []string) {
		input, _ := cmd.Flags().GetString("input")
		if input == "" {
			fmt.Println("input file is required")
			os.Exit(1)
		}
		pois, err := mapDataToJson(input)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		poisByte, err := json.Marshal(pois)
		if err != nil {
			fmt.Println(string(poisByte))
		}
		os.WriteFile("map.json", poisByte, 0644)
		fmt.Printf("%+v\n", string(poisByte))
	},
}

func mapDataToJson(file string) (Pois, error) {
	var pois []Point
	fByte, err := os.ReadFile(file)
	if err != nil {
		return Pois{}, err
	}
	fStrArr := strings.Split(string(fByte), "\n")
	for _, fStr := range fStrArr {
		poiArr := strings.Split(fStr, " ")
		if len(poiArr) < 3 {
			// fmt.Printf("invalid input: %v\n\n", poiArr)
			continue
		}
		var poi Point
		ll := strings.Split(poiArr[0], ",")
		if len(ll) != 2 {
			// fmt.Printf("invalid input: %v\n\n", ll)
			continue
		}
		lat, err := strconv.ParseFloat(ll[0], 64)
		if err != nil {
			// fmt.Printf("invalid input: %v\n\n", ll[0])
			continue
		}
		lon, err := strconv.ParseFloat(ll[1], 64)
		if err != nil {
			// fmt.Printf("invalid input: %v\n\n", ll[1])
			continue
		}
		poi.Lat = lat
		poi.Lon = lon
		poi.Name = poiArr[1]
		if poi.Id, err = strconv.Atoi(strings.TrimSpace(poiArr[2])); err != nil {
			poi.Id = 0
		}
		if len(poiArr) >= 4 {
			c, err := NewColor(strings.TrimSpace(poiArr[3]))
			if err != nil {
				c = Red
				// return Pois{}, fmt.Errorf("invalid color: %s", poiArr[3])
			}
			poi.Color = c
		}
		pois = append(pois, poi)
	}
	// fmt.Printf("len pois: %v\n", len(pois))
	return Pois{pois}, nil
}
