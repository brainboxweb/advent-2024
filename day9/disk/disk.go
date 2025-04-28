// Package disk is the disk drive
package disk

// Disk represents a disk drive
type Disk interface {
	Compress()
	Checksum() int
}
