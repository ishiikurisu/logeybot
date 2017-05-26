package model

import "testing"

func TestGenerationOfTargetFilesBasedOnIDs(t *testing.T) {
    var id int64 = 123456
    var calculated string = getIdFile(id)
    if expected := "./data/logeybot/123456.txt"; expected != calculated {
        t.Fatalf("expected %s != calculated %s\n", expected, calculated)
    }
}
