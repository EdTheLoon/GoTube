package server

import (
	"os/exec"
)

func getDetails(url string) ([]byte, error) {
	out, err := exec.Command("youtube-dl",
		"--print-json",
		url).Output()

	if err != nil {
		return out, err
	}
	return out, nil
}

func getFormats(url string) {

}

func downloadVideo(url string, format string, s *Server, filename string) {
	out, err := exec.Command("youtube-dl",
		"-f "+format,
		"--audio-quality", "0",
		"--metadata-from-title", "'%(artist)s - %(title)s'",
		// "--embed-thumbnail",
		"-o", s.GetDownloadsDir()+"/"+filename,
		"--add-metadata",
		url).CombinedOutput()

	if err != nil {
		s.Log("Error: " + err.Error())
		s.Log("youtube-dl: " + string(out))
	} else {
		s.Log(string(out))
		s.Log("Successfully downloaded: " + s.GetDownloadsDir() + "/" + filename)
	}
}

func downloadAudio(url string, format string, s *Server, filename string) {
	out, err := exec.Command("youtube-dl",
		"-f bestaudio/best",
		"-x",
		"--audio-format", format,
		"--audio-quality", "0",
		"--metadata-from-title", "'%(artist)s - %(title)s'",
		// "--embed-thumbnail",
		"-o", s.GetDownloadsDir()+"/"+filename,
		"--add-metadata",
		url).CombinedOutput()

	if err != nil {
		s.Log("Error: " + err.Error())
		s.Log("youtube-dl: " + string(out))
	} else {
		s.Log(string(out))
		s.Log("Successfully downloaded: " + s.GetDownloadsDir() + "/" + filename)
	}
}
