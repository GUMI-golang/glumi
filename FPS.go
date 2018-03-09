package glumi

import "time"

type FPS interface {
	Start()
	Wait() time.Time
	Stop()
} 
type IntervalFPS  struct {
	interval time.Duration
	ticker   *time.Ticker
}
func (s *IntervalFPS) Start()  {
	s.ticker = time.NewTicker(s.interval)
}
func (s *IntervalFPS) Wait() time.Time {
	return <- s.ticker.C
}
func (s *IntervalFPS) Stop()  {
	s.ticker.Stop()
	s.ticker = nil
}
func (s *IntervalFPS) SetInterval(interval time.Duration)  {
	s.interval = interval
}
func (s *IntervalFPS) GetInterval()(interval time.Duration)  {
	return s.interval
}

type LimitlessFPS struct {

}
func (s *LimitlessFPS) Start()  {
}
func (s *LimitlessFPS) Wait() time.Time {
	return time.Now()
}
func (s *LimitlessFPS) Stop()  {
}