// Convert a time in seconds to years given a planet
package space

//type Planet string


func Age(ageSec float64, planet string) float64{
	year := 31557600.0
	period := make(map[string]float64)
	period["Mercury"] = 0.2408467
	period["Venus"] = 0.61519726
	period["Earth"] = 1.0
	period["Mars"] = 1.8808158
	period["Jupiter"] = 11.862615
	period["Saturn"] = 29.447498
	period["Uranus"] = 84.016846
	period["Neptune"] = 164.79132

	return float64(ageSec) / year / period[planet]
}
