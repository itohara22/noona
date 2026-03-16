package main


func GetUdpTracker(trackers []string) []string{
	res := []string{}
	for _,t := range trackers {
		a := t[0];
		if a == 'u'{
			res = append(res, t)
		}
	}

	return res
}
