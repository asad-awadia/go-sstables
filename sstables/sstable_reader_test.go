package sstables

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/thomasjungblut/go-sstables/skiplist"
	"testing"
)

func TestSimpleHappyPathRead(t *testing.T) {
	reader, err := NewSSTableReader(
		ReadBasePath("test_files/SimpleWriteHappyPathSSTable"),
		ReadWithKeyComparator(skiplist.BytesComparator))
	assert.Nil(t, err)
	defer reader.Close()

	// 0 because there was no metadata file
	assert.Equal(t, 0, int(reader.metaData.NumRecords))
	assert.Equal(t, 0, len(reader.metaData.MinKey))
	assert.Equal(t, 0, len(reader.metaData.MaxKey))
	skipListMap := TEST_ONLY_NewSkipListMapWithElements([]int{1, 2, 3, 4, 5, 6, 7,})
	assertContentMatchesSkipList(t, reader, skipListMap)
}

func TestSimpleHappyPathBloomRead(t *testing.T) {
	reader, err := NewSSTableReader(
		ReadBasePath("test_files/SimpleWriteHappyPathSSTableWithBloom"),
		ReadWithKeyComparator(skiplist.BytesComparator))
	assert.Nil(t, err)
	defer reader.Close()

	// 0 because there was no metadata file
	assert.Equal(t, 0, int(reader.metaData.NumRecords))
	assert.Equal(t, 0, len(reader.metaData.MinKey))
	assert.Equal(t, 0, len(reader.metaData.MaxKey))
	skipListMap := TEST_ONLY_NewSkipListMapWithElements([]int{1, 2, 3, 4, 5, 6, 7,})
	assertContentMatchesSkipList(t, reader, skipListMap)
}

func TestSimpleHappyPathWithMetaData(t *testing.T) {
	reader, err := NewSSTableReader(
		ReadBasePath("test_files/SimpleWriteHappyPathSSTableWithMetaData"),
		ReadWithKeyComparator(skiplist.BytesComparator))
	assert.Nil(t, err)
	defer reader.Close()

	assert.Equal(t, 7, int(reader.metaData.NumRecords))
	assert.Equal(t, []byte{0, 0, 0, 1}, reader.metaData.MinKey)
	assert.Equal(t, []byte{0, 0, 0, 7}, reader.metaData.MaxKey)
	skipListMap := TEST_ONLY_NewSkipListMapWithElements([]int{1, 2, 3, 4, 5, 6, 7,})
	assertContentMatchesSkipList(t, reader, skipListMap)
}

func TestNegativeContainsHappyPath(t *testing.T) {
	reader, err := NewSSTableReader(
		ReadBasePath("test_files/SimpleWriteHappyPathSSTable"),
		ReadWithKeyComparator(skiplist.BytesComparator))
	assert.Nil(t, err)
	defer reader.Close()

	assertNegativeContains(t, reader)
}

func TestNegativeContainsHappyPathBloom(t *testing.T) {
	reader, err := NewSSTableReader(
		ReadBasePath("test_files/SimpleWriteHappyPathSSTableWithBloom"),
		ReadWithKeyComparator(skiplist.BytesComparator))
	assert.Nil(t, err)
	defer reader.Close()

	assertNegativeContains(t, reader)
}

func TestScanStartingAt(t *testing.T) {
	reader, err := NewSSTableReader(
		ReadBasePath("test_files/SimpleWriteHappyPathSSTableWithMetaData"),
		ReadWithKeyComparator(skiplist.BytesComparator))
	assert.Nil(t, err)
	defer reader.Close()

	expected := []int{1, 2, 3, 4, 5, 6, 7,}
	// whole sequence when out of bounds to the left
	it, err := reader.ScanStartingAt(intToByteSlice(0))
	assert.Nil(t, err)
	assertIteratorMatchesSlice(t, it, expected)

	// staggered test
	for i, start := range expected {
		sliced := expected[i:]
		it, err := reader.ScanStartingAt(intToByteSlice(start))
		assert.Nil(t, err)
		assertIteratorMatchesSlice(t, it, sliced)
	}

	// test out of range iteration, which should yield an empty iterator
	it, err = reader.ScanStartingAt(intToByteSlice(10))
	assert.Nil(t, err)
	k, v, err := it.Next()
	assert.Nil(t, k)
	assert.Nil(t, v)
	assert.Equal(t, Done, err)
}

func TestScanRange(t *testing.T) {
	reader, err := NewSSTableReader(
		ReadBasePath("test_files/SimpleWriteHappyPathSSTableWithMetaData"),
		ReadWithKeyComparator(skiplist.BytesComparator))
	assert.Nil(t, err)
	defer reader.Close()

	expected := []int{1, 2, 3, 4, 5, 6, 7,}
	// whole sequence when out of bounds to the left and right
	it, err := reader.ScanRange(intToByteSlice(0), intToByteSlice(10))
	assert.Nil(t, err)
	assertIteratorMatchesSlice(t, it, expected)

	// whole sequence when in bounds for inclusiveness
	it, err = reader.ScanRange(intToByteSlice(1), intToByteSlice(7))
	assert.Nil(t, err)
	assertIteratorMatchesSlice(t, it, expected)

	// only 4 when requesting between 4 and 4
	it, err = reader.ScanRange(intToByteSlice(4), intToByteSlice(4))
	assert.Nil(t, err)
	assertIteratorMatchesSlice(t, it, []int{4})

	// error when higher key and lower key are inconsistent
	_, err = reader.ScanRange(intToByteSlice(1), intToByteSlice(0))
	assert.NotNil(t, err)

	// staggered test with end outside of range
	for i, start := range expected {
		sliced := expected[i:]
		it, err := reader.ScanRange(intToByteSlice(start), intToByteSlice(10))
		assert.Nil(t, err)
		assertIteratorMatchesSlice(t, it, sliced)
	}

	// staggered test with end crossing to the left
	for i, start := range expected {
		it, err := reader.ScanRange(intToByteSlice(start), intToByteSlice(expected[len(expected)-i-1]))
		if i <= (len(expected) / 2) {
			assert.Nil(t, err)
			sliced := expected[i : len(expected)-i]
			assertIteratorMatchesSlice(t, it, sliced)
		} else {
			assert.NotNil(t, err)
		}

	}

	// test out of range iteration, which should yield an empty iterator
	it, err = reader.ScanRange(intToByteSlice(10), intToByteSlice(100))
	assert.Nil(t, err)
	k, v, err := it.Next()
	assert.Nil(t, k)
	assert.Nil(t, v)
	assert.Equal(t, Done, err)
}

func assertNegativeContains(t *testing.T, reader *SSTableReader) {
	assert.False(t, reader.Contains([]byte{}))
	assert.False(t, reader.Contains([]byte{1}))
	assert.False(t, reader.Contains([]byte{1, 2, 3}))
	_, err := reader.Get([]byte{})
	assert.Equal(t, errors.New("key was not found"), err)
	_, err = reader.Get([]byte{1})
	assert.Equal(t, errors.New("key was not found"), err)
	_, err = reader.Get([]byte{1, 2, 3})
	assert.Equal(t, errors.New("key was not found"), err)
}
