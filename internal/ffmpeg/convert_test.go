package ffmpeg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAvcConvertCommand(t *testing.T) {
	t.Run("empty filename", func(t *testing.T) {
		Options := Options{
			Bin:      "",
			Encoder:  "intel",
			Size:     1500,
			Bitrate:  "50",
			MapVideo: "",
			MapAudio: "",
		}
		_, _, err := AvcConvertCommand("", "", Options)

		assert.Equal(t, err.Error(), "empty input filename")
	})
	t.Run("avc name empty", func(t *testing.T) {
		Options := Options{
			Bin:      "",
			Encoder:  "intel",
			Size:     1500,
			Bitrate:  "50",
			MapVideo: "",
			MapAudio: "",
		}
		_, _, err := AvcConvertCommand("VID123.mov", "", Options)

		assert.Equal(t, err.Error(), "empty output filename")
	})
	t.Run("animated file", func(t *testing.T) {
		Options := Options{
			Bin:      "",
			Encoder:  "intel",
			Size:     1500,
			Bitrate:  "50",
			MapVideo: "",
			MapAudio: "",
		}
		r, _, err := AvcConvertCommand("VID123.gif", "VID123.gif.avc", Options)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, " -i VID123.gif -movflags faststart -pix_fmt yuv420p -vf scale=trunc(iw/2)*2:trunc(ih/2)*2 -f mp4 -y VID123.gif.avc", r.String())
	})
	t.Run("libx264", func(t *testing.T) {
		Options := Options{
			Bin:      "",
			Encoder:  "libx264",
			Size:     1500,
			Bitrate:  "50",
			MapVideo: "",
			MapAudio: "",
		}
		r, _, err := AvcConvertCommand("VID123.mov", "VID123.mov.avc", Options)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, " -i VID123.mov -c:v libx264 -map  -map  -c:a aac -vf scale='if(gte(iw,ih), min(1500, iw), -2):if(gte(iw,ih), -2, min(1500, ih))',format=yuv420p -max_muxing_queue_size 1024 -crf 23 -vsync vfr -r 30 -b:v 50 -f mp4 -y VID123.mov.avc", r.String())
	})
	t.Run("h264_qsv", func(t *testing.T) {
		Options := Options{
			Bin:      "",
			Encoder:  "h264_qsv",
			Size:     1500,
			Bitrate:  "50",
			MapVideo: "",
			MapAudio: "",
		}
		r, _, err := AvcConvertCommand("VID123.mov", "VID123.mov.avc", Options)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, " -qsv_device /dev/dri/renderD128 -i VID123.mov -c:a aac -vf scale='if(gte(iw,ih), min(1500, iw), -2):if(gte(iw,ih), -2, min(1500, ih))',format=rgb32 -c:v h264_qsv -map  -map  -vsync vfr -r 30 -b:v 50 -bitrate 50 -f mp4 -y VID123.mov.avc", r.String())
	})
	t.Run("h264_videotoolbox", func(t *testing.T) {
		Options := Options{
			Bin:      "",
			Encoder:  "h264_videotoolbox",
			Size:     1500,
			Bitrate:  "50",
			MapVideo: "",
			MapAudio: "",
		}
		r, _, err := AvcConvertCommand("VID123.mov", "VID123.mov.avc", Options)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, " -i VID123.mov -c:v h264_videotoolbox -map  -map  -c:a aac -vf scale='if(gte(iw,ih), min(1500, iw), -2):if(gte(iw,ih), -2, min(1500, ih))',format=yuv420p -profile high -level 51 -vsync vfr -r 30 -b:v 50 -f mp4 -y VID123.mov.avc", r.String())
	})
	t.Run("h264_vaapi", func(t *testing.T) {
		Options := Options{
			Bin:      "",
			Encoder:  "h264_vaapi",
			Size:     1500,
			Bitrate:  "50",
			MapVideo: "",
			MapAudio: "",
		}
		r, _, err := AvcConvertCommand("VID123.mov", "VID123.mov.avc", Options)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, " -hwaccel vaapi -i VID123.mov -c:a aac -vf scale='if(gte(iw,ih), min(1500, iw), -2):if(gte(iw,ih), -2, min(1500, ih))',format=nv12,hwupload -c:v h264_vaapi -map  -map  -vsync vfr -r 30 -b:v 50 -f mp4 -y VID123.mov.avc", r.String())
	})
	t.Run("h264_nvenc", func(t *testing.T) {
		Options := Options{
			Bin:      "",
			Encoder:  "h264_nvenc",
			Size:     1500,
			Bitrate:  "50",
			MapVideo: "",
			MapAudio: "",
		}
		r, _, err := AvcConvertCommand("VID123.mov", "VID123.mov.avc", Options)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, " -hwaccel auto -i VID123.mov -pix_fmt yuv420p -c:v h264_nvenc -map  -map  -c:a aac -preset 15 -pixel_format yuv420p -gpu any -vf scale='if(gte(iw,ih), min(1500, iw), -2):if(gte(iw,ih), -2, min(1500, ih))',format=yuv420p -rc:v constqp -cq 0 -tune 2 -r 30 -b:v 50 -profile:v 1 -level:v auto -coder:v 1 -f mp4 -y VID123.mov.avc", r.String())
	})
	t.Run("h264_v4l2m2m", func(t *testing.T) {
		Options := Options{
			Bin:      "",
			Encoder:  "h264_v4l2m2m",
			Size:     1500,
			Bitrate:  "50",
			MapVideo: "",
			MapAudio: "",
		}
		r, _, err := AvcConvertCommand("VID123.mov", "VID123.mov.avc", Options)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, " -i VID123.mov -c:v h264_v4l2m2m -map  -map  -c:a aac -vf scale='if(gte(iw,ih), min(1500, iw), -2):if(gte(iw,ih), -2, min(1500, ih))',format=yuv420p -num_output_buffers 72 -num_capture_buffers 64 -max_muxing_queue_size 1024 -crf 23 -vsync vfr -r 30 -b:v 50 -f mp4 -y VID123.mov.avc", r.String())
	})
}
