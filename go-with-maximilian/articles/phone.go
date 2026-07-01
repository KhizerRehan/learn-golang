package articles

import "fmt"

type Phone struct {
	brand   string
	battery int // battery percentage
	apps    []string
}

func New(b string, bty int, apps []string) (Phone, error) {

	if b == "" {
		return Phone{}, fmt.Errorf("brand is required")
	}

	if bty < 0 || bty > 100 {
		return Phone{}, fmt.Errorf("battery must be 0–100")
	}

	return Phone{
		brand:   b,
		battery: bty,
		apps:    apps,
	}, nil
}

func NewPhone(b string, bty int, apps []string) Phone {

	if b == "" {
		b = "Unknown"
	}

	if bty < 0 {
		bty = 0
	}
	if bty > 100 {
		bty = 100
	}

	return Phone{
		brand:   b,
		battery: bty,
		apps:    apps,
	}
}

// Define Individual Methods
func (p Phone) GetBrand() string {
	return p.brand // return the brand of the phone
}

func (p Phone) GetBattery() int {
	return p.battery // return the battery percentage of the phone
}

func (p Phone) GetApps() []string {
	return p.apps // return the apps of the phone
}

func GetPhoneDetails(p Phone) string {
	return fmt.Sprintf("Brand: %s, Battery: %d, Apps: %v", p.GetBrand(), p.GetBattery(), p.GetApps())
}

func GetNewPhones() {
	var phones []Phone

	samsungPhone, err1 := New("Samsung", 100, []string{"WhatsApp", "Facebook"})
	if err1 != nil {
		fmt.Println("Error creating Samsung phone:", err1)
		return
	}
	phones = append(phones, samsungPhone)

	applePhone, err2 := New("Apple", 90, []string{"Instagram", "Snapchat"})
	if err2 != nil {
		fmt.Println("Error creating Apple phone:", err2)
		return
	}
	phones = append(phones, applePhone)

	for _, phone := range phones {
		fmt.Println(GetPhoneDetails(phone))
	}
}
