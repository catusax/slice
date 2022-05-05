package slice_tool

import (
	"github.com/ysmood/got"
	"testing"
)

func TestIsEmpty(t *testing.T) {
	var sli []string

	got.T(t).Equal(IsEmpty(sli), true)
}
func TestIsNotEmpty(t *testing.T) {
	var sli = []string{"12e31", "ada"}

	got.T(t).Equal(IsNotEmpty(sli), true)
}

func TestRemove(t *testing.T) {
	var sli = []string{"12e31", "ada"}

	sli = Remove(sli, 1)

	got.T(t).Equal(len(sli), 1)
}

func TestInsert(t *testing.T) {
	var sli = []string{"12e31", "ada"}

	sli = Insert(sli, 1, "bc")

	got.T(t).Equal(len(sli), 3)
	got.T(t).Equal(sli[1], "bc")
}

type testStruct struct {
	Name string
	Age  int
}

func TestSearch(t *testing.T) {
	var sli = []testStruct{{
		Name: "a",
	}, {}}

	got.T(t).Equal(Search(sli, testStruct{Name: "a"}), 0)
}

func TestContains(t *testing.T) {
	var sli = []string{"12e31", "ada"}

	sli = Insert(sli, 1, "bc")

	got.T(t).Equal(len(sli), 3)
	got.T(t).Equal(sli[1], "bc")
}

func TestUnique(t *testing.T) {
	var sli = []string{"12e31", "ada", "ada"}

	sli = Unique(sli)

	got.T(t).Equal(len(sli), 2)
}

func TestFilter(t *testing.T) {
	var sli = []string{"12e31", "ada", "ada"}

	sli = Filter(sli, func(item string, index int) bool {
		if item == "ada" {
			return true
		}
		return false
	})

	got.T(t).Equal(len(sli), 2)
}

func TestReject(t *testing.T) {
	var sli = []string{"12e31", "ada", "ada"}

	sli = Reject(sli, func(item string, index int) bool {
		if item == "ada" {
			return true
		}
		return false
	})

	got.T(t).Equal(len(sli), 1)
}

func TestCopy(t *testing.T) {
	var sli = []string{"12e31", "ada", "ada"}

	sli2 := Copy(sli)
	sli[1] = "bc"

	got.T(t).Neq(sli2[1], "bc")
}

func TestEach(t *testing.T) {
	var sli = []string{"12e31", "ada", "ada", "c"}
	var i = 0

	Each(sli, func(item string, index int) {
		i += index
	})

	got.T(t).Equal(i, 6)
}

func TestMap(t *testing.T) {
	var sli = []string{"12e31", "ada", "ada"}

	sli2 := Map[string, int](sli, func(item string, index int) int {
		return index
	})

	got.T(t).Equal(sli2[2], 2)
}

func TestEvery(t *testing.T) {
	var sli = []string{"12e31", "ada", "ada"}

	ev := Every(sli, func(item string, index int) bool {
		return true
	})

	got.T(t).Equal(ev, true)

	ev2 := Every(sli, func(item string, index int) bool {
		if item == "ada" {
			return false
		}
		return true
	})

	got.T(t).Equal(ev2, false)

}

func TestForPage(t *testing.T) {
	var sli = []string{"12e31", "adas", "ada"}

	page2 := ForPage(sli, 1, 2)

	got.T(t).Equal(len(page2), 1)
	got.T(t).Equal(page2[0], "ada")
}

func TestNth(t *testing.T) {
	var sli = []string{"12e31", "adas", "ada"}

	sli2 := Nth(sli, 2, 0)
	got.T(t).Equal(len(sli2), 2)
	got.T(t).Equal(sli2[1], "ada")
}

func TestPad(t *testing.T) {
	var sli = []string{"12e31", "adas", "ada"}

	sli2 := Pad(sli, 5, "a")
	got.T(t).Equal(len(sli2), 5)
	got.T(t).Equal(sli2[3], "a")
}

func TestPop(t *testing.T) {
	var sli = []string{"12e31", "ada"}

	got.T(t).Equal(Pop(&sli), "ada")
	got.T(t).Equal(len(sli), 1)
}

func TestPush(t *testing.T) {
	var sli = []string{"12e31", "ada"}

	got.T(t).Equal(Push(&sli, "bc"), 3)
	got.T(t).Equal(sli[2], "bc")
}

func TestShift(t *testing.T) {
	var sli = []string{"12e31", "ada"}

	got.T(t).Equal(Shift(&sli), "12e31")
	got.T(t).Equal(len(sli), 1)
}

func TestUnShift(t *testing.T) {
	var sli = []string{"12e31", "ada"}

	got.T(t).Equal(UnShift(&sli, "bc"), 3)
	got.T(t).Equal(sli[0], "bc")
}
