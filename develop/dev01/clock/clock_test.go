package clock

import "testing"

func TestAll(t *testing.T) {
	t.Run("Test Good Url", TestGoodUrl)
	t.Run("Test Bad Url", TestBadUrl)
}

func TestGoodUrl(t *testing.T) {
	_, err := GetDate("0.beevik-ntp.pool.ntp.org")
	if err != nil{
		t.Errorf("Unexpected error: %s", err)
	}
}

func TestBadUrl(t *testing.T) {
	_, err := GetDate("bad")
	if err == nil{
		t.Errorf("Expected error, but got nil")
	}
}