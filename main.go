package main
func main() {

	// create list of image file from path
	// run darknet from darknet path
	// parse results
	imgvldr := New("/Users/henryxu/Documents/darknet")
	imgvldr.ValidateImages("/Users/henryxu/Documents/ImageSet/corvette", ".jpg")
}
