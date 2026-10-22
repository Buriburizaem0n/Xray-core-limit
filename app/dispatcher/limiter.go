package dispatcher

import (
	"context"
	"errors"
	"io"
	"sync"
	"time"

	"golang.org/x/time/rate"

	"github.com/xtls/xray-core/common"
	"github.com/xtls/xray-core/common/buf"
	"github.com/xtls/xray-core/features/stats"
)

var errInvalidLimiterBurst = errors.New("rate limiter burst must be positive")

func waitForBytes(ctx context.Context, limiter *rate.Limiter, size int) error {
	if limiter == nil || limiter.Limit() <= 0 || size <= 0 {
		return nil
	}
	burst := limiter.Burst()
	if burst <= 0 {
		return errInvalidLimiterBurst
	}
	for size > 0 {
		n := min(size, burst)
		if err := limiter.WaitN(ctx, n); err != nil {
			return err
		}
		size -= n
	}
	return nil
}

// RateLimitedWriter wraps a buf.Writer with rate limiting using golang.org/x/time/rate
type RateLimitedWriter struct {
	Writer  buf.Writer
	Counter stats.Counter
	limiter *rate.Limiter

	ctx    context.Context
	cancel context.CancelFunc
	mu     sync.Mutex
	closed bool
}

// NewRateLimitedWriter creates a rate-limited writer.
func NewRateLimitedWriter(w buf.Writer, limiter *rate.Limiter, counter stats.Counter) *RateLimitedWriter {
	ctx, cancel := context.WithCancel(context.Background())

	return &RateLimitedWriter{
		Writer:  w,
		Counter: counter,
		limiter: limiter,
		ctx:     ctx,
		cancel:  cancel,
	}
}

// WriteMultiBuffer implements buf.Writer with rate limiting
func (w *RateLimitedWriter) WriteMultiBuffer(mb buf.MultiBuffer) error {
	w.mu.Lock()
	if w.closed {
		w.mu.Unlock()
		buf.ReleaseMulti(mb)
		return io.ErrClosedPipe
	}
	w.mu.Unlock()

	if err := waitForBytes(w.ctx, w.limiter, int(mb.Len())); err != nil {
		buf.ReleaseMulti(mb)
		return err
	}

	if w.Counter != nil {
		w.Counter.Add(int64(mb.Len()))
	}
	return w.Writer.WriteMultiBuffer(mb)
}

// Close implements common.Closable
func (w *RateLimitedWriter) Close() error {
	w.mu.Lock()
	if !w.closed {
		w.closed = true
		w.cancel()
	}
	w.mu.Unlock()
	return common.Close(w.Writer)
}

// Interrupt implements common.Interruptible
func (w *RateLimitedWriter) Interrupt() {
	w.mu.Lock()
	if !w.closed {
		w.closed = true
		w.cancel()
	}
	w.mu.Unlock()
	common.Interrupt(w.Writer)
}

// RateLimitedReader wraps a buf.Reader with rate limiting using golang.org/x/time/rate
type RateLimitedReader struct {
	Reader  buf.Reader
	Counter stats.Counter
	limiter *rate.Limiter

	ctx    context.Context
	cancel context.CancelFunc
	mu     sync.Mutex
	closed bool
}

// NewRateLimitedReader creates a rate-limited reader.
func NewRateLimitedReader(r buf.Reader, limiter *rate.Limiter, counter stats.Counter) *RateLimitedReader {
	ctx, cancel := context.WithCancel(context.Background())

	return &RateLimitedReader{
		Reader:  r,
		Counter: counter,
		limiter: limiter,
		ctx:     ctx,
		cancel:  cancel,
	}
}

// ReadMultiBuffer implements buf.Reader with rate limiting
func (r *RateLimitedReader) ReadMultiBuffer() (buf.MultiBuffer, error) {
	r.mu.Lock()
	if r.closed {
		r.mu.Unlock()
		return nil, io.ErrClosedPipe
	}
	r.mu.Unlock()

	mb, err := r.Reader.ReadMultiBuffer()
	if err != nil && mb.IsEmpty() {
		return mb, err
	}

	if waitErr := waitForBytes(r.ctx, r.limiter, int(mb.Len())); waitErr != nil {
		buf.ReleaseMulti(mb)
		return nil, waitErr
	}

	if r.Counter != nil {
		r.Counter.Add(int64(mb.Len()))
	}
	return mb, err
}

// ReadMultiBufferTimeout implements buf.TimeoutReader with rate limiting
func (r *RateLimitedReader) ReadMultiBufferTimeout(timeout time.Duration) (buf.MultiBuffer, error) {
	r.mu.Lock()
	if r.closed {
		r.mu.Unlock()
		return nil, io.ErrClosedPipe
	}
	r.mu.Unlock()

	var mb buf.MultiBuffer
	var err error

	if tr, ok := r.Reader.(buf.TimeoutReader); ok {
		mb, err = tr.ReadMultiBufferTimeout(timeout)
	} else {
		mb, err = r.Reader.ReadMultiBuffer()
	}

	if err != nil && mb.IsEmpty() {
		return mb, err
	}

	if waitErr := waitForBytes(r.ctx, r.limiter, int(mb.Len())); waitErr != nil {
		buf.ReleaseMulti(mb)
		return nil, waitErr
	}

	if r.Counter != nil {
		r.Counter.Add(int64(mb.Len()))
	}
	return mb, err
}

// Interrupt implements common.Interruptible
func (r *RateLimitedReader) Interrupt() {
	r.mu.Lock()
	if !r.closed {
		r.closed = true
		r.cancel()
	}
	r.mu.Unlock()
	common.Interrupt(r.Reader)
}

// Close implements common.Closable
func (r *RateLimitedReader) Close() error {
	r.mu.Lock()
	if !r.closed {
		r.closed = true
		r.cancel()
	}
	r.mu.Unlock()
	return common.Close(r.Reader)
}
