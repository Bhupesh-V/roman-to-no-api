package controllers

import (
	// "fmt"
	"fmt"
	"github.com/gin-gonic/gin"
	// "errors"
	"net/http"
)

type RomanNumeral struct {
	Roman string `json:"roman,omitempty"`
}

type romanError struct {
	char string
}

func (e *romanError) Error() string {
	return fmt.Sprintf("illegal character %s in roman", e.char)
}

// POST /convert
// Convert Roman Numeral
func ConvertRoman(c *gin.Context) {
	// Validate input
	var input RomanNumeral
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	integral, err := RomanToInt(input.Roman)
	if err != nil {
        fmt.Println(err)
        c.JSON(http.StatusOK, gin.H{"roman": input.Roman, "error": err})
    } else {
        c.JSON(http.StatusOK, gin.H{"roman": input.Roman, "integer": integral})
    }
}

func RomanToInt(roman string) (int, error) {
	romanSymbols := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	total := 0
	for index := 0; index < len(roman); index++ {
		if value, ok := romanSymbols[string(roman[index])]; !ok {
			return value, &romanError{string(roman[index])}
		}
		// prevent index out of bound panic
		if index < len(roman)-1 {
			if romanSymbols[string(roman[index])] >= romanSymbols[string(roman[index+1])] {
				total += romanSymbols[string(roman[index])]
			} else {
				total -= romanSymbols[string(roman[index])]
			}
		} else {
			total += romanSymbols[string(roman[index])]
		}
	}
	return total, nil
}
