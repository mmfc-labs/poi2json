package main

import (
	"encoding/json"
	"errors"
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

type Role string

const (
	User    Role = "user"
	Manager Role = "manager"
)

type Status string

const (
	Open    Status = "open"
	Close   Status = "close"
	Unknown Status = "unknown"
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

type DirectionPoint struct {
	Lat    float64 `json:"lat"`
	Lon    float64 `json:"lon"`
	Auth   Role    `json:"auth"`
	Status Status  `json:"status"`
}

type Point struct {
	Id      int              `json:"id"`
	Name    string           `json:"name"`
	Color   Color            `json:"color"`
	Lat     float64          `json:"lat"`
	Lon     float64          `json:"lon"`
	Towards []DirectionPoint `json:"towards"`
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
		// fmt.Printf("%+v\n", string(poisByte))
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
		poi, err := ParseLine(fStr)
		if err != nil {
			continue
		}
		pois = append(pois, poi)
	}
	// fmt.Printf("len pois: %v\n", len(pois))
	return Pois{pois}, nil
}

func ParseLine(l string) (Point, error) {
	poiArr := strings.Split(l, " ")
	if len(poiArr) < 3 {
		// fmt.Printf("invalid input: %v\n\n", poiArr)
		return Point{}, errors.New("invalid input")
	}
	poi, err := GetPointFromString(poiArr[0])
	if err != nil {
		return Point{}, err
	}
	poi.Name = poiArr[1]
	if poi.Id, err = strconv.Atoi(strings.TrimSpace(poiArr[2])); err != nil {
		poi.Id = 0
	}
	if len(poiArr) >= 4 {
		// fmt.Printf("poiArr[3]: %v\n", poiArr[3])
		c, err := NewColor(strings.TrimSpace(poiArr[3]))
		if err != nil {
			c = Red
			// return Pois{}, fmt.Errorf("invalid color: %s", poiArr[3])
		}
		poi.Color = c
	}
	if len(poiArr) >= 5 {
		// fmt.Printf("poiArr[4]: %s\n", poiArr[4])
		poi.Towards, _ = GetTowards(strings.TrimSpace(poiArr[4]))
	}
	return poi, nil
}

func GetTowards(poiss string) ([]DirectionPoint, error) {
	var towards []DirectionPoint
	if poiss == "" {
		return towards, nil
	}
	poiArrS := strings.Split(poiss, ";")
	for _, poiS := range poiArrS {
		poiS, role, status, err := GetTowardStringAndAuthStatus(poiS)
		if err != nil {
			continue
		}
		poi, err := GeDrectionPointFromString(poiS)
		if err != nil {
			continue
		}
		poi.Auth = role
		poi.Status = status
		towards = append(towards, poi)
	}
	return towards, nil
}

func GetTowardStringAndAuthStatus(s string) (string, Role, Status, error) {
	var poiString string
	var role Role
	var status Status
	if s == "" {
		return poiString, role, status, errors.New("invalid input: empty string")
	}
	poiString = s
	sArr := strings.Split(s, "-")
	poiString = sArr[0]
	if len(sArr) >= 2 {
		switch strings.TrimSpace(sArr[1]) {
		case "m":
			role = Manager
			status = Open
		case "u":
			role = User
			status = Open
		case "uc":
			role = User
			status = Close
		case "U":
			role = Manager
			status = Unknown
		case "c":
			role = Manager
			status = Close
		default:
			role = Manager
			status = Unknown
		}
	}
	return poiString, role, status, nil
}

func GeDrectionPointFromString(s string) (DirectionPoint, error) {
	var dp DirectionPoint
	var err error
	sArr := strings.Split(s, ",")
	if len(sArr) < 2 {
		return dp, errors.New("invalid input: " + s)
	}
	if dp.Lat, err = strconv.ParseFloat(sArr[0], 64); err != nil {
		return dp, err
	}
	if dp.Lon, err = strconv.ParseFloat(sArr[1], 64); err != nil {
		return dp, err
	}
	return dp, nil
}

func GetPointFromString(pString string) (Point, error) {
	if pString == "" {
		return Point{}, errors.New("point string is empty")
	}
	var poi Point
	ll := strings.Split(pString, ",")
	if len(ll) != 2 {
		// fmt.Printf("invalid input: %v\n\n", ll)
		return poi, errors.New("invalid input")
	}
	lat, err := strconv.ParseFloat(ll[0], 64)
	if err != nil {
		// fmt.Printf("invalid input: %v\n\n", ll[0])
		return poi, errors.New("invalid input")
	}
	lon, err := strconv.ParseFloat(ll[1], 64)
	if err != nil {
		// fmt.Printf("invalid input: %v\n\n", ll[1])
		return poi, errors.New("invalid input")
	}
	poi.Lat = lat
	poi.Lon = lon
	return poi, nil
}
