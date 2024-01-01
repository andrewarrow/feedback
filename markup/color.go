package markup

import (
	"fmt"
	"math/rand"
)

var Colors = []string{"gray",
	"red",
	"yellow",
	"green",
	"blue",
	"indigo",
	"purple",
	"pink",
	"rose",
	"teal",
}

func RandomColor() string {
	randInt := rand.Intn(8) + 1
	text := "gray-900"
	if randInt > 5 {
		text = "white"
	}
	return fmt.Sprintf("bg-%s-%d00 text-%s",
		Colors[rand.Intn(len(Colors))], randInt, text)
}

/*
    bg-yellow-100 - #fffff0
    bg-yellow-200 - #fefcbf
    bg-yellow-300 - #faf089
    bg-yellow-400 - #f6e05e
    bg-yellow-500 - #ecc94b
    bg-yellow-600 - #d69e2e
    bg-yellow-700 - #b7791f
    bg-yellow-800 - #975a16
    bg-yellow-900 - #744210

   bg-red-100 - #fef2f2
   bg-red-200 - #fee2e2
   bg-red-300 - #fecaca
   bg-red-400 - #fca5a5
   bg-red-500 - #f87171
   bg-red-600 - #dc2626
   bg-red-700 - #b91c1c
   bg-red-800 - #991b1b
   bg-red-900 - #7f1d1d

    bg-teal-100 - #e6fffa
    bg-teal-200 - #b2f5ea
    bg-teal-300 - #81e6d9
    bg-teal-400 - #4fd1c5
    bg-teal-500 - #38b2ac
    bg-teal-600 - #319795
    bg-teal-700 - #2c7a7b
    bg-teal-800 - #285e61
    bg-teal-900 - #234e52

    bg-rose-100 - #fff1f2
    bg-rose-200 - #ffe4e6
    bg-rose-300 - #fecdd3
    bg-rose-400 - #fda4af
    bg-rose-500 - #fb7185
    bg-rose-600 - #f4338b
    bg-rose-700 - #d946ef
    bg-rose-800 - #c026d3
    bg-rose-900 - #a21caf

    bg-pink-100 - #fff5f7
    bg-pink-200 - #fed7e2
    bg-pink-300 - #fbb6ce
    bg-pink-400 - #f687b3
    bg-pink-500 - #ed64a6
    bg-pink-600 - #d53f8c
    bg-pink-700 - #b83280
    bg-pink-800 - #97266d
    bg-pink-900 - #702459

    bg-purple-100 - #faf5ff
    bg-purple-200 - #e9d8fd
    bg-purple-300 - #d6bcfa
    bg-purple-400 - #b794f4
    bg-purple-500 - #9f7aea
    bg-purple-600 - #805ad5
    bg-purple-700 - #6b46c1
    bg-purple-800 - #553c9a
    bg-purple-900 - #44337a

    bg-indigo-100 - #ebf4ff
    bg-indigo-200 - #c3dafe
    bg-indigo-300 - #a3bffa
    bg-indigo-400 - #7f9cf5
    bg-indigo-500 - #667eea
    bg-indigo-600 - #5a67d8
    bg-indigo-700 - #4c51bf
    bg-indigo-800 - #434190
    bg-indigo-900 - #3c366b

    bg-blue-100 - #ebf8ff
    bg-blue-200 - #bee3f8
    bg-blue-300 - #90cdf4
    bg-blue-400 - #63b3ed
    bg-blue-500 - #4299e1
    bg-blue-600 - #3182ce
    bg-blue-700 - #2b6cb0
    bg-blue-800 - #2c5282
    bg-blue-900 - #2a4365

    bg-green-100 - #f0fff4
    bg-green-200 - #c6f6d5
    bg-green-300 - #9ae6b4
    bg-green-400 - #68d391
    bg-green-500 - #48bb78
    bg-green-600 - #38a169
    bg-green-700 - #2f855a
    bg-green-800 - #276749
    bg-green-900 - #22543d

    bg-gray-100 - #f7fafc
    bg-gray-200 - #edf2f7
    bg-gray-300 - #e2e8f0
    bg-gray-400 - #cbd5e0
    bg-gray-500 - #a0aec0
    bg-gray-600 - #718096
    bg-gray-700 - #4a5568
    bg-gray-800 - #2d3748
    bg-gray-900 - #1a202c
*/
