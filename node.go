package xco

import "encoding/xml"

// Node represents an arbitrary XML Node.
type Node struct {
	XMLName    xml.Name
	Attributes []xml.Attr `xml:",any,attr"`
	Data       string     `xml:",cdata"`
	Nodes      []Node     `xml:",any"`
}
