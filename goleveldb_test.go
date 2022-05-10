package db

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGoLevelDBNewDB(t *testing.T) {
	name := fmt.Sprintf("test_%x", randStr(12))
	dir := os.TempDir()
	db, err := NewDB(name, GoLevelDBBackend, dir)
	defer closeDBWithCleanupDBDir(db, dir, name)
	require.NoError(t, err)

	_, ok := db.(*GoLevelDB)
	assert.True(t, ok)
}

func TestGoLevelDBStats(t *testing.T) {
	name := fmt.Sprintf("test_%x", randStr(12))
	dir := os.TempDir()
	db, err := NewDB(name, GoLevelDBBackend, dir)
	defer closeDBWithCleanupDBDir(db, dir, name)
	require.NoError(t, err)

	assert.NotEmpty(t, db.Stats())
}

func BenchmarkGoLevelDBRangeScans1M(b *testing.B) {
	name := fmt.Sprintf("test_%x", randStr(12))
	dir := os.TempDir()
	db, err := NewDB(name, GoLevelDBBackend, dir)
	defer closeDBWithCleanupDBDir(db, dir, name)
	require.NoError(b, err)

	benchmarkRangeScans(b, db, int64(1e6))
}

func BenchmarkGoLevelDBRangeScans10M(b *testing.B) {
	name := fmt.Sprintf("test_%x", randStr(12))
	dir := os.TempDir()
	db, err := NewDB(name, GoLevelDBBackend, dir)
	defer closeDBWithCleanupDBDir(db, dir, name)
	require.NoError(b, err)

	benchmarkRangeScans(b, db, int64(10e6))
}

func BenchmarkGoLevelDBRandomReadsWrites(b *testing.B) {
	name := fmt.Sprintf("test_%x", randStr(12))
	dir := os.TempDir()
	db, err := NewGoLevelDB(name, dir)
	defer closeDBWithCleanupDBDir(db, dir, name)
	require.NoError(b, err)

	benchmarkRandomReadsWrites(b, db)
}
