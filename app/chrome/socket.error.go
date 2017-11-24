/*
Package chrome provides an interface to a headless Chrome instance.
*/
package chrome

/*
SocketError represents an error
*/
type SocketError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
