package addons

import "fmt"

func IsBlackListed(acnt string, dest string) (bool, error) {

	qryStr := fmt.Sprintf("SELECT ID FROM blackList WHERE msisdn = '%s' OR msisdn = '%s' ", acnt, dest)
	qOut, err := runQuery(qryStr)
	if err != nil {
		return true, err
	}
	return len(qOut) > 0, err
}
