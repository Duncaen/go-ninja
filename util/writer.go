// Copyright 2019 Duncan Overbruck
//
// Permission to use, copy, modify, and/or distribute this software
// for any purpose with or without fee is hereby granted, provided
// that the above copyright notice and this permission notice appear
// in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL
// WARRANTIES WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE
// AUTHOR BE LIABLE FOR ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL
// DAMAGES OR ANY DAMAGES WHATSOEVER RESULTING FROM LOSS OF USE, DATA
// OR PROFITS, WHETHER IN AN ACTION OF CONTRACT, NEGLIGENCE OR OTHER
// TORTIOUS ACTION, ARISING OUT OF OR IN CONNECTION WITH THE USE OR
// PERFORMANCE OF THIS SOFTWARE.
//
// SPDX-License-Identifier: ISC

package util

import (
	"io"
)

// PrefixWriter implements the io.Writer interface and adds a prefix to
// every written line
type PrefixWriter struct {
	wr io.Writer
	err error
	pre []byte
	nl bool
}

// NewPrefixWriter creates and initializes a new PrefixWriter
func NewPrefixWriter(w io.Writer, prefix []byte) *PrefixWriter {
	return &PrefixWriter{w, nil, prefix, true}
}

// Write adds a prefix to every line in buf and writes it to
// the underlying io.Writer
func (w *PrefixWriter) Write(buf []byte) (int, error) {
	var n int
	for _, c := range buf {
		if w.nl {
			nn, err := w.wr.Write(w.pre)
			if err != nil {
				return n, err
			}
			n += nn
		}
		if _, err := w.wr.Write([]byte{c}); err != nil {
			return n, err
		}
		n++
		w.nl = c == '\n'
	}
	return n, nil
}

// SeqWriter counts and checks a sequence of writes for errors.
//
// If a write to the underlying io.Writer returns an error,
// SeqWriter records it and ignores every further Write.
type SeqWriter struct {
	wr io.Writer
	err error
	n int64
}

// NewSeqWriter creates and initializes a new SeqWriter
func NewSeqWriter(w io.Writer) *SeqWriter {
	return &SeqWriter{w, nil, 0}
}

// Write writes buf to the underlying Writer if there was no previous error
// and and records the number of written bytes and a possible error
func (s *SeqWriter) Write(buf []byte) (int, error) {
	if s.err != nil {
		return 0, s.err
	}
	n, err := s.wr.Write(buf)
	s.n += int64(n)
	s.err = err
	return n, err
}

// Done resets and returns the number of written bytes and/or an error
func (s *SeqWriter) Done() (int64, error) {
	n, err := s.n, s.err
	s.n = 0
	s.err = nil
	return n, err
}
