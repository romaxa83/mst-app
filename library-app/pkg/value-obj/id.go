package value_obj

import (
	"fmt"
	"strconv"
)

type ID string

func (i ID) AsString() string {
	return fmt.Sprintf("%s", i)
}

func (i ID) AsInt() (int, error) {
	id, err := strconv.Atoi(i.AsString())
	if err != nil {
		return 0, err
	}

	return id, nil
}
