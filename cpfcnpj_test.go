package brdoc

import (
	. "testing"
)

func TestIsCPF(t *T) {
	t.Run("Invalid CPF format", func(t *T) {
		v := "3467875434578764345789654"
		r := IsCPF(v)
		assert(t, v, r, false)

		v = ""
		r = IsCPF(v)
		assert(t, v, r, false)

		v = "#$%¨&*(ABCDEF"
		r = IsCPF(v)
		assert(t, v, r, false)
	})
	t.Run("Invalid digits in CPF", func(t *T) {
		v := "000.000.000-11"
		r := IsCPF(v)
		assert(t, v, r, false)

		v = "111.111.111-22"
		r = IsCPF(v)
		assert(t, v, r, false)
	})
	t.Run("Valid CPF", func(t *T) {
		v := "000.000.000-00"
		r := IsCPF(v)
		assert(t, v, r, true)

		v = "111.111.111-11"
		r = IsCPF(v)
		assert(t, v, r, true)
	})
}

func TestIsCNPJ(t *T) {
	t.Run("Invalid CNPJ format", func(t *T) {
		v := "3467875434578764345789654"
		r := IsCNPJ(v)
		assert(t, v, r, false)

		v = ""
		r = IsCNPJ(v)
		assert(t, v, r, false)

		v = "#$%¨&*(ABCDEF"
		r = IsCNPJ(v)
		assert(t, v, r, false)
	})
	t.Run("Invalid digits in CNPJ", func(t *T) {
		v := "00.000.000/0000-11"
		r := IsCNPJ(v)
		assert(t, v, r, false)

		v = "11.111.111/1111-00"
		r = IsCNPJ(v)
		assert(t, v, r, false)
	})
	t.Run("Valid CNPJ", func(t *T) {
		v := "00.000.000/0000-00"
		r := IsCNPJ(v)
		assert(t, v, r, true)

		v = "11.111.111/1111-80"
		r = IsCNPJ(v)
		assert(t, v, r, true)
	})
}

func assert(t *T, value string, result bool, expected bool) {
	if result == expected {
		t.Logf("The result of %s should be \"%v\": ja!", value, expected)
	} else {
		t.Errorf("The result of %s should be \"%v\": nein!", value, expected)
	}
}
