package math101

import "testing"

func TestNegate(t *testing.T) {
	v := Vec3{1, 1, 1}
	v.Negate()
	if v != (Vec3{-1, -1, -1}) {
		t.Errorf("TestNegate[0] failed")
	}

	v = Vec3{-1, 0, 1}
	v.Negate()
	v.Negate()
	if v != (Vec3{-1, 0, 1}) {
		t.Errorf("TestNegate[1] failed")
	}
}

func TestLength(t *testing.T) {
	v := Vec3{1, 0, 0}
	if v.Length() != 1 {
		t.Errorf("TestLength[0] failed")
	}

	v = Vec3{1, -2, 2}
	if v.Length() != 3 {
		t.Errorf("TestLength[1] failed")
	}
}

func TestAdd(t *testing.T) {
	if Add(Vec3{0, 1, 2}, Vec3{0, -1, -2}) != (Vec3{0, 0, 0}) {
		t.Errorf("TestAdd[0] failed")
	}

	v := Vec3{5, 3, 7}
	if Add(v, Vec3{0, 0, 0}) != v {
		t.Errorf("TestAdd[1] failed")
	}
}

func TestDot(t *testing.T) {
	if Dot(Vec3{1, 2, 3}, Vec3{0, 0, 0}) != 0 {
		t.Errorf("TestDot[0] failed")
	}

	if Dot(Vec3{1, 2, 3}, Vec3{1, 1, 1}) != 6 {
		t.Errorf("TestDot[1] failed")
	}

	if Dot(Vec3{-1, -2, -5}, Vec3{1, 2, 3}) != -20 {
		t.Errorf("TestDot[2] failed")
	}
}

func TestLerp(t *testing.T) {
	if Lerp(Vec3{0, 0, 0}, Vec3{1, 1, 1}, 0.5) != (Vec3{0.5, 0.5, 0.5}) {
		t.Errorf("TestLerp[0] failed")
	}

	if Lerp(Vec3{-1, -1, -1}, Vec3{1, 1, 1}, 0.5) != (Vec3{0, 0, 0}) {
		t.Errorf("TestLerp[1] failed")
	}
}
