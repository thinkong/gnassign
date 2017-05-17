package gnassign

// this is based off filepath/path/match.go since i only needed the star

func scanChunk(pattern string) (star bool, chunk, rest string) {
	for len(pattern) > 0 && pattern[0] == '*' {
		pattern = pattern[1:]
		star = true
	}
	var i int
	Scan:
	for i = 0; i < len(pattern); i++ {
		if pattern[i] == '*' {
			break Scan
		}
	}
	return star, pattern[0:i], pattern[i:]
}

func matchChunk(chunk, s string) (rest string, ok bool, err error) {
	for len(chunk) > 0 {
		if len(s) == 0 {
			return
		}
		if chunk[0] != s[0] {
			return
		}
		s = s[1:]
		chunk = chunk[1:]
	}
	return s, true, nil
}

func Match(pattern, name string) (matched bool, err error) {
	Pattern:
	for len(pattern) > 0 {
		var star bool
		var chunk string
		star, chunk, pattern = scanChunk(pattern)
		t, ok, err := matchChunk(chunk, name)
		if ok && (len(t) == 0 || len(pattern) > 0) {
			name = t
			continue
		}
		if err != nil {
			return false, err
		}
		if star {
			for i := 0; i < len(name); i++ {
				t, ok, err := matchChunk(chunk, name[i + 1:])
				if ok {
					if len(pattern) == 0 && len(t) > 0 {
						continue
					}
					name = t
					continue Pattern
				}
				if err != nil {
					return false, err
				}
			}
		}
		return false, nil
	}
	return len(name) == 0, nil
}
