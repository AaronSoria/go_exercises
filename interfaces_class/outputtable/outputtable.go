package outputtable

import "taquion.com/interfaces/saver"

type Outputtable interface {
	saver.Saver
	Display()
}
