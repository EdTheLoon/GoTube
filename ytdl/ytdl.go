package ytdl

import (
	"os/exec"

	"github.com/EdTheLoon/GoTube/server"
)

func getDetails(url string) {

}

func getFormats(url string) {

}

func downloadVideo(url string, format string, s *server.Server) {
	out, err := exec.Command("youtube-dl",
		"-x",
		"-f "+format,
		"--metadata-from-title '%(artist)s - %(title)s'",
		"--audio-quality 0",
		"--embed-thumbnail",
		"-o '"+s.GetDownloadsDir()+"%(artist)s - %(title)s.%(ext)s'",
		"--add-metadata",
		url).Output()

	if err != nil {
		s.Log(err.Error())
	}
	s.Log(string(out))
}

func downloadAudio(url string, format string, s *server.Server) {
	out, err := exec.Command("youtube-dl",
		"-x",
		"-f bestaudio/best",
		"--audio-format "+format,
		"--metadata-from-title '%(artist)s - %(title)s'",
		"--audio-quality 0",
		"--embed-thumbnail",
		"-o '"+s.GetDownloadsDir()+"%(artist)s - %(title)s.%(ext)s'",
		"--add-metadata",
		url).Output()

	if err != nil {
		s.Log(err.Error())
	}
	s.Log(string(out))
}
