package ryanair

import (
	"encoding/json"
	"log"
	"fmt"
)

type Travel_t struct {
	Currency string
	Origin string
	OriginName string
	Destination string
	DestinationName string
	AvaiableDates []Date_t
}

type Date_t struct {
	DateOut string
	Flights []Flight_t
}

type Flight_t struct {
	FlightNumber string
	Price float64
	DepartureTime string
	ArrivalTime string
}

///Converts date from mm/dd/yyyy to yyyy-mm-dd
func EncodeDate(date string) string {
	return fmt.Sprintf("%s-%s-%s", date[6:10], date[0:2], date[3:5])
}

///Converts date from yyyy-mm-dd to mm/dd/yyyy
func DecodeDate(date string) string {
	return fmt.Sprintf("%s/%s/%s", date[5:7], date[8:10], date[0:4])
}

func getFlight(in map[string]interface{}) Flight_t {
	var ans Flight_t
	time := in["time"].([]interface{})
	fare := in["regularFare"].(map[string]interface{})["fares"].([]interface{})[0].(map[string]interface{})

	ans.FlightNumber = in["flightNumber"].(string)
	ans.DepartureTime = time[0].(string)[11:16]
	ans.ArrivalTime = time[1].(string)[11:16]
	ans.Price = fare["amount"].(float64)
	return ans
}

func getFlights(in []interface{}) []Flight_t {
	size := len(in)
	ans := make([]Flight_t, size)

	for i := 0; i < size; i++ {
		ans[i] = getFlight(in[i].(map[string]interface{}))
	} 

	return ans
}

func getDate(in map[string]interface{}) Date_t {
	var ans Date_t
	ans.DateOut = DecodeDate(in["dateOut"].(string)[0:10])
	ans.Flights = getFlights(in["flights"].([]interface{}))
	return ans
}

func getDates(in []interface{}) []Date_t {
	size := len(in)
	ans := make([]Date_t, size)

	for i := 0; i < size; i++ {
		ans[i] = getDate(in[i].(map[string]interface{}))
	} 

	return ans
}

func getTravel(in map[string]interface{}) *Travel_t {
	ans := new(Travel_t)
	ans.Currency = in["currency"].(string)
	trip := in["trips"].([]interface{})[0].(map[string]interface{})
	ans.Origin = trip["origin"].(string)
	ans.OriginName = trip["originName"].(string)
	ans.Destination = trip["destination"].(string)
	ans.DestinationName = trip["destinationName"].(string)
	ans.AvaiableDates = getDates(trip["dates"].([]interface{}))

	return ans
}

func parseJson(body []byte) *Travel_t {
	var s map[string]interface{}
	err := json.Unmarshal(body, &s)

	if err != nil {
		log.Println(err)
	}

	return getTravel(s)
}