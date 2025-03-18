package disk

const space = -1

func NewAdvancedDisk(code []int) Disk {
	blocks := make(map[int]block)
	disc := AdvancedDisk{blocks: blocks}

	bytes := buildAdvancedDisc(code)
	name := 0 // Use ints as map keys
	length := 0
	index := 0
	for k, val := range bytes {
		if val == space {
			continue
		}
		if val == name {
			length++
			continue
		}
		// Store...
		disc.AddBlock(index, name, length)
		name = val
		length = 1
		index = k
	}
	disc.AddBlock(index, name, length) // finish it off
	disc.maxKey = name

	return &disc
}

type AdvancedDisk struct {
	blocks map[int]block // index is the position
	maxKey int
}

func (d *AdvancedDisk) Compress() {
	maximum := d.maxKey
	for i := maximum; i > 0; i-- {
		d.moveLeft(i)
	}
}

func (d *AdvancedDisk) Checksum() int {
	checksum := 0
	for k, block := range d.blocks {
		for i := range block.length {
			sub := (k + i) * block.name
			checksum += sub
		}
	}

	return checksum
}

func buildAdvancedDisc(code []int) []int {
	disk := []int{}
	id := -1
	isFile := false
	next := 0
	for _, val := range code {
		isFile = !isFile
		if isFile {
			id++
			next = id
		} else {
			next = space
		}
		for range val {
			disk = append(disk, next)
		}
	}

	return disk
}

type block struct {
	name   int
	length int
}

func (d *AdvancedDisk) AddBlock(index, name, length int) {
	block := block{name, length}
	d.blocks[index] = block
}

func (d *AdvancedDisk) moveLeft(name int) {
	index := d.find(name)
	block := d.blocks[index]
	spaceIndex := d.findSpace(block.length, index)
	if spaceIndex < 0 {
		return // Do nothing
	}
	// do the move
	delete(d.blocks, index)
	d.AddBlock(spaceIndex, block.name, block.length)
}

func (d *AdvancedDisk) find(name int) int {
	for k, block := range d.blocks {
		if name == block.name {
			return k
		}
	}

	return -1
}

//revive:disable:cognitive-complexity

func (d *AdvancedDisk) findSpace(size, maxIndex int) int {
	inSpace := false
	spaceIndex := 0
	spaceLength := 0
	i := 0 // go through them all
	for {
		if i >= maxIndex {
			return -1
		}
		block, ok := d.blocks[i]
		if ok {
			inSpace = false
			spaceLength = 0
			i += block.length // jump ahead
			continue
		}
		if !ok { // it's a space
			if !inSpace {
				spaceIndex = i // grab the index
			}
			spaceLength++
			if spaceLength >= size {
				return spaceIndex
			}
			inSpace = true
		}
		i++
	}
}

//revive:enable:cognitive-complexity

// for debug
// func (d *Disk) Dump() {
// 	keys := make([]int, 0, len(d.Blocks))
// 	for k := range d.Blocks {
// 		keys = append(keys, k)
// 	}
// 	sort.Ints(keys)
// 	for _, key := range keys {
// 		fmt.Println(key, d.Blocks[key])
// 	}
// }
