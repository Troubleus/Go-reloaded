package main

import (
	"testing"
	"go-reloaded/internal"
	"reflect"
)

func TestFormatPunctuate(t *testing.T) {
	cas := []struct {
		nom        string   // la famille du cas : normal, limite, erreur, piège
		entree     [][]rune // l'entrée
		veut      [][]rune // la sortie attendue
	}{
		{"normal : text",[][]rune{{'C','e','c','i'},{'e','s','t'},{'\''},{'u','n'},{'t','e','s','t'},{'\''},{'.'}},[][]rune{{'C','e','c','i'},{'e','s','t'},{'\'','u','n'},{'t','e','s','t','\'','.'}}},
		{"limite : aucun ponctuation",[][]rune{{'C','e','c','i'},{'e','s','t'},{'u','n'},{'t','e','s','t'}},[][]rune{{'C','e','c','i'},{'e','s','t'},{'u','n'},{'t','e','s','t'}}},
	}
	for _, c := range cas {
		t.Run(c.nom, func(t *testing.T) {
			got := internal.FormatPunctuate(c.entree)
			if !reflect.DeepEqual(got, c.veut) {
				t.Fatalf("résultat = %v, résultat attendu : %v", got, c.veut)
			}
		})
	}
}