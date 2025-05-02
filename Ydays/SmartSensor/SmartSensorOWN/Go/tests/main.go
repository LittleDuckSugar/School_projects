/*package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v)
	v.Scale(10)
	fmt.Println(v)
	fmt.Println(v.Abs())
}
*/
/*
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	Scale(v, 10)
	fmt.Println(Abs(v))
}

*/
/*
package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	// a = v

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
*/
/*
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}

*/
/*
package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ip IPAddr) String() string {
	back := ""
	for _, value := range ip {
		back += string(value)
	}
	return back
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}


*/
package main

import (
	"fmt"
)

type Room struct {
	RoomName string
	Sensors  map[string]Sensor
	Area     uint
}

type Sensor struct {
	Name        string
	SensorType  string
	Temperature bool
	Humidity    bool
	Position    string
	IPAddress   string
	MACAddress  string
	Delay       uint
	Allowed     bool
}

type Place struct {
	Name  string
	Rooms map[string]Room
}

func (room Room) GetElem() *Room {
	return &room
}

// AddSensor add a new sensor to the current room
func (room *Room) AddSensor(newSensor Sensor) {
	if room.Sensors == nil {
		room.Sensors = make(map[string]Sensor)
	}

	if _, found := room.Sensors[newSensor.Name]; found {
		fmt.Println("There is already a room with this name, please set an other name")
	} else {
		room.Sensors[newSensor.Name] = newSensor
		fmt.Println("The following is sensor now added to", room.RoomName)
		fmt.Println(newSensor)
	}
}

// AddSensor add a new sensor to the current room
func (room *Room) Test() {
	fmt.Println("Test")
}

// addRoom add an existing room to the current place
func (place *Place) AddNewRoom(newRoom Room) {

	if place.Rooms == nil {
		place.Rooms = make(map[string]Room)
	}

	if _, found := place.Rooms[newRoom.RoomName]; found {
		fmt.Println("There is already a room with this name, please set an other name to add this room to your place")
	} else {
		place.Rooms[newRoom.RoomName] = newRoom
		fmt.Println("The following room is now added :")
		fmt.Println(newRoom)
	}
}

// CreateNewRoom create a new room and add it to the current place
func CreateNewRoom(roomName string, roomArea uint) Room {
	return Room{RoomName: roomName, Area: roomArea}
}

var myPlace Place = Place{Name: "Ynov - Demo", Rooms: nil}

func main() {
	DHT22_Salon := Sensor{Name: "ESP_Salon_01", SensorType: "DHT22", Temperature: true, Humidity: true, Position: "Mi hauteur", IPAddress: "192.168.1.28", MACAddress: "80:7D:3A:F3:9A:E0", Delay: 60, Allowed: true}

	fmt.Println(myPlace)
	myPlace.AddNewRoom(CreateNewRoom("Salon", 20))
	fmt.Println(myPlace)

	for roomName, room := range myPlace.Rooms {
		if roomName == "Salon" {
			room.AddSensor(DHT22_Salon)
			myPlace.Rooms["Salon"] = room
			break
		}
	}

	// myPlace.GetRooms("Salon").GetElem().AddSensor(DHT22_Salon)
	fmt.Println(myPlace)
}
