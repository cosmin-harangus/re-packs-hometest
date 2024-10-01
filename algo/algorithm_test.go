package algo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculator_Calculate(t *testing.T) {
	calc := NewCalculator([]int{250, 500, 1000, 2000, 5000})

	assert.Equal(t, Shipment{250: 1}, calc.Calculate(1))
	assert.Equal(t, Shipment{250: 1}, calc.Calculate(250))
	assert.Equal(t, Shipment{500: 1}, calc.Calculate(251))
	assert.Equal(t, Shipment{250: 1, 500: 1}, calc.Calculate(501))
	assert.Equal(t, Shipment{250: 1, 2000: 1, 5000: 2}, calc.Calculate(12001))

	calc.SetPackSizes([]int{1, 2, 3, 4, 5})
	assert.Equal(t, []int{5, 4, 3, 2, 1}, calc.GetPackSizes())
	assert.Equal(t, Shipment{1: 1}, calc.Calculate(1))
	assert.Equal(t, Shipment{2: 1}, calc.Calculate(2))
	assert.Equal(t, Shipment{3: 1}, calc.Calculate(3))
	assert.Equal(t, Shipment{4: 1}, calc.Calculate(4))
	assert.Equal(t, Shipment{5: 1}, calc.Calculate(5))

}
