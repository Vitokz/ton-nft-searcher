package logger

import (
	"io"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/diode"
)

var _ zerolog.LevelWriter = &FilteredWriter{}

// FilteredWriter defines writer for specified levels to filter.
type FilteredWriter struct {
	writer io.Writer
	levels []zerolog.Level
}

// Write writes len(p) bytes from p to the underlying data stream.
func (w *FilteredWriter) Write(p []byte) (int, error) {
	return w.writer.Write(p)
}

// WriteLevel implements LevelWriter interface.
func (w *FilteredWriter) WriteLevel(level zerolog.Level, p []byte) (int, error) {
	if len(w.levels) == 0 {
		return w.writer.Write(p)
	}

	for _, filteredLevel := range w.levels {
		if level == filteredLevel {
			return w.writer.Write(p)
		}
	}

	return len(p), nil
}

// AsyncWrap wraps writer with diode to write messages in async mode.
func (w *FilteredWriter) AsyncWrap() {
	w.writer = diode.NewWriter(
		w.writer,
		defaultWriterBufferSize,
		defaultWriterPollInterval,
		nil,
	)
}
