package txt

import (
	"testing"
)

var word = "testing"

func TestGroupWord(t *testing.T) {
	s := groupWord(word)
	correct := []string{"t", "e", "st", "i", "ng"}

	for i, v := range s {
		if correct[i] != v {
			t.Fail()
		}
	}
}

func TestEncode(t *testing.T) {
	s := groupWord(word)
	encode(s)
}

func TestEndsWithCC(t *testing.T) {
	tru := stemEndsCC(word)
	fal := stemEndsCC("control")

	if !tru && fal {
		t.Fail()
	}
}

func TestEndsWithCVC(t *testing.T) {
	a := stemEndsCVC(groupWord(word))
	b := stemEndsCVC(groupWord("tree"))
	c := stemEndsCVC(groupWord("prolyx"))

	if !a || b || c {
		t.Fail()
	}
}

func TestPorterOne(t *testing.T) {
	responses := []string{
		porter_one("caresses"),
		porter_one("flies"),
		porter_one("traps"),
		porter_one("feed"),
		porter_one("motoring"),
		porter_one("plastered"),
	}
	values := []string{
		"caress",
		"fli",
		"trap",
		"feed",
		"motor",
		"plaster",
	}

	for i, v := range responses {
		if values[i] != v {
			t.Fail()
		}
	}
}

func TestPorterTwo(t *testing.T) {
	responses := []string{
		porter_two("relational"),
		porter_two("rational"),
		porter_two("predication"),
		porter_two("callousness"),
		porter_two("decisiveness"),
		porter_two("feudalism"),
	}

	values := []string{
		"relate",
		"rational",
		"predicate",
		"callous",
		"decisive",
		"feudal",
	}

	for i, v := range responses {
		if values[i] != v {
			t.Logf("expected %v, got %v\n", values[i], v)
			t.Fail()
		}
	}
}

func TestPorterThree(t *testing.T) {
	responses := []string{
		porter_three("triplicate"),
		porter_three("formative"),
		porter_three("formalize"),
		porter_three("electriciti"),
		porter_three("electrical"),
		porter_three("hopeful"),
		porter_three("goodness"),
	}

	values := []string{
		"triplic",
		"form",
		"formal",
		"electric",
		"electric",
		"hope",
		"good",
	}

	for i, v := range responses {
		if values[i] != v {
			t.Logf("expected %v, got %v\n", values[i], v)
			t.Fail()
		}
	}
}

func TestPorterFour(t *testing.T) {
	responses := []string{
		porter_four("revival"),
		porter_four("allowance"),
		porter_four("airliner"),
		porter_four("irritant"),
		porter_four("communism"),
		porter_four("activate"),
		porter_four("effective"),
		porter_four("adoption"),
	}

	values := []string{
		"reviv",
		"allow",
		"airlin",
		"irrit",
		"commun",
		"activ",
		"effect",
		"adopt",
	}

	for i, v := range responses {
		if values[i] != v {
			t.Logf("expected %v, got %v\n", values[i], v)
			t.Fail()
		}
	}
}

func TestPorterFive(t *testing.T) {
	responses := []string{
		porter_five("probate"),
		porter_five("rate"),
		porter_five("cease"),
		porter_five("controll"),
		porter_five("roll"),
	}

	values := []string{
		"probat",
		"rate",
		"ceas",
		"control",
		"roll",
	}

	for i, v := range responses {
		if values[i] != v {
			t.Logf("expected %v, got %v\n", values[i], v)
			t.Fail()
		}
	}
}

func TestStem(t *testing.T) {
	responses := []string{
		Stem("relational"),
		Stem("rational"),
		Stem("predication"),
		Stem("callousness"),
		Stem("decisiveness"),
		Stem("feudalism"),
	}

	values := []string{
		"relat",
		"ration",
		"predic",
		"callous",
		"decis",
		"feudal",
	}

	for i, v := range responses {
		if values[i] != v {
			t.Logf("expected %v, got %v\n", values[i], v)
			t.Fail()
		}
	}
}

func BenchmarkStem(b *testing.B) {
	Stem("antidisestablishmentarianism")
}
