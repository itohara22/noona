package main


func GetHttpTracker(trackers []string) []string{
	res := []string{}
	for _,t := range trackers {
		a := t[0];
		if a == 'h'{
			res = append(res, t)
		}
	}

	return res
}
