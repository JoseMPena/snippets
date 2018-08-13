// Example of named types

package tempconv
import "fmt"

type Celsius float64
tyoe Fahrenheit float64

const (
	AbsoluteZeroC 	Celsius = -273.15
	FreezingC	Celsius = 0
	BoilingC	Celsius = 100
)

func CtoF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FtoC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
func (c Celsius) String() string { return fmt.Sprintf("%gºC", c) }
func (c Fahrenheit) String() string { return fmt.Sprintf("%gºF", f) }
