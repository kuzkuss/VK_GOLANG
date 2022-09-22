package uniq

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

	assert.Equal(t, expectedRes, receivedRes)
	assert.Equal(t, nil, err)
}

func TestUniqCOption(t *testing.T) {
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
		"3 I love music.",
		"1 ",
		"2 I love music of Kartik.",
		"1 Thanks.",
		"2 I love music of Kartik.",
	}

	receivedRes, err := Uniq(strs, opts)

	assert.Equal(t, expectedRes, receivedRes)
	assert.Equal(t, nil, err)
}

