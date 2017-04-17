package main

import (
	"testing"

	"gopkg.in/check.v1"
)

func Test(t *testing.T) {
	check.TestingT(t)
}

type CalculatorSuite struct {
	calculator *Calculator
}

var _ = check.Suite(&CalculatorSuite{})

func (s *CalculatorSuite) SetUpSuite(c *check.C) {
	s.calculator = &Calculator{}
}

func (s *CalculatorSuite) TestAddSuccess(c *check.C) {
	result := s.calculator.Add(42, 10)
	c.Check(result, check.Equals, 52)
}

// func (s *CalculatorSuite) TestAddFail(c *check.C) {
// 	result := s.calculator.Add(42, 10)
// 	c.Check(result, check.Equals, 10)
// }

func (s *CalculatorSuite) TestSubtract(c *check.C) {
	result := s.calculator.Subtract(30, 20)
	c.Check(result, check.Equals, 10)
}

func (s *CalculatorSuite) TestMultiply(c *check.C) {
	result := s.calculator.Multiply(3, 10)
	c.Check(result, check.Equals, 30)
}
