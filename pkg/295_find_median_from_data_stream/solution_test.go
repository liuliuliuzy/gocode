package findmedianfromdatastream

import "testing"

func TestSolution(t *testing.T) {
	obj := Constructor()
	obj.AddNum(2)
	if obj.FindMedian() != 2.0 {
		t.Error("wrong answer")
		t.Errorf("expected: %f, got: %f", 2.0, obj.FindMedian())
	}
	obj.AddNum(1)
	if obj.FindMedian() != 1.5 {
		t.Error("wrong answer")
		t.Errorf("expected: %f, got: %f", 1.5, obj.FindMedian())
	}
	obj.AddNum(3)
	if obj.FindMedian() != 2.0 {
		t.Error("wrong answer")
		t.Errorf("expected: %f, got: %f", 2.0, obj.FindMedian())
	}
	obj.AddNum(4)
	if obj.FindMedian() != 2.5 {
		t.Error("wrong answer")
		t.Errorf("expected: %f, got: %f", 2.5, obj.FindMedian())
	}
}
