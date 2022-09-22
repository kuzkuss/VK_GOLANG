package uniq

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUniqDefault(t *testing.T) {
	opts := Options {
		count: false,
		double: false,
		unique: false,
		fields: 0,
		chars: 0,
		insensitive: false,
	}

	strs := []string {
		"I love music.",
		"I love music.",
		"I love music.",
		"",
		"I love music of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
		"I love music of Kartik.",
		"I love music of Kartik.",
	}

	expectedRes := []string {
		"I love music.",
		"",
		"I love music of Kartik.",
		"Thanks.",
		"I love music of Kartik.",
	}

	receivedRes, err := Uniq(strs, opts)

	require.Equal(t, expectedRes, receivedRes)
	require.Equal(t, nil, err)
}

func TestUniqCOption(t *testing.T) {
	opts := Options {
		count: true,
		double: false,
		unique: false,
		fields: 0,
		chars: 0,
		insensitive: false,
	}

	strs := []string {
		"I love music.",
		"I love music.",
		"I love music.",
		"",
		"I love music of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
		"I love music of Kartik.",
		"I love music of Kartik.",
	}

	expectedRes := []string {
		"3 I love music.",
		"1 ",
		"2 I love music of Kartik.",
		"1 Thanks.",
		"2 I love music of Kartik.",
	}

	receivedRes, err := Uniq(strs, opts)

	require.Equal(t, expectedRes, receivedRes)
	require.Equal(t, nil, err)
}

func TestUniqDOption(t *testing.T) {
	opts := Options {
		count: false,
		double: true,
		unique: false,
		fields: 0,
		chars: 0,
		insensitive: false,
	}

	strs := []string {
		"I love music.",
		"I love music.",
		"I love music.",
		"",
		"I love music of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
		"I love music of Kartik.",
		"I love music of Kartik.",
	}

	expectedRes := []string {
		"I love music.",
		"I love music of Kartik.",
		"I love music of Kartik.",
	}

	receivedRes, err := Uniq(strs, opts)

	require.Equal(t, expectedRes, receivedRes)
	require.Equal(t, nil, err)
}

func TestUniqUOption(t *testing.T) {
	opts := Options {
		count: false,
		double: false,
		unique: true,
		fields: 0,
		chars: 0,
		insensitive: false,
	}

	strs := []string {
		"I love music.",
		"I love music.",
		"I love music.",
		"",
		"I love music of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
		"I love music of Kartik.",
		"I love music of Kartik.",
	}

	expectedRes := []string {
		"",
		"Thanks.",
	}

	receivedRes, err := Uniq(strs, opts)

	require.Equal(t, expectedRes, receivedRes)
	require.Equal(t, nil, err)
}

func TestUniqIOption(t *testing.T) {
	opts := Options {
		count: false,
		double: false,
		unique: false,
		fields: 0,
		chars: 0,
		insensitive: true,
	}

	strs := []string {
		"I LOVE MUSIC.",
		"I love music.",
		"I LoVe MuSiC.",
		"",
		"I love MuSIC of Kartik.",
		"I love music of kartik.",
		"Thanks.",
		"I love music of kartik.",
		"I love MuSIC of Kartik.",
	}

	expectedRes := []string {
		"I LOVE MUSIC.",
		"",
		"I love MuSIC of Kartik.",
		"Thanks.",
		"I love music of kartik.",
	}

	receivedRes, err := Uniq(strs, opts)

	require.Equal(t, expectedRes, receivedRes)
	require.Equal(t, nil, err)
}

func TestUniqFOption(t *testing.T) {
	opts := Options {
		count: false,
		double: false,
		unique: false,
		fields: 1,
		chars: 0,
		insensitive: false,
	}

	strs := []string {
		"We love music.",
		"I love music.",
		"They love music.",
		"",
		"I love music of Kartik.",
		"We love music of Kartik.",
		"Thanks.",
	}

	expectedRes := []string {
		"We love music.",
		"",
		"I love music of Kartik.",
		"Thanks.",
	}

	receivedRes, err := Uniq(strs, opts)

	require.Equal(t, expectedRes, receivedRes)
	require.Equal(t, nil, err)
}

func TestUniqSOption(t *testing.T) {
	opts := Options {
		count: false,
		double: false,
		unique: false,
		fields: 0,
		chars: 1,
		insensitive: false,
	}

	strs := []string {
		"I love music.",
		"A love music.",
		"C love music.",
		"",
		"I love music of Kartik.",
		"We love music of Kartik.",
		"Thanks.",
	}

	expectedRes := []string {
		"I love music.",
		"",
		"I love music of Kartik.",
		"We love music of Kartik.",
		"Thanks.",
	}

	receivedRes, err := Uniq(strs, opts)

	require.NoError(t, err)
	require.Equal(t, expectedRes, receivedRes)
}

