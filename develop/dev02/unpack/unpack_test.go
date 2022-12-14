package unpack

import "testing"


func TestAll(t *testing.T) {
    t.Run("Test escapes", Test1)
    t.Run("Test bad start with digit", Test2)
    t.Run("Test bad eol after \\", Test3)
    t.Run("Test repeats", Test4)
}

func Test1(t *testing.T) {
    str := `qwe\1`
    got, _ := Unpack(str)
    want := "qwe1"
    if got != want {
        t.Errorf("Unpack('qwe\\1') = %s; want qwe1", got)
    }
}

func Test2(t *testing.T) {
    str := `1qwe\1`
    _, err := Unpack(str)
    if err == nil {
        t.Errorf("Expected error, got nil")
    }
}

func Test3(t *testing.T) {
    str := `1qwe\`
    _, err := Unpack(str)
    if err == nil {
        t.Errorf("Expected error, got nil")
    }
}

func Test4(t *testing.T) {
    str := `a5b10`
    got, _ := Unpack(str)
    want := "aaaaabbbbbbbbbb"
    if got != want {
        t.Errorf("Unpack('a5b10') = %s; want aaaaabbbbbbbbbb", got)
    }
}

