package main


//sync.Once
var icons maps[string]image.Image
func loadIcons(){
	icons = map[string]image.Image{
		"left":loadIcons("left.png")
		"right":loadIcons("right.png")
		"up":loadIcon("up.png")
		"down":loadIcon("down.png")
	}
}
//Icon被多个goroutine调用时不是并发安全的
func Icon(name string)image.Image{
	if name == nil{
		loadIcons()
	}
	return icons[name]
}
func main() {
	
}