package main

import (
	"go-reloaded/internal"
	"reflect"
	"testing"
)

func TestHex(t *testing.T) {
	cas := []struct {
		nom        string // la famille du cas : normal, limite, erreur, piège
		entree     []rune // l'entrée
		veut       []rune // la sortie attendue
		veutErreur bool   // si on attend une erreur
	}{
		{"normal : nombre hexa", []rune{'F', '1', '4'}, []rune{'3', '8', '6', '0'}, false},
		{"erreur : pas un nombre hexadecimal", []rune{'A', 'r', 'b', 'r', 'e'}, []rune{}, true},
		{"limite : nombre hexa vide", []rune{}, []rune{'0'}, false},
		{"piège : nombre hexa avec minuscules", []rune{'a', 'b', 'c'}, []rune{'2', '7', '4', '8'}, false},
	}

	for _, c := range cas {
		t.Run(c.nom, func(t *testing.T) {
			got, err := internal.Hex(c.entree)
			if (err != nil) != c.veutErreur {
				t.Fatalf("erreur = %v, erreur attendue : %v", err, c.veutErreur)
			}
			if !reflect.DeepEqual(got, c.veut) {
				t.Errorf("Hex(%q) = %q, attendu %q", c.entree, got, c.veut)
			}
		})
	}
}

func TestBin(t *testing.T) {
	cas := []struct {
		nom        string // la famille du cas : normal, limite, erreur, piège
		entree     []rune // l'entrée
		veut       []rune // la sortie attendue
		veutErreur bool   // si on attend une erreur
	}{
		{"normal : nombre binaire", []rune{'1', '0', '1', '1', '1', '0'}, []rune{'4', '6'}, false},
		{"erreur : pas nombre binaire", []rune{'A', 'r', 'b', 'r', 'e'}, []rune{}, true},
		{"limite : nombre binaire vide", []rune{}, []rune{'0'}, false},
	}

	for _, c := range cas {
		t.Run(c.nom, func(t *testing.T) {
			got, err := internal.Bin(c.entree)
			if (err != nil) != c.veutErreur {
				t.Fatalf("erreur = %v, erreur attendue : %v", err, c.veutErreur)
			}
			if !reflect.DeepEqual(got, c.veut) {
				t.Errorf("Bin(%q) = %q, attendu %q", c.entree, got, c.veut)
			}
		})
	}
}

func TestUp(t *testing.T) {
	cas := []struct {
		nom    string // la famille du cas : normal, limite, erreur, piège
		entree []rune // l'entrée
		veut   []rune // la sortie attendue
	}{
		{"normal : mot minuscule", []rune{'a', 'r', 'b', 'r', 'e'}, []rune{'A', 'R', 'B', 'R', 'E'}},
		{"normal : mot mixte", []rune{'a', 'R', 'B', 'r', 'E'}, []rune{'A', 'R', 'B', 'R', 'E'}},
		{"limite : mot avec nombre ou caractères spéciaux", []rune{'O', 'u', 'i', '6', '5', '.'}, []rune{'O', 'U', 'I', '6', '5', '.'}},
		{"piège : mot déjà en majuscule", []rune{'A', 'R', 'B', 'R', 'E'}, []rune{'A', 'R', 'B', 'R', 'E'}},
	}

	for _, c := range cas {
		t.Run(c.nom, func(t *testing.T) {
			got := internal.Up(c.entree)
			if !reflect.DeepEqual(got, c.veut) {
				t.Errorf("Up(%q) = %q, attendu %q", c.entree, got, c.veut)
			}
		})
	}
}
func TestLow(t *testing.T) {
	cas := []struct {
		nom    string // la famille du cas : normal, limite, erreur, piège
		entree []rune // l'entrée
		veut   []rune // la sortie attendue
	}{
		{"normal : mot majuscule", []rune{'A', 'R', 'B', 'R', 'E'}, []rune{'a', 'r', 'b', 'r', 'e'}},
		{"normal : mot mixte", []rune{'a', 'R', 'B', 'r', 'E'}, []rune{'a', 'r', 'b', 'r', 'e'}},
		{"limite : mot avec nombre ou caractères spéciaux", []rune{'O', 'u', 'i', '6', '5', '.'}, []rune{'o', 'u', 'i', '6', '5', '.'}},
		{"piège : mot déjà en minuscule", []rune{'a', 'r', 'b', 'r', 'e'}, []rune{'a', 'r', 'b', 'r', 'e'}},
	}

	for _, c := range cas {
		t.Run(c.nom, func(t *testing.T) {
			got := internal.Low(c.entree)
			if !reflect.DeepEqual(got, c.veut) {
				t.Errorf("Low(%q) = %q, attendu %q", c.entree, got, c.veut)
			}
		})
	}
}

func TestCap(t *testing.T) {
	cas := []struct {
		nom    string // la famille du cas : normal, limite, erreur, piège
		entree []rune // l'entrée
		veut   []rune // la sortie attendue
	}{
		{"normal : mot minuscule", []rune{'a', 'r', 'b', 'r', 'e'}, []rune{'A', 'r', 'b', 'r', 'e'}},
		{"normal : mot mixte", []rune{'a', 'R', 'B', 'r', 'E'}, []rune{'A', 'R', 'B', 'r', 'E'}},
		{"limite : mot avec nombre", []rune{'3', 'a', 'r', 'b', 'r', 'e'}, []rune{'3', 'a', 'r', 'b', 'r', 'e'}},
		{"limite : mot avec caractère spécial", []rune{'\'', 'o', 'u', 'i'}, []rune{'\'', 'o', 'u', 'i'}},
	}

	for _, c := range cas {
		t.Run(c.nom, func(t *testing.T) {
			got := internal.Cap(c.entree)
			if !reflect.DeepEqual(got, c.veut) {
				t.Errorf("Cap(%q) = %q, attendu %q", c.entree, got, c.veut)
			}
		})
	}
}
