package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Mode interface{
	Convert(c string) (string,error)
}
type R2H struct {}
type H2R struct {}

func main(){
	var color string
	var mode string
	colorUsage := "The color you want to change, color can be rgb or hex. RGB can be in following formats: \n[1]: rgb(a,b,c)\n[2]: a,b,c"
	flag.StringVar(&mode, "mode", "r2h" , "[1]: Rgb to Hex = r2h\n[2]: Hex to Rgb = h2r")
	flag.StringVar(&color, "c", "", colorUsage)
	flag.Parse()
	if color == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	instance, err := createConcrete(mode)
	errorHandle(err)
	res, err := instance.Convert(color)
	errorHandle(err)
	print(res)
}

func errorHandle(err error){
	if err != nil{
		fmt.Println(err)
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func createConcrete(mode string) (Mode,error) {
	if mode == "r2h"{
		i := R2H{}
		return i, nil
	} else if mode == "h2r" {
		i := H2R{}
		return i, nil
	}
	return nil, fmt.Errorf("ERROR: This %s,  mode is not valid", mode)
}

func (_ R2H) Convert(rgb string) (string,error){
	var r,g,b int
	var err error

	if strings.HasPrefix(rgb, "rgb"){
		rgb = rgb[4:len(rgb) - 1]
	}

	if len(rgb) < 5 {
		return "", fmt.Errorf("ERROR: %s is not a valid input. Should be like rgb(a,b,c) or a,b,c", rgb)
	}

	rgbArray := strings.Split(rgb, ",")

	if r,err = strconv.Atoi(rgbArray[0]); err != nil {
		return "", fmt.Errorf("%s", err.Error()[14:])
	}

	if g,err = strconv.Atoi(rgbArray[1]); err != nil {
		return "", fmt.Errorf("%s", err.Error()[14:])
	}

	if b,err = strconv.Atoi(rgbArray[2]); err != nil{
		return "", fmt.Errorf("%s", err.Error()[14:])
	}

	if r < 0 || r > 255 {
		return "", fmt.Errorf("ERROR: r should be greater or equal than 0 and less than 256")
	}

	if g < 0 || g > 255 {
		return "", fmt.Errorf("ERROR: g should be greater or equal than 0 and less than 256")
	}

	if b < 0 || b > 255 {
		return "", fmt.Errorf("ERROR: b should be greater or equal than 0 and less than 256")
	}

	return fmt.Sprintf("#%02x%02x%02x", r, g, b), nil
}

func (_ H2R) Convert(hex string) (string,error){
	if len(hex) < 6 {
		return "", fmt.Errorf("ERROR: %s is not a valid input. Should be like #000000", hex)
	}

	if strings.HasPrefix(hex, "#"){
		hex = hex[1:]
	}

	r,err := strconv.ParseInt(hex[:2], 16, 64)
	if err != nil{
		return "", fmt.Errorf("%s", err.Error()[18:])
	}

	g,err := strconv.ParseInt(hex[2:4], 16, 64)
	if err != nil{
		return "", fmt.Errorf("ERROR: %s", err.Error()[18:])
	}

	b,err := strconv.ParseInt(hex[4:6], 16, 64)
	if err != nil{
		return "", fmt.Errorf("ERROR %s", err.Error()[18:])
	}

	return fmt.Sprintf("rgb(%d,%d,%d)", r,g,b), nil
}