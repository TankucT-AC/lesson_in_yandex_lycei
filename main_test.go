package main

import "testing"

type Test struct{
    in []byte
    ans int
    err error
}

var tests = []Test{
    {[]byte{1, 12, 27}, 3, nil},
    {[]byte{255}, 0, ErrInvalidUTF8},
    {[]byte{127, 125, 11}, 3, nil},
    {[]byte{0}, 1, nil},
}

func TestGetUTFLength(t *testing.T) {
    for i, test := range tests {
        got, err := GetUTFLength(test.in)
        if got != test.ans || err != test.err {
            t.Errorf("#%d: Ответ не совпал", i)
        }
    }
}