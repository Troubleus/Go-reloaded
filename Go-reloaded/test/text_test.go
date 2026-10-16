package main

import (
	"reflect"
	"go-reloaded/internal"
	"testing"
)

func TestTextToSlice(t *testing.T) {
	cas := []struct {
		nom        string // la famille du cas : normal, limite, erreur, piège
		entree     string // l'entrée
		veut      [][][]rune // la sortie attendue
	}{
		{"normal : text", "ceci (up) est un test .\nComment ça va ?", [][][]rune{{{'c','e','c','i'},{'(','u','p',')'},{'e','s','t'},{'u','n'},{'t','e','s','t'},{'.'}}, {{'C','o','m','m','e','n','t'},{'ç','a'},{'v','a'},{'?'}}}},
		{"limite : text vide", "", [][][]rune{{{}}}},
		{"limite : text sans espace", "ARBRE", [][][]rune{{{'A','R','B','R','E'}}}},
		{"limite : text avec plusieurs retour à la ligne", "ceci (up) est un test .\n\nComment ça va ?", [][][]rune{{{'c','e','c','i'},{'(','u','p',')'},{'e','s','t'},{'u','n'},{'t','e','s','t'},{'.'}},{{}},{{'C','o','m','m','e','n','t'},{'ç','a'},{'v','a'},{'?'}}}},
	}
	for _, c := range cas {
		t.Run(c.nom, func(t *testing.T) {
			got := internal.TextToSlice(c.entree)
			if !reflect.DeepEqual(got, c.veut) {
				t.Fatalf("résultat = %v, résultat attendu : %v", got, c.veut)
			}
		})
	}
}
func TestSliceToText(t *testing.T) {
	cas := []struct {
		nom        string // la famille du cas : normal, limite, erreur, piège
		entree     [][]rune // l'entrée
		veut      string // la sortie attendue
	}{
		{"normal : text", [][]rune{{'C','e','c','i'},{'e','s','t'},{'\'','u','n'},{'t','e','s','t','\'','.'}}, "Ceci est 'un test'."},
	}
	for _, c := range cas {
		t.Run(c.nom, func(t *testing.T) {
			got := internal.SliceToText(c.entree)
			if !reflect.DeepEqual(got, c.veut) {
				t.Fatalf("résultat = %v, résultat attendu : %v", got, c.veut)
			}
		})
	}
}