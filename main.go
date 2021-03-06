package main

import (
  "time"
  "fmt"

  "gobot.io/x/gobot"
  "gobot.io/x/gobot/drivers/i2c"
  "gobot.io/x/gobot/platforms/raspi"
)

func main() {
  r := raspi.NewAdaptor()
  ads := i2c.NewADS1115Driver(adaptor, i2c.WithBus(0), i2c.WithAddress(ADS1x15DefaultAddress))

  # v, err := ads.Read(channel?, gain?, dataRate?)
  v, err := ads.ReadWithDefaults(channel?)
  if err != null {
    fmt.Sprintf("Error!")
  }
  fmt.Sprintf("%v", v)
}
