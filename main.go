package main

import (
    "log"
    "os"

    "github.com/urfave/cli"
    "github.com/IrishWhiskey/ryanair_cli/ryanair"
)

func SetInfo(app *cli.App) {
    (*app).Name = "ryanair-cli"
    (*app).Usage = "get ryanair info"
}

func SetFlags(app *cli.App) {
    (*app).Flags = []cli.Flag {
        &cli.StringFlag {
            Name: "dep",
            Usage: "departure airport code",
            Required: true,
        },
        &cli.StringFlag {
            Name: "arr",
            Usage: "arrival airport code",
            Required: true,
        },
        &cli.StringFlag {
            Name: "date",
            Usage: "travel date in mm/dd/yyyy format",
            Required: true,
        },
    }
}

func getCurrencySymbol(currency string) string {
    if currency == "GBP" {
        return "£"
    }
    if currency == "USD" {
        return "$"
    }
    return "€"
}

func Action(c *cli.Context) {
    travel, err := ryanair.Query((*c).String("dep"), (*c).String("arr"), (*c).String("date"))

    if err != nil {
        log.Println(err)
    }

    s := getCurrencySymbol(travel.Currency)
    log.Printf("Departure: %s (%s), Arrival: %s (%s)", travel.OriginName, travel.Origin, travel.DestinationName, travel.Destination)
    for _, date := range travel.AvaiableDates {
        log.Printf(" * Date: %s\n", date.DateOut)
        for idx, flight := range date.Flights {
            log.Printf("    %d. %s DepartureTime: %s, ArrivalTime: %s, Price: %.2f%s\n", idx+1, flight.FlightNumber, flight.DepartureTime, flight.ArrivalTime, flight.Price, s)
        }
    }
    
}

func Init(app *cli.App) {
    SetInfo(app)
    SetFlags(app)
    (*app).Action = func(c *cli.Context) error {
        Action(c)
        return nil
    }
}

func main() {
    log.SetFlags(0)

    app := cli.NewApp()
    Init(app)

    err := app.Run(os.Args)
    if err != nil {
        log.Fatal(err)
    }
}
