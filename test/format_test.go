package main

import (
	"go-reloaded/internal"
	"reflect"
	"testing"
)

func TestParameters(t *testing.T) {
	cas := []struct {
		nom        string // la famille du cas : normal, limite, erreur, piège
		entree     []rune // l'entrée
		veutString string // la sortie attendue
		veutNumber int    // la sortie attendue
		veutErreur bool   // si on attend une erreur
	}{
		// normal : parametre sans nombre
		{"normal : parametre sans nombre", []rune{'(', 'b', 'i', 'n', ')'}, "bin", 0, false},
		// normal : parametre avec nombre
		{"normal : parametre avec nombre", []rune{'(', 'u', 'p', ',', '2', ')'}, "up", 2, false},
		// erreur : pas un parametre
		{"erreur : pas un parametre", []rune{'A', 'r', 'b', 'r', 'e'}, "", 0, true},
	}
	for _, c := range cas {
		t.Run(c.nom, func(t *testing.T) {
			gotString, gotNumber, err := internal.Parameters(c.entree)
			if (err != nil) != c.veutErreur {
				t.Fatalf("erreur = %v, erreur attendue : %v", err, c.veutErreur)
			}
			if !reflect.DeepEqual(gotString, c.veutString) {
				t.Fatalf("mot = %v, mot attendu : %v", gotString, c.veutString)
			}
			if gotNumber != c.veutNumber {
				t.Fatalf("nombre = %v, nombre attendu : %v", gotNumber, c.veutNumber)
			}
		})
	}
}

func TestRemoveParameters(t *testing.T) {
	cas := []struct {
		nom        string   // la famille du cas : normal, limite, erreur, piège
		entree     [][]rune // l'entrée
		veut      [][]rune // la sortie attendue
	}{
		{"normal : text", [][]rune{{'C','o','m','m','e','n','t'}, {'(','u','p',')'}, {'ç','a'}, {'v','a'}, {'?'}}, [][]rune{{'C','o','m','m','e','n','t'}, {'ç','a'}, {'v','a'}, {'?'}}},
		{"limite : aucun paramètres", [][]rune{{'C','o','m','m','e','n','t'}, {'ç','a'}, {'v','a'}, {'?'}}, [][]rune{{'C','o','m','m','e','n','t'}, {'ç','a'}, {'v','a'}, {'?'}}},
	}
	for _, c := range cas {
		t.Run(c.nom, func(t *testing.T) {
			got := internal.RemoveParameters(c.entree)
			if !reflect.DeepEqual(got, c.veut) {
				t.Fatalf("résultat = %v, résultat attendu : %v", got, c.veut)
			}
		})
	}
}

func TestFormatParameters(t *testing.T) {
	cas := []struct {
		nom        string   // la famille du cas : normal, limite, erreur, piège
		entree     [][]rune // l'entrée
		veut      [][]rune // la sortie attendue

	}{
		{"normal : text", [][]rune{{'c','e','c','i'}, {'(','c','a','p',')'}, {'e','s','t'}, {'u','n'}, {'t','e','s','t'}, {'.'}}, [][]rune{{'C','e','c','i'}, {'e','s','t'}, {'u','n'}, {'t','e','s','t'}, {'.'}}},
	}
	for _, c := range cas {
		t.Run(c.nom, func(t *testing.T) {
			got, _ := internal.FormatParameters(c.entree)
			if !reflect.DeepEqual(got, c.veut) {
				t.Fatalf("résultat = %v, résultat attendu : %v", got, c.veut)
			}
		})
	}
}
