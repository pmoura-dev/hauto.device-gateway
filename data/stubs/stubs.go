package stubs

var (
	DevicesData = []byte(`
	[
		{
			"id": 1,
			"natural_id": "12345",
			"name": "Color Light Bulb",
			"type": "color_light",
			"room": "Living Room",
			"controller": "shelly_color_bulb",
			"state": {
				"status": "online",
				"color": {
					"hue": 120,
					"saturation": 75,
					"lightness": 50,
					"alpha": 0.5
				}
			}
		},
		{
			"id": 2,
			"natural_id": "54321",
			"name": "Air Conditioner",
			"type": "air_conditioner",
			"room": "Living Room",
			"controller": "hisense_ac",
			"state": {
				"status": "online",
				"current_temperature": 23,
				"mode": "heating",
				"current_threshold_temperature": 25
			}
		}
	]
	`)
)
