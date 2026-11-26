package main

import (
	"testing"
)

func panicCheck_test_functions(t *testing.T, panicExpect string, fn func()) {
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
	res0tc0 := test1(12)

	if res0tc0 != 13 {
		t.Errorf("res0tc0 = %v, but predicted 13", res0tc0)
	}
	res0tc1 := test1(-9223372036854775796)

	if res0tc1 != -9223372036854775797 {
		t.Errorf("res0tc1 = %v, but predicted -9223372036854775797", res0tc1)
	}
}

func TestGenerated_test2(t *testing.T) {
	res0tc0 := test2(false, false)

	if res0tc0 != false {
		t.Errorf("res0tc0 = %v, but predicted false", res0tc0)
	}
	res0tc1 := test2(true, true)

	if res0tc1 != true {
		t.Errorf("res0tc1 = %v, but predicted true", res0tc1)
	}
	res0tc2 := test2(true, false)

	if res0tc2 != false {
		t.Errorf("res0tc2 = %v, but predicted false", res0tc2)
	}
}

func TestGenerated_testArithmetic(t *testing.T) {
	res0tc0 := testArithmetic(-8919549788968174172, 3004641839422955821)

	if res0tc0 != 1 {
		t.Errorf("res0tc0 = %v, but predicted 1", res0tc0)
	}
	res0tc1 := testArithmetic(3180548182006436739, 6623606982543203939)

	if res0tc1 != 6 {
		t.Errorf("res0tc1 = %v, but predicted 6", res0tc1)
	}
}

func TestGenerated_testUnary(t *testing.T) {
	res0tc0 := testUnary(9223372036854775806)

	if res0tc0 != 9223372036854775805 {
		t.Errorf("res0tc0 = %v, but predicted 9223372036854775805", res0tc0)
	}
	res0tc1 := testUnary(-2)

	if res0tc1 != 2 {
		t.Errorf("res0tc1 = %v, but predicted 2", res0tc1)
	}
}

func TestGenerated_testComparisons(t *testing.T) {
	res0tc0 := testComparisons(0, 0)

	if res0tc0 != true {
		t.Errorf("res0tc0 = %v, but predicted true", res0tc0)
	}
	res0tc1 := testComparisons(9223372036854775806, -9223372036854775807)

	if res0tc1 != true {
		t.Errorf("res0tc1 = %v, but predicted true", res0tc1)
	}
	res0tc2 := testComparisons(0, -1)

	if res0tc2 != true {
		t.Errorf("res0tc2 = %v, but predicted true", res0tc2)
	}
	res0tc3 := testComparisons(-8853654483146637056, 8853654483146637055)

	if res0tc3 != true {
		t.Errorf("res0tc3 = %v, but predicted true", res0tc3)
	}
}

func TestGenerated_testLogicalOps(t *testing.T) {
	res0tc0 := testLogicalOps(-9223372036854775806, 0)

	if res0tc0 != true {
		t.Errorf("res0tc0 = %v, but predicted true", res0tc0)
	}
	res0tc1 := testLogicalOps(0, 0)

	if res0tc1 != false {
		t.Errorf("res0tc1 = %v, but predicted false", res0tc1)
	}
	res0tc2 := testLogicalOps(2, 0)

	if res0tc2 != false {
		t.Errorf("res0tc2 = %v, but predicted false", res0tc2)
	}
}

func TestGenerated_testWhileLoop(t *testing.T) {
	res0tc0 := testWhileLoop(-9223372036854775806)

	if res0tc0 != 0 {
		t.Errorf("res0tc0 = %v, but predicted 0", res0tc0)
	}
	res0tc1 := testWhileLoop(1)

	if res0tc1 != 0 {
		t.Errorf("res0tc1 = %v, but predicted 0", res0tc1)
	}
	res0tc2 := testWhileLoop(2)

	if res0tc2 != 1 {
		t.Errorf("res0tc2 = %v, but predicted 1", res0tc2)
	}
	res0tc3 := testWhileLoop(3)

	if res0tc3 != 3 {
		t.Errorf("res0tc3 = %v, but predicted 3", res0tc3)
	}
	res0tc4 := testWhileLoop(4)

	if res0tc4 != 6 {
		t.Errorf("res0tc4 = %v, but predicted 6", res0tc4)
	}
	res0tc5 := testWhileLoop(5)

	if res0tc5 != 10 {
		t.Errorf("res0tc5 = %v, but predicted 10", res0tc5)
	}
	res0tc6 := testWhileLoop(6)

	if res0tc6 != 15 {
		t.Errorf("res0tc6 = %v, but predicted 15", res0tc6)
	}
	res0tc7 := testWhileLoop(7)

	if res0tc7 != 21 {
		t.Errorf("res0tc7 = %v, but predicted 21", res0tc7)
	}
	res0tc8 := testWhileLoop(8)

	if res0tc8 != 28 {
		t.Errorf("res0tc8 = %v, but predicted 28", res0tc8)
	}
	res0tc9 := testWhileLoop(9)

	if res0tc9 != 36 {
		t.Errorf("res0tc9 = %v, but predicted 36", res0tc9)
	}
	res0tc10 := testWhileLoop(10)

	if res0tc10 != 45 {
		t.Errorf("res0tc10 = %v, but predicted 45", res0tc10)
	}
	res0tc11 := testWhileLoop(11)

	if res0tc11 != 55 {
		t.Errorf("res0tc11 = %v, but predicted 55", res0tc11)
	}
	res0tc12 := testWhileLoop(12)

	if res0tc12 != 66 {
		t.Errorf("res0tc12 = %v, but predicted 66", res0tc12)
	}
	res0tc13 := testWhileLoop(13)

	if res0tc13 != 78 {
		t.Errorf("res0tc13 = %v, but predicted 78", res0tc13)
	}
	res0tc14 := testWhileLoop(14)

	if res0tc14 != 91 {
		t.Errorf("res0tc14 = %v, but predicted 91", res0tc14)
	}
}

func TestGenerated_testForLoop(t *testing.T) {
	res0tc0 := testForLoop(0)

	if res0tc0 != 1 {
		t.Errorf("res0tc0 = %v, but predicted 1", res0tc0)
	}
	res0tc1 := testForLoop(1)

	if res0tc1 != 2 {
		t.Errorf("res0tc1 = %v, but predicted 2", res0tc1)
	}
	res0tc2 := testForLoop(2)

	if res0tc2 != 4 {
		t.Errorf("res0tc2 = %v, but predicted 4", res0tc2)
	}
	res0tc3 := testForLoop(3)

	if res0tc3 != 7 {
		t.Errorf("res0tc3 = %v, but predicted 7", res0tc3)
	}
	res0tc4 := testForLoop(4)

	if res0tc4 != 28 {
		t.Errorf("res0tc4 = %v, but predicted 28", res0tc4)
	}
	res0tc5 := testForLoop(5)

	if res0tc5 != 33 {
		t.Errorf("res0tc5 = %v, but predicted 33", res0tc5)
	}
	res0tc6 := testForLoop(6)

	if res0tc6 != 198 {
		t.Errorf("res0tc6 = %v, but predicted 198", res0tc6)
	}
	res0tc7 := testForLoop(7)

	if res0tc7 != 205 {
		t.Errorf("res0tc7 = %v, but predicted 205", res0tc7)
	}
	res0tc8 := testForLoop(8)

	if res0tc8 != 1640 {
		t.Errorf("res0tc8 = %v, but predicted 1640", res0tc8)
	}
	res0tc9 := testForLoop(9)

	if res0tc9 != 1649 {
		t.Errorf("res0tc9 = %v, but predicted 1649", res0tc9)
	}
	res0tc10 := testForLoop(10)

	if res0tc10 != 16490 {
		t.Errorf("res0tc10 = %v, but predicted 16490", res0tc10)
	}
	res0tc11 := testForLoop(11)

	if res0tc11 != 16501 {
		t.Errorf("res0tc11 = %v, but predicted 16501", res0tc11)
	}
	res0tc12 := testForLoop(12)

	if res0tc12 != 198012 {
		t.Errorf("res0tc12 = %v, but predicted 198012", res0tc12)
	}
	res0tc13 := testForLoop(13)

	if res0tc13 != 198025 {
		t.Errorf("res0tc13 = %v, but predicted 198025", res0tc13)
	}
	res0tc14 := testForLoop(14)

	if res0tc14 != 2772350 {
		t.Errorf("res0tc14 = %v, but predicted 2772350", res0tc14)
	}
}

func TestGenerated_testInfiniteLoopBreak(t *testing.T) {
	res0tc0 := testInfiniteLoopBreak(-9223372036854775806)

	if res0tc0 != 0 {
		t.Errorf("res0tc0 = %v, but predicted 0", res0tc0)
	}
	res0tc1 := testInfiniteLoopBreak(1)

	if res0tc1 != 1 {
		t.Errorf("res0tc1 = %v, but predicted 1", res0tc1)
	}
	res0tc2 := testInfiniteLoopBreak(2)

	if res0tc2 != 2 {
		t.Errorf("res0tc2 = %v, but predicted 2", res0tc2)
	}
	res0tc3 := testInfiniteLoopBreak(3)

	if res0tc3 != 3 {
		t.Errorf("res0tc3 = %v, but predicted 3", res0tc3)
	}
	res0tc4 := testInfiniteLoopBreak(4)

	if res0tc4 != 4 {
		t.Errorf("res0tc4 = %v, but predicted 4", res0tc4)
	}
	res0tc5 := testInfiniteLoopBreak(5)

	if res0tc5 != 5 {
		t.Errorf("res0tc5 = %v, but predicted 5", res0tc5)
	}
	res0tc6 := testInfiniteLoopBreak(6)

	if res0tc6 != 6 {
		t.Errorf("res0tc6 = %v, but predicted 6", res0tc6)
	}
	res0tc7 := testInfiniteLoopBreak(7)

	if res0tc7 != 7 {
		t.Errorf("res0tc7 = %v, but predicted 7", res0tc7)
	}
	res0tc8 := testInfiniteLoopBreak(8)

	if res0tc8 != 8 {
		t.Errorf("res0tc8 = %v, but predicted 8", res0tc8)
	}
	res0tc9 := testInfiniteLoopBreak(9)

	if res0tc9 != 9 {
		t.Errorf("res0tc9 = %v, but predicted 9", res0tc9)
	}
	res0tc10 := testInfiniteLoopBreak(10)

	if res0tc10 != 10 {
		t.Errorf("res0tc10 = %v, but predicted 10", res0tc10)
	}
	res0tc11 := testInfiniteLoopBreak(11)

	if res0tc11 != 11 {
		t.Errorf("res0tc11 = %v, but predicted 11", res0tc11)
	}
	res0tc12 := testInfiniteLoopBreak(12)

	if res0tc12 != 12 {
		t.Errorf("res0tc12 = %v, but predicted 12", res0tc12)
	}
	res0tc13 := testInfiniteLoopBreak(13)

	if res0tc13 != 13 {
		t.Errorf("res0tc13 = %v, but predicted 13", res0tc13)
	}
	res0tc14 := testInfiniteLoopBreak(14)

	if res0tc14 != 14 {
		t.Errorf("res0tc14 = %v, but predicted 14", res0tc14)
	}
}

func TestGenerated_testLoopWithConcreteBoundAndSymbolicBranching(t *testing.T) {
	res0tc0 := testLoopWithConcreteBoundAndSymbolicBranching(false)

	if res0tc0 != 0 {
		t.Errorf("res0tc0 = %v, but predicted 0", res0tc0)
	}
	res0tc1 := testLoopWithConcreteBoundAndSymbolicBranching(true)

	if res0tc1 != 20 {
		t.Errorf("res0tc1 = %v, but predicted 20", res0tc1)
	}
}

func TestGenerated_testLoopWithSymbolicBoundAndSymbolicBranching(t *testing.T) {
	panicCheck_test_functions(t, "Assumption violated: n should be less than or equal to 10", func() { testLoopWithSymbolicBoundAndSymbolicBranching(12, false) })
	res0tc1 := testLoopWithSymbolicBoundAndSymbolicBranching(-9223372036854775794, false)

	if res0tc1 != 0 {
		t.Errorf("res0tc1 = %v, but predicted 0", res0tc1)
	}
	res0tc2 := testLoopWithSymbolicBoundAndSymbolicBranching(1, false)

	if res0tc2 != 0 {
		t.Errorf("res0tc2 = %v, but predicted 0", res0tc2)
	}
	res0tc3 := testLoopWithSymbolicBoundAndSymbolicBranching(1, true)

	if res0tc3 != 0 {
		t.Errorf("res0tc3 = %v, but predicted 0", res0tc3)
	}
	res0tc4 := testLoopWithSymbolicBoundAndSymbolicBranching(2, false)

	if res0tc4 != 0 {
		t.Errorf("res0tc4 = %v, but predicted 0", res0tc4)
	}
	res0tc5 := testLoopWithSymbolicBoundAndSymbolicBranching(3, false)

	if res0tc5 != 0 {
		t.Errorf("res0tc5 = %v, but predicted 0", res0tc5)
	}
	res0tc6 := testLoopWithSymbolicBoundAndSymbolicBranching(2, true)

	if res0tc6 != 0 {
		t.Errorf("res0tc6 = %v, but predicted 0", res0tc6)
	}
	res0tc7 := testLoopWithSymbolicBoundAndSymbolicBranching(4, false)

	if res0tc7 != 0 {
		t.Errorf("res0tc7 = %v, but predicted 0", res0tc7)
	}
	res0tc8 := testLoopWithSymbolicBoundAndSymbolicBranching(3, true)

	if res0tc8 != 2 {
		t.Errorf("res0tc8 = %v, but predicted 2", res0tc8)
	}
	res0tc9 := testLoopWithSymbolicBoundAndSymbolicBranching(5, false)

	if res0tc9 != 0 {
		t.Errorf("res0tc9 = %v, but predicted 0", res0tc9)
	}
	res0tc10 := testLoopWithSymbolicBoundAndSymbolicBranching(6, false)

	if res0tc10 != 0 {
		t.Errorf("res0tc10 = %v, but predicted 0", res0tc10)
	}
	res0tc11 := testLoopWithSymbolicBoundAndSymbolicBranching(4, true)

	if res0tc11 != 2 {
		t.Errorf("res0tc11 = %v, but predicted 2", res0tc11)
	}
	res0tc12 := testLoopWithSymbolicBoundAndSymbolicBranching(7, false)

	if res0tc12 != 0 {
		t.Errorf("res0tc12 = %v, but predicted 0", res0tc12)
	}
	res0tc13 := testLoopWithSymbolicBoundAndSymbolicBranching(5, true)

	if res0tc13 != 6 {
		t.Errorf("res0tc13 = %v, but predicted 6", res0tc13)
	}
	res0tc14 := testLoopWithSymbolicBoundAndSymbolicBranching(8, false)

	if res0tc14 != 0 {
		t.Errorf("res0tc14 = %v, but predicted 0", res0tc14)
	}
	res0tc15 := testLoopWithSymbolicBoundAndSymbolicBranching(9, false)

	if res0tc15 != 0 {
		t.Errorf("res0tc15 = %v, but predicted 0", res0tc15)
	}
	res0tc16 := testLoopWithSymbolicBoundAndSymbolicBranching(6, true)

	if res0tc16 != 6 {
		t.Errorf("res0tc16 = %v, but predicted 6", res0tc16)
	}
	res0tc17 := testLoopWithSymbolicBoundAndSymbolicBranching(10, false)

	if res0tc17 != 0 {
		t.Errorf("res0tc17 = %v, but predicted 0", res0tc17)
	}
	res0tc18 := testLoopWithSymbolicBoundAndSymbolicBranching(7, true)

	if res0tc18 != 12 {
		t.Errorf("res0tc18 = %v, but predicted 12", res0tc18)
	}
	res0tc19 := testLoopWithSymbolicBoundAndSymbolicBranching(8, true)

	if res0tc19 != 12 {
		t.Errorf("res0tc19 = %v, but predicted 12", res0tc19)
	}
	res0tc20 := testLoopWithSymbolicBoundAndSymbolicBranching(9, true)

	if res0tc20 != 20 {
		t.Errorf("res0tc20 = %v, but predicted 20", res0tc20)
	}
	res0tc21 := testLoopWithSymbolicBoundAndSymbolicBranching(10, true)

	if res0tc21 != 20 {
		t.Errorf("res0tc21 = %v, but predicted 20", res0tc21)
	}
}

func TestGenerated_testComplexConditions(t *testing.T) {
	res0tc0 := testComplexConditions(2, 0, 1)

	if res0tc0 != 1 {
		t.Errorf("res0tc0 = %v, but predicted 1", res0tc0)
	}
	res0tc1 := testComplexConditions(459419962931421214, 459419962931421215, 0)

	if res0tc1 != 1 {
		t.Errorf("res0tc1 = %v, but predicted 1", res0tc1)
	}
	res0tc2 := testComplexConditions(4532105087670112556, -4532105087670112557, -9143172115396383018)

	if res0tc2 != 3 {
		t.Errorf("res0tc2 = %v, but predicted 3", res0tc2)
	}
	res0tc3 := testComplexConditions(-2647831093471200330, 2647831093471200329, 4611686018427387904)

	if res0tc3 != 3 {
		t.Errorf("res0tc3 = %v, but predicted 3", res0tc3)
	}
	res0tc4 := testComplexConditions(0, 0, 0)

	if res0tc4 != 2 {
		t.Errorf("res0tc4 = %v, but predicted 2", res0tc4)
	}
	res0tc5 := testComplexConditions(-34, -34, 33)

	if res0tc5 != 3 {
		t.Errorf("res0tc5 = %v, but predicted 3", res0tc5)
	}
}

func TestGenerated_testNestedIf(t *testing.T) {
	res0tc0 := testNestedIf(2, 1)

	if res0tc0 != 1 {
		t.Errorf("res0tc0 = %v, but predicted 1", res0tc0)
	}
	res0tc1 := testNestedIf(2, -9223372036854775807)

	if res0tc1 != 2 {
		t.Errorf("res0tc1 = %v, but predicted 2", res0tc1)
	}
	res0tc2 := testNestedIf(-9223372036854775806, 1)

	if res0tc2 != 3 {
		t.Errorf("res0tc2 = %v, but predicted 3", res0tc2)
	}
	res0tc3 := testNestedIf(-9223372036854775806, -9223372036854775807)

	if res0tc3 != 4 {
		t.Errorf("res0tc3 = %v, but predicted 4", res0tc3)
	}
}

func TestGenerated_testBitwise(t *testing.T) {
	res0tc0 := testBitwise(0, 1)

	if res0tc0 != 2 {
		t.Errorf("res0tc0 = %v, but predicted 2", res0tc0)
	}
	res0tc1 := testBitwise(-9223372036854775808, 1)

	if res0tc1 != -4611686018427387904 {
		t.Errorf("res0tc1 = %v, but predicted -4611686018427387904", res0tc1)
	}
}

func TestGenerated_testCombined(t *testing.T) {
	res0tc0 := testCombined(4611686018427387906)

	if res0tc0 != 0 {
		t.Errorf("res0tc0 = %v, but predicted 0", res0tc0)
	}
	res0tc1 := testCombined(1)

	if res0tc1 != 1 {
		t.Errorf("res0tc1 = %v, but predicted 1", res0tc1)
	}
	res0tc2 := testCombined(2)

	if res0tc2 != 1 {
		t.Errorf("res0tc2 = %v, but predicted 1", res0tc2)
	}
	res0tc3 := testCombined(3)

	if res0tc3 != 6 {
		t.Errorf("res0tc3 = %v, but predicted 6", res0tc3)
	}
	res0tc4 := testCombined(4)

	if res0tc4 != 7 {
		t.Errorf("res0tc4 = %v, but predicted 7", res0tc4)
	}
	res0tc5 := testCombined(5)

	if res0tc5 != 5 {
		t.Errorf("res0tc5 = %v, but predicted 5", res0tc5)
	}
	res0tc6 := testCombined(6)

	if res0tc6 != 30 {
		t.Errorf("res0tc6 = %v, but predicted 30", res0tc6)
	}
	res0tc7 := testCombined(7)

	if res0tc7 != 31 {
		t.Errorf("res0tc7 = %v, but predicted 31", res0tc7)
	}
}

func TestGenerated_testEdgeCases(t *testing.T) {
	res0tc0 := testEdgeCases(0)

	if res0tc0 != 0 {
		t.Errorf("res0tc0 = %v, but predicted 0", res0tc0)
	}
	res0tc1 := testEdgeCases(-9223372036854775808)

	if res0tc1 != -1 {
		t.Errorf("res0tc1 = %v, but predicted -1", res0tc1)
	}
	res0tc2 := testEdgeCases(2)

	if res0tc2 != 1 {
		t.Errorf("res0tc2 = %v, but predicted 1", res0tc2)
	}
}

func TestGenerated_testMultipleReturns(t *testing.T) {
	res0tc0 := testMultipleReturns(-9223372036854775808)

	if res0tc0 != -9223372036854775808 {
		t.Errorf("res0tc0 = %v, but predicted -9223372036854775808", res0tc0)
	}
	res0tc1 := testMultipleReturns(0)

	if res0tc1 != 0 {
		t.Errorf("res0tc1 = %v, but predicted 0", res0tc1)
	}
	res0tc2 := testMultipleReturns(2)

	if res0tc2 != 4 {
		t.Errorf("res0tc2 = %v, but predicted 4", res0tc2)
	}
}

func TestGenerated_testSimpleSum(t *testing.T) {
	res0tc0 := testSimpleSum(-1.0999997294968582, 0)

	if res0tc0 != 1.2 {
		t.Errorf("res0tc0 = %v, but predicted 1.2", res0tc0)
	}
	res0tc1 := testSimpleSum(-513.0377868715245, 523.0000171663282)

	if res0tc1 != 1.1 {
		t.Errorf("res0tc1 = %v, but predicted 1.1", res0tc1)
	}
	res0tc2 := testSimpleSum(-2.1000000000656613, 673.0937500000001)

	if res0tc2 != 1.2 {
		t.Errorf("res0tc2 = %v, but predicted 1.2", res0tc2)
	}
}
