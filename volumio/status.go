package volumio

import "strings"

// Returns the currently playing MPC track
func GetCurrentTrack() (string, error) {
	track, err := Mpc("current")
	if err != nil {
		return "", err
	}

	return string(track), nil
}

// Returns current volume percentage
func GetCurrentVolume() (string, error) {
	out, err := Mpc("volume")
	if err != nil {
		return "", err
	}

	volume := strings.ReplaceAll(string(out), "volume:", "")

	return volume, nil
}
