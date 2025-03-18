package disk

type Disk interface {
	Compress()
	Checksum() int
}
