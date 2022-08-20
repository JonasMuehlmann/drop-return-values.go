// Copyright © 2022 Jonas Muehlmann
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the "Software"),
// to deal in the Software without restriction, including without limitation
// the rights to use, copy, modify, merge, publish, distribute, sublicense,
// and/or sell copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included
// in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
// DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
// TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
// OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package drop

func From2To1[TIn1, TIn2 any](in1 TIn1, in2 TIn2) TIn1 {
	return in1
}

func From3To1[TIn1, TIn2, TIn3 any](in1 TIn1, in2 TIn2, in3 TIn3) TIn1 {
	return in1
}

func From4To1[TIn1, TIn2, TIn3, TIn4 any](in1 TIn1, in2 TIn2, in3 TIn3, in4 TIn4) TIn1 {
	return in1
}

func From5To1[TIn1, TIn2, TIn3, TIn4, TIn5 any](in1 TIn1, in2 TIn2, in3 TIn3, in4 TIn4, in5 TIn5) TIn1 {
	return in1
}

func From6To1[TIn1, TIn2, TIn3, TIn4, TIn5, TIn6 any](in1 TIn1, in2 TIn2, in3 TIn3, in4 TIn4, in5 TIn5, in6 TIn6) TIn1 {
	return in1
}

func From3To2[TIn1, TIn2, TIn3 any](in1 TIn1, in2 TIn2, in3 TIn3) (TIn1, TIn2) {
	return in1, in2
}

func From4To2[TIn1, TIn2, TIn3, TIn4 any](in1 TIn1, in2 TIn2, in4 TIn4) (TIn1, TIn2) {
	return in1, in2
}

func From5To2[TIn1, TIn2, TIn3, TIn4, TIn5 any](in1 TIn1, in2 TIn2) (TIn1, TIn2) {
	return in1, in2
}

func From6To2[TIn1, TIn2, TIn3, TIn4, TIn5, TIn6 any](in1 TIn1, in2 TIn2, in3 TIn3, in4 TIn4, in5 TIn5, in6 TIn6) (TIn1, TIn2) {
	return in1, in2
}

func From4To3[TIn1, TIn2, TIn3, TIn4 any](in1 TIn1, in2 TIn2, in3 TIn3, in4 TIn4) (TIn1, TIn2, TIn3) {
	return in1, in2, in3
}

func From5To3[TIn1, TIn2, TIn3, TIn4, TIn5 any](in1 TIn1, in2 TIn2, in3 TIn3, in4 TIn4, in5 TIn5) (TIn1, TIn2, TIn3) {
	return in1, in2, in3
}

func From6To3[TIn1, TIn2, TIn3, TIn4, TIn5, TIn6 any](in1 TIn1, in2 TIn2, in3 TIn3, in4 TIn4, in5 TIn5, in6 TIn6) (TIn1, TIn2, TIn3) {
	return in1, in2, in3
}

func From5To4[TIn1, TIn2, TIn3, TIn4, TIn5 any](in1 TIn1, in2 TIn2, in3 TIn3, in4 TIn4, in5 TIn5) (TIn1, TIn2, TIn3, TIn4) {
	return in1, in2, in3, in4
}

func From6To4[TIn1, TIn2, TIn3, TIn4, TIn5, TIn6 any](in1 TIn1, in2 TIn2, in3 TIn3, in4 TIn4, in5 TIn5, in6 TIn6) (TIn1, TIn2, TIn3, TIn4) {
	return in1, in2, in3, in4
}

func From6To5[TIn1, TIn2, TIn3, TIn4, TIn5, TIn6 any](in1 TIn1, in2 TIn2, in3 TIn3, in4 TIn4, in5 TIn5, in6 TIn6) (TIn1, TIn2, TIn3, TIn4, TIn5) {
	return in1, in2, in3, in4, in5
}
