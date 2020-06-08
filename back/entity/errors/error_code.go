package errors

import "fmt"

type PrimaryErrorCode int

func (c *PrimaryErrorCode) String() string {
	return fmt.Sprintf("%04d", c)
}

type SecondaryErrorCode int

func (c *SecondaryErrorCode) String() string {
	return fmt.Sprintf("%04d", c)
}
