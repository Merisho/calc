package calc

import "testing"

func BenchmarkCalcSinglePriority(b *testing.B) {
    c := NewCalc()
    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        c.Calc("1 + 2 + 3 + 4 - 5 + 6 - 7 + 8 + 9 + 10")
    }
}

func BenchmarkCalcMultiPriority(b *testing.B) {
    c := NewCalc()
    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        c.Calc("1 + 2 * 3 * 4 / 8 - 6 + 7 + 8 * 9 / 10 - 2 + 123 - 32 + 4523 * 100 / 200")
    }
}

func BenchmarkCalcComplex(b *testing.B) {
    c := NewCalc()

    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        c.Calc("234 + 5135 - (42341 + 12354) + 131142 * 4212 - (6542 + 3513 - (3534 / 125 + 51312 - 1243 * 54123 - (51341 - 43214 - 78765 - (86345 * 624572 + 523452 - 1245609 / 806959 - 7950934) + 624 - 1 * 62323)))")
    }
}
