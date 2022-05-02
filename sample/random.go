package sample

import (
	"grpc_app/pb"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Generate random keyboard layout.
func randomKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return pb.Keyboard_QWERTY
	case 2:
		return pb.Keyboard_QWERTZ
	default:
		return pb.Keyboard_AZERTY
	}
}

// Generate random CPU brand.
func randomCPUBrand() string {
	return randomStringFromSet("Intel", "AMD")
}

// Generate random CPU model based on brand.
func randomCPUName(brand string) string {
	if brand == "Intel" {
		return randomStringFromSet(
			"Intel Core i9-12900H",
			"Intel Core i9-11900H",
			"Intel Core i7-12700H",
			"Intel Core i7-11800H",
			"Intel Core i7-10510U",
		)
	}
	return randomStringFromSet(
		"AMD Ryzen 9 5900HS",
		"AMD Ryzen 7 5800U",
		"AMD Ryzen 7 3700U",
		"AMD Ryzen 5 5500U",
		"AMD Ryzen 5 3500U",
	)
}

// Generate random GPU brand.
func randomGPUBrand() string {
	return randomStringFromSet("Nvidia", "AMD")
}

// Generate random GPU model based on brand.
func randomGPUName(brand string) string {
	if brand == "Nvidia" {
		return randomStringFromSet(
			"NVIDIA RTX 3080",
			"NVIDIA RTX 3070",
			"NVIDIA RTX 3060",
			"NVIDIA RTX 3050",
			"NVIDIA RTX 2070",
		)
	}
	return randomStringFromSet(
		"RX 590",
		"RX 580",
		"RX 550",
		"RX 5700-XT",
		"RX Vega-56",
	)
}

// Generate random laptop brand.
func randomLaptopBrand() string {
	return randomStringFromSet("Apple", "Dell", "Lenovo", "Asus", "HP")
}

// Generate random laptop model based on brand.
func randomLaptopName(brand string) string {
	switch brand {
	case "Apple":
		return randomStringFromSet("Macbook Air", "Macboor Pro")
	case "Dell":
		return randomStringFromSet("Inspirion", "XPS", "Alienware", "G Series")
	case "Lenovo":
		return randomStringFromSet("ThinkPad", "ThinkBook", "ideapad", "Lenovo Legion")
	case "Asus":
		return randomStringFromSet("ProArt Studiobook", "Zenbook", "Vivobook", "ROG")
	case "HP":
		return randomStringFromSet("Spectre", "Envy", "Omen", "Pavilion")
	default:
		return ""
	}
}

// Generate random screen resolution.
func randomScreenResolution() *pb.Screen_Resolution {
	height := randomInt(1080, 4320)
	width := height * 16 / 9
	resolution := &pb.Screen_Resolution{
		Height: uint32(height),
		Width:  uint32(width),
	}
	return resolution
}

// Generate random IPS/OLED panel.
func randomScreenPanel() pb.Screen_Panel {
	if rand.Intn(2) == 1 {
		return pb.Screen_IPS
	}
	return pb.Screen_OLED
}

// Generate random string from Set of strings.
func randomStringFromSet(a ...string) string {
	n := len(a)
	if n == 0 {
		return ""
	}
	return a[rand.Intn(n)]
}

// Generate random bool.
func randomBool() bool {
	return rand.Intn(2) == 1
}

// Generate random int.
func randomInt(min int, max int) int {
	return min + rand.Intn(max-min+1)
}

// Generate random float64.
func randomFloat64(min float64, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

// Generate random float32.
func randomFloat32(min float32, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

// Generate random ID.
func randomID() string {
	return uuid.New().String()
}
