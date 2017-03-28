package util

import "fmt"

func CheckError(err error,info string) (res bool) {
	if (err != nil) {
		fmt.Println("info : " + info + " | " + "error : " + err.Error())
		return false;
		//panic(err);
	}
	return true;
}
