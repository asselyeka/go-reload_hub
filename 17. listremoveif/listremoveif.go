package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	if l.Head == nil {
		return
	}
	current := l.Head
	var prev *NodeL = nil
	for current != nil {
		if current.Data == data_ref {
			if prev == nil {
				l.Head = current.Next
				current = l.Head
				prev = nil
			} else {
				prev.Next = current.Next
				current = current.Next
			}
		} else {
			prev = current
			current = current.Next
		}
	}
}
