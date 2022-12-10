package uniq

import "flag"

type Options struct {
	count       bool
	double      bool
	unique      bool
	fields      int
	chars       int
	insensitive bool
}

func Init(opts *Options) {
	const (
		usageCount       = "the number of repetitions of the line"
		usageDouble      = "print only repeating rows"
		usageUnique      = "print only unique rows in input data"
		usageNumFields   = "ignore the first num_fields of fields in a row"
		usageNumChars    = "ignore the first num_chars characters in the string"
		usageInsensitive = "ignore the case of letters"
	)
	flag.BoolVar(&opts.count, "c", false, usageCount)
	flag.BoolVar(&opts.double, "d", false, usageDouble)
	flag.BoolVar(&opts.unique, "u", false, usageUnique)
	flag.IntVar(&opts.fields, "f", 0, usageNumFields)
	flag.IntVar(&opts.chars, "s", 0, usageNumChars)
	flag.BoolVar(&opts.insensitive, "i", false, usageInsensitive)
}
