package main

import (
	"testing"
)

func panicCheck_main(t *testing.T, panicExpect string, fn func()) {
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("Expected panic, but no panic occurred")
		}
		if r != panicExpect {
			t.Errorf("Expected panic with message='%s', but no panic occurred", panicExpect)
		}
	}()

	fn()
}
func TestGenerated_test1(t *testing.T) {
	res0TC0 := test1(12)

	if res0TC0 != 13 {
		t.Errorf("res0TC0 = %v, but predicted 13", res0TC0)
	}
	res0TC1 := test1(-9223372036854775796)

	if res0TC1 != -9223372036854775797 {
		t.Errorf("res0TC1 = %v, but predicted -9223372036854775797", res0TC1)
	}
}

func TestGenerated_test2(t *testing.T) {
	res0TC0 := test2(false, false)

	if res0TC0 != false {
		t.Errorf("res0TC0 = %v, but predicted false", res0TC0)
	}
	res0TC1 := test2(true, true)

	if res0TC1 != true {
		t.Errorf("res0TC1 = %v, but predicted true", res0TC1)
	}
	res0TC2 := test2(true, false)

	if res0TC2 != false {
		t.Errorf("res0TC2 = %v, but predicted false", res0TC2)
	}
}

func TestGenerated_testArithmetic(t *testing.T) {
	res0TC0 := testArithmetic(-8919549788968174172, 3004641839422955821)

	if res0TC0 != 1 {
		t.Errorf("res0TC0 = %v, but predicted 1", res0TC0)
	}
	res0TC1 := testArithmetic(3180548182006436739, 6623606982543203939)

	if res0TC1 != 6 {
		t.Errorf("res0TC1 = %v, but predicted 6", res0TC1)
	}
}

func TestGenerated_testUnary(t *testing.T) {
	res0TC0 := testUnary(9223372036854775806)

	if res0TC0 != 9223372036854775805 {
		t.Errorf("res0TC0 = %v, but predicted 9223372036854775805", res0TC0)
	}
	res0TC1 := testUnary(-2)

	if res0TC1 != 2 {
		t.Errorf("res0TC1 = %v, but predicted 2", res0TC1)
	}
}

func TestGenerated_testComparisons(t *testing.T) {
	res0TC0 := testComparisons(0, 0)

	if res0TC0 != true {
		t.Errorf("res0TC0 = %v, but predicted true", res0TC0)
	}
	res0TC1 := testComparisons(9223372036854775806, -9223372036854775807)

	if res0TC1 != true {
		t.Errorf("res0TC1 = %v, but predicted true", res0TC1)
	}
	res0TC2 := testComparisons(0, -1)

	if res0TC2 != true {
		t.Errorf("res0TC2 = %v, but predicted true", res0TC2)
	}
	res0TC3 := testComparisons(-8853654483146637056, 8853654483146637055)

	if res0TC3 != true {
		t.Errorf("res0TC3 = %v, but predicted true", res0TC3)
	}
}

func TestGenerated_testLogicalOps(t *testing.T) {
	res0TC0 := testLogicalOps(-9223372036854775806, 0)

	if res0TC0 != true {
		t.Errorf("res0TC0 = %v, but predicted true", res0TC0)
	}
	res0TC1 := testLogicalOps(0, 0)

	if res0TC1 != false {
		t.Errorf("res0TC1 = %v, but predicted false", res0TC1)
	}
	res0TC2 := testLogicalOps(2, 0)

	if res0TC2 != false {
		t.Errorf("res0TC2 = %v, but predicted false", res0TC2)
	}
}

func TestGenerated_testWhileLoop(t *testing.T) {
	res0TC0 := testWhileLoop(-9223372036854775806)

	if res0TC0 != 0 {
		t.Errorf("res0TC0 = %v, but predicted 0", res0TC0)
	}
	res0TC1 := testWhileLoop(1)

	if res0TC1 != 0 {
		t.Errorf("res0TC1 = %v, but predicted 0", res0TC1)
	}
	res0TC2 := testWhileLoop(2)

	if res0TC2 != 1 {
		t.Errorf("res0TC2 = %v, but predicted 1", res0TC2)
	}
	res0TC3 := testWhileLoop(3)

	if res0TC3 != 3 {
		t.Errorf("res0TC3 = %v, but predicted 3", res0TC3)
	}
	res0TC4 := testWhileLoop(4)

	if res0TC4 != 6 {
		t.Errorf("res0TC4 = %v, but predicted 6", res0TC4)
	}
	res0TC5 := testWhileLoop(5)

	if res0TC5 != 10 {
		t.Errorf("res0TC5 = %v, but predicted 10", res0TC5)
	}
	res0TC6 := testWhileLoop(6)

	if res0TC6 != 15 {
		t.Errorf("res0TC6 = %v, but predicted 15", res0TC6)
	}
	res0TC7 := testWhileLoop(7)

	if res0TC7 != 21 {
		t.Errorf("res0TC7 = %v, but predicted 21", res0TC7)
	}
	res0TC8 := testWhileLoop(8)

	if res0TC8 != 28 {
		t.Errorf("res0TC8 = %v, but predicted 28", res0TC8)
	}
	res0TC9 := testWhileLoop(9)

	if res0TC9 != 36 {
		t.Errorf("res0TC9 = %v, but predicted 36", res0TC9)
	}
	res0TC10 := testWhileLoop(10)

	if res0TC10 != 45 {
		t.Errorf("res0TC10 = %v, but predicted 45", res0TC10)
	}
	res0TC11 := testWhileLoop(11)

	if res0TC11 != 55 {
		t.Errorf("res0TC11 = %v, but predicted 55", res0TC11)
	}
	res0TC12 := testWhileLoop(12)

	if res0TC12 != 66 {
		t.Errorf("res0TC12 = %v, but predicted 66", res0TC12)
	}
	res0TC13 := testWhileLoop(13)

	if res0TC13 != 78 {
		t.Errorf("res0TC13 = %v, but predicted 78", res0TC13)
	}
	res0TC14 := testWhileLoop(14)

	if res0TC14 != 91 {
		t.Errorf("res0TC14 = %v, but predicted 91", res0TC14)
	}
}

func TestGenerated_testForLoop(t *testing.T) {
	res0TC0 := testForLoop(0)

	if res0TC0 != 1 {
		t.Errorf("res0TC0 = %v, but predicted 1", res0TC0)
	}
	res0TC1 := testForLoop(1)

	if res0TC1 != 2 {
		t.Errorf("res0TC1 = %v, but predicted 2", res0TC1)
	}
	res0TC2 := testForLoop(2)

	if res0TC2 != 4 {
		t.Errorf("res0TC2 = %v, but predicted 4", res0TC2)
	}
	res0TC3 := testForLoop(3)

	if res0TC3 != 7 {
		t.Errorf("res0TC3 = %v, but predicted 7", res0TC3)
	}
	res0TC4 := testForLoop(4)

	if res0TC4 != 28 {
		t.Errorf("res0TC4 = %v, but predicted 28", res0TC4)
	}
	res0TC5 := testForLoop(5)

	if res0TC5 != 33 {
		t.Errorf("res0TC5 = %v, but predicted 33", res0TC5)
	}
	res0TC6 := testForLoop(6)

	if res0TC6 != 198 {
		t.Errorf("res0TC6 = %v, but predicted 198", res0TC6)
	}
	res0TC7 := testForLoop(7)

	if res0TC7 != 205 {
		t.Errorf("res0TC7 = %v, but predicted 205", res0TC7)
	}
	res0TC8 := testForLoop(8)

	if res0TC8 != 1640 {
		t.Errorf("res0TC8 = %v, but predicted 1640", res0TC8)
	}
	res0TC9 := testForLoop(9)

	if res0TC9 != 1649 {
		t.Errorf("res0TC9 = %v, but predicted 1649", res0TC9)
	}
	res0TC10 := testForLoop(10)

	if res0TC10 != 16490 {
		t.Errorf("res0TC10 = %v, but predicted 16490", res0TC10)
	}
	res0TC11 := testForLoop(11)

	if res0TC11 != 16501 {
		t.Errorf("res0TC11 = %v, but predicted 16501", res0TC11)
	}
	res0TC12 := testForLoop(12)

	if res0TC12 != 198012 {
		t.Errorf("res0TC12 = %v, but predicted 198012", res0TC12)
	}
	res0TC13 := testForLoop(13)

	if res0TC13 != 198025 {
		t.Errorf("res0TC13 = %v, but predicted 198025", res0TC13)
	}
	res0TC14 := testForLoop(14)

	if res0TC14 != 2772350 {
		t.Errorf("res0TC14 = %v, but predicted 2772350", res0TC14)
	}
}

func TestGenerated_testInfiniteLoopBreak(t *testing.T) {
	res0TC0 := testInfiniteLoopBreak(-9223372036854775806)

	if res0TC0 != 0 {
		t.Errorf("res0TC0 = %v, but predicted 0", res0TC0)
	}
	res0TC1 := testInfiniteLoopBreak(1)

	if res0TC1 != 1 {
		t.Errorf("res0TC1 = %v, but predicted 1", res0TC1)
	}
	res0TC2 := testInfiniteLoopBreak(2)

	if res0TC2 != 2 {
		t.Errorf("res0TC2 = %v, but predicted 2", res0TC2)
	}
	res0TC3 := testInfiniteLoopBreak(3)

	if res0TC3 != 3 {
		t.Errorf("res0TC3 = %v, but predicted 3", res0TC3)
	}
	res0TC4 := testInfiniteLoopBreak(4)

	if res0TC4 != 4 {
		t.Errorf("res0TC4 = %v, but predicted 4", res0TC4)
	}
	res0TC5 := testInfiniteLoopBreak(5)

	if res0TC5 != 5 {
		t.Errorf("res0TC5 = %v, but predicted 5", res0TC5)
	}
	res0TC6 := testInfiniteLoopBreak(6)

	if res0TC6 != 6 {
		t.Errorf("res0TC6 = %v, but predicted 6", res0TC6)
	}
	res0TC7 := testInfiniteLoopBreak(7)

	if res0TC7 != 7 {
		t.Errorf("res0TC7 = %v, but predicted 7", res0TC7)
	}
	res0TC8 := testInfiniteLoopBreak(8)

	if res0TC8 != 8 {
		t.Errorf("res0TC8 = %v, but predicted 8", res0TC8)
	}
	res0TC9 := testInfiniteLoopBreak(9)

	if res0TC9 != 9 {
		t.Errorf("res0TC9 = %v, but predicted 9", res0TC9)
	}
	res0TC10 := testInfiniteLoopBreak(10)

	if res0TC10 != 10 {
		t.Errorf("res0TC10 = %v, but predicted 10", res0TC10)
	}
	res0TC11 := testInfiniteLoopBreak(11)

	if res0TC11 != 11 {
		t.Errorf("res0TC11 = %v, but predicted 11", res0TC11)
	}
	res0TC12 := testInfiniteLoopBreak(12)

	if res0TC12 != 12 {
		t.Errorf("res0TC12 = %v, but predicted 12", res0TC12)
	}
	res0TC13 := testInfiniteLoopBreak(13)

	if res0TC13 != 13 {
		t.Errorf("res0TC13 = %v, but predicted 13", res0TC13)
	}
	res0TC14 := testInfiniteLoopBreak(14)

	if res0TC14 != 14 {
		t.Errorf("res0TC14 = %v, but predicted 14", res0TC14)
	}
}

func TestGenerated_testLoopWithConcreteBoundAndSymbolicBranching(t *testing.T) {
	res0TC0 := testLoopWithConcreteBoundAndSymbolicBranching(false)

	if res0TC0 != 0 {
		t.Errorf("res0TC0 = %v, but predicted 0", res0TC0)
	}
	res0TC1 := testLoopWithConcreteBoundAndSymbolicBranching(true)

	if res0TC1 != 20 {
		t.Errorf("res0TC1 = %v, but predicted 20", res0TC1)
	}
}

func TestGenerated_testLoopWithSymbolicBoundAndSymbolicBranching(t *testing.T) {
	panicCheck(t, "Assumption violated: n should be less than or equal to 10", func() { testLoopWithSymbolicBoundAndSymbolicBranching(12, false) })
	res0TC1 := testLoopWithSymbolicBoundAndSymbolicBranching(-9223372036854775794, false)

	if res0TC1 != 0 {
		t.Errorf("res0TC1 = %v, but predicted 0", res0TC1)
	}
	res0TC2 := testLoopWithSymbolicBoundAndSymbolicBranching(1, false)

	if res0TC2 != 0 {
		t.Errorf("res0TC2 = %v, but predicted 0", res0TC2)
	}
	res0TC3 := testLoopWithSymbolicBoundAndSymbolicBranching(1, true)

	if res0TC3 != 0 {
		t.Errorf("res0TC3 = %v, but predicted 0", res0TC3)
	}
	res0TC4 := testLoopWithSymbolicBoundAndSymbolicBranching(2, false)

	if res0TC4 != 0 {
		t.Errorf("res0TC4 = %v, but predicted 0", res0TC4)
	}
	res0TC5 := testLoopWithSymbolicBoundAndSymbolicBranching(3, false)

	if res0TC5 != 0 {
		t.Errorf("res0TC5 = %v, but predicted 0", res0TC5)
	}
	res0TC6 := testLoopWithSymbolicBoundAndSymbolicBranching(2, true)

	if res0TC6 != 0 {
		t.Errorf("res0TC6 = %v, but predicted 0", res0TC6)
	}
	res0TC7 := testLoopWithSymbolicBoundAndSymbolicBranching(4, false)

	if res0TC7 != 0 {
		t.Errorf("res0TC7 = %v, but predicted 0", res0TC7)
	}
	res0TC8 := testLoopWithSymbolicBoundAndSymbolicBranching(3, true)

	if res0TC8 != 2 {
		t.Errorf("res0TC8 = %v, but predicted 2", res0TC8)
	}
	res0TC9 := testLoopWithSymbolicBoundAndSymbolicBranching(5, false)

	if res0TC9 != 0 {
		t.Errorf("res0TC9 = %v, but predicted 0", res0TC9)
	}
	res0TC10 := testLoopWithSymbolicBoundAndSymbolicBranching(6, false)

	if res0TC10 != 0 {
		t.Errorf("res0TC10 = %v, but predicted 0", res0TC10)
	}
	res0TC11 := testLoopWithSymbolicBoundAndSymbolicBranching(4, true)

	if res0TC11 != 2 {
		t.Errorf("res0TC11 = %v, but predicted 2", res0TC11)
	}
	res0TC12 := testLoopWithSymbolicBoundAndSymbolicBranching(7, false)

	if res0TC12 != 0 {
		t.Errorf("res0TC12 = %v, but predicted 0", res0TC12)
	}
	res0TC13 := testLoopWithSymbolicBoundAndSymbolicBranching(5, true)

	if res0TC13 != 6 {
		t.Errorf("res0TC13 = %v, but predicted 6", res0TC13)
	}
	res0TC14 := testLoopWithSymbolicBoundAndSymbolicBranching(8, false)

	if res0TC14 != 0 {
		t.Errorf("res0TC14 = %v, but predicted 0", res0TC14)
	}
	res0TC15 := testLoopWithSymbolicBoundAndSymbolicBranching(9, false)

	if res0TC15 != 0 {
		t.Errorf("res0TC15 = %v, but predicted 0", res0TC15)
	}
	res0TC16 := testLoopWithSymbolicBoundAndSymbolicBranching(6, true)

	if res0TC16 != 6 {
		t.Errorf("res0TC16 = %v, but predicted 6", res0TC16)
	}
	res0TC17 := testLoopWithSymbolicBoundAndSymbolicBranching(10, false)

	if res0TC17 != 0 {
		t.Errorf("res0TC17 = %v, but predicted 0", res0TC17)
	}
	res0TC18 := testLoopWithSymbolicBoundAndSymbolicBranching(7, true)

	if res0TC18 != 12 {
		t.Errorf("res0TC18 = %v, but predicted 12", res0TC18)
	}
	res0TC19 := testLoopWithSymbolicBoundAndSymbolicBranching(8, true)

	if res0TC19 != 12 {
		t.Errorf("res0TC19 = %v, but predicted 12", res0TC19)
	}
	res0TC20 := testLoopWithSymbolicBoundAndSymbolicBranching(9, true)

	if res0TC20 != 20 {
		t.Errorf("res0TC20 = %v, but predicted 20", res0TC20)
	}
	res0TC21 := testLoopWithSymbolicBoundAndSymbolicBranching(10, true)

	if res0TC21 != 20 {
		t.Errorf("res0TC21 = %v, but predicted 20", res0TC21)
	}
}

func TestGenerated_testComplexConditions(t *testing.T) {
	res0TC0 := testComplexConditions(2, 0, 1)

	if res0TC0 != 1 {
		t.Errorf("res0TC0 = %v, but predicted 1", res0TC0)
	}
	res0TC1 := testComplexConditions(459419962931421214, 459419962931421215, 0)

	if res0TC1 != 1 {
		t.Errorf("res0TC1 = %v, but predicted 1", res0TC1)
	}
	res0TC2 := testComplexConditions(4532105087670112556, -4532105087670112557, -9143172115396383018)

	if res0TC2 != 3 {
		t.Errorf("res0TC2 = %v, but predicted 3", res0TC2)
	}
	res0TC3 := testComplexConditions(-2647831093471200330, 2647831093471200329, 4611686018427387904)

	if res0TC3 != 3 {
		t.Errorf("res0TC3 = %v, but predicted 3", res0TC3)
	}
	res0TC4 := testComplexConditions(0, 0, 0)

	if res0TC4 != 2 {
		t.Errorf("res0TC4 = %v, but predicted 2", res0TC4)
	}
	res0TC5 := testComplexConditions(-34, -34, 33)

	if res0TC5 != 3 {
		t.Errorf("res0TC5 = %v, but predicted 3", res0TC5)
	}
}

func TestGenerated_testNestedIf(t *testing.T) {
	res0TC0 := testNestedIf(2, 1)

	if res0TC0 != 1 {
		t.Errorf("res0TC0 = %v, but predicted 1", res0TC0)
	}
	res0TC1 := testNestedIf(2, -9223372036854775807)

	if res0TC1 != 2 {
		t.Errorf("res0TC1 = %v, but predicted 2", res0TC1)
	}
	res0TC2 := testNestedIf(-9223372036854775806, 1)

	if res0TC2 != 3 {
		t.Errorf("res0TC2 = %v, but predicted 3", res0TC2)
	}
	res0TC3 := testNestedIf(-9223372036854775806, -9223372036854775807)

	if res0TC3 != 4 {
		t.Errorf("res0TC3 = %v, but predicted 4", res0TC3)
	}
}

func TestGenerated_testBitwise(t *testing.T) {
	res0TC0 := testBitwise(0, 1)

	if res0TC0 != 2 {
		t.Errorf("res0TC0 = %v, but predicted 2", res0TC0)
	}
	res0TC1 := testBitwise(-9223372036854775808, 1)

	if res0TC1 != -4611686018427387904 {
		t.Errorf("res0TC1 = %v, but predicted -4611686018427387904", res0TC1)
	}
}

func TestGenerated_testCombined(t *testing.T) {
	res0TC0 := testCombined(4611686018427387906)

	if res0TC0 != 0 {
		t.Errorf("res0TC0 = %v, but predicted 0", res0TC0)
	}
	res0TC1 := testCombined(1)

	if res0TC1 != 1 {
		t.Errorf("res0TC1 = %v, but predicted 1", res0TC1)
	}
	res0TC2 := testCombined(2)

	if res0TC2 != 1 {
		t.Errorf("res0TC2 = %v, but predicted 1", res0TC2)
	}
	res0TC3 := testCombined(3)

	if res0TC3 != 6 {
		t.Errorf("res0TC3 = %v, but predicted 6", res0TC3)
	}
	res0TC4 := testCombined(4)

	if res0TC4 != 7 {
		t.Errorf("res0TC4 = %v, but predicted 7", res0TC4)
	}
	res0TC5 := testCombined(5)

	if res0TC5 != 5 {
		t.Errorf("res0TC5 = %v, but predicted 5", res0TC5)
	}
	res0TC6 := testCombined(6)

	if res0TC6 != 30 {
		t.Errorf("res0TC6 = %v, but predicted 30", res0TC6)
	}
	res0TC7 := testCombined(7)

	if res0TC7 != 31 {
		t.Errorf("res0TC7 = %v, but predicted 31", res0TC7)
	}
}

func TestGenerated_testEdgeCases(t *testing.T) {
	res0TC0 := testEdgeCases(0)

	if res0TC0 != 0 {
		t.Errorf("res0TC0 = %v, but predicted 0", res0TC0)
	}
	res0TC1 := testEdgeCases(-9223372036854775808)

	if res0TC1 != -1 {
		t.Errorf("res0TC1 = %v, but predicted -1", res0TC1)
	}
	res0TC2 := testEdgeCases(2)

	if res0TC2 != 1 {
		t.Errorf("res0TC2 = %v, but predicted 1", res0TC2)
	}
}

func TestGenerated_testMultipleReturns(t *testing.T) {
	res0TC0 := testMultipleReturns(-9223372036854775808)

	if res0TC0 != -9223372036854775808 {
		t.Errorf("res0TC0 = %v, but predicted -9223372036854775808", res0TC0)
	}
	res0TC1 := testMultipleReturns(0)

	if res0TC1 != 0 {
		t.Errorf("res0TC1 = %v, but predicted 0", res0TC1)
	}
	res0TC2 := testMultipleReturns(2)

	if res0TC2 != 4 {
		t.Errorf("res0TC2 = %v, but predicted 4", res0TC2)
	}
}

func TestGenerated_testSimpleSum(t *testing.T) {
	res0TC0 := testSimpleSum(-1.0999997294968582, 0)

	if res0TC0 != 1.2 {
		t.Errorf("res0TC0 = %v, but predicted 1.2", res0TC0)
	}
	res0TC1 := testSimpleSum(-513.0377868715245, 523.0000171663282)

	if res0TC1 != 1.1 {
		t.Errorf("res0TC1 = %v, but predicted 1.1", res0TC1)
	}
	res0TC2 := testSimpleSum(-2.1000000000656613, 673.0937500000001)

	if res0TC2 != 1.2 {
		t.Errorf("res0TC2 = %v, but predicted 1.2", res0TC2)
	}
}
