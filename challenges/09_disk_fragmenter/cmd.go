package _9_disk_fragmenter

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func A() error {
	head, tail := parseFragments(input)
	sentinel := &Fragment{
		id:   -2,
		pos:  0,
		len:  0,
		prev: nil,
		next: head,
	}
	head.prev = sentinel

	head, tail = advanceToSpace(head), retreatToFile(tail)

	// loop until we run out of spaces or files
	for ; head != nil && tail != nil; head, tail = advanceToSpace(head), retreatToFile(tail) {
		switch {
		case head.len > tail.len:
			// create a new empty space for the remainder
			space := &Fragment{
				id:   -1,
				pos:  head.pos + tail.len,
				len:  head.len - tail.len,
				prev: head,
				next: head.next,
			}
			// shrink and re-link the head
			head.len = tail.len
			head.next = space
			// head length is now equal to tail length
			fallthrough
		case head.len == tail.len:
			// overwrite head
			head.id = tail.id
			// unlink tail
			tail = tail.prev
			tail.next = nil
		case head.len < tail.len:
			// overwrite head
			head.id = tail.id
			// shrink tail
			tail.len -= head.len
		}
	}

	var checksum int

	for peek := sentinel.next; peek != nil; peek = peek.next {
		for i := peek.pos; i < peek.pos+peek.len; i++ {
			checksum += i * peek.id
		}
	}

	fmt.Println(checksum)

	return nil
}

func B() error {
	head, tail := parseFragments(input)
	sentinel := &Fragment{
		id:   -2,
		pos:  0,
		len:  0,
		prev: nil,
		next: head,
	}
	head.prev = sentinel

	head, tail = advanceToSpace(head), retreatToFile(tail)

	// loop until we run out of files
	for ; tail != nil; tail = prevFile(tail) {
		// try to find a space large enough for the tail in front of it
		searchHead := advanceToSpace(head)
		for ; searchHead != nil && searchHead.pos < tail.pos && searchHead.len < tail.len; searchHead = nextSpace(searchHead) {
		}

		// if there's no space big enough we can't move the file
		if searchHead == nil || searchHead.pos >= tail.pos {
			continue
		}

		switch {
		case searchHead.len > tail.len:
			// create a new empty space for the remainder
			space := &Fragment{
				id:   -1,
				pos:  searchHead.pos + tail.len,
				len:  searchHead.len - tail.len,
				prev: searchHead,
				next: searchHead.next,
			}
			// shrink and re-link the searchHead
			searchHead.len = tail.len
			searchHead.next = space
			// searchHead length is now equal to tail length
			fallthrough
		case searchHead.len == tail.len:
			// overwrite searchHead
			searchHead.id = tail.id
			// erase tail
			tail.id = -1
		}
	}

	var checksum int

	for peek := sentinel.next; peek != nil; peek = peek.next {
		if peek.id < 0 {
			continue
		}

		for i := peek.pos; i < peek.pos+peek.len; i++ {
			checksum += i * peek.id
		}
	}

	fmt.Println(checksum)

	return nil
}

type Fragment struct {
	id, pos, len int
	prev, next   *Fragment
}

func advanceToSpace(head *Fragment) *Fragment {
	for ; head != nil && head.id >= 0; head = head.next {
	}

	return head
}

func nextSpace(head *Fragment) *Fragment {
	return advanceToSpace(head.next)
}

func retreatToFile(tail *Fragment) *Fragment {
	for ; tail != nil && tail.id < 0; tail = tail.prev {
	}

	return tail
}

func prevFile(tail *Fragment) *Fragment {
	return retreatToFile(tail.prev)
}

func parseFragments(fileMap string) (head, tail *Fragment) {
	// create a sentinel fragment from which we will grow our list
	sentinel := &Fragment{
		id:   -2,
		pos:  0,
		len:  0,
		prev: nil,
		next: nil,
	}

	last := sentinel
	var pos int

	for i, char := range []rune(strings.TrimSpace(fileMap)) {
		// skip zero-length fragments
		length := int(char - '0')
		if length == 0 {
			continue
		}

		// only every second entry is a file
		id := i / 2
		// every other entry is free space (id -1)
		if i%2 == 1 {
			id = -1
		}

		// make a new fragment linked to the last
		current := &Fragment{
			id:   id,
			pos:  pos,
			len:  length,
			prev: last,
			next: nil,
		}

		// update the position of the next fragment
		pos += length

		// double-link the last fragment
		if last != nil {
			last.next = current
		}

		// move forwards
		last = current
	}

	head = sentinel.next
	head.prev = nil
	tail = last
	return
}
