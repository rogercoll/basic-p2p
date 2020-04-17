package messages

import (
	"fmt"
	"time"
    "strconv"
)

//The idea is to use interfaces
func Parse(v *Version) (VersionReadble, error) {
	var tmp VersionReadble
	tmp.Version = fmt.Sprint(v.Version)
	i, err := strconv.ParseInt(fmt.Sprint(v.Timestamp), 10, 64)
    if err != nil {
        return VersionReadble{}, err
    }
	tmp.Timestamp = time.Unix(i, 0)
	return tmp, nil
}