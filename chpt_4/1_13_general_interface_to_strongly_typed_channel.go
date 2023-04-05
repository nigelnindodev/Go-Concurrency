package main

func main() {
	toString := func(done <-chan interface{}, valueStream <-chan interface{}) <-chan string {
		stringStream := make(chan string)
		go func() {
			defer close(stringStream)
			for v := range valueStream {
				select {
				case <-done:
					return
				case stringStream <- v.(string):
				}
			}
		}()
		return stringStream
	}

	_ = toString // suppress not used warning [https://stackoverflow.com/a/21744129]
	/**
	- Will not compile since not completed, this is justa trivial example of how this could potentially work. See if you
		can complete it if you have been following.
	*/
}
