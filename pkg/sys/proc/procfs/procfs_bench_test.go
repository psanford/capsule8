// Copyright 2018 Capsule8, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package procfs

import "testing"

func BenchmarkNewFileSystem(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := NewFileSystem("testdata/proc"); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBootID(b *testing.B) {
	fs, err := NewFileSystem("testdata/proc")
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fs.BootID()
	}
}

func BenchmarkMaxPID(b *testing.B) {
	fs, err := NewFileSystem("testdata/proc")
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fs.MaxPID()
	}
}

func BenchmarkNumCPU(b *testing.B) {
	fs, err := NewFileSystem("testdata/proc")
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fs.NumCPU()
	}
}

func BenchmarkReadFile(b *testing.B) {
	fs, err := NewFileSystem("testdata/proc")
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err = fs.ReadFile("sys/kernel/pid_max"); err != nil {
			b.Fatal(err)
		}
	}
}
