package pixiv

import (
    "testing"
)

func TestGetNovel(t *testing.T){
    input := []string{"17576664", "1299320", "10290316", "10539710", "1268983"}

    for _, id := range input {
        _, err := GetNovelByID(id)
        if err != nil {
            t.Log(id)
            t.Error(err)
        }
    }
}
