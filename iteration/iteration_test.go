package iteration
import "testing"
import "fmt"

func TestRepeat(t *testing.T){
	repeated := Repeat("a", 10)
	expected := "aaaaaaaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func ExampleRepeat(){
	repeated := Repeat("r", 3)
	fmt.Println(repeated)
	// Output: rrr
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i< b.N; i++ {
		Repeat("a", 5)
	}
}