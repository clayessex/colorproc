
# Sort ColorNames into similar colors by hue

Expects colornames.go built by ../parsecolornames

### ColorNames should look like
```
var ColorNames = []struct {
	name string
	rgb  string
}{
	{name: "100 Mph", rgb: "#c93f38"},
	{name: "18th Century Green", rgb: "#a59344"},
	{name: "1975 Earth Red", rgb: "#7b463b"},
	{name: "1989 Miami Hotline", rgb: "#dd3366"},
	{name: "20000 Leagues Under the Sea", rgb: "#191970"},
	{name: "24 Carrot", rgb: "#e56e24"},
	{name: "24 Karat", rgb: "#dfc685"},
	{name: "3AM in Shibuya", rgb: "#225577"},
	{name: "3am Latte", rgb: "#c0a98e"},
	{name: "400XT Film", rgb: "#d2d2c0"},
	{name: "5-Masted Preußen", rgb: "#9bafad"},
	{name: "8 Bit Eggplant", rgb: "#990066"},
	{name: "90% Cocoa", rgb: "#3d1c02"},
	{name: "99 Years Blue", rgb: "#000099"},
    ...
```

