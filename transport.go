package main

type Transport struct {
	locations       []string
	tFare, tip      int
	pickUp, dropOff string
}

func (this *Transport) setPickUpLocation(pickUp string) {

	this.pickUp = pickUp
}

func (this *Transport) setDropOffLocation(dropOff string) {

	this.dropOff = dropOff
}

func (this *Transport) setTip(tip int) {

	this.tip = tip

}
func (this Transport) indexOf(el string) int {

	for i, v := range this.locations {
		if v == el {
			return i
		}
	}
	return -1
}

func (this *Transport) setTFare() {

	pickUpIndex := this.indexOf(this.pickUp)
	dropOffIndex := this.indexOf(this.dropOff)

	if dropOffIndex > pickUpIndex {
		this.tFare = this.tFare * (dropOffIndex - pickUpIndex)
	} else {
		this.tFare = this.tFare * (pickUpIndex - dropOffIndex)
	}

}
