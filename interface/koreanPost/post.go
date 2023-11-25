package koreanPost

import "fmt"

type PostSender struct {
	// ...
}

func (f *PostSender) Send(parcel string) {
	fmt.Printf("Korean Post sends %s parcel\n", parcel)
}
